package gpg

import (
	"errors"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
)

type Client struct {
	passphrase string
	privateKey string
	publicKey  string
}

func NewClient(passphrase, privateKey string) *Client {
	return &Client{
		passphrase: passphrase,
		privateKey: privateKey,
	}
}

func (c *Client) WithPublicKey(publicKey string) *Client {
	c.publicKey = publicKey
	return c
}

func (c *Client) Sign(data []byte) (string, error) {
	message := crypto.NewPlainMessage(data)

	signingKeyRing, err := c.privateKeyRing()
	if err != nil {
		return "", err
	}

	pgpSignature, err := signingKeyRing.SignDetached(message)
	if err != nil {
		return "", err
	}

	armoredSignature, err := pgpSignature.GetArmored()
	if err != nil {
		return "", err
	}

	return armoredSignature, nil
}

func (c *Client) Encrypt(data []byte) ([]byte, error) {
	if c.publicKey == "" {
		return nil, errors.New("public key is not set")
	}

	message := crypto.NewPlainMessage(data)

	privateKeyRing, err := c.privateKeyRing()
	if err != nil {
		return nil, err
	}

	publicKeyObj, err := crypto.NewKeyFromArmored(c.publicKey)
	if err != nil {
		return nil, err
	}
	publicKeyRing, err := crypto.NewKeyRing(publicKeyObj)
	if err != nil {
		return nil, err
	}

	// encrypt and sign
	encryptedMessage, err := publicKeyRing.Encrypt(message, privateKeyRing)
	if err != nil {
		return nil, err
	}

	privateKeyRing.ClearPrivateParams()

	return encryptedMessage.GetBinary(), nil
}

func (c *Client) privateKeyRing() (*crypto.KeyRing, error) {
	privateKeyObj, err := crypto.NewKeyFromArmored(c.privateKey)
	if err != nil {
		return nil, err
	}
	unlockedKeyObj, err := privateKeyObj.Unlock([]byte(c.passphrase))
	if err != nil {
		return nil, err
	}
	privateKeyRing, err := crypto.NewKeyRing(unlockedKeyObj)
	if err != nil {
		return nil, err
	}

	return privateKeyRing, nil
}
