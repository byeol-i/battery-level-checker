package auth

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/byeol-i/battery-level-checker/pkg/config"
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/api/option"
)

type FirebaseApp struct {
	app *firebase.App
}

func NewFirebaseApp() (*FirebaseApp, error) {
	path := config.GetFirebaseCredFilePath()

	opt := option.WithCredentialsFile(path)
	firebaseConfig := &firebase.Config{ProjectID: "worker-51312"}

	app, err := firebase.NewApp(context.Background(), firebaseConfig, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
		return nil, err
	}

	return &FirebaseApp{
		app: app,
	}, nil
}

func (hdl *FirebaseApp) GetUser(ctx context.Context, uid string) (*auth.UserRecord, error) {

	// [START get_user_golang]
	// Get an auth client from the firebase.App
	client, err := hdl.app.Auth(ctx)
	if err != nil {
		logger.Error("error getting Auth client", zap.Error(err))
		return nil, err
	}

	u, err := client.GetUser(ctx, uid)
	if err != nil {
		logger.Error("error getting user", zap.Error(err), zap.String("uid",uid))
		return nil, err
	}
	log.Printf("Successfully fetched user data: %v\n", u)
	// [END get_user_golang]
	return u, nil
}

func (hdl *FirebaseApp) CreateCustomToken(ctx context.Context, uid string) (string, error) {
	client, err := hdl.app.Auth(ctx)
	if err != nil {
		logger.Error("error getting Auth client", zap.Error(err))
		return "", err
	}

	token, err := client.CustomToken(ctx, uid)
	if err != nil {
		logger.Error("Can't create custom token!", zap.Error(err))
		return "", err
	}

	return token, nil
	
}

func (hdl *FirebaseApp) VerifyIDToken(ctx context.Context, idToken string) (string, error) {
	client, err := hdl.app.Auth(ctx)
	if err != nil {
		logger.Error("error getting Auth client", zap.Error(err))
		return "", err
	}

	decodedToken, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		logger.Error("can't verify token", zap.Error(err))
		return "", err
	}

	return decodedToken.UID, nil
}