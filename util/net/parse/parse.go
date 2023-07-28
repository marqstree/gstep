package util_net_parse

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Body2dto(r *http.Request, dto interface{}) error {
	err := json.NewDecoder(r.Body).Decode(&dto)
	return err
}

func Body2map(r *http.Request) (map[string]interface{}, error) {
	s, _ := io.ReadAll(r.Body)
	if len(s) == 0 {
		return nil, nil
	}
	m := make(map[string]interface{})
	err := json.Unmarshal(s, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func PrintRequest(r *http.Request) {
	s, _ := io.ReadAll(r.Body)
	body := make(map[string]interface{})
	json.Unmarshal(s, &body)
	fmt.Println("body:", body)
}
