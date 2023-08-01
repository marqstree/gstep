package route

import (
	"fmt"
	"github.com/marqstree/gstep/config"
	"github.com/marqstree/gstep/route/handler/NotifyHandler"
	"github.com/marqstree/gstep/route/handler/ProcessHandler"
	"github.com/marqstree/gstep/route/handler/TaskHandler"
	"github.com/marqstree/gstep/route/handler/TemplateHandler"
	"github.com/marqstree/gstep/util/net/AjaxJson"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

var Mux = http.NewServeMux()

func middleware(h http.HandlerFunc) http.HandlerFunc {
	handler := crossOrigin(h)
	handler = jsonResponseHead(h)
	handler = errorHandle(handler)
	return handler
}

func jsonResponseHead(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		h(w, r)
	}
}

func errorHandle(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()

			if nil != err {
				debug.PrintStack()
				AjaxJson.Fail(fmt.Sprintf("%s", err)).Response(w)
			}
		}()

		h(w, r)
	}
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
	//流程模板
	Mux.HandleFunc("/template/save", middleware(TemplateHandler.Save))
	//流程实例
	Mux.HandleFunc("/process/start", middleware(ProcessHandler.Start))
	//任务
	Mux.HandleFunc("/task/pass", middleware(TaskHandler.Pass))
	Mux.HandleFunc("/task/refuse", middleware(TaskHandler.Refuse))
	Mux.HandleFunc("/task/cease", middleware(TaskHandler.Cease))

	//查询
	Mux.HandleFunc("/task/pending", middleware(TaskHandler.Pending))

	//+++ 测试接口 ++++++++++++++++++
	//接收通知
	Mux.HandleFunc("/notify/task_state_change", middleware(NotifyHandler.TaskStateChange))
}
