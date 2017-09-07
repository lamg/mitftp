package mitftp

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	cont = "Hola coco"
)

func TestTFTP(t *testing.T) {
	var fr, fw, dw *bytes.Buffer
	fr, fw, dw = bytes.NewBufferString(cont),
		bytes.NewBufferString(""),
		bytes.NewBufferString("")
	go ServeFile(":1069", fr, fw)
	var e error
	e = DownloadFile("localhost:1069", "puah", dw)
	require.NoError(t, e)
	require.True(t, dw.String() == cont)
}
