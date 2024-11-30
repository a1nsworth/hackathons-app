SQLC_CONFIG_PATH=configs/sqlc.yaml

generate:
	@echo "Генерация кода с помощью sqlc..."
	sqlc generate -f $(SQLC_CONFIG_PATH)

clean:
	@echo "Очистка сгенерированных файлов..."
	rm -rf internal/db

