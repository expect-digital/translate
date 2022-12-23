VERSION 0.6

proto:
  FROM bufbuild/buf
  ENV BUF_CACHE_DIR /.cache/buf_cache
  COPY --dir api/translate .
  WORKDIR translate
  RUN --mount=type=cache,target=$BUF_CACHE_DIR buf mod update
  RUN --mount=type=cache,target=$BUF_CACHE_DIR buf build
  RUN --mount=type=cache,target=$BUF_CACHE_DIR buf generate
  SAVE ARTIFACT gen/proto/go/translate/v1 AS LOCAL pkg/server/translate/v1
