package config

import "gopkg.in/alecthomas/kingpin.v2"

var (
	firebaseCredFilePath = kingpin.Flag("firebaseCredFilePath", "./conf/firebase/key.json").Default("./conf/firebase/key.json").String()
	firebaseProjectID = kingpin.Flag("firebaseProjectID", "worker-51312").Default("worker-51312").String()
)

func GetFirebaseCredFilePath() string {
	kingpin.Parse()

	return *firebaseCredFilePath
}

func GetFirebaseProjectID() string {
	kingpin.Parse()

	return *firebaseProjectID
}