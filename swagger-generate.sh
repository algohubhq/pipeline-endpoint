#!/bin/bash

# Local instance must be running to pull the swagger.json file
java -jar ./swagger-codegen-cli.jar generate -i http://localhost:5000/swagger/v1-beta1/swagger.json -Dio.swagger.parser.util.RemoteUrl.trustAll=true -l go -o algorun-go-client

mkdir -p ./swagger/
cp ./algorun-go-client/file_reference.go ./swagger/
cp ./algorun-go-client/api_bad_request_response.go ./swagger/
cp ./algorun-go-client/error_model.go ./swagger/
cp ./algorun-go-client/endpoint_path_model.go ./swagger/
cp ./algorun-go-client/content_type_model.go ./swagger/

rm -rf ./algorun-go-client/

