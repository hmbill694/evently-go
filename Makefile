templ:
	@templ generate -watch -proxy=http:localhost:42069

tailwind:
	@npx tailwindcss -i ./assets/tailwind.css -o ./dist/styles.css --watch

run:
	@templ generate
	@go run main.go
