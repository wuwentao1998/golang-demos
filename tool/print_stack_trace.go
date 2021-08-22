package tool

import (
	"fmt"
	"runtime"
	"strings"
)

func PrintStackTrace(err interface{}) string {
	builder := new(strings.Builder)

	_, _ = fmt.Fprintf(builder, "%v\n", err)
	for i := 1; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		_, _ = fmt.Fprintf(builder, "%s:%d (0x%x)\n", file, line, pc)
	}

	return builder.String()

}
