package config

import "gopkg.in/alecthomas/kingpin.v2"

var (
	firebaseCredFilePath = kingpin.Flag("firebaseCredFilePath", "path").Default("/run/secrets/firebase-key").String()
	firebaseProjectID = kingpin.Flag("firebaseProjectID", "worker-51312").Default("worker-51312").String()
)

func GetFirebaseCredFilePath() string {
	kingpin.Parse()

	// return *firebaseCredFilePath
	return "conf/firebase/key.json"
}

func GetFirebaseProjectID() string {
	kingpin.Parse()

	return *firebaseProjectID
}