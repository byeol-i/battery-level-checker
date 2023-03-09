package config

import (
	"flag"
	"os"
)


var (
	local = flag.Bool("test",  false, "using local key")
	credFilePath = flag.String("firebaseCredFilePath", "/run/secrets/firebase-key", "cred path")
	firebaseProjectID = flag.String("firebaseProjectID", "worker-51312", "firebaseProjectID")
	
)

func GetFirebaseCredFilePath() string {
	flag.Parse()

	if *local {
		return "conf/firebase/key.json"
	} else {
		return *credFilePath
	}
}

func GetFirebaseProjectID() string {
	flag.Parse()
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
