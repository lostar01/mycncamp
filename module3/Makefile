root:
	export ROOT=lostar.com/mycncamp/module3
build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 *.go

release: build
	echo "build httpserver container"
	docker build -t lostar01/mycncamp-httpserver:${tag} .
push: release
	echo "pushing mycncamp-httpserver"
	docker push lostar01/mycncamp-httpserver:${tag}
help:
	@echo "make [build|release|push]"
	@echo "make build		build httpserver code"
	@echo "make release		build code and create docker image"
	@echo "make push		build code and crate image and push , need docker already login docker hub"
