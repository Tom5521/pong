run:
    CC=gcc go run -v -tags rgfw .

compile:
    just build windows amd64
    just build linux amd64

build os arch:
    #!/bin/bash
    if [[ "{{ os }}" == "windows" ]]; then
        CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS={{ os }} GOARCH={{ arch }} \
        go build -ldflags "-s -w -H=windowsgui" -v -tags rgfw -o ./builds/pong-{{ os }}-{{ arch }}.exe .
        exit 0
    fi
    CGO_ENABLED=1 CC=gcc GOOS={{ os }} GOARCH={{ arch }} \
    go build -ldflags "-s -w" -v -tags rgfw -o ./builds/pong-{{ os }}-{{ arch }}
