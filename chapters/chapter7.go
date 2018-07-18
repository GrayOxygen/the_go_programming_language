package chapters

import (
	"io"
)

func F(out io.Writer) {
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}

