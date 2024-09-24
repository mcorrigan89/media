package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port int
	Env  string
	DB   struct {
		DSN     string
		Logging bool
	}
	CientURL string
	Cors     struct {
		TrustedOrigins []string
	}
	ServiceApis struct {
		Idenitity struct {
			URL string
		}
	}
	Storage struct {
		Endpoint        string
		BucketName      string
		AccessKeyID     string
		SecretAccessKey string
	}
}

func LoadConfig(cfg *Config) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Load ENV
	env := os.Getenv("ENV")
	if env == "" {
		cfg.Env = "local"
	} else {
		cfg.Env = env
	}

	// Load PORT
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("PORT not available in .env")
	}

	cfg.Port = port

	// Load CLIENT_URL
	client_url := os.Getenv("CLIENT_URL")
	if client_url == "" {
		log.Fatalf("CLIENT_URL not available in .env")
	}

	cfg.CientURL = client_url

	// Load DATABASE_URL
	postgres_url := os.Getenv("POSTGRES_URL")
	if postgres_url == "" {
		log.Fatalf("POSTGRES_URL not available in .env")
	}

	cfg.DB.DSN = postgres_url

	cfg.Cors.TrustedOrigins = []string{"http://localhost:3000"}

	identity_url := os.Getenv("IDENTITY_URL")
	if identity_url == "" {
		log.Fatalf("IDENTITY_URL not available in .env")
	}

	cfg.ServiceApis.Idenitity.URL = identity_url

	// Load STORAGE
	storage_endpoint := os.Getenv("STORAGE_ENDPOINT")
	if storage_endpoint == "" {
		log.Fatalf("STORAGE_ENDPOINT not available in .env")
	}

	cfg.Storage.Endpoint = storage_endpoint

	storage_bucket_name := os.Getenv("STORAGE_BUCKET_NAME")
	if storage_bucket_name == "" {
		log.Fatalf("STORAGE_BUCKET_NAME not available in .env")
	}

	cfg.Storage.BucketName = storage_bucket_name

	storage_access_key_id := os.Getenv("STORAGE_ACCESS_KEY_ID")
	if storage_access_key_id == "" {
		log.Fatalf("STORAGE_ACCESS_KEY_ID not available in .env")
	}

	cfg.Storage.AccessKeyID = storage_access_key_id

	storage_secret_access_key := os.Getenv("STORAGE_SECRET_ACCESS_KEY")
	if storage_secret_access_key == "" {
		log.Fatalf("STORAGE_SECRET_ACCESS_KEY not available in .env")
	}

	cfg.Storage.SecretAccessKey = storage_secret_access_key
}
