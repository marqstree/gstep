package RequestParsUtil

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Body2dto(r *http.Request, dto interface{}) {
	err := json.NewDecoder(r.Body).Decode(&dto)
	if nil != err {
		panic(err)
	}
}

func Body2map(r *http.Request) map[string]interface{} {
	s, _ := io.ReadAll(r.Body)
	if len(s) == 0 {
		return nil
	}
	m := make(map[string]interface{})
	err := json.Unmarshal(s, &m)
	if err != nil {
		panic(err)
	}
	return m
}

func PrintRequest(r *http.Request) {
	s, _ := io.ReadAll(r.Body)
	body := make(map[string]interface{})
	json.Unmarshal(s, &body)
	fmt.Println("body:", body)
}
