# GOHUNS — Corrector Ortográfico Estático CLI

> Un corrector interactivo, ultra rápido y **MINIMALISTA AL EXTREMO** que vive en tus pipelines y flujos de terminal 🖥️
> *Porque ver errores ortográficos en tu CLI debe doler a la vista, directo al grano.*

**Preview**

<img width="712" height="375" alt="image" src="https://github.com/user-attachments/assets/892c4cd9-1b64-4f67-94d9-8a49b78d6033" />

---

## Características Clave ⚡

* **✓ Corrección Raw:** Escupe tu texto idéntico por `stdout`, pero resalta los typos en rojo brillante. Al toque.
* **✓ Caché Local Integrada:** Descarga diccionarios Hunspell públicos y los indexa sin rodeos en `~/.gohuns/`.
* **✓ Cero Basura Visual:** Formato limpio `lang: [nombre]` o `err: [causa]` ideal para automatizar con scripts de bash.
* **✓ Optimización Silenciosa:** Parsea los archivos `.aff` pesados quitando líneas rotas (`MAP`) en milisegundos.
* **✓ Manual de Sistema:** Documentación integrada nativamente con el comando `man` tradicional de Linux.

---

## 🚀 Instalación y Empaquetado
**Si quieres instalarlo facil hazlo de aqui:** [Releases (Binarios listos!)](https://github.com/Qmaker-programmer/gohuns/releases)
**O si eres hardcore aqui la forma manual: **
### Prerrequisitos

* **Go 1.21+** y **Git**. No necesitas nada más.

### Setup Automatizado con Makefile

Clona el repositorio, descarga las dependencias y deja que el `Makefile` inteligente haga el trabajo pesado en base a lo que busques:

```bash
git clone https://github.com/Qmaker-programmer/gohuns.git
cd gohuns
go mod tidy

```

#### Comandos Disponibles:

* `make run`: Ejecuta el código fuente `src/main.go` en vivo para pruebas rápidas.
* `make build`: Genera el ejecutable optimizado (`-ldflags="-s -w"`) adaptado únicamente a tu arquitectura actual.
* `make all`: Compila por cruzado los 6 binarios estáticos limpios (Linux, Windows y macOS en arquitecturas AMD64 y ARM64) dentro del directorio `bin/`.
* `make all-debs`: Compila los binarios y empaqueta de forma 100% automática los instaladores nativos `.deb` (para arquitecturas `amd64` e `arm64`) sincronizando la versión global del proyecto en sus metadatos internos y en las páginas del manual.
* `make clean`: Borra de un plumazo los binarios compilados, instaladores y residuos temporales para dejar el repositorio reluciente.

---

### Instalación Nativa en Linux (.deb)

Si generaste los instaladores usando `make all-debs`, puedes instalar `gohuns` de forma global en distribuciones basadas en Debian/Ubuntu corriendo:

```bash
sudo apt install ./bin/gohuns_0.0.2_amd64.deb

```

*Esto configurará el binario directamente en tu `/usr/bin/` e inyectará las páginas de ayuda oficiales en el sistema.*

---

## ⌨️ Banderas y Comandos

Sin interfaces lentas ni configuraciones complejas. Solo flags directas:

### 📥 1. Descarga e Indexación (`-d`)

Bájate el `.aff` y `.dic` de internet asignándole un identificador corto con la bandera `-n`:

```bash
gohuns -d -n es_ES https://url.com/es_ES.aff https://url.com/es_ES.dic

```

*Output esperado:* `ok: 'es_ES' instalado en caché`

### 📚 2. Listar lo que tienes (`-list`)

Imprime rápido los lenguajes listos en tu sistema:

```bash
gohuns -list

```

### 🔍 3. Corrección Interactiva

Usa pipes (`|`) o redirecciones clásicas de entrada (`<`). Si omites `-l`, el sistema usará `es_ES` por defecto.

**Frase rápida por stream:**

```bash
echo "Esta frase tiene un herror" | gohuns -l es_ES

```

**Documento o logs completos:**

```bash
gohuns -l es_ES < logs_o_codigo.txt

```

### ❓ 4. Ayuda y Documentación Nativa

Si solo quieres un recordatorio rápido de las banderas:

```bash
gohuns -help

```

Si deseas ver la documentación completa estructurada del proyecto directo desde el gestor de manuales de Linux:

```bash
man gohuns

```

---

## 🛠️ Estructura del Proyecto

```text
.
├── bin/                        # Distribución final ordenada
│   ├── gohuns_0.0.2_amd64.deb  # Instalador Debian/Ubuntu AMD64
│   ├── gohuns_0.0.2_arm64.deb  # Instalador Debian/Ubuntu ARM64
│   ├── gohuns-darwin-amd64     # Binario macOS Intel
│   ├── gohuns-darwin-arm64     # Binario macOS Apple Silicon
│   ├── gohuns-linux-amd64      # Binario Linux AMD64
│   ├── gohuns-linux-arm64      # Binario Linux ARM64
│   ├── gohuns-windows-amd64.exe# Binario Windows AMD64
│   └── gohuns-windows-arm64.exe# Binario Windows ARM64
├── examples/
│   └── bad_Text                # Archivo con typos para realizar pruebas
├── gohuns.1                    # Plantilla base del manual para el comando 'man'
├── go.mod / go.sum             # Módulos y dependencias de Go
├── LICENSE                     # Licencia del proyecto
├── Makefile                    # Automatizador multiplataforma y empaquetador inteligente
└── src/
    └── main.go                 # Todo el código fuente (sin vueltas)

```

---

## 📊 Comparación de Peso Pluma

| Característica | GOHUNS | Otras Herramientas |
| --- | --- | --- |
| **Peso del Binario** | Ultra ligero y estático (`-ldflags="-s -w"`) | Paquetes gigantescos basados en Node/Electron |
| **Integración Bash** | Perfecto para pipelines nativos (` | `/`<`) |
| **Output Formateado** | Minimalista (`lang:`, `err:`) | Formatos visuales pesados o ruidosos |
| **Documentación** | Integrada en el sistema mediante `man` | Sitios web externos obligatorios |

---

## 🐛 Errores Comunes (`err:`)

**`err: falta diccionario 'es_US'`**

* **Causa:** El diccionario solicitado no ha sido descargado. Corre: `gohuns -d -n es_US <url_aff> <url_dic>`.

**`err: json corrupto`**

* **Causa:** Modificaste manualmente el índice `~/.gohuns/map.json` y se rompió el formato. Elimina el archivo y vuelve a descargar tus diccionarios con la flag `-d`.

---

Desarrollado de forma limpia, minimalista y con **mucho Té** por [Qmaker](https://github.com/Qmaker-programmer). 🚩
