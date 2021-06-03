package infrastructure

import "os"

// Env has environment stored
type Env struct {
	ServerPort        string
	Environment       string
	LogOutput         string
	DBUsername        string
	DBPassword        string
	DBHost            string
	DBPort            string
	DBName            string
	StorageBucketName string
	SentryDNS         string
	MailClientID      string
	MailClientSecret  string
	MailTokenType     string
	AdminEmail        string
	AdminPassword     string
	AdminDisplay      string
	AdminURI          string
	ClientURI         string
	StripeKey         string
	AccessToken       string
}

// NewEnv creates a new environment
func NewEnv() Env {
	env := Env{}
	env.LoadEnv()
	return env
}

// LoadEnv loads environment
func (env *Env) LoadEnv() {
	env.ServerPort = os.Getenv("ServerPort")
	env.Environment = os.Getenv("Environment")
	env.LogOutput = os.Getenv("LogOutput")
	env.DBUsername = os.Getenv("DBUsername")
	env.DBPassword = os.Getenv("DBPassword")
	env.DBHost = os.Getenv("DBHost")
	env.DBPort = os.Getenv("DBPort")
	env.DBName = os.Getenv("DBName")
	env.StorageBucketName = os.Getenv("StorageBucketName")
	env.MailClientID = os.Getenv("MailClientID")
	env.MailClientSecret = os.Getenv("MailClientSecret")
	env.MailTokenType = os.Getenv("MailTokenType")
	env.AdminEmail = os.Getenv("AdminEmail")
	env.AdminPassword = os.Getenv("AdminPassword")
	env.AdminDisplay = os.Getenv("AdminDisplay")
	env.AdminURI = os.Getenv("AdminURI")
	env.ClientURI = os.Getenv("ClientURI")
	env.StripeKey = os.Getenv("StripeKey")
	env.AccessToken = os.Getenv("AccessToken")
}
