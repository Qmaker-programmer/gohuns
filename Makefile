# Nombre del ejecutable base
BINARY_NAME=gohuns

# Rutas de origen y destino
SRC=src/main.go
BIN_DIR=bin

.PHONY: all run build build-all clean

run:
	go run $(SRC)

build:
	go build -ldflags="-s -w" -o $(BIN_DIR)/$(BINARY_NAME) $(SRC)

# Compilación cruzada total (6 binarios: Windows, Linux y macOS en amd64 y arm64)
all: clean
	@mkdir -p $(BIN_DIR)
	@echo "🚀 Iniciando compilación multiplataforma..."
	
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

clean:
	rm -rf $(BIN_DIR)
