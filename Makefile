clean:
	@echo INITIALISING DB CLEAN
	@test -f test.db && rm test.db || echo "NOTHING CLEANED"
run: clean
	go run cmd/main.go