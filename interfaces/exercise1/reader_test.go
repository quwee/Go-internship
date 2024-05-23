package exercise1

import (
	"bytes"
	"io"
	"strconv"
	"testing"
)

type testcase struct {
	id          int
	data        []byte
	readSize    int64
	expectedN   int
	expectedErr error
}

var testcases = []testcase{
	{1, []byte("Hello World!"), 5, 5, nil},
	{2, []byte(""), 5, 0, io.EOF},
	{3, []byte("Hello"), 10, 5, nil},
}

func TestCustomLimitReader(t *testing.T) {
	for _, c := range testcases {
		t.Run("testcase_"+strconv.Itoa(c.id), func(t *testing.T) {
			buf := bytes.NewBuffer(c.data)
			reader := CustomLimitReader(buf, c.readSize)
			readData := make([]byte, c.readSize)

			gotN, gotErr := reader.Read(readData)

			if gotN != c.expectedN || gotErr != c.expectedErr {
				t.Fatalf("data: %v, readSize: %v, expectedN: %v, expectedErr: %v, gotN: %v, gotErr: %v",
					c.data, c.readSize, c.expectedN, c.expectedErr, gotN, gotErr)
			}
		})
	}
}
