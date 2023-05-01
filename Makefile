MAIN_PKG    := "insomnia_backend"
GIT_HASH    := $(shell git rev-parse --short HEAD)
DATE        := $(shell date +%Y%m%d)
DOCKER_TAG  := $(TAG).$(GIT_HASH).$(DATE)

.PHONY:  build

build:
	rm -rf ./cmd/swagger/auto_generate
	swag init -o ./cmd/swagger/auto_generate
	rm -rf build/$(MAIN_PKG)
	go build -o build/$(MAIN_PKG)

doc:
	./build/$(MAIN_PKG) doc

test:
	export GIN_MODE=release && go test -v -count=1 ./pkg/rule_engine ./cmd/server
