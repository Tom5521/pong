
skip-compress := env_var_or_default("SKIP_COMPRESS","0")

run:
    CC=gcc go run -v -tags rgfw .

release:
    just compile
    gh release create $(git describe --tags --abbrev=0) --generate-notes ./builds/*

clean:
    rm -rf builds ./*.exe ./pong

compile:
    just build windows amd64
    just build linux amd64

build os arch:
    #!/bin/bash

    bin_name=./builds/pong-{{os}}-{{arch}}
    compiler=gcc
    cgo_enabled=1
    ldflags="-s -w"
    tags="-tags 'rgfw'"

    if [[ "{{os}}" == "windows" ]]; then
        bin_name="$bin_name.exe"
        compiler=x86_64-w64-mingw32-gcc
        ldflags="$ldflags -H=windowsgui"
        tags=""
    fi
    CGO_ENABLED=$cgo_enabled CC=$compiler GOOS={{os}} GOARCH={{arch}} \
    go build -ldflags "$ldflags" -v $tags -o $bin_name

    if [[ {{skip-compress}} == 1 ]]; then
        exit 0
    fi

    just compress $bin_name 

compress bin:
    #!/bin/bash

    which upx > /dev/null 2>&1
    if [[ $? != 0 ]]; then
        exit 0
    fi

    upx -9 {{bin}}
