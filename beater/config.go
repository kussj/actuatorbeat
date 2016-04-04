package beater

// Defaults for config variables which are not set
const (
	DefaultUrl string = "http://localhost:8080/metrics"
)

type ActuatorbeatConfig struct {
	Period 			*int64
	URLs			[]string
}

type ConfigSettings struct {
	Input ActuatorbeatConfig
}
