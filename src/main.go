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

type RegistroIdioma struct {
	Nombre string `json:"nombre"`
	UrlAff string `json:"url_aff"`
	UrlDic string `json:"url_dic"`
}

func obtenerRutasBase() (string, string) {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: no se obtuvo el directorio HOME: %v\n", err)
		os.Exit(1)
	}
	
	dirLenguajes := filepath.Join(home, ".gohuns", "lenguajes")
	rutaJson := filepath.Join(home, ".gohuns", "map.json")
	
	if err := os.MkdirAll(dirLenguajes, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "err: fallo al crear carpetas de configuración: %v\n", err)
		os.Exit(1)
	}
	
	return dirLenguajes, rutaJson
}

func guardarEnMapaJson(rutaJson, nombre, urlAff, urlDic string) {
	var registros []RegistroIdioma

	if _, err := os.Stat(rutaJson); err == nil {
		data, err := os.ReadFile(rutaJson)
		if err == nil {
			json.Unmarshal(data, &registros)
		}
	}

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

	data, err := json.MarshalIndent(registros, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: fallo al codificar JSON: %v\n", err)
		return
	}

	err = os.WriteFile(rutaJson, data, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: no se pudo escribir map.json: %v\n", err)
	}
}

func descargarYLimpiarAff(url, caminoDestino string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP %s", resp.Status)
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
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP %s", resp.Status)
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
		// Minimalismo puro: imprime en rojo el error tal cual
		color.Set(color.FgRed, color.Bold)
		fmt.Print(palabra)
		color.Unset()
	}
}

func main() {
	dirLenguajes, rutaJson := obtenerRutasBase()

	langPtr := flag.String("l", "es_ES", "Idioma para corregir")
	downloadPtr := flag.Bool("d", false, "Modo descarga")
	namePtr := flag.String("n", "", "Nombre del idioma")
	listPtr := flag.Bool("list", false, "Lista diccionarios disponibles")

	flag.Usage = func() {
		color.Cyan("🛡️ GOHUNS — Corrector Ortográfico Minimalista CLI\n")
		fmt.Println("Uso:")
		fmt.Println("  gohuns -l <idioma> < texto.txt")
		fmt.Println("  echo \"texto\" | gohuns -l <idioma>\n")
		fmt.Println("Banderas:")
		flag.PrintDefaults()
		fmt.Println("\nEjemplos:")
		fmt.Println("  gohuns -list")
		fmt.Println("  gohuns -d -n es_CL <url_aff> <url_dic>")
	}

	flag.Parse()

	// MODO LISTAR DICCIONARIOS (Ultra minimalista)
	if *listPtr {
		if _, err := os.Stat(rutaJson); os.IsNotExist(err) {
			fmt.Println("cache: vacía (no hay diccionarios en map.json)")
			return
		}

		data, err := os.ReadFile(rutaJson)
		if err != nil {
			fmt.Fprintf(os.Stderr, "err: no se pudo leer el índice: %v\n", err)
			os.Exit(1)
		}

		var registros []RegistroIdioma
		if err := json.Unmarshal(data, &registros); err != nil {
			fmt.Fprintf(os.Stderr, "err: json corrupto: %v\n", err)
			os.Exit(1)
		}

		if len(registros) == 0 {
			fmt.Println("cache: vacía")
			return
		}
		for _, r := range registros {
			fmt.Printf("lang: %s\n", color.CyanString(r.Nombre))
		}
		return
	}

	// MODO DESCARGA
	if *downloadPtr {
		urls := flag.Args()
		if *namePtr == "" || len(urls) < 2 {
			fmt.Fprintln(os.Stderr, "err: faltan parámetros de descarga.")
			fmt.Fprintln(os.Stderr, "uso: gohuns -d -n <nombre> <url_aff> <url_dic>")
			os.Exit(1)
		}

		urlAff := urls[0]
		urlDic := urls[1]

		caminoAff := filepath.Join(dirLenguajes, *namePtr+".aff")
		caminoDic := filepath.Join(dirLenguajes, *namePtr+".dic")

		if err := descargarYLimpiarAff(urlAff, caminoAff); err != nil {
			fmt.Fprintf(os.Stderr, "err: aff falló (%v)\n", err)
			os.Exit(1)
		}
		
		if err := descargarDic(urlDic, caminoDic); err != nil {
			fmt.Fprintf(os.Stderr, "err: dic falló (%v)\n", err)
			os.Exit(1)
		}

		guardarEnMapaJson(rutaJson, *namePtr, urlAff, urlDic)
		fmt.Printf("ok: '%s' instalado en caché\n", *namePtr)
		return
	}

	// MODO CORRECTOR
	affPath := filepath.Join(dirLenguajes, *langPtr+".aff")
	dicPath := filepath.Join(dirLenguajes, *langPtr+".dic")

	speller, err := gospell.NewGoSpell(affPath, dicPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: falta diccionario '%s'\n", *langPtr)
		fmt.Fprintf(os.Stderr, "descarga: gohuns -d -n %s <url_aff> <url_dic>\n", *langPtr)
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
