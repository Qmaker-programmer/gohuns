# Nombre del ejecutable base y versión
BINARY_NAME=gohuns
VERSION=0.0.2

# Rutas de origen y destino
SRC=src/main.go
BIN_DIR=bin

.PHONY: all run build clean all-debs

run:
	go run $(SRC)

build:
	go build -ldflags="-s -w" -o $(BIN_DIR)/$(BINARY_NAME) $(SRC)

# Compilación cruzada total (6 binarios limpios en /bin)
all: clean
	@mkdir -p $(BIN_DIR)
	@echo "🚀 Iniciando compilación multiplataforma (v$(VERSION))..."
	
	@echo "🐧 Compilando para Linux (amd64)..."
	@GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(BIN_DIR)/$(BINARY_NAME)-linux-amd64 $(SRC)
	
	@echo "🐧 Compilando para Linux (arm64)..."
	@GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o $(BIN_DIR)/$(BINARY_NAME)-linux-arm64 $(SRC)
	
	@echo "🪟 Compilando para Windows (amd64)..."
	@GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o $(BIN_DIR)/$(BINARY_NAME)-windows-amd64.exe $(SRC)
	
	@echo "🪟 Compilando para Windows (arm64)..."
	@GOOS=windows GOARCH=arm64 go build -ldflags="-s -w" -o $(BIN_DIR)/$(BINARY_NAME)-windows-arm64.exe $(SRC)
	
	@echo "🍏 Compilando para macOS (Intel - amd64)..."
	@GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o $(BIN_DIR)/$(BINARY_NAME)-darwin-amd64 $(SRC)
	
	@echo "🍏 Compilando para macOS (Apple Silicon - arm64)..."
	@GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o $(BIN_DIR)/$(BINARY_NAME)-darwin-arm64 $(SRC)
	
	@echo "✨ ¡Los 6 binarios estáticos listos en la carpeta /$(BIN_DIR)!"

# Automatiza los .deb guardándolos en /bin y sincronizando la versión automáticamente (en control y manpage)
# Automatiza los .deb guardándolos en /bin y sincronizando la versión automáticamente (en control y manpage)
all-debs: all
	@# Validación inteligente del entorno
	@if [ -z "$$(which dpkg-deb 2>/dev/null)" ]; then \
		echo "❌ Error: 'dpkg-deb' no está instalado. No se pueden crear paquetes .deb en este sistema."; \
		exit 1; \
	fi
	
	@echo "📦 Generando instaladores .deb para Linux (v$(VERSION))..."
	
	@# --- PAQUETE AMD64 ---
	@echo "  -> Estructurando debian_build_amd64..."
	@mkdir -p debian_build_amd64/DEBIAN
	@mkdir -p debian_build_amd64/usr/bin
	@mkdir -p debian_build_amd64/usr/share/man/man1
	@cp $(BIN_DIR)/$(BINARY_NAME)-linux-amd64 debian_build_amd64/usr/bin/$(BINARY_NAME)
	@# Modifica la versión dentro del manual sobre la marcha usando sed
	@if [ -f gohuns.1 ]; then \
		sed 's/gohuns [0-9]\+\.[0-9]\+\.[0-9]\+/gohuns $(VERSION)/g' gohuns.1 > debian_build_amd64/usr/share/man/man1/gohuns.1; \
	fi
	@echo "Package: $(BINARY_NAME)" > debian_build_amd64/DEBIAN/control
	@echo "Version: $(VERSION)" >> debian_build_amd64/DEBIAN/control
	@echo "Section: utils" >> debian_build_amd64/DEBIAN/control
	@echo "Priority: optional" >> debian_build_amd64/DEBIAN/control
	@echo "Architecture: amd64" >> debian_build_amd64/DEBIAN/control
	@echo "Maintainer: Qmaker" >> debian_build_amd64/DEBIAN/control
	@echo "Description: Corrector Ortografico Minimalista CLI" >> debian_build_amd64/DEBIAN/control
	@echo " Un corrector estatico, ultra rapido y minimalista para la CLI." >> debian_build_amd64/DEBIAN/control
	@dpkg-deb --build debian_build_amd64 $(BIN_DIR)/$(BINARY_NAME)_$(VERSION)_amd64.deb
	
	@# --- PAQUETE ARM64 ---
	@echo "  -> Estructurando debian_build_arm64..."
	@mkdir -p debian_build_arm64/DEBIAN
	@mkdir -p debian_build_arm64/usr/bin
	@mkdir -p debian_build_arm64/usr/share/man/man1
	@cp $(BIN_DIR)/$(BINARY_NAME)-linux-arm64 debian_build_arm64/usr/bin/$(BINARY_NAME)
	@# Modifica la versión dentro del manual sobre la marcha usando sed
	@if [ -f gohuns.1 ]; then \
		sed 's/gohuns [0-9]\+\.[0-9]\+\.[0-9]\+/gohuns $(VERSION)/g' gohuns.1 > debian_build_arm64/usr/share/man/man1/gohuns.1; \
	fi
	@echo "Package: $(BINARY_NAME)" > debian_build_arm64/DEBIAN/control
	@echo "Version: $(VERSION)" >> debian_build_arm64/DEBIAN/control
	@echo "Section: utils" >> debian_build_arm64/DEBIAN/control
	@echo "Priority: optional" >> debian_build_arm64/DEBIAN/control
	@echo "Architecture: arm64" >> debian_build_arm64/DEBIAN/control
	@echo "Maintainer: Qmaker" >> debian_build_arm64/DEBIAN/control
	@echo "Description: Corrector Ortografico Minimalista CLI" >> debian_build_arm64/DEBIAN/control
	@echo " Un corrector estatico, ultra rapido y minimalista para la CLI." >> debian_build_arm64/DEBIAN/control
	@dpkg-deb --build debian_build_arm64 $(BIN_DIR)/$(BINARY_NAME)_$(VERSION)_arm64.deb
	
	@# Limpieza estricta de residuos locales
	@rm -rf debian_build_amd64 debian_build_arm64 debian_build
	@echo "✨ ¡Paquetes debs guardados en /$(BIN_DIR): $(BINARY_NAME)_$(VERSION)_amd64.deb y $(BINARY_NAME)_$(VERSION)_arm64.deb!"

clean:
	rm -rf $(BIN_DIR)
	rm -rf debian_build_amd64 debian_build_arm64 debian_build
	rm -f *.deb
