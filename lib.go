package mitftp

import (
	"github.com/pin/tftp"
	"io"
)

func DownloadFile(addr, file string, w io.Writer) (e error) {
	var c *tftp.Client
	c, e = tftp.NewClient(addr)
	var wt io.WriterTo
	if e == nil {
		wt, e = c.Receive(file, "octet")
	}
	if e == nil {
		_, e = wt.WriteTo(w)
	}
	return
}

func ServeFile(addr string, r io.Reader, w io.Writer) (e error) {
	var s *tftp.Server
	s = tftp.NewServer(
		func(fn string, rf io.ReaderFrom) (re error) {
			_, re = rf.ReadFrom(r)
			return
		},
		func(fn string, wt io.WriterTo) (we error) {
			_, we = wt.WriteTo(w)
			return
		},
	)
	e = s.ListenAndServe(addr)
	if e != nil {
		println(e.Error())
	}
	return
}
