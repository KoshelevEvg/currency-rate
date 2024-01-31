build:
	docker-compose build currency-rate
run:
	docker-compose up currency-rate
test:
	go test -v ./..