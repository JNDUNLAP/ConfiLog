package logger

const (
	colorRed     = "\033[91m"
	colorGreen   = "\033[92m"
	colorYellow  = "\033[93m"
	colorMagenta = "\033[95m"
	colorCyan    = "\033[36m"
	colorWhite   = "\033[17m"
	colorReset   = "\033[0m"
	timeColor    = "\033[1;30m"
)

var levelColors = map[Level]string{
	DEBUG:   colorCyan,
	INFO:    colorGreen,
	WARNING: colorYellow,
	ERROR:   colorRed,
	FATAL:   colorMagenta,
	PANIC:   colorRed,
}
