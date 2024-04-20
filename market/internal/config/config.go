package config

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ilyakaznacheev/cleanenv"
)

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

var cfg config

type config struct {
	EnvType      string             `yaml:"env_type"`
	JWT          jwtConfig          `yaml:"jwt"`
	Storage      storageConfig      `yaml:"storage"`
	Server       serverConfig       `yaml:"server"`
	EventsClient eventsClientConfig `yaml:"events_client"`
}

type jwtConfig struct {
	PathPublicKey string `yaml:"public_key_path"` // Ignore this field
	PublicKey     *rsa.PublicKey
}

type serverConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type storageConfig struct {
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
}

type eventsClientConfig struct {
	URL string
}

// Cfg return copy of cfg (line 18)
func Config() config {
	return cfg
}

func LoadConfig() {
	envType := getEnvType()
	path := getConfigFilePath(envType)
	cleanenv.ReadConfig(path, &cfg)
	readKeys()
}

func getConfigFilePath(envType string) string {
	path := fmt.Sprintf("./config/%s.yaml", envType)
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		log.Fatalf("%s file not found", path)
	}
	return path
}

func getEnvType() string {
	envType := os.Getenv("ENV_TYPE")
	if envType == "" {
		log.Fatal("Empty ENV_TYPE variable")
	}
	if envType != EnvProd {
		log.Printf("!!! Using %s env type. Not for production !!!", envType)
		log.Printf("!!! Using %s env type. Not for production !!!", envType)
		log.Printf("!!! Using %s env type. Not for production !!!", envType)
	}
	return envType
}

func readKeys() {
	// Read public
	pem_key, err := os.ReadFile(cfg.JWT.PathPublicKey)
	if err != nil {
		panic(err)
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(pem_key)
	if err != nil {
		panic(err)
	}
	cfg.JWT.PublicKey = key
}
