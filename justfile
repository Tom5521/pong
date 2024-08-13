

build:
  just compile windows amd64
  just compile linux amd64

run:
  CC=gcc go run -v .

[private]
compile goos goarch:
  #!/bin/bash

  if [[ {{goos}} == "windows" ]]; then
    CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOOS={{goos}} GOARCH={{goarch}} go build -ldflags "-s -w" -v -o builds/pong-{{goos}}-{{goarch}}.exe
    exit
  fi

  CC=gcc GOOS={{goos}} GOARCH={{goarch}} go build -ldflags "-s -w" -v -o builds/pong-{{goos}}-{{goarch}}
