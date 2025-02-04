ARG GOLANG_VERSION=1.23.3
FROM golang:${GOLANG_VERSION} as build
ARG CLIENT_GEN_VERSION=0.26.1
ARG LISTER_GEN_VERSION=0.26.1
ARG INFORMER_GEN_VERSION=0.26.1
ARG CONTROLLER_GEN_VERSION=0.17.1
ARG CI_LINT_VERSION=1.52.0
ARG MOQ_VERSION=0.3.4

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v${CI_LINT_VERSION}     && go install github.com/matryer/moq@v${MOQ_VERSION}     && go install sigs.k8s.io/controller-tools/cmd/controller-gen@v${CONTROLLER_GEN_VERSION}     && go install k8s.io/code-generator/cmd/client-gen@v${CLIENT_GEN_VERSION}     && go install k8s.io/code-generator/cmd/lister-gen@v${LISTER_GEN_VERSION}     && go install k8s.io/code-generator/cmd/informer-gen@v${INFORMER_GEN_VERSION}
RUN git config --file=/.gitconfig --add safe.directory /work

WORKDIR /build
COPY . .
RUN pwd
RUN ls -l


#RUN mkdir /artifacts
#RUN make PREFIX=/artifacts cmds

FROM registry.access.redhat.com/ubi9/ubi-minimal:9.4

LABEL io.k8s.display-name="IBM Power DRA Driver"
LABEL name="IBM Power DRA Driver"
LABEL vendor="IBM"
LABEL version="0.1.0"
LABEL release="N/A"
LABEL summary="Automate the management and monitoring of addition of specific Power devices to a Pod."
LABEL description="Automate the management and monitoring of addition of specific Power devices to a Pod."

COPY power-dra-kubeletplugin /usr/bin/power-dra-kubeletplugin
