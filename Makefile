BINARY = daemon-pinguino

install-daemon:
	go install ./cmd/...

deps:
	go get .

clean:
	rm $(BINARY)

test:
	go test .
	golint

servedoc:
	godoc -http=:6060
