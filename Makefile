.PHONY: css clean run

css:
	npx @tailwindcss/cli -i ./static/css/input.css -o ./static/css/output.css --minify

clean:
	@echo INITIALISING DB CLEAN
	@test -f test.db && rm test.db || echo "NOTHING CLEANED"

run: clean css
	go run cmd/main.go