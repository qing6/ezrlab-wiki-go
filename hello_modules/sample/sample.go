package sample

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"golang.org/x/image/font"
)

// JSONUnmarshal read json file by using path, and unmarshal the content bytes to v
func JSONUnmarshal(file string, v interface{}) error {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return errors.WithMessage(err, "open file fail")
	}
	if err = json.NewDecoder(f).Decode(v); err != nil {
		return errors.WithMessage(err, "do decode fail")
	}
	return nil
}

// JSONMarshal marshal v to json bytes, write bytes to file
func JSONMarshal(file string, v interface{}) error {
	f, err := os.Create(file)
	defer f.Close()
	if err != nil {
		return errors.WithMessage(err, "open file fail")
	}
	if err = json.NewEncoder(f).Encode(v); err != nil {
		return errors.WithMessage(err, "do encode fail")
	}
	return nil
}

func mock() {
	fmt.Println(font.HintingNone)
}
