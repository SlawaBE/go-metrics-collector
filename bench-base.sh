#!/bin/bash

if [[ "$(go env GOOS)" == "windows" ]]; then
    EXE=".exe"
else
    EXE=""
fi
HEY_BINARY="hey${EXE}"

if [ ! -e "./${HEY_BINARY}" ] ; then
    echo "Not found hey"
    exit 1
fi

./${HEY_BINARY} -n 1000 -T 'application/json' -d '{"id":"test136","type":"gauge","value":1.1}' -m POST 'http://localhost:8080/update'

./${HEY_BINARY} -n 1000 -T 'application/json' -d '[{"id":"test132","type":"gauge","value":1.1}]' -m POST 'http://localhost:8080/updates'

./${HEY_BINARY} -n 1000 -T 'application/json' -d '{"id":"test136","type":"gauge"}' -m POST 'http://localhost:8080/value'

./${HEY_BINARY} -n 1000 -m GET 'http://localhost:8080/'

curl -s -v http://localhost:8085/debug/pprof/heap > profiles/base.pprof 

go tool pprof -http=":9091" profiles/base.pprof
