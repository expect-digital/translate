package convert

import "github.com/expect-digital/translate/pkg/model"

type To func(model.Messages) ([]byte, error)
