package config

type Config struct {
	Array   *array
	HashMap *hashmap

	DefaultFeature   string
	multipleFeatures bool
}

func New(multipleFeatures bool) *Config {
	return &Config{
		Array: &array{
			storage: make(map[string][]interface{}),
		},
		HashMap: &hashmap{
			storage: make(map[string]map[string]interface{}),
		},

		multipleFeatures: multipleFeatures,
	}
}
