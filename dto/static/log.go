package static

type ILogger interface {
	Info(function string, msg string)
	Warn(function string, msg string)
	Error(function string, msg string)
	Fatal(function string, msg string)
}