all: pack docker
	go build ./...

pack:
	go-bindata static/ init.sql

docker:
	docker build -t mlesniak/budget-tracker . 

# For documentation and testing purposes.
run:
	docker run --rm -p 8080:8080 -e BUDGET_PASSWORD=foo mlesniak/budget-tracker

clean:
	rm -f budget-tracker

