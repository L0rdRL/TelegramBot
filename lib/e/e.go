package e

import "fmt"

func Wrap(msg string, err error) error {
	return nil, fmt.Errorf("%s: %w", msg, err)
}

func WrapIfErr(msg string, err error) error {
	if err == nil {
		return nil
	}
	return Wrap(msg, err)
}