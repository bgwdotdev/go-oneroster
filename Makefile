build:
	go build \
		-o goors \
		cmd/goors/main.go
deps: 
	cd $(CURDIR)/cmd/goors; \
	go get -v
build-linux:
	CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64
	go build \
		-o goors-linux-amd64 \
		cmd/goors/main.go
build-windows:
	CGO_ENABLED=0 \
	GOOS=windows\
	GOARCH=amd64
	go build \
		-o goors-windows-amd64.exe \
		cmd/goors/main.go
