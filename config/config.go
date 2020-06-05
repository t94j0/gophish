package config

import (
	"encoding/json"
	"io/ioutil"

	log "github.com/gophish/gophish/logger"
)

// AdminServer represents the Admin server configuration details
type AdminServer struct {
	ListenURL string `json:"listen_url"`
	UseTLS    bool   `json:"use_tls"`
	CertPath  string `json:"cert_path"`
	KeyPath   string `json:"key_path"`
	CSRFKey   string `json:"csrf_key"`
}

// PhishServer represents the Phish server configuration details
type PhishServer struct {
	ListenURL  string `json:"listen_url"`
	UseTLS     bool   `json:"use_tls"`
	CertPath   string `json:"cert_path"`
	KeyPath    string `json:"key_path"`
	ServerName string `json:"server_name"`
}

// Config represents the configuration information.
type Config struct {
	AdminConf      AdminServer `json:"admin_server"`
	PhishConf      PhishServer `json:"phish_server"`
	DBName         string      `json:"db_name"`
	DBPath         string      `json:"db_path"`
	DBSSLCaPath    string      `json:"db_sslca_path"`
	MigrationsPath string      `json:"migrations_prefix"`
	TestFlag       bool        `json:"test_flag"`
	ContactAddress string      `json:"contact_address"`
	Logging        *log.Config `json:"logging"`
	// Controlls the X-Mailer address
	ServerName string `json:"server_name"`
	// The hostname is returned to the reciever. Change this option dynamically
	HostName string `json:"host_name"`
}

// Version contains the current gophish version
var Version = ""

// LoadConfig loads the configuration from the specified filepath
func LoadConfig(filepath string) (*Config, error) {
	// Get the config file
	configFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = json.Unmarshal(configFile, config)
	if err != nil {
		return nil, err
	}
	if config.PhishConf.ServerName == "" {
		config.PhishConf.ServerName = config.ServerName
	}
	// Choosing the migrations directory based on the database used.
	config.MigrationsPath = config.MigrationsPath + config.DBName
	// Explicitly set the TestFlag to false to prevent config.json overrides
	config.TestFlag = false
	return config, nil
}
