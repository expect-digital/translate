VERSION 0.6

build:
    FROM bufbuild/buf
    COPY api/translate api/translate
    RUN rm -rf api/translate/gen
    WORKDIR api/translate
    RUN buf mod update
    RUN buf lint
    RUN buf build
    RUN buf generate
    SAVE ARTIFACT gen AS LOCAL api/translate/gen