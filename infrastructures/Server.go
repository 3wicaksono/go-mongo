package infrastructures

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	config "github.com/spf13/viper"
)

var tmpl = `
----------------------------
  %s
----------------------------
  - Port      : %s
----------------------------
`

// ServerListen the app server
func ServerListen(handler http.Handler) {

	fmt.Println(fmt.Sprintf(tmpl,
		config.GetString("app.name"),
		config.GetString("app.port"),
	))

	server := &http.Server{
		Addr:           ":" + config.GetString("app.port"),
		Handler:        handler,
		ReadTimeout:    config.GetDuration("app.read_timeout") * time.Second,
		WriteTimeout:   config.GetDuration("app.write_timeout") * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	log.Error(err)
}
