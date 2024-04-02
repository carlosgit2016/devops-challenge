package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	zapConfig := []byte(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout"],
		"errorOutputPaths": ["stderr"],
		"encoderConfig": {
		  "messageKey": "message",
		  "levelKey": "level",
		  "levelEncoder": "lowercase"
		}
	}`)
	var cfg zap.Config
	if err := json.Unmarshal(zapConfig, &cfg); err != nil {
		fmt.Print(fmt.Errorf("zap Logger configuration error %s", err.Error()))
	}
	logger = zap.Must(cfg.Build())
	defer logger.Sync()
}

func Handler(dir string) http.Handler {
	handler := http.FileServer(http.Dir(dir))
	return handler
}

func main() {
	var staticdir string = "./web/public/"

	port := os.Getenv("PORT")

	logger.Info(fmt.Sprintf("Listening on :%s...", port))
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), Handler(staticdir))
	if err != nil {
		logger.Fatal(err.Error())
	}
}
