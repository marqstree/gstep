package LogUtil

import (
	"github.com/marqstree/gstep/util/JsonUtil"
	"log"
)

func PrintPretty(obj any) {
	jsonStr := JsonUtil.Obj2PrettyJson(obj)
	log.Printf("%s", jsonStr)
}
