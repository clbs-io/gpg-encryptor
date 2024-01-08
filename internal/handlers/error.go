package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	err            error `json:"-"`
	httpStatusCode int   `json:"-"`

	ErrorMessage string    `json:"error_message"`
	CreatedAt    time.Time `json:"created_at"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.httpStatusCode)
	log.Println(e.err)
	return nil
}

func ErrGPGSign(err error) render.Renderer {
	return &ErrResponse{
		err:            err,
		httpStatusCode: http.StatusInternalServerError,
		ErrorMessage:   "Error while signing data",
		CreatedAt:      time.Now(),
	}
}

func ErrGPGEncrypt(err error) render.Renderer {
	return &ErrResponse{
		err:            err,
		httpStatusCode: http.StatusInternalServerError,
		ErrorMessage:   "Error while encrypting data",
		CreatedAt:      time.Now(),
	}
}

func ErrBase64Decode(err error) render.Renderer {
	return &ErrResponse{
		err:            err,
		httpStatusCode: http.StatusBadRequest,
		ErrorMessage:   "Error while decoding base64 data",
		CreatedAt:      time.Now(),
	}
}

func ErrBase64Encode(err error) render.Renderer {
	return &ErrResponse{
		err:            err,
		httpStatusCode: http.StatusInternalServerError,
		ErrorMessage:   "Error while encoding base64 data",
		CreatedAt:      time.Now(),
	}
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		err:            err,
		httpStatusCode: http.StatusBadRequest,
		ErrorMessage:   "Invalid request",
		CreatedAt:      time.Now(),
	}
}
