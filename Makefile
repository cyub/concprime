APP_NAME=concprime
IMAGE_NAME=concprime

build: clean
	go build -v -o $(APP_NAME)

build-image:
	docker build --tag $(IMAGE_NAME) .

run: clean create-trace-file
	docker run --rm -it -v $(PWD)/trace.out:/app/trace.out concprime -s=0

run-withoutput: clean create-trace-file
	docker run --rm -it -v $(PWD)/trace.out:/app/trace.out concprime

benchmark:
	go test -benchmem -bench=. -benchtime=10s

clean:
	go clean
	rm -rf trace.out

create-trace-file:
	touch trace.out