package rand_util

import (
	"github.com/bwmarrin/snowflake"
)

var f, _ = snowflake.NewNode(1)

func Reset(node int64) error {
	n, err := snowflake.NewNode(node)
	if err != nil {
		return err
	}
	f = n
	return nil
}

func Snowflake() int64 {
	return f.Generate().Int64()
}
