package config_dto

type Config struct {
	Debug bool   `yaml:"debug"`
	Port  string `yaml:"port"`
	Log   struct {
		AppLogFile string `yaml:"appLogFile"`
	} `yaml:"log"`
	AppConfig AppConfig `yaml:"appConfig"`
	MongoDB   struct {
		Database string `yaml:"database"`
		Master   struct {
			URI      string `yaml:"uri"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		} `yaml:"master"`
		Slave struct {
			URI      string `yaml:"uri"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		} `yaml:"slave"`
	} `yaml:"mongoDB"`
	AuthKeys struct {
		JwtSecret struct {
			EncryptionKey       string `yaml:"encryptionKey"`
			TokenExpirationTime int    `yaml:"tokenExpirationTime"`
		} `yaml:"jwtSecret"`
		Snowflake struct {
			TokenExpirationTime int `yaml:"tokenExpirationTime"`
		} `yaml:"snowflake"`
	} `yaml:"authKeys"`
	Crypto struct {
		Aes256 struct {
			EncryptionKey string `yaml:"encryptionKey"`
		} `yaml:"aes-256"`
		Chacha20 struct {
			EncryptionKey string `yaml:"encryptionKey"`
		} `yaml:"chacha20"`
	} `yaml:"crypto"`
}
