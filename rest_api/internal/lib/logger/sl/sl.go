package sl // сокращенно от slog

import (
	"log/slog"
)

// функция для добавление error в виде параметра лога
func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
