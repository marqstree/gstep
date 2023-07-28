package util_json

import (
	"encoding/json"
	"fmt"
)

func Obj2json(o any) (string, error) {
	result, err := json.Marshal(o)
	return fmt.Sprintf("%s", result), err
}
