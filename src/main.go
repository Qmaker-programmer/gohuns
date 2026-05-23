package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/client9/gospell"
	"github.com/fatih/color"
)

// Estructura para el mapa JSON
type RegistroIdioma struct {
	Nombre string `json:"nombre"`
	UrlAff string `json:"url_aff"`
	UrlDic string `json:"url_dic"`
}

func obtenerRutasBase() (string, string) {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error obteniendo el directorio HOME: %v\n", err)
		os.Exit(1)
	}
	
	// Carpeta global: ~/.gohuns/lenguajes
	dirLenguajes := filepath.Join(home, ".gohuns", "lenguajes")
	rutaJson := filepath.Join(home, ".gohuns", "map.json")
	
	// Crear la estructura de carpetas si no existe
	if err := os.MkdirAll(dirLenguajes, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creando carpetas de configuración: %v\n", err)
		os.Exit(1)
	}
	
	return dirLenguajes, rutaJson
}

func guardarEnMapaJson(rutaJson, nombre, urlAff, urlDic string) {
	var registros []RegistroIdioma

	// Si el archivo ya existe, leer su contenido actual
	if _, err := os.Stat(rutaJson); err == nil {
		data, err := os.ReadFile(rutaJson)
		if err == nil {
			json.Unmarshal(data, &registros)
		}
	}

	// Crear o actualizar el registro
	nuevoRegistro := RegistroIdioma{Nombre: nombre, UrlAff: urlAff, UrlDic: urlDic}
	actualizado := false
	for i, r := range registros {
		if r.Nombre == nombre {
			registros[i] = nuevoRegistro
			actualizado = true
			break
		}
	}
	if !actualizado {
		registros = append(registros, nuevoRegistro)
	}

	// Guardar el JSON con formato legible
	data, err := json.MarshalIndent(registros, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error codificando JSON: %v\n", err)
		return
	}

	err = os.WriteFile(rutaJson, data, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error escribiendo map.json: %v\n", err)
	}
}

func descargarYLimpiarAff(url, caminoDestino string) error {
	fmt.Printf("Descargando y optimizando: %s...\n", filepath.Base(caminoDestino))
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error de servidor: %s", resp.Status)
	}

	out, err := os.Create(caminoDestino)
	if err != nil {
		return err
	}
	defer out.Close()

	scanner := bufio.NewScanner(resp.Body)
	writer := bufio.NewWriter(out)

	for scanner.Scan() {
		linea := scanner.Text()
		if strings.HasPrefix(linea, "MAP") {
			continue
		}
		_, err = writer.WriteString(linea + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

func descargarDic(url, caminoDestino string) error {
	fmt.Printf("Descargando: %s...\n", filepath.Base(caminoDestino))
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error de servidor: %s", resp.Status)
	}

	out, err := os.Create(caminoDestino)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func procesarPalabra(palabra string, speller *gospell.GoSpell) {
	if speller.Spell(palabra) {
		fmt.Print(palabra)
	} else {
		color.Set(color.FgRed, color.Bold)
		fmt.Print(palabra)
		color.Unset()
	}
}

func main() {
	dirLenguajes, rutaJson := obtenerRutasBase()

	langPtr := flag.String("l", "es_ES", "Idioma para corregir (ej: es_CL)")
	downloadPtr := flag.Bool("d", false, "Activa el modo descarga")
	namePtr := flag.String("n", "", "Nombre personalizado para el idioma a descargar")

	flag.Parse()

	// MODO DESCARGA
	if *downloadPtr {
		urls := flag.Args()
		if *namePtr == "" || len(urls) < 2 {
			fmt.Fprintln(os.Stderr, "Error: Faltan parámetros para la descarga.")
			fmt.Fprintln(os.Stderr, "Uso: gohuns -d -n <nombre> <url_aff> <url_dic>")
			os.Exit(1)
		}

		urlAff := urls[0]
		urlDic := urls[1]

		caminoAff := filepath.Join(dirLenguajes, *namePtr+".aff")
		caminoDic := filepath.Join(dirLenguajes, *namePtr+".dic")

		if err := descargarYLimpiarAff(urlAff, caminoAff); err != nil {
			fmt.Fprintf(os.Stderr, "Error procesando .aff: %v\n", err)
			os.Exit(1)
		}
		
		if err := descargarDic(urlDic, caminoDic); err != nil {
			fmt.Fprintf(os.Stderr, "Error procesando .dic: %v\n", err)
			os.Exit(1)
		}

		// Registrar la descarga en el map.json de forma automática
		guardarEnMapaJson(rutaJson, *namePtr, urlAff, urlDic)
		
		fmt.Printf("¡Éxito! Idioma '%s' instalado globalmente en ~/.gohuns/lenguajes/\n", *namePtr)
		return
	}

	// MODO CORRECTOR (Global)
	affPath := filepath.Join(dirLenguajes, *langPtr+".aff")
	dicPath := filepath.Join(dirLenguajes, *langPtr+".dic")

	speller, err := gospell.NewGoSpell(affPath, dicPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: No se encontró el diccionario '%s' en el almacenamiento global.\n", *langPtr)
		fmt.Fprintf(os.Stderr, "Usa: gohuns -d -n %s <url_aff> <url_dic> para instalarlo.\n", *langPtr)
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	var wordBuffer []rune

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			break
		}

		if unicode.IsLetter(r) {
			wordBuffer = append(wordBuffer, r)
		} else {
			if len(wordBuffer) > 0 {
				procesarPalabra(string(wordBuffer), speller)
				wordBuffer = wordBuffer[:0]
			}
			fmt.Print(string(r))
		}
	}

	if len(wordBuffer) > 0 {
		procesarPalabra(string(wordBuffer), speller)
	}
}
