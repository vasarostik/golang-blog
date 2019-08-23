
package config

import (
"fmt"
"io/ioutil"

yaml "gopkg.in/yaml.v2"
)

// Load_Manager returns ConfigManager struct
func Load_Manager(path string) (*ConfigManager, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}
	var cfg = new(ConfigManager)
	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}
	return cfg, nil
}

// Configuration holds data necessery for configuring application
type Configuration struct {
	API_ms  *API_ms `yaml:"api_ms,omitempty"`
	GRPC_ms *GRPC_ms `yaml:"grpc_ms,omitempty"`
	NATS_ms *NATS_ms `yaml:"nats_ms,omitempty"`
}

type API_ms struct {
	Server *Server      `yaml:"server,omitempty"`
	DB     *Database    `yaml:"database,omitempty"`
	Redis  *Redis    	`yaml:"redisdb,omitempty"`
	GRPC   *GRPC		`yaml:"jrpc,omitempty"`
	NATS_Server   	*NATS_Server		`yaml:"nats_server,omitempty"`
	JWT    *JWT         `yaml:"jwt,omitempty"`
	App    *Application `yaml:"application,omitempty"`
}

type GRPC_ms struct {
	Redis  *Redis    	`yaml:"redisdb,omitempty"`
	GRPC   *GRPC		`yaml:"jrpc,omitempty"`
}

type NATS_ms struct {
	NATS_Subscriber	*NATS_Subscriber `yaml:"nats_subscriber,omitempty"`
	NATS_Server   	*NATS_Server		`yaml:"nats_server,omitempty"`
}
// Database holds data necessery for database configuration
type Database struct {
	PSN        string `yaml:"psn,omitempty"`
	LogQueries bool   `yaml:"log_queries,omitempty"`
	Timeout    int    `yaml:"timeout_seconds,omitempty"`
}

// Server holds data necessery for server configuration
type Server struct {
	Port         string `yaml:"port,omitempty"`
	Debug        bool   `yaml:"debug,omitempty"`
	ReadTimeout  int    `yaml:"read_timeout_seconds,omitempty"`
	WriteTimeout int    `yaml:"write_timeout_seconds,omitempty"`
}

// JWT holds data necessery for JWT configuration
type JWT struct {
	Secret           string `yaml:"secret,omitempty"`
	Duration         int    `yaml:"duration_minutes,omitempty"`
	RefreshDuration  int    `yaml:"refresh_duration_minutes,omitempty"`
	MaxRefresh       int    `yaml:"max_refresh_minutes,omitempty"`
	SigningAlgorithm string `yaml:"signing_algorithm,omitempty"`
}

// Application holds application configuration details
type Application struct {
	MinPasswordStr int    `yaml:"min_password_strength,omitempty"`
	SwaggerUIPath  string `yaml:"swagger_path,omitempty"`
}

type Redis struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type GRPC struct {
	Addr string `json:"addr"`
}

type NATS_Server struct {
	Addr string `json:"addr"`
}

type NATS_Subscriber struct {
	Addr string `json:"addr"`
	Subject string `json:"subject"`
	LogFile string `json:"logfile"`
	PSN        string `yaml:"psn,omitempty"`
}

type ConfigManager struct {
	Addr string `yaml:"configManagerAddr"`
}