# GOHUNS — Corrector Ortográfico CLI & Gestor de Diccionarios

> Un corrector ortográfico interactivo, ultrarrápido y **FACHERO** que vive en tu terminal 🖥️  
> *Porque enviar correos con faltas de ortografía en pleno 2026 es de código espagueti, amigo.*

---

## Características Clave 🔥

* **✓ Corrección en Flujo (Stream):** Procesa texto en tiempo real desde la entrada estándar (`stdin`) y te escupe los errores en rojo brillante.
* **✓ Gestor de Idiomas DIY:** ¿No tienes el diccionario? Le pasas dos URLs y se descarga, optimiza e instala el idioma globalmente.
* **✓ Sanitización al Vuelo:** Limpia las líneas conflictivas (`MAP`) de los archivos `.aff` para que el motor no explote.
* **✓ Cero Base de Datos:** Todo indexado en un humilde y confiable archivo JSON (`map.json`).

---

## 🚀 Instalación rápida

### Prerrequisitos
* **Go 1.21+** (si no lo tienes, estás viviendo en el pasado)
* **Git** (para clonar esta obra de arte)
* Un terminal que soporte colores (si usas la consola de MS-DOS de 1995, se va a ver feo)

### Setup en 3 pasos
**Si te da flojera compilar, ve directo a Releases y bájate el binario ya masticado: [Instalación desde Releases](https://github.com/Qmaker-programmer/gohuns/releases)**

---

**Si eres un verdadero programador y quieres meterte las manos en la masa:**

```bash
# 1️⃣ Clona el juguete y entra al directorio
git clone [https://github.com/Qmaker-programmer/gohuns.git](https://github.com/Qmaker-programmer/gohuns.git)
cd gohuns

# 2️⃣ Descarga las dependencias (Go hace toda la magia negra)
go mod tidy

# 3️⃣ Pruébalo en tiempo de desarrollo
go run src/main.go

```

O compila un binario permanente como dios manda:

```bash
# Usando el maravilloso Makefile incluido (si tienes make instalado)
make run      # Ejecución directa para ver si no rompiste nada
make build    # Compila el binario optimizado para tu sistema actual
make clean    # Borra los binarios por si quieres empezar de cero

# ¿Quieres lanzar tu app al mundo? Compila para Windows, Linux y macOS en AMD64 y ARM64 en un solo comando:
make all

# ¡Todo quedará empaquetado en bin/ listo para producción!

```

---

## 📦 Almacenamiento Global — Dónde viven los diccionarios

Para no ensuciar tu espacio de trabajo, `gohuns` crea una fortaleza en tu directorio `HOME` (`~/.gohuns/`):

* `map.json` — El registro civil de tus idiomas. Guarda qué has bajado y de dónde (`url_aff`, `url_dic`).
* `lenguajes/` — La carpeta donde se almacenan los archivos `.aff` y `.dic` extraídos de internet.

---

## ⌨️ Modo de Uso

`gohuns` no tiene una interfaz gráfica aburrida; se controla con el poder de las `flags`:

### 📥 1. Modo Descarga e Instalación (`-d`)

Antes de corregir, necesitas munición (diccionarios). Tienes que darle un nombre con `-n` y pasarle las URLs del `.aff` y `.dic`.

```bash
gohuns -d -n es_ES [https://ejemplo.com/es_ES.aff](https://ejemplo.com/es_ES.aff) [https://ejemplo.com/es_ES.dic](https://ejemplo.com/es_ES.dic)

```

*¡Boom! Guardado automáticamente en `~/.gohuns/lenguajes/es_ES.aff` e indexado en el JSON.*

### 🔍 2. Modo Corrector (Por Defecto)

Canaliza cualquier texto usando tuberías (`|`) o redirecciones (`<`). Si no especificas el idioma con `-l`, por defecto buscará `es_ES`.

**Corregir una frase rápida (Input directo):**

```bash
echo "Este texto tiene una palabra herronea" | gohuns -l es_ES

```

*(Verás "herronea" brillando en un hermoso color rojo de advertencia).*

**Corregir un archivo entero:**

```bash
gohuns -l es_ES < mi_documento_infestado_de_errores.txt

```

---

## 🛠️ Estructura del Proyecto

El repositorio está más limpio que tu historial de navegación:

```text
.
├── bin/                        # 🚀 ¡Los 6 binarios estáticos viven aquí tras hacer make all!
│   ├── gohuns-darwin-amd64
│   ├── gohuns-darwin-arm64
│   ├── gohuns-linux-arm64
│   ├── gohuns-windows-amd64.exe
│   └── gohuns-windows-arm64.exe
├── examples/
│   └── bad_Text                # Archivo de prueba para que veas cómo fallas en ortografía
├── go.mod                      # El manifiesto de dependencias
├── go.sum                      # Los hashes para que no te metan virus
├── LICENSE                     # Licencia del proyecto
├── Makefile                    # El automatizador definitivo de comandos
├── README.md                   # El archivo que estás leyendo ahora mismo 👋
└── src/
    └── main.go                 # TODO el código (sí, todo en un archivo porque somos MUY eficientes 😅)

```

---

## 🧰 Dependencias Utilizadas

Aquí no usamos librerías de 50GB. Solo lo justo y necesario:

```text
[github.com/client9/gospell](https://github.com/client9/gospell)   → El motor que parsea diccionarios Hunspell sin llorar
[github.com/fatih/color](https://github.com/fatih/color)       → Para pintar la terminal de rojo cuando escribes mal

```

---

## 📊 Comparación Inteligente

| Feature | GOHUNS | El corrector de Word | Pasar el texto por ChatGPT |
| --- | --- | --- | --- |
| Vive en la terminal | ✅ 🎉 | ❌ | ❌ |
| Funciona sin internet (Local-first) | ✅ | ✅ | ❌ (f por tu privacidad) |
| Multiplataforma nativo | ✅ (6 arquitecturas) | ❌ | ✅ (vía navegador pesado) |
| Consumo de RAM | Casi 0 | Una barbaridad | Tu tarjeta gráfica sufre |
| Estilo | Neutro | 🥱 Aburrido | 🤖 Artificial |

---

## 🐛 Troubleshooting (Resolución de dramas)

**"Error: No se encontró el diccionario 'es_CL'..."**

* **Solución:** No seas despistado, ejecútalo primero con el modo descarga: `gohuns -d -n es_CL <url_aff> <url_dic>` para que el sistema lo registre de forma global.

**"¡Los colores no se ven en mi Windows!"**

* **Solución:** Deja de usar el `cmd.exe` de Windows XP. Pásate a Windows Terminal o PowerShell para disfrutar de la experiencia a todo color.

---

## 🤝 Contribuciones

¿Encontraste un bug? ¿Quieres que limpie más líneas raras del `.aff`?

1. Haz un Fork de este repositorio.
2. Crea una rama fachera: `git checkout -b feature/mi-mejora-loca`
3. Haz tus cambios y reza para que compile.
4. Abre un Pull Request y lo revisamos con una taza de té.

---

Desarrollado con mucho código limpio y **mucho Té** por [Qmaker](https://github.com/Qmaker-programmer). 🚩
