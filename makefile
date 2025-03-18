run:
	go run src/main.go --input="docs/exemplo.be" --output="exemplo.go"

build:
	go run cmd/main.go --input="docs/exemplo.emoji" --output="exemplo.go"

run-emoji:
	go run exemplo.go

build-run:
	go run cmd/main.go --input="docs/exemplo.emoji" --output="exemplo.go" && go run exemplo.go