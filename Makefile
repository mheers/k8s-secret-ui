all_unit_tests:
	@go test ./... -v -count=1

docker: ##  Builds the application
	docker buildx build --platform linux/amd64 -t mheers/k8s-secret-ui --output type=docker .

docker-arm64: ##  Builds the application for arm64
	docker buildx build --platform linux/arm64 -t mheers/k8s-secret-ui --output type=docker .

docker-multi: ##  Builds the application for amd64 and arm64
	docker buildx build --platform linux/amd64,linux/arm64 -t mheers/k8s-secret-ui --push .
