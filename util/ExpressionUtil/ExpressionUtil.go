package ExpressionUtil

import (
	"fmt"
	"github.com/dop251/goja"
	"strings"
)

func ExecuteExpression(jsTemplate string, params *map[string]any) bool {
	exp := Template2jsExpression(jsTemplate, params)
	return RunJsExpression(exp)
}

func RunJsExpression(jsExp string) bool {
	vm := goja.New()
	v, err := vm.RunString(jsExp)
	if err != nil {
		panic(err)
	}
	return v.ToBoolean()
}

func Template2jsExpression(jsTemplate string, params *map[string]any) string {
	exp := jsTemplate
	for k, v := range *params {
		strV := fmt.Sprintf("%v", v)
		exp = strings.Replace(exp, k, strV, 1)
	}
	return exp
}
