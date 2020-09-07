install: config
	go install

run: config
	go run

config:
	./checkConfig

uninstall:
	go clean -i
