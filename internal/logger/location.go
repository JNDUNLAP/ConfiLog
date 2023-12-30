package logger

import "runtime"

func GetFunctionNameAndLine() (string, int) {
	pc, _, line, ok := runtime.Caller(3)
	if !ok {
		return "unknown", 0
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "unknown", line
	}

	return fn.Name(), line
}
