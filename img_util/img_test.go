package img_util_test

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/grin-ch/winter-common/img_util"
)

func TestNewImg(t *testing.T) {
	src, err := img_util.NewImg("ZERO")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	str := base64.StdEncoding.EncodeToString(src)
	fmt.Println(str)
}
