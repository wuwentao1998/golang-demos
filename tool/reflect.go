package tool

import (
	"reflect"
	"runtime"
	"strings"
)

func FuncName(f interface{}) string {
	path := FuncPath(f)
	if path == "" {
		return ""
	}

	const funcSeparator = "/"
	pathList := strings.Split(path, funcSeparator)
	return pathList[len(pathList)-1]
}

func FuncPath(f interface{}) string {
	t := reflect.ValueOf(f)
	if t.Kind() != reflect.Func {
		return ""
	}

	return runtime.FuncForPC(t.Pointer()).Name()
}
