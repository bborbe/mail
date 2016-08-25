install:
	GOBIN=$(GOPATH)/bin GO15VENDOREXPERIMENT=1 go install bin/mail_smtp/mail_smtp.go
test:
	GO15VENDOREXPERIMENT=1 go test `glide novendor`
check:
	golint ./...
	errcheck -ignore '(Close|Write)' ./...
run:
	mail_smtp \
	-loglevel=trace
format:
	find . -name "*.go" -exec gofmt -w "{}" \;
	goimports -w=true .
prepare:
	go get -u golang.org/x/tools/cmd/goimports
	go get -u github.com/Masterminds/glide
	go get -u github.com/golang/lint/golint
	go get -u github.com/kisielk/errcheck
	glide install
update:
	glide up
clean:
	rm -rf vendor
