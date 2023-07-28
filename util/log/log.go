package util_log

import (
	"encoding/json"
	"log"
)

func PrintPretty(obj any) {
	jsonStr, err := json.MarshalIndent(obj, "  ", "  ")
	if nil == err {
		log.Printf("%s", jsonStr)
	}
}
