package route

import (
	"fmt"
	"github.com/marqstree/gstep/config"
	"github.com/marqstree/gstep/route/handler/DepartmentHandler"
	"github.com/marqstree/gstep/route/handler/NotifyHandler"
	"github.com/marqstree/gstep/route/handler/PositionHandler"
	"github.com/marqstree/gstep/route/handler/ProcessHandler"
	"github.com/marqstree/gstep/route/handler/TaskHandler"
	"github.com/marqstree/gstep/route/handler/TemplateHandler"
	"github.com/marqstree/gstep/util/ServerError"
	"github.com/marqstree/gstep/util/net/AjaxJson"
	"github.com/marqstree/gstep/util/net/RequestParsUtil"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

var Mux = http.NewServeMux()

func middleware(h http.HandlerFunc) http.HandlerFunc {
	handler := authHandle(h)
	handler = errorHandle(handler)
	handler = jsonResponseHead(handler)
	handler = crossOrigin(handler)
	return handler
}

func noAuthMiddleware(h http.HandlerFunc) http.HandlerFunc {
	handler := errorHandle(h)
	handler = jsonResponseHead(handler)
	handler = crossOrigin(handler)
	return handler
}

func jsonResponseHead(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		h(w, r)
	}
}

func authHandle(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		secret := config.Config.Auth.Secret
		token := RequestParsUtil.GetAuthorizationToken(r)

		if secret != token {
			panic(ServerError.New("无访问权限"))
		}

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
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,x-requested-with,Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		//注意:Access-Control-Allow-Origin不能设置成*
		if len(r.Header.Get("Origin")) > 0 {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		} else if len(r.Header.Get("Referer")) > 0 {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Referer"))
		}

		//预请求只返回响应头
		if "OPTIONS" == r.Method {
			//注意:w.WriteHeader(http.StatusAccepted)之后的w.Header().Set代码无效
			w.WriteHeader(http.StatusAccepted)
			return
		}

		h(w, r)
	}
}

func Setup() {
	setupRoutes()

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", config.Config.Port),
		Handler:      Mux,
		ReadTimeout:  time.Duration(30 * int(time.Second)),
		WriteTimeout: time.Duration(30 * int(time.Second)),
		//MaxHeaderBytes: 1 << 20,
	}
	log.Printf("web server start up at port%s", server.Addr)
	err := server.ListenAndServe()

	if nil != err {
		log.Printf("server startup fail: %v", err)
	}
}

// define route
func setupRoutes() {
	//1.流程模板
	//保存
	Mux.HandleFunc("/template/save", noAuthMiddleware(TemplateHandler.Save))
	//查询
	Mux.HandleFunc("/template/query", noAuthMiddleware(TemplateHandler.Query))
	//详情
	Mux.HandleFunc("/template/detail", noAuthMiddleware(TemplateHandler.Detail))
	//基本信息
	Mux.HandleFunc("/template/info", noAuthMiddleware(TemplateHandler.Info))
	//保存基本信息
	Mux.HandleFunc("/template/save_info", noAuthMiddleware(TemplateHandler.SaveInfo))

	//2.流程实例
	//启动流程
	Mux.HandleFunc("/process/start", noAuthMiddleware(ProcessHandler.Start))
	//任务审核
	Mux.HandleFunc("/task/pass", noAuthMiddleware(TaskHandler.Pass))
	//任务回退
	Mux.HandleFunc("/task/retreat", noAuthMiddleware(TaskHandler.Retreat))
	//任务终止
	Mux.HandleFunc("/task/cease", noAuthMiddleware(TaskHandler.Refuse))

	//查询我的任务
	Mux.HandleFunc("/task/pending", noAuthMiddleware(TaskHandler.Pending))

	//+++ 测试接口 ++++++++++++++++++
	//接收通知
	Mux.HandleFunc("/notify/task_state_change", noAuthMiddleware(NotifyHandler.TaskStateChange))

	//部门查询
	Mux.HandleFunc("/department/get_child_department", noAuthMiddleware(DepartmentHandler.GetChildDepartments))
	Mux.HandleFunc("/department/get_users", noAuthMiddleware(DepartmentHandler.GetUsers))

	//职位查询
	Mux.HandleFunc("/position/positions", noAuthMiddleware(PositionHandler.GetPositions))
}
