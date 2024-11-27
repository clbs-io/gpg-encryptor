package handlers

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/render"

	"github.com/cybroslabs/gpg-encryptor/internal/gpg"
)

// Sign example
//
//	@Summary		  Sign data
//	@Description	Sign a file supplied in the request body using the public key supplied in the request body.
//	@Tags         GPG
//	@Accept			  json
//	@Produce		  json
//	@Param  		  signOnlyRequest body      SignRequest   true "A JSON struct with base64-encoded file (bin) and strings per private key and passphrase"
//	@Success		  200		          {object}	SignResponse "Returns JSON with GPG signature (in plain text)"
//	@Failure		  400		          {object}  ErrResponse	"Error response, bad request"
//	@Failure		  500		          {object}  ErrResponse	"Error while signing file"
//	@Router			  /v1/sign        [post]
func Sign(w http.ResponseWriter, r *http.Request) {
	var req SignRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Error decoding request: %v", err)
		if err = render.Render(w, r, ErrInvalidRequest(err)); err != nil {
			log.Printf("Error rendering invalid request: %v", err)
		}
		return
	}

	decodedData, err := base64.StdEncoding.DecodeString(req.Data)
	if err != nil {
		log.Printf("Error decoding base64 data: %v", err)
		if err = render.Render(w, r, ErrBase64Decode(err)); err != nil {
			log.Printf("Error rendering base64 decode error: %v", err)
		}
		return
	}

	gpg := gpg.NewClient(req.GPGPassphrase, req.GPGPrivKey)

	signature, err := gpg.Sign(decodedData)
	if err != nil {
		log.Printf("Error signing data: %v", err)
		if err = render.Render(w, r, ErrGPGSign(err)); err != nil {
			log.Printf("Error rendering gpg sign error: %v", err)
		}
		return
	}

	if err = render.Render(w, r, NewSignResponse(signature)); err != nil {
		log.Printf("Error rendering sign response: %v", err)
	}
}

type SignRequest struct {
	Data          string `json:"data"`
	GPGPrivKey    string `json:"gpg_private_key"`
	GPGPassphrase string `json:"gpg_passphrase"`
}

type SignResponse struct {
	GPGSignature string `json:"gpg_signature"`
}

func (s *SignResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)
	return nil
}

func NewSignResponse(signature string) render.Renderer {
	return &SignResponse{
		GPGSignature: signature,
	}
}
