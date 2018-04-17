package browser

import (
	"github.com/ryomak/browser/src/transmitter"
	"fmt"
)

func Browse(url string) {
	res ,err:= transmitter.Transmit(url)
	fmt.Println(err)
	fmt.Println(res.Body)
}
