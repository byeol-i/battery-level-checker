package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/byeol-i/battery-level-checker/pkg/controllers"
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"github.com/byeol-i/battery-level-checker/pkg/router"

	_ "github.com/byeol-i/battery-level-checker/docs" // echo-swagger middleware
	"golang.org/x/sync/errgroup"
)

var (
	noAuth = flag.Bool("noAuth", false, "for testing")
	cacheExpirationTime = 3 * time.Second
)

type ConnectionWatcher struct {
    n int64
}

func (cw *ConnectionWatcher) OnStateChange(conn net.Conn, state http.ConnState) {
    switch state {
    case http.StateNew:
        cw.Add(1)
    case http.StateHijacked, http.StateClosed:
        cw.Add(-1)
    }
}

func (cw *ConnectionWatcher) Count() int {
    return int(atomic.LoadInt64(&cw.n))
}

func (cw *ConnectionWatcher) Add(c int64) {
    atomic.AddInt64(&cw.n, c)
}

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
	flag.Parse()
	if err := realMain(); err != nil {
		log.Printf("err :%s", err)
		os.Exit(1)
	}
}

func realMain() error {
	wg, _ := errgroup.WithContext(context.Background())

	notFoundCtrl := &controllers.NotFoundController{}
	batteryCtrl := controllers.NewBatteryController()
	deviceCtrl := controllers.NewDeviceController()
	authCtrl := controllers.NewAuthController()
	userCtrl := controllers.NewUserController()
	rtr := router.NewRouter(notFoundCtrl, "v1")

	var cw ConnectionWatcher

	if (!*noAuth) {
		rtr.Use(authCtrl.VerifyToken)
	} else {
		logger.Info("Didn't using auth server")
	}

	rtr.AddRule("Battery", "GET", `/battery/history/`, batteryCtrl.GetAllBattery)
	rtr.AddRule("Battery", "GET", `/battery/`, batteryCtrl.GetBattery)
	rtr.AddRule("Battery", "POST", `/battery/`, batteryCtrl.UpdateBattery)
	
	rtr.AddRule("Device", "POST", `/device$`, deviceCtrl.AddNewDevice)
	rtr.AddRule("Device", "DELETE", `/device/`, deviceCtrl.DeleteDevice)

	rtr.AddRule("User", "POST", "/user$", userCtrl.AddNewUser)
	rtr.AddRule("User", "DELETE", "/user/", userCtrl.DeleteUser)


	rtr.AddRule("Server", "GET", "/stress", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, strconv.Itoa(cw.Count()))
	})

	wg.Go(func() error {
		var err error
		
		ln, err := net.Listen("tcp", "0.0.0.0:80")
		if err != nil {
			return err
		}

		defer ln.Close()

		srv := http.Server{
			Handler: rtr,
			ConnState: cw.OnStateChange,
		}
		err = srv.Serve(ln)

		return err
	})

	return wg.Wait()
}
