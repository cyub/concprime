APP_NAME=concprime
IMAGE_NAME=concprime

build: clean
	go build -v -o $(APP_NAME)

build-image:
	docker build --tag $(IMAGE_NAME) .

run:
	docker run --rm -it concprime -s=0

run-withoutput:
	docker run --rm -it concprime

benchmark:
	go test -benchmem -bench=. -benchtime=10s

clean:
	go clean