VERSION 0.6
ARG go_version=1.19.4
ARG golangci_lint_version=1.50.1

proto:
  FROM bufbuild/buf
  ENV BUF_CACHE_DIR /.cache/buf_cache
  COPY --dir api/translate .
  WORKDIR translate
  RUN --mount=type=cache,target=$BUF_CACHE_DIR buf mod update
  RUN --mount=type=cache,target=$BUF_CACHE_DIR buf build
  RUN --mount=type=cache,target=$BUF_CACHE_DIR buf generate
  SAVE ARTIFACT gen/proto/go/translate/v1 AS LOCAL pkg/server/translate/v1

lint-go:
  FROM golang:$go_version-alpine
  RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v$golangci_lint_version
  RUN apk add build-base
  WORKDIR /translate
  COPY --dir pkg .
  COPY .golangci.yml .
  COPY go.mod go.sum .
  RUN go mod download
  RUN golangci-lint run

lint-proto:
  FROM bufbuild/buf
  ENV BUF_CACHE_DIR /.cache/buf_cache
  COPY --dir api/translate .
  WORKDIR translate
  RUN --mount=type=cache,target=$BUF_CACHE_DIR buf mod update
  RUN --mount=type=cache,target=$BUF_CACHE_DIR buf lint

lint:
  BUILD +lint-proto
  BUILD +lint-go
