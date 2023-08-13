package RequestParsUtil

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Body2dto(r *http.Request, dto interface{}) {
	if (http.NoBody == r.Body) {
		return
	}

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

func GetAuthorizationToken(r *http.Request) string {
	value := r.Header.Get("Authorization")
	list := strings.Fields(value)
	if len(list) != 2 {
		return ""
	}

	return list[1]
}
