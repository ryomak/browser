package transmitter

import (
	"net/http"
)

func Transmit(url string) (string, error) {
	_, err := http.Get(url)
	if err != nil {
		return "", err
	}
	return "jeu", nil
}
