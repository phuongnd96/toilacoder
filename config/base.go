package config

import (
	"log"
	"os"
	"toilacoder/helper"

	gocache "github.com/patrickmn/go-cache"
)

type Config struct {
	SecureCookie bool `json:"secure_cookie"`

	Host string `json:"host"`

	Cache *gocache.Cache

	ImageFolder string `json:"image_folder"`

	CORS struct {
		AllowedOrigins []string `json:"allowed_origins"`
		MaxAgeInHours  int32    `json:"max_age_in_hours"`
	} `json:"cors"`

	DomainSendMail        string `json:"domain_send_mail"`
	SMTPMailServerAddress string `json:"smtp_email_server_address"`
	SMTPMailServerPort    int    `json:"smtp_email_server_port"`

	UserDefault struct {
		RootID        string `json:"root_id"`
		RootPassword  string `json:"root_password"`
		AdminID       string `json:"admin_id"`
		AdminPassword string `json:"admin_password"`
	} `json:"user_default"`

	Database struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
		Address  string `json:"address"`
	} `json:"database"`
}

func SetupConfig() Config {
	var conf Config

	// Đọc file config.local.json
	configFile, err := os.Open("config.local.json")
	if err != nil {
		// Nếu không có file config.local.json thì đọc file config.default.json
		configFile, err = os.Open("config.default.json")
		if err != nil {
			panic(err)
		}
		defer configFile.Close()
	}
	defer configFile.Close()

	// Parse dữ liệu JSON và bind vào Controller
	err = helper.DecodeDataFromJsonFile(configFile, &conf)
	if err != nil {
		log.Println("Không đọc được file config.")
		panic(err)
	}

	return conf
}
