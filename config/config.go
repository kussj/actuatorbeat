// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

type Config struct {
	Actuatorbeat ActuatorbeatConfig
}

type ActuatorbeatConfig struct {
	Period string `yaml:"period"`
}
