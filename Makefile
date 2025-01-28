# This project applies to ppc64le only
ARCH ?= ppc64le

REGISTRY ?= quay.io/jcho0
REPOSITORY ?= power-dra-driver
TAG ?= first-test

CONTAINER_RUNTIME ?= $(shell command -v podman 2> /dev/null || echo docker)

########################################################################
# Go Targets

#.PHONY: build
#build: fmt vet
#	GOOS=linux GOARCH=$(ARCH) go build -o bin/power-dra-driver cmd/power-dra-driver/main.go

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: clean
clean:
	rm -f ./bin/power-dra-driver
	rm -rf vendor

########################################################################
# Container Targets

.PHONY: image
image: build
	$(CONTAINER_RUNTIME) buildx build \
		-t $(REGISTRY)/$(REPOSITORY):$(TAG) \
		--platform linux/$(ARCH) -f Dockerfile .

.PHONY: push
push:
	$(info push Container image...)
	$(CONTAINER_RUNTIME) push $(REGISTRY)/$(REPOSITORY):$(TAG)

########################################################################
# Deployment Targets

.PHONY: dep-plugin
dep-plugin:
	kustomize build manifests | oc apply -f -

.PHONY: dep-examples
dep-examples:
	kustomize build examples | oc apply -f -
