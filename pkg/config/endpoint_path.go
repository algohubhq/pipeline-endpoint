package config

type EndpointPath struct {
	Name string `json:"name,omitempty"`

	IsDefault bool `json:"isDefault,omitempty"`

	Description string `json:"description,omitempty"`

	ConnectionString string `json:"connectionString,omitempty"`

	EndpointType string `json:"endpointType,omitempty"`

	MessageDataType string `json:"messageDataType,omitempty"`
}
