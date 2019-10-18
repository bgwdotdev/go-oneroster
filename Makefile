build:
	go build \
		-o goors \
		cmd/goors/main.go
static:
	CGO_ENABLED=1 \
	GOOS=linux \
	go build \
		-a \
		-ldflags '-linkmode external -extldflags "-static"' \
		-o goors \
	cmd/goors/main.go
deps: 
	cd $(CURDIR)/cmd/goors; \
	go get -v
