## Pipeline Endpoint

The Pipeline Endpoint is a HTTP / gRPC to Kafka gateway proxy that is deployed by the Algo.Run Pipeline Operator.

An endpoint defines how the pipeline will accept requests and data from outside the Kubernetes cluster. An endpoint can have one or more paths, which will be used to form the url. Each path can be independently piped to any compatible input within the pipeline. An endpoint path can be used to segregate incoming data streams from IoT devices, enable external systems to integrate with the pipeline, create parallel processing pipes within a single pipeline and other creative uses.

The endpoint path can be configured with the following options:

- Name - A short, url friendly name for the endpoint path. Endpoint path name can only contain alphanumeric characters or single hyphens, and cannot begin or end with a hyphen with a maximum length of 50 characters.
- Description - A human friendly description of the endpoint path and how it is used.
- Endpoint Type - The endpoint type determines which protocol will be accepted for this path. Currently the possible protocols are HTTP(S) and gRPC, with Kafka and RTMP on the roadmap and coming soon.
- Message Data Type - The message data type enables you to choose how the incoming data is stored in the pipeline. The options are:

	- Embedded - If embedded, the incoming data will be added directly as the value for the message that will be delivered to Kafka.
	- FileReference - If FileReference, the incoming data will be saved to shared storage and a json message will be generated containing the location information for the file. The json file reference message will then be delivered to Kafka.
- Content Type - The accepted content type can be defined for this endpoint path. By defining the content type you gain additional features:

	- Ensure only compatible outputs and inputs can be piped to each other
	- Validation of the data being delivered to ensure it matches the content type

Let's take a look at how the endpoint is implemented.

![](https://content.algohub.com/assets/Endpoint-Deployment.png)

As you can see, a deployed endpoint will have the following resources created:

- An Ambassador mapping will be generated for the path, which routes ingress traffic from the endpoint path to the appropriate container.
- A container for the endpoint type is created to handle the appropriate protocol.

Data can then be sent to the endpoint path following using these Url conventions:

|Endpoint Type|Url|
| ------------ | ------------ |
|HTTP(S)|http(s)://{AlgoRun IP}:{HTTP Port}/{Deployment Owner}/{Deployment Name}/{Endpoint Path}|
|gRPC|http(s)://{AlgoRun IP}:{gRPC Port}/|
|Kafka|Kafka Broker: {AlgoRun IP}:9092<br>Topic: {Deployment Owner}.{Deployment Name}.{Endpoint Path}|
|RTMP|Coming Soon|

## Fork

This repository was forked from [Kafka Ambassador]https://github.com/AnchorFree/kafka-ambassador and modified for the Algo.Run Pipeline Deployment semantics.