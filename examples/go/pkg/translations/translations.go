package translations

import (
	_ "golang.org/x/text/message"
)

//nolint:all
//go:generate gotext -srclang en-GB update -out catalog.go -lang=en-GB,ru-RU,de-DE,lv-LV example.com/translate-example/cmd/main
