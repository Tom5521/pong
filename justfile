run:
    CC=gcc go run -v -tags rgfw .

compile:
    just build windows amd64
    just build linux amd64

build os arch:
    #!/bin/bash

    bin_name=./builds/pong-{{os}}-{{arch}}
    compiler=gcc
    cgo_enabled=1
    ldflags="-s -w"

    if [[ "{{os}}" == "windows" ]]; then
        bin_name="$bin_name.exe"
        compiler=x86_64-w64-mingw32-gcc
        ldflags="$ldflags -H=windowsgui" 
    fi
    CGO_ENABLED=$cgo_enabled CC=$compiler GOOS={{os}} GOARCH={{arch}} \
    go build -ldflags "$ldflags" -v -tags rgfw -o $bin_name

    just compress $bin_name

compress bin:
    #!/bin/bash

    which upx > /dev/null 2>&1
    if [[ $? != 0 ]]; then
        exit 0
    fi

    upx --best {{bin}}