#!/bin/bash



function generate {
    echo "Generating"
    .build/swagger_linux_amd64 generate server --name=restapi --spec=.assets/schema.yaml --target=./
}



function build {
    echo "Building"
    go build -v -o .bin/linux/restapi cmd/restapi-server/main.go
}



function run {
    echo "Run"
    .bin/linux/restapi --host "" --port 8000
}



# if we're doing a dev build
if [[ "$*" == *"--gen"* ]]; then
	generate
elif [[ "$*" == *"--build"* ]]; then
    build
elif [[ "$*" == *"--run"* ]]; then
    run
else 
    generate
    build
    run
fi

github.com/david-macpherson/go-swagger-custom