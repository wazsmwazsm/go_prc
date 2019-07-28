package config

type configs map[string]interface{}

var config configs = map[string]interface{}{
	"clusterID":        "eventbus",
	"natsStreamingUrl": "nats://192.168.5.103:6222",
}

func Get(key string) interface{} {
	return config[key]
}
