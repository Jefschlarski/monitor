package configs

import "github.com/spf13/viper"

var cfg *config

type config struct {
	KAFKA KAFKAConfig
}

type KAFKAConfig struct {
	host  string
	port  string
	topic string
}

func (k *KAFKAConfig) Host() string {
	return k.host
}

func (k *KAFKAConfig) Port() string {
	return k.port
}

func (k *KAFKAConfig) Topic() string {
	return k.topic
}

func (k *KAFKAConfig) HostPort() string {
	return k.host + ":" + string(k.port)
}

func init() {
	viper.SetDefault("kafka.host", "localhost")
	viper.SetDefault("kafka.port", "9092")
	viper.SetDefault("kafka.drive", "orbis_monitor_v1")
}

// Load reads the configuration file named "config.toml" located in the root directory.
// It populates the global variable "cfg" with the values read from the file.
//
// It returns an error if there was a problem reading the configuration file.
func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}
	cfg = new(config)

	cfg.KAFKA = KAFKAConfig{
		host:  viper.GetString("kafka.host"),
		port:  viper.GetString("kafka.port"),
		topic: viper.GetString("kafka.topic"),
	}

	return nil
}

// GetKafkaConfig retrieves the kafka configuration.
//
// Returns:
//
//	KafkaConfig: struct of the kafka configuration
func GetKafkaConfig() KAFKAConfig {
	return cfg.KAFKA
}
