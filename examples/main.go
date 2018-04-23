package main

import (
	"flag"
	"fmt"

	"github.com/17media/emoji"
)

var message = flag.String("m", ":beer: is great!", "the message to emojitize")

func main() {
	flag.Parse()

	fmt.Println(emoji.Emojitize(*message))
}
