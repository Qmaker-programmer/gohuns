# GOHUNS — Corrector Ortográfico Estático CLI

> Un corrector interactivo, ultra rápido y **MINIMALISTA AL EXTREMO** que vive en tus pipelines 🖥️  
> *Porque ver errores ortográficos en tu terminal debe doler a la vista, directo al grano.*

---

## Características Clave ⚡

* **✓ Corrección Raw:** Escupe tu texto idéntico por `stdout`, pero resalta los typos en rojo brillante. Al toque.
* **✓ Caché Local Integrada:** Descarga diccionarios Hunspell públicos y los indexa sin rodeos en `~/.gohuns/`.
* **✓ Cero Basura Visual:** Formato limpio `lang: [nombre]` o `err: [causa]` ideal para automatizar con scripts de bash.
* **✓ Optimización Silenciosa:** Parsea los archivos `.aff` pesados quitando líneas rotas (`MAP`) en milisegundos.

---

## 🚀 Instalación al Toque

### Prerrequisitos
* **Go 1.21+** y **Git**. No necesitas nada más.

### Setup en 3 patadas
**Si quieres el binario directo sin compilar:** [Releases en GitHub](https://github.com/Qmaker-programmer/gohuns/releases)

---

**Si lo haces tú mismo con Make (Recomendado):**

```bash
git clone [https://github.com/Qmaker-programmer/gohuns.git](https://github.com/Qmaker-programmer/gohuns.git)
cd gohuns
go mod tidy

# El Makefile hace el trabajo pesado:
make run      # Ejecuta el src/main.go directo
make build    # Te genera el ejecutable para tu arquitectura
make all      # Compila los 6 binarios estáticos (Linux, Win, Mac - AMD64/ARM64)

```

---

## ⌨️ Banderas y Comandos

Sin interfaces lentas. Solo `flags` nativas:

### 📥 1. Descarga e Indexación (`-d`)

Bájate el `.aff` y `.dic` de internet asignándole un identificador con `-n`:

```bash
gohuns -d -n es_ES [https://url.com/es_ES.aff](https://url.com/es_ES.aff) [https://url.com/es_ES.dic](https://url.com/es_ES.dic)

```

*Output esperado:* `ok: 'es_ES' instalado en caché`

### 📚 2. Listar lo que tienes (`-list`)

Imprime rápido los lenguajes listos en tu sistema:

```bash
gohuns -list

```

*Output esperado:*

```text
lang: es_ES
lang: es_CL

```

### 🔍 3. Corrección Interactiva

Usa pipes (`|`) o redirecciones (`<`). Si omites `-l`, usa `es_ES` por defecto.

**Frase rápida:**

```bash
echo "Esta frase tiene un herror" | gohuns -l es_ES

```

**Documento completo:**

```bash
gohuns -l es_ES < logs_o_codigo.txt

```

### ❓ 4. Ayuda Rápida

```bash
gohuns -help

```

---

## 🛠️ Estructura Limpia del Proyecto

```text
.
├── bin/                        # Binarios listos (generados con make all)
│   ├── gohuns-darwin-amd64
│   ├── gohuns-darwin-arm64
│   ├── gohuns-linux-arm64
│   ├── gohuns-linux-amd64
│   ├── gohuns-windows-amd64.exe
│   └── gohuns-windows-arm64.exe
├── examples/
│   └── bad_Text                # Archivo con typos horribles para probar
├── go.mod / go.sum             # Módulos de Go
├── Makefile                    # Multi-compilador rápido
└── src/
    └── main.go                 # TODO el código fuente (un solo archivo, sin vueltas)
```

---

## 📊 Comparación de Peso Pluma

| Feature | GOHUNS | Otras Herramientas Pesadas |
| --- | --- | --- |
| Peso del Binario | Ligero y estático (`-ldflags="-s -w"`) | Gigante (Node/Electron) |
| Integración Bash | ✅ Perfecto para pipes (` | `) |
| Output formateado | Minimalista (`lang:`, `err:`) | Mensajes largos y molestos |

---

## 🐛 Erres y Dramas (`err:`)

**`err: falta diccionario 'es_US'`**

* **Causa:** No lo has descargado. Corre: `gohuns -d -n es_US <url_aff> <url_dic>`.

**`err: json corrupto`**

* **Causa:** Modificaste el archivo `~/.gohuns/map.json` a mano y rompiste algo. Bórralo y vuelve a descargar.

---

Desarrollado de forma limpia, minimalista y con **mucho Té** por [Qmaker](https://github.com/Qmaker-programmer). 🚩
