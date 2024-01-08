package handlers

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/cybroslabs/gpg-encryptor/internal/gpg"

	"github.com/go-chi/render"
)

// Encrypt example
//
//	@Summary		Encrypt and sign a file
//	@Description	Encrypt and sign a file supplied in the request body using the public key supplied in the request body.
//	@Tags       GPG
//	@Accept			json
//	@Produce		json
//	@Param  		encryptRequest body      EncryptRequest   true  "A JSON struct with base64 encoded file (bin) and strings per public key, private key and passphrase"
//	@Success		200			       {object}	EncryptResponse  "Returns JSON with base64-encoded encrypted and signed file"
//	@Failure		400			       {object}	ErrResponse	            "Error response, bad request"
//	@Failure		500			       {object}	ErrResponse	            "Error response, error while processing request"
//	@Router			/v1/encrypt    [post]
func Encrypt(w http.ResponseWriter, r *http.Request) {
	var req EncryptRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	decodedData, err := base64.StdEncoding.DecodeString(req.Data)
	if err != nil {
		render.Render(w, r, ErrBase64Decode(err))
		return
	}

	gpg := gpg.NewClient(req.GPGPassphrase, req.GPGPrivKey).WithPublicKey(req.GPGPubKey)

	encryptedData, err := gpg.Encrypt(decodedData)
	if err != nil {
		render.Render(w, r, ErrGPGEncrypt(err))
		return
	}

	render.Render(w, r, NewEncryptResponse(encryptedData))
}

type EncryptRequest struct {
	Data          string `json:"data"`
	GPGPubKey     string `json:"gpg_public_key"`
	GPGPrivKey    string `json:"gpg_private_key"`
	GPGPassphrase string `json:"gpg_passphrase"`
}

type EncryptResponse struct {
	Data string `json:"data"` // base64-encoded binary data
}

func (e *EncryptResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)
	return nil
}

func NewEncryptResponse(data []byte) render.Renderer {
	encodedData := base64.StdEncoding.EncodeToString(data)

	return &EncryptResponse{
		Data: encodedData,
	}
}
