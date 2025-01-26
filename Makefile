# Makefile для создания миграций

# Переменные которые будут использоваться в наших командах (Таргетах)
# Строка подключения к БД
DB_DSN := "postgres://postgres:yourpassword@localhost:5432/main?sslmode=disable"
# Команда для работы с миграциями
# Утилитиа путь к миграциям и строка подключения к БД
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Таргет для создания новой миграции
migrate-new:
    # Утилитиа migrate создаёт новую миграцию в папке migrations
    # -ext sql расширение переданного файла
    # -dir ./migrations - путь к папке с миграциями
    # ${NAME} - имя миграции
	migrate create -ext sql -dir ./migrations ${NAME}

# Применение миграций
migrate:
    # Утилитиа migrate применяет миграции
	$(MIGRATE) up

# Откат миграций
migrate-down:
	$(MIGRATE) down

# для удобства добавим команду run, которая будет запускать наше приложение
run:
    # Теперь при вызове make run мы запустим наш сервер
	go run cmd/app/main.go 

lint:
    # Команда для запуска линтера
    # --out-format=colored-line-number - формат вывода
	golangci-lint run --out-format=colored-line-number

gen:
    #Утилита oapi-codegen генерирует код на основе openapi.yaml
    # Тег пути -config openapi/.openapi
    # Тег включения -include-tags tasks 
    # -package tasks - пакет для генерации
    # openapi/openapi.yaml - файл с описанием API
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go
	
	oapi-codegen -config openapi/.openapi -include-tags users -package users openapi/openapi.yaml > ./internal/web/users/api.gen.go
