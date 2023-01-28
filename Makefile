run:
	go run main.go

docker-build:
	docker-compose build

docker-up:
	@printf 'Enter the number of network bits: (a value between 1 and 32): '
	@read bits && BITS=$$bits docker-compose up