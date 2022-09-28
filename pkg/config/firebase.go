package config

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	firebaseCredFilePath = kingpin.Flag("firebaseCredFilePath", "path").Default("/run/secrets/firebase-key").String()
	firebaseProjectID = kingpin.Flag("firebaseProjectID", "worker-51312").Default("worker-51312").String()
)

func GetFirebaseCredFilePath() string {
	kingpin.Parse()

	deploy := os.Getenv("SWARM")
	if deploy == "true" {
		return *firebaseCredFilePath
	} else {
		return "conf/firebase/key.json"
	}
}

func GetFirebaseProjectID() string {
	kingpin.Parse()

	return *firebaseProjectID
}

func GetApiKey() string {
	key := os.Getenv("API_KEY")

	return key
}

func GetAppID() string {
	appID := os.Getenv("APP_ID")

	return appID
}
