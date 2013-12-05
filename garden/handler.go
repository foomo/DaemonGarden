package garden

import (
	"net/http"
	"net/url"
	"strings"
)

type CallExecuter interface {
	Execute(rawCall []string) (reply string, statusCode int)
}

type Handler struct {
	callExecuter CallExecuter
}

func NewHandler(callExecuter CallExecuter) *Handler {
	handler := new(Handler)
	handler.callExecuter = callExecuter
	return handler
}

func decodeURI(uri string) (decoded []string, err error) {
	parts := strings.Split(uri, "/")
	for _, part := range parts {
		decodedPart, decodingErr := url.QueryUnescape(part)
		if decodingErr == nil {
			if len(decodedPart) > 0 {
				decoded = append(decoded, decodedPart)
			}
		} else {
			err = decodingErr
			return
		}
	}
	return
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoded, err := decodeURI(r.RequestURI)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("url decoding error"))
		return
	} else {
		reply, statusCode := h.callExecuter.Execute(decoded)
		w.WriteHeader(statusCode)
		w.Write([]byte(reply))
	}
}
