#!/bin/bash

if [ ! -e "./hey.exe" ] ; then
    echo "Not found hey"
    exit 1
fi

./hey.exe -n 1000 -T 'application/json' -d '{"id":"test136","type":"gauge","value":1.1}' -m POST 'http://localhost:8080/update' 

./hey.exe -n 1000 -T 'application/json' -d '[{"id":"test132","type":"gauge","value":1.1}]' -m POST 'http://localhost:8080/updates'

./hey.exe -n 1000 -T 'application/json' -d '{"id":"test136","type":"gauge"}' -m POST 'http://localhost:8080/value'

./hey.exe -n 1000 -m GET 'http://localhost:8080/'

curl -s -v http://localhost:8085/debug/pprof/heap > profiles/base.pprof 

go tool pprof -http=":9091" profiles/base.pprof
