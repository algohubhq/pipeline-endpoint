#!/bin/bash

# Local instance must be running to pull the swagger.json file
java -jar ./swagger-codegen-cli.jar generate -i https://localhost:5443/swagger/v1/swagger.json -Dio.swagger.parser.util.RemoteUrl.trustAll=true -l go -o algorun-go-client

mkdir -p ./swagger/
cp ./algorun-go-client/file_reference.go ./swagger/
cp ./algorun-go-client/api_bad_request_response.go ./swagger/
cp ./algorun-go-client/error_model.go ./swagger/

rm -rf ./algorun-go-client/

