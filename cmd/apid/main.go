package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/byeol-i/battery-level-checker/pkg/controllers"
	"github.com/byeol-i/battery-level-checker/pkg/router"

	_ "github.com/byeol-i/battery-level-checker/docs" // echo-swagger middleware
	"golang.org/x/sync/errgroup"
)

// @title Battery level checker API
// @version 0.0.1
// @description This is a simple Battery level checker server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email aglide100@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost
// @BasePath /api/v1
func main() {
	if err := realMain(); err != nil {
		log.Printf("err :%s", err)
		os.Exit(1)
	}
}

func realMain() error {
	wg, ctx := errgroup.WithContext(context.Background())

	notFoundCtrl := &controllers.NotFoundController{}
	batteryCtrl := controllers.NewBatteryController()
	authCtrl := controllers.NewAuthController()
	rtr := router.NewRouter(notFoundCtrl, "v1")

	rtr.Use(authCtrl.VerifyToken)

	rtr.AddRule("Battery", "GET", `/battery$`, batteryCtrl.GetBatteryList)	
	rtr.AddRule("Battery", "GET", `/battery/[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12}$`, batteryCtrl.GetBattery)
	rtr.AddRule("Battery", "POST", `/battery/[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12}$`, batteryCtrl.UpdateBattery)	
	
	// rtr.AddRule("Auth", "POST", `/auth/login$`, authCtrl.CreateCustom)	
	rtr.AddRule("Auth", "GET", `/auth/test$`, authCtrl.LoginTest)	

	_ = ctx

	wg.Go(func () error  {
		var err error

		ln, err := net.Listen("tcp", "0.0.0.0:80")
		if err != nil {
			return err
		}

		defer ln.Close()

		srv := http.Server{Handler: rtr}

		err = srv.Serve(ln)

		return err
	})

	return wg.Wait()
}