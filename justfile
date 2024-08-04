

build:
  just compile windows amd64
  just compile linux amd64

[private]
compile goos goarch:
  #!/bin/bash

  if [[ {{goos}} == "windows" ]]; then
    CGO_ENABLED=0 GOOS={{goos}} GOARCH={{goarch}} go build -ldflags "-s -w" -v -o builds/pong-{{goos}}-{{goarch}}.exe
    exit
  fi

  CC=gcc GOOS={{goos}} GOARCH={{goarch}} go build -ldflags "-s -w" -v -o builds/pong-{{goos}}-{{goarch}}
