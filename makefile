run:
	go run cmd/exec-lex.go --input="docs/exemplo.emoji" --output="exemplo.go"

run-emoji:
	go run cmd/exec-lex.go --input="docs/exemplo.emoji" --output="exemplo.go" && go run exemplo.go