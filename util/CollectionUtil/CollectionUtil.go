package CollectionUtil

import (
	"container/list"
	"encoding/json"
	"sort"
)

func List2Array(list *list.List) []interface{} {
	var len = list.Len()
	if len == 0 {
		return nil
	}
	var arr []interface{}
	for e := list.Front(); e != nil; e = e.Next() {
		arr = append(arr, e.Value)
	}
	return arr
}

// ExistsDuplicateInStringsArr 字符串数组中是否存在重复元素
func ExistsDuplicateInStringsArr(arr []string) bool {
	length := len(arr)
	sort.Strings(arr)
	for i := 1; i < length; i++ {
		if arr[i-1] == arr[i] {
			return true
		}
	}
	return false
}

func Obj2map(obj any) (map[string]any, error) {
	jsonStr, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	if err := json.Unmarshal(jsonStr, &result); err != nil {
		return nil, err
	}
	return result, nil
}
