#!/bin/bash

# Local instance must be running to pull the swagger.json file

java -jar ./openapi-generator-cli-4.2.3.jar generate \
    -i http://localhost:5000/swagger/v1-beta1/swagger.json \
    -g go \
    -t openapi-template \
    -o algorun-go-client

mkdir -p ./openapi/
cp ./algorun-go-client/model_file_reference.go ./openapi/
cp ./algorun-go-client/model_api_bad_request_response.go ./openapi/
cp ./algorun-go-client/model_error_model.go ./openapi/
cp ./algorun-go-client/model_endpoint_path_model.go ./openapi/
cp ./algorun-go-client/model_content_type_model.go ./openapi/

cp ./algorun-go-client/model_endpoint_types.go ./openapi/
cp ./algorun-go-client/model_message_data_types.go ./openapi/

rm -rf ./algorun-go-client/

