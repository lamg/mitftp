package main

import (
	"bytes"
	. "github.com/lamg/mitftp"
)

func main() {
	ServeFile(":1069", bytes.NewBufferString("Hola mundo"),
		bytes.NewBuffer(""))
}
