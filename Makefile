all_unit_tests:
	@go test ./... -v -count=1

docker: ##  Builds the application
	docker buildx build --platform linux/amd64 -t mheers/k8s-secret-ui --output type=docker .

docker-arm64: ##  Builds the application for arm64
	docker buildx build --platform linux/arm64 -t mheers/k8s-secret-ui --output type=docker .

docker-multi: ##  Builds the application for amd64 and arm64
	$(eval VERSION := $(shell cat .VERSION | grep VERSION | cut -d'=' -f2))
	docker buildx build --platform linux/amd64,linux/arm64 -t mheers/k8s-secret-ui:$(VERSION) --push .
	docker buildx build --platform linux/amd64,linux/arm64 -t mheers/k8s-secret-ui:latest --push .

helm-package-k8s-secret-ui:
	$(eval VERSION := $(shell cat .VERSION | grep VERSION | cut -d'=' -f2))
	yq eval '.image.tag = "$(VERSION)"' -i  ./deployment/k8s-secret-ui/values.yaml
	helm package --version $(VERSION)-helm deployment/k8s-secret-ui

helm-publish-k8s-secret-ui:
	$(eval VERSION := $(shell cat .VERSION | grep VERSION | cut -d'=' -f2))
	helm push k8s-secret-ui-$(VERSION)-helm.tgz oci://registry-1.docker.io/mheers
