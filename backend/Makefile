swag-gen:
	swag init --parseDependency --parseInternal -g ./internal/app/app.go -o ./docs

start-app:
	make swag-gen && air -c .air.toml