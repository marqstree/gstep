package route

import (
	"fmt"
	"github.com/marqstree/gstep/config"
	"github.com/marqstree/gstep/route/handler"
	"log"
	"net/http"
	"time"
)

var Mux = http.NewServeMux()

func middleware(h http.HandlerFunc) http.HandlerFunc {
	return crossOrigin(h)
}

func crossOrigin(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		h(w, r)
	}
}

func Setup() {
	setupRoutes()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", config.Config.Port),
		Handler:        Mux,
		ReadTimeout:    time.Duration(30 * int(time.Second)),
		WriteTimeout:   time.Duration(30 * int(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("web server start up at port%s", server.Addr)
	err := server.ListenAndServe()

	if nil != err {
		log.Printf("server startup fail: %v", err)
	}
}

// define route
func setupRoutes() {
	Mux.HandleFunc("/template/save", middleware(handler.SaveWorkflowTemplate))
}
