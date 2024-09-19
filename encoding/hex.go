package encoding

import (
	"encoding/hex"
)

type HexCommand struct{}

func (h *HexCommand) Encode(input string) string {
	return hex.EncodeToString([]byte(input))
}

func (h *HexCommand) Decode(input string) (string, error) {
	decoded, err := hex.DecodeString(input)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

func (h *HexCommand) Help() string {
	return "Hex commands: encode <text>, decode <hex_string>"
}
