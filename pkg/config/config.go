package config

import (
	"os"
	"time"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var SocketClients = map[string]*websocket.Conn{}

type Config struct {
	API_URL            string        `json:"api_url" binding:"required"`
	UPLOAD_PATH        string        `json:"upload_path" binding:"required"`
	DB_HOST            string        `json:"db_host" binding:"required"`
	DB_PORT            string        `json:"db_port" binding:"required"`
	DB_USER            string        `json:"db_user" binding:"required"`
	DB_PASSWORD        string        `json:"db_password" binding:"required"`
	DB_NAME            string        `json:"db_name" binding:"required"`
	ACCESS_KEY         string        `json:"access_key" binding:"required"`
	ACCESS_TIME        time.Duration `json:"access_time" binding:"required"`
	REFRESH_KEY        string        `json:"refresh_key" binding:"required"`
	REFRESH_TIME       time.Duration `json:"refresh_time" binding:"required"`
	SMS_API            string        `json:"sms_api" binding:"required"`
	SMS_API_KEY        string        `json:"sms_api_key" binding:"required"`
	PAYMENT_API        string        `json:"payment_api" binding:"required"`
	PAYMENT_SESSION    int           `json:"payment_session" binding:"required"`
	SENAGAT_USERNAME   string        `json:"senagat_username" binding:"required"`
	SENAGAT_PASSWORD   string        `json:"senagat_password" binding:"required"`
	HALKBANK_USERNAME  string        `json:"halkbank_username" binding:"required"`
	HALKBANK_PASSWORD  string        `json:"halkbank_password" binding:"required"`
	RYSGAL_USERNAME    string        `json:"rysgal_username" binding:"required"`
	RYSGAL_PASSWORD    string        `json:"rysgal_password" binding:"required"`
	APP_VERSION        string        `json:"app_ve rsion" binding:"required"`
	API_PORT           string        `json:"api_port" binding:"required"`
	GIN_MODE           string        `json:"gin_mode" binding:"required"`
	LOGGER_FOLDER_PATH string        `json:"logger_folder_path" binding:"required"`
	LOGGER_FILENAME    string        `json:"logger_filename" binding:"required"`
	HLS_RUN_ON         string        `json:"hls_run_on" binding:"required"`
}

var ENV Config

func InitConfig() {
	godotenv.Load()

	ENV.API_URL = os.Getenv("API_URL")
	ENV.UPLOAD_PATH = os.Getenv("UPLOAD_PATH")

	ENV.DB_HOST = os.Getenv("DB_HOST")
	ENV.DB_PORT = os.Getenv("DB_PORT")
	ENV.DB_USER = os.Getenv("DB_USER")
	ENV.GIN_MODE = os.Getenv("GIN_MODE")
	ENV.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	ENV.DB_NAME = os.Getenv("DB_NAME")

	ENV.LOGGER_FOLDER_PATH = os.Getenv("LOGGER_FOLDER_PATH")
	ENV.LOGGER_FILENAME = os.Getenv("LOGGER_FILENAME")
	ENV.ACCESS_KEY = os.Getenv("ACCESS_KEY")
	AT, _ := time.ParseDuration(os.Getenv(("ACCESS_TIME")))
	ENV.ACCESS_TIME = AT

	ENV.REFRESH_KEY = os.Getenv("REFRESH_KEY")
	RT, _ := time.ParseDuration(os.Getenv(("REFRESH_TIME")))
	ENV.REFRESH_TIME = RT

	ENV.APP_VERSION = os.Getenv("APP_VERSION")
	ENV.HLS_RUN_ON = os.Getenv("HLS_RUN_ON")
}
