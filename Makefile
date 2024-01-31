build:
	docker-compose build currency-rate
run:
	docker-compose up currency-rate
test:
	go test -v -count=1 ./...
cover:
	go test -short -count=1 -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out