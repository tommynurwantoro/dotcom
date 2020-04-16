pretty:
	gofmt -s -w .

mod:
	go mod tidy

build:
	go build -o bin/dotcom cmd/web/main.go

run: pretty mod build
	./bin/dotcom

deploy:
	sudo cp dotcom.service /lib/systemd/system/dotcom.service

# Only for development
dev: build
	./bin/dotcom
