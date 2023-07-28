package pwd

import (
	"crypto/rand"
	"math/big"
	"pwdmix/internal/cli"
)

// PasswordGenerator generates a password based on the options passed in
func PasswordGenerator(opts *cli.Options) ([]byte, error) {
	passBytes := make([]byte, opts.GetLength())

	_, err := rand.Read(passBytes)
	if err != nil {
		return nil, err
	}

	charsets := createCharsets(opts)

	// Convert each byte to a random index in the charset
	for i := range passBytes {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charsets))))
		if err != nil {
			return nil, err
		}
		passBytes[i] = charsets[index.Int64()]
	}
	return passBytes, nil
}

// createCharsets creates the password generator's charset based on the options passed in
func createCharsets(opts *cli.Options) string {
	var charsets string
	if opts.GetMode() == "n" {
		charsets = numeric
	} else {
		charsets = alphabetic
		if opts.GetMode() == "an" {
			charsets += numeric
		}
	}

	if opts.GetSpecial() {
		charsets += special
	}
	return charsets
}
