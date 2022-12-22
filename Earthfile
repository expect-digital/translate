VERSION 0.6

proto:
  FROM bufbuild/buf
  COPY --dir api/translate .
  WORKDIR translate
  RUN buf mod update
  RUN buf build
  RUN buf generate
  SAVE ARTIFACT gen/proto/go/translate/v1 AS LOCAL pkg/server/translate/v1
