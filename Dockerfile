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
