package convert

import (
	"testing"

	"github.com/expect-digital/translate/pkg/model"
	"github.com/stretchr/testify/assert"
)

func Test_fromNgxTranslate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		serialized   []byte
		wantMessages model.Messages
		wantErr      bool
	}{
		{
			name:       "Not nested",
			serialized: []byte(`{"hello":"world"}`),
			wantMessages: model.Messages{
				Messages: []model.Message{
					{
						ID:      "hello",
						Message: "world",
					},
				},
			},
			wantErr: false,
		},
		{
			name:       "Nested normally",
			serialized: []byte(`{"hello":{"beautiful":"world"}}`),
			wantMessages: model.Messages{
				Messages: []model.Message{
					{
						ID:      "hello.beautiful",
						Message: "world",
					},
				},
			},
			wantErr: false,
		},
		{
			name:       "Nested with dot",
			serialized: []byte(`{"hello.beautiful":"world"}`),
			wantMessages: model.Messages{
				Messages: []model.Message{
					{
						ID:      "hello.beautiful",
						Message: "world",
					},
				},
			},
			wantErr: false,
		},
		{
			name:       "Nested mixed",
			serialized: []byte(`{"hello.beautiful":"world","hello":{"beautiful":"world"}}`),
			wantMessages: model.Messages{
				Messages: []model.Message{
					{
						ID:      "hello.beautiful",
						Message: "world",
					},
					{
						ID:      "hello.beautiful",
						Message: "world",
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, err := FromNgxTranslate(tt.serialized)

			assert.NoError(t, err)
			assert.Equal(t, tt.wantMessages, result)
		})
	}
}

func Test_toNgxTranslate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		messages model.Messages
		wantB    []byte
		wantErr  bool
	}{
		{
			name: "Messages to NGX-Translate",
			messages: model.Messages{
				Messages: []model.Message{
					{
						ID:      "hello",
						Message: "world",
					},
					{
						ID:      "hello.beautiful",
						Message: "world",
					},
				},
			},
			wantB:   []byte(`{"hello":"world","hello.beautiful":"world"}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, err := ToNgxTranslate(tt.messages)
			assert.NoError(t, err)

			assert.Equal(t, tt.wantB, result)
		})
	}
}
