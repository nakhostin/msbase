package jitsi

type (
	Config struct {
		Host      string     `yaml:"host"`
		TLS       bool       `yaml:"tls"`
		AppID     string     `yaml:"app_id"`
		JwtSecret string     `yaml:"jwt_secret"`
		Echo      EchoConfig `yaml:"echo"`
	}
	EchoConfig struct {
		JWTSecret string `yaml:"jwt_secret"`
		Prefix    string `yaml:"prefix"`
	}
)

var DefualtConfig = Config{
	Host:      "127.0.0.1",
	JwtSecret: "9273EC80FC11B9B80A32ED3673D52316",
	TLS:       true,
	AppID:     "nima",
	Echo: EchoConfig{
		JWTSecret: "5059BC4940913D21C5A60B39F11079CA",
		Prefix:    "/jitsi",
	},
}
