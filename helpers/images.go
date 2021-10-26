package helpers

import (
	"encoding/base64"
	"os"
)

func Base64toImage(base64str, filename string) error {
	dec, err := base64.StdEncoding.DecodeString(base64str)
	if err != nil {
		return err
	}

	f, err := os.Create("assets/product/" + filename)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		return err
	}
	if err := f.Sync(); err != nil {
		return err
	}

	return nil
}
