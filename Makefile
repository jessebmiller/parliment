# Dependency managemnt lifecycle

NAME = $(shell basename $$(git rev-parse --show-toplevel))
LOCAL_BUILD_TAG = $(NAME):local

.PHONY: fmt
fmt: build
	docker run -v $(shell pwd):$(shell docker run $(LOCAL_BUILD_TAG) pwd) $(LOCAL_BUILD_TAG) go fmt ./...


.PHONY: ensure
ensure: build
	touch Gopkg.toml
	touch Gopkg.lock
	docker run -v $(shell pwd):$(shell docker run $(LOCAL_BUILD_TAG) pwd) $(LOCAL_BUILD_TAG) dep ensure


.PHONY: test
test: build
	docker run -v $(shell pwd):$(shell docker run $(LOCAL_BUILD_TAG) pwd) $(LOCAL_BUILD_TAG) go test ./...


.PHONY: build
build:
	docker build -t $(LOCAL_BUILD_TAG) .

.PHONY: start
start: build
	-docker stop $(NAME)
	docker run --name $(NAME) -p 8080:8080 $(LOCAL_BUILD_TAG)
