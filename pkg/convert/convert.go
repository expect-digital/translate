package convert

import "github.com/expect-digital/translate/pkg/model"

type From = func([]byte) (model.Messages, error)
