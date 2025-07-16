# Имя итогового бинарника
BINARY_NAME=proxy-api
BUILD_DIR=build

# Папка для Swagger-документации
SWAGGER_DIR=./docs

# Порт приложения
PORT=8080

# Цель по умолчанию
all: build

# Сборка бинарника
build:
	@echo "Building the binary..."
	go build -o $(BUILD_DIR)/$(BINARY_NAME) main.go

# Генерация Swagger документации
swagger:
	@echo "Generating Swagger docs..."
	swag init --output $(SWAGGER_DIR) --parseDependency --parseInternal

# Запуск тестов
test:
	@echo "Running tests..."
	go test -v ./...

# Запуск приложения
run:
	@echo "Starting the app on port $(PORT)..."
	go run main.go

# Очистка артефактов сборки
clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)
	rm -rf $(SWAGGER_DIR)

# Docker: сборка образа
docker-build:
	@echo "Building Docker image..."
	docker build -t $(BINARY_NAME) .

# Docker: сборка без кэша
docker-build-nocache:
	@echo "Building Docker image without cache..."
	docker build --no-cache -t $(BINARY_NAME) .

# Docker: запуск контейнера
docker-run:
	@echo "Running Docker container on port $(PORT)..."
	docker run -p $(PORT):$(PORT) $(BINARY_NAME)

# Docker Compose: сборка
dc-build:
	@echo "Docker Compose build..."
	docker-compose build

# Docker Compose: сборка без кэша
dc-build-nocache:
	@echo "Docker Compose build without cache..."
	docker-compose build --no-cache

# Docker Compose: запуск
dc-up:
	@echo "Docker Compose up..."
	docker-compose up

# Docker Compose: остановка
dc-down:
	@echo "Docker Compose down..."
	docker-compose down

.PHONY: all build swagger test run clean \
        docker-build docker-build-nocache docker-run \
        dc-up dc-down dc-build dc-build-nocache