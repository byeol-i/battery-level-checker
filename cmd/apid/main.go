package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/aglide100/battery-level-checker/pkg/controllers"
	"github.com/aglide100/battery-level-checker/pkg/router"

	// _ "github.com/aglide100/battery-level-checker/docs"
	// echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
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
// @BasePath /v1
func main() {
	if err := realMain(); err != nil {
		log.Printf("err :%s", err)
		os.Exit(1)
	}
}

func realMain() error {
	wg, ctx := errgroup.WithContext(context.Background())
	rtr := router.NewRouter()

	testCtrl := controllers.NewTestController()
	deviceCtrl := controllers.NewDeviceController()

	// Swagger Ui middleware server
	// e := echo.New()
	// e.GET("/swagger/*", echoSwagger.WrapHandler)
	// e.Logger.Fatal(e.Start(":1323"))

	rtr.AddRule("", "GET", "^/$", testCtrl.PrintHello)
	// uuid
	
	rtr.AddRule("Battery", "GET", "^/api/v1/battery/[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12}$", deviceCtrl.GetBattery)
	rtr.AddRule("Battery", "POST", "^/api/v1/battery/[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12}$", deviceCtrl.UpdateBattery)
	
	// rtr.AddRule("")
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