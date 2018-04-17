package browser

import "github.com/ryomak/browser/src/transmitter"

func Browse(url string) {
	transmitter.Transmit(url)
}
