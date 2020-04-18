package config

import (
	"os"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

// Config is a handler for possible configuration parameters, such as RemoteConfig and ConfigWatch.
type T struct {
	Filename  string
	EnvPrefix string
}

func (c *T) ReadConfig(defaults map[string]interface{}) (*viper.Viper, error) {

	v := viper.New()
	v.SetEnvPrefix(c.EnvPrefix)

	for key, value := range defaults {
		v.SetDefault(key, value)
	}

	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if c.Filename != "" {
		v.SetConfigFile(c.Filename)
		err := v.ReadInConfig()
		return v, err
	}

	// Get full config from the envar
	endpointConfig := os.Getenv("ENDPOINT_CONFIG")
	if endpointConfig != "" {
		v.SetConfigType("json") // or viper.SetConfigType("yaml")
		err := v.ReadConfig(strings.NewReader(endpointConfig))
		return v, err
	}

	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error
			return v, nil
		}
	}

	return v, err

}

func Flatten(value interface{}) map[string]interface{} {
	return flattenPrefixed(value, "")
}

func flattenPrefixed(value interface{}, prefix string) map[string]interface{} {
	m := make(map[string]interface{})
	flattenPrefixedToResult(value, prefix, m)
	return m
}

func flattenPrefixedToResult(value interface{}, prefix string, m map[string]interface{}) {
	base := ""
	if prefix != "" {
		base = prefix + "."
	}

	original := reflect.ValueOf(value)
	kind := original.Kind()
	if kind == reflect.Ptr || kind == reflect.Interface {
		original = reflect.Indirect(original)
		kind = original.Kind()
	}
	t := original.Type()

	switch kind {
	case reflect.Map:
		if t.Key().Kind() != reflect.String {
			break
		}
		for _, childKey := range original.MapKeys() {
			childValue := original.MapIndex(childKey)
			flattenPrefixedToResult(childValue.Interface(), base+childKey.String(), m)
		}
	case reflect.Struct:
		for i := 0; i < original.NumField(); i++ {
			childValue := original.Field(i)
			childKey := t.Field(i).Name
			flattenPrefixedToResult(childValue.Interface(), base+childKey, m)
		}
	default:
		if prefix != "" {
			m[prefix] = value
		}
	}
}
