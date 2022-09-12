package auth

import (
	"github.com/byeol-i/battery-level-checker/pkg/grpcSvc/client"
)

func CheckToken(token string) error {
	// firebase admin sdk
	
	err := client.CallAuth(token)
	if err != nil {
		return err
	}
	

	return nil
}