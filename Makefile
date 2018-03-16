all: pack
	go build ./...

pack:
	go-bindata static/

clean:
	rm -f budget-tracker
