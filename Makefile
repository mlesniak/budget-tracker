all:  build docker

build: pack
	go get
	go build ./...

pack:
	go-bindata static/ init.sql

docker:
	cd deployment && ./build.sh && cd ..

# For documentation and testing purposes.
run:
	docker run --rm -p 8080:8080 -e BUDGET_PASSWORD=foo mlesniak/budget-tracker

clean:
	rm -f budget-tracker

