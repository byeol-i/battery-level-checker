package auth

import (
	"context"
	"errors"
	"log"
	"time"

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

type GetResult struct {
	Result interface{}
	Error  error
}

func GetUserFromFirebase(app *FirebaseApp, ctx context.Context, uid string) GetResult {
	client, err := app.app.Auth(ctx)
	if err != nil {
		logger.Error("error getting Auth client", zap.Error(err))
		return GetResult{
			Result: nil,
			Error:  errors.New("error getting Auth client"),
		}
	}

	u, err := client.GetUser(ctx, uid)
	if err != nil {
		logger.Error("error getting user", zap.Error(err), zap.String("uid", uid))
		return GetResult{
			Result: nil,
			Error:  errors.New("error getting user"),
		}
	}

	return GetResult{
		Result: u,
		Error:  nil,
	}
}

func VerifyIDTokenFromFirebase(app *FirebaseApp, ctx context.Context, idToken string) GetResult {
	client, err := app.app.Auth(ctx)
	if err != nil {
		logger.Error("error getting Auth client", zap.Error(err))
		return GetResult{
			Result: nil,
			Error:  errors.New("error getting Auth client"),
		}
	}

	decodedToken, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		logger.Error("Can't verify token", zap.Error(err))
		return GetResult{
			Result: nil,
			Error:  err,
		}
	}

	return GetResult{
		Result: decodedToken,
		Error:  nil,
	}
}

func (hdl *FirebaseApp) GetUser(ctx context.Context, uid string) (*auth.UserRecord, error) {
	result := make(chan GetResult)

	go func() {
		result <- GetUserFromFirebase(hdl, ctx, uid)
	}()
	select {
	case <-time.After(5 * time.Second):
		return nil, errors.New("timed out")
	case result := <-result:
		if result.Error != nil {
			return nil, result.Error
		}

		if _, ok := result.Result.(*auth.UserRecord); ok {
			return nil, errors.New("Type error!")
		}

		return result.Result.(*auth.UserRecord), nil
	}
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

func (hdl *FirebaseApp) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	result := make(chan GetResult)

	go func() {
		result <- VerifyIDTokenFromFirebase(hdl, ctx, idToken)
	}()
	select {
	case <-time.After(5 * time.Second):
		return nil, errors.New("timed out")
	case result := <-result:
		if result.Error != nil {
			return nil, result.Error
		}

		if _, ok := result.Result.(*auth.Token); ok {
			return nil, errors.New("Type error")
		}

		log.Println(result.Result)
		return result.Result.(*auth.Token), nil
	}
}
