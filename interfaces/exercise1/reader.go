package exercise1

import (
	"fmt"
	"io"
)

type CustomLimitedReader struct {
	R io.Reader
	N int64
}

func (c CustomLimitedReader) Read(p []byte) (int, error) {
	if c.N <= 0 {
		return 0, fmt.Errorf("invalid read byte size: %d", c.N)
	}
	if int64(len(p)) > c.N {
		p = p[:c.N]
	}
	return c.R.Read(p)
}

func CustomLimitReader(r io.Reader, n int64) io.Reader {
	return CustomLimitedReader{r, n}
}
