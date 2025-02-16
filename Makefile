.PHONY: gen-all
gen-all:
	go run github.com/99designs/gqlgen generate

.PHONY: start-webapp
start-webapp:
	go run server.go

.PHONY: watch-webapp
watch-webapp:
	go run github.com/cosmtrek/air

.PHONY: migrate-apply
migrate-apply:
	go run cmd/migrate/main.go