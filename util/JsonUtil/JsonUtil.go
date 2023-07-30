package JsonUtil

import (
	"encoding/json"
	"fmt"
)

func Obj2json(o any) string {
	result, err := json.Marshal(o)

	if nil != err {
		panic(err)
	}

	return fmt.Sprintf("%s", result)
}
