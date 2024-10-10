package conf

type AppConfig struct {
	AppPort         int    `json:"APP_PORT"`
	AppMode         string `json:"APP_MODE"`
	DatabaseDSN     string `json:"DATABASE_DSN"`
	TracerName      string `json:"TRACER_NAME"`
	RedisClientName string `json:"REDIS_CLIENT_NAME"`
}
