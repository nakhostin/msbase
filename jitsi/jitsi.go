package jitsi

import (
	"fmt"
	"io/ioutil"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v2"
)

type JitsiSDK struct {
	config Config     // config : jitsi required configs
	echo   *echo.Echo // echo : optional to update echo
}

// Init : load sdk with config struct and optional echo
func Init(cfg Config, ec *echo.Echo) *JitsiSDK {
	sdk := JitsiSDK{
		config: cfg,
		echo:   ec,
	}

	if ec != nil {
		gp := ec.Group(sdk.config.Echo.Prefix)
		sdk.echoRoutes(gp)
	}

	return &sdk
}

// InitWithConfigFile : load sdk with config file and optional echo
func InitWithConfigFile(path string, ec *echo.Echo) (*JitsiSDK, error) {
	cfg, err := loadYamlConfig(path)
	if err != nil {
		return nil, err
	}
	return Init(cfg, ec), nil
}

// CreateMeet: Create Meeting Url
// room = "*" => access to all rooms
func (v JitsiSDK) CreateMeetWithClaims(claims map[string]interface{}) (string, error) {
	// generate host url
	jitsiURL := v.generateHostURL()
	token, err := v.generateJwtTokenWithMap(claims)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(jitsiURL, "/", claims["room"], "?jwt=", token), nil
}

type Payload struct {
	Room      string  `json:"room"`      // enter your room name
	Context   Context `json:"context"`   // enter user information
	Moderator bool    `json:"moderator"` // set admin and guest mode
	Aud       string  `json:"aud"`       // this value set default, recommend don't change value
	Iss       string  `json:"iss"`       // this value set default, recommend don't change value
	Sub       string  `json:"sub"`       // this value set default, recommend don't change value
	jwt.StandardClaims
}

func (v JitsiSDK) NewJitsiPayload() *Payload {
	return &Payload{Aud: v.config.AppID, Iss: v.config.AppID, Sub: v.config.Host}
}

type Context struct {
	User User `json:"user"`
}
type User struct {
	Name   string `json:"name"`
	Email  string `json:"emial"`
	Avatar string `json:"avatar"`
}

func (v JitsiSDK) CreateMeet(payload Payload) (string, error) {
	jitsiURL := v.generateHostURL()
	token, err := v.generateJwtToken(payload)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(jitsiURL, "/", payload.Room, "?jwt=", token), nil
}

// generateJwtToken : generate jitsi required token
func (v JitsiSDK) generateJwtToken(claims Payload) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(v.config.JwtSecret))
}

// generateJwtTokenWithMap : generate jitsi required token with map claims
func (v JitsiSDK) generateJwtTokenWithMap(claims jwt.MapClaims) (string, error) {
	claims["aud"], claims["iss"], claims["sub"] = v.config.AppID, v.config.AppID, v.config.Host
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(v.config.JwtSecret))
}

// generateHostURL : generate host url using jitsi config
func (v JitsiSDK) generateHostURL() string {
	protocol := "http://"
	if v.config.TLS {
		protocol = "https://"
	}
	return protocol + v.config.Host
}

func loadYamlConfig(path string) (Config, error) {
	config := new(Config)
	file, err := ioutil.ReadFile(path)
	if err == nil {
		if err := yaml.Unmarshal(file, &config); err != nil {
			panic(err)
		}
	}
	return *config, nil
}
