#!/bin/bash

rm -rf ./bin
PORT=8080

go build -o ./bin/qube

docker kill qube

docker build -t qube_img .

docker run -d -it --rm --name qube \
    -p $PORT:$PORT \
    -e PORT=8080 \
    -e ENV=production \
    qube_img

echo "container 'qube' is now up and running..."