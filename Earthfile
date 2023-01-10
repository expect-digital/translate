VERSION 0.6
ARG go_version=1.19.4
ARG golangci_lint_version=1.50.1
ARG openapitools_version=6.2.1
FROM golang:$go_version-alpine

init:
  RUN printf "#!/bin/sh\nearthly +lint" > pre-push
  RUN chmod ug+x pre-push
  SAVE ARTIFACT pre-push AS LOCAL .git/hooks/pre-push

proto:
  FROM bufbuild/buf
  ENV BUF_CACHE_DIR=/.cache/buf_cache
  COPY --dir api/translate .
  WORKDIR translate
  RUN --mount=type=cache,target=$BUF_CACHE_DIR buf mod update
  RUN --mount=type=cache,target=$BUF_CACHE_DIR buf build
  RUN --mount=type=cache,target=$BUF_CACHE_DIR buf generate --template buf.gen.go.yaml
  SAVE ARTIFACT gen/proto/go/translate/v1 translate/v1 AS LOCAL pkg/server/translate/v1

deps:
  ENV CGO_ENABLED=0
  RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v$golangci_lint_version
  WORKDIR /translate
  COPY go.mod go.sum .
  RUN go mod download
  SAVE ARTIFACT go.mod AS LOCAL go.mod
  SAVE ARTIFACT go.sum AS LOCAL go.sum

lint-go:
  FROM +deps
  COPY --dir pkg .
  COPY --dir +proto/translate/v1 pkg/server/translate/v1
  COPY .golangci.yml .
  RUN golangci-lint run

lint-proto:
  FROM bufbuild/buf
  ENV BUF_CACHE_DIR=/.cache/buf_cache
  COPY --dir api/translate .
  WORKDIR translate
  RUN --mount=type=cache,target=$BUF_CACHE_DIR buf mod update
  RUN --mount=type=cache,target=$BUF_CACHE_DIR buf lint

lint:
  BUILD +lint-proto
  BUILD +lint-go

openapi-swagger:
  FROM bufbuild/buf
  ENV BUF_CACHE_DIR=/.cache/buf_cache
  COPY --dir api/translate .
  WORKDIR translate
  RUN --mount=type=cache,target=$BUF_CACHE_DIR buf mod update
  RUN --mount=type=cache,target=$BUF_CACHE_DIR buf build
  RUN --mount=type=cache,target=$BUF_CACHE_DIR buf generate --template buf.gen.openapi.yaml
  SAVE ARTIFACT gen/openapiv2/translate/v1 translate/v1/translate.swagger.json

openapi-angular-client:
  FROM openapitools/openapi-generator-cli:v$openapitools_version
  WORKDIR /translate
  COPY +openapi-swagger/translate/v1/translate.swagger.json .
  RUN docker-entrypoint.sh generate \
    --input-spec translate.swagger.json \
    --generator-name typescript-angular \
    --output client
  SAVE ARTIFACT client AS LOCAL website/src/app/services/api
