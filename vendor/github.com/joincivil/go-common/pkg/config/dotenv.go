package config

import (
	"os"
	"strings"

	log "github.com/golang/glog"

	"github.com/joho/godotenv"
)

const (
	errNoSuchFileDirectory = "no such file or directory"
)

func isNoSuchFileErr(err error) bool {
	return strings.Contains(err.Error(), errNoSuchFileDirectory)
}

// PopulateFromDotEnv attempts to retrieve env vars from .env files to populate
// into the environment.
// envEnvVar passed in indicates the env var that specifies the environment to
// use "test", "development", "production" for the .env file
func PopulateFromDotEnv(envEnvVar string) error {
	env := os.Getenv(envEnvVar)
	if "" == env {
		env = "development"
	}

	err := godotenv.Load(".env." + env + ".local")
	if err != nil && !isNoSuchFileErr(err) {
		log.Errorf("Did not load .env.%v.local: err: %v", env, err)
		return err
	} else if err == nil {
		log.Infof("Loaded .env.%v.local", env)
	}

	if "test" != env {
		err := godotenv.Load(".env.local")
		if err != nil && !isNoSuchFileErr(err) {
			log.Errorf("Did not load .env.local: err: %v", err)
			return err
		} else if err == nil {
			log.Info("Loaded .env.local")
		}
	}

	err = godotenv.Load(".env." + env)
	if err != nil && !isNoSuchFileErr(err) {
		log.Errorf("Did not load .env.%v: err: %v", env, err)
		return err
	} else if err == nil {
		log.Infof("Loaded .env.%v", env)
	}

	err = godotenv.Load()
	if err != nil && !isNoSuchFileErr(err) {
		log.Errorf("Did not load .env: err: %v", err)
		return err
	} else if err == nil {
		log.Infof("Loaded .env")
	}

	return nil
}
