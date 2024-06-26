package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBase(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}

	tests := []struct {
		name string
		want want
	}{
		{
			name: "google.com",
			want: want{
				code:        200,
				response:    `{"url":"google.com_short"}`,
				contentType: "application/json",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			req := ShortUrlRequest{Url: "google.com"}

			reqJson, err := json.Marshal(req)

			if err != nil {
				log.Fatal(err)
			}
			reader := bytes.NewReader(reqJson)

			request := httptest.NewRequest(http.MethodGet, "/short", reader)
			// создаём новый Recorder
			w := httptest.NewRecorder()
			NewBaseController().shortUrl(w, request)

			res := w.Result()
			// проверяем код ответа
			assert.Equal(t, res.StatusCode, test.want.code)
			// получаем и проверяем тело запроса
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)

			require.NoError(t, err)
			assert.JSONEq(t, string(resBody), test.want.response)
			assert.Equal(t, res.Header.Get("Content-Type"), test.want.contentType)
		})
	}

}
