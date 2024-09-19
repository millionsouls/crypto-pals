package encoding

import (
	"encoding/base64"
)

type Base64Command struct{}

func (b *Base64Command) Encode(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

func (b *Base64Command) Decode(input string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

func (b *Base64Command) Help() string {
	return "Base64 commands: encode <text>, decode <base64_string>"
}
