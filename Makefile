DOCKER_REGISTRY_SERVER := registry.cn-shenzhen.aliyuncs.com
VERSION := v0.0.1

.PHONY: build
build: export GO111MODULE=on
build: export GOSUMDB=off
build:
	@echo "buiding feishuBot"
	go mod tidy
	go build

.PHONY: image
image:
	@echo "building image"
	docker build -t feishubot:$(VERSION) .
	docker tag feishubot:$(VERSION) $(DOCKER_REGISTRY_SERVER)/virgocoder/feishubot:$(VERSION)
	docker push $(DOCKER_REGISTRY_SERVER)/virgocoder/feishubot:$(VERSION)
