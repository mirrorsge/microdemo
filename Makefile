.PHONY: build_gw_docker
build_gw_docker:
	docker build -t microdemogw:latest -f ./gateway/Dockerfile .
.PHONY: push_gw_docker
push_gw_docker:
	docker tag microdemogw:latest jianjunli/microdemogw:latest
	docker push jianjunli/microdemogw:latest

.PHONY: build_api_docker
build_api_docker:
	docker build -t microdemoapi:latest -f ./api/
