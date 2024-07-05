package compress_util_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/grin-ch/winter-common/compress_util"
)

func TestGzipCompress(t *testing.T) {
	var src = []byte(strings.Repeat("abcdefg", 100))
	fmt.Println(len(src))
	got, err := compress_util.GzipCompress(src)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(len(got))
	data, err := compress_util.GzipDecompress(got)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(len(data))
}
