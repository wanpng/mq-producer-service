package config

import "os"

// EnvType enum type for environment settings
type EnvType string

const (
	production  EnvType = "PRODUCTION"
	uat         EnvType = "UAT"
	development EnvType = "Environment"
)

// InitSettings initialize settings based on environment
func InitSettings() {
	switch env := GetEnv(); env {
	case production:
		loadProdSettings()
	case uat:
		loadUATSettings()
	default:
		loadDevSettings()
	}
}

// GetEnv get current environment settings
func GetEnv() EnvType {
	switch env := os.Getenv("ENV"); env {
	case "PRODUCTION":
		return production
	case "UAT":
		return uat
	default:
		return development
	}
}
