package main

import (
	"os"
	"fmt"
	"net/http"
	"github.com/grellyd/grellyddotcom/handlers"
	"github.com/grellyd/grellyddotcom/router"
	"github.com/grellyd/filelogging/globallogger"
	"github.com/grellyd/filelogging/state"
)

func setup() (err error) {
	err = registerRoutes()
	if err != nil {
		return fmt.Errorf("setup failed: %s", err.Error())
	}
	err = globallogger.NewGlobalLogger("grellyd.com Server", state.DEBUGGING)
	if err != nil {
		return fmt.Errorf("setup failed: %s", err.Error())
	}
	return nil
}

func main() {
	err := setup()
	if err != nil {
		// Top Level Err Handle
		globallogger.Fatal(err.Error())
		os.Exit(1)
	}
	globallogger.Debug("Setup Complete")
	http.ListenAndServe(":3000", nil)
}

func registerRoutes() (err error) {
	err = router.Register("/", "(^/$)|(^/(status|about|quote)$)", handlers.Static)
	if err != nil {
		return fmt.Errorf("unable to register static: %s", err.Error())
	}
	err = router.Register("/blog/", "^/blog/([a-zA-Z0-9]*)$", handlers.Blog)
	if err != nil {
		return fmt.Errorf("unable to register blog: %s", err.Error())
	}
	err = router.Register("/css/", "^/css/([a-zA-Z0-9_]*).css$", handlers.CSS)
	if err != nil {
		return fmt.Errorf("unable to register css: %s", err.Error())
	}
	err = router.Register("/images/", "^/images/([a-zA-Z0-9_]*).jpg$", handlers.Images)
	if err != nil {
		return fmt.Errorf("unable to register images: %s", err.Error())
	}
	return nil
}
