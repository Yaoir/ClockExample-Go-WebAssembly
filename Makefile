# Makefile for WebAssembly/Go version of digital clock

# Files needed for deployment. These must be on the web server.
DEPLOY=styles.css backgnd_tile.gif favicon.ico index.html main.wasm wasm_exec.js

# Files for development.
DEVFILES=*.go *.js deploy/upload* backgnd_tile.gif favicon.ico index.html main.wasm Makefile push README.md styles.css TODO

# build main.wasm (WebAssembly) from main.go
main:
	GOOS=js GOARCH=wasm go build -o main.wasm main.go

vet:
	GOOS=js GOARCH=wasm go vet main.go

# build the web server for development testing
webserver:
	go build webserver.go

# 'make test' runs the web server for local testing
test:
	./webserver

# copy files for deployment into deploy directory
dep:
	@cp -a $(DEPLOY) deploy

# make backup in .bak directory
backup back:
	@cp -a $(DEVFILES) .bak
