package utils

import (
	"aixinge/global"
	"aixinge/utils/snowflake"
	uuid "github.com/satori/go.uuid"
	"sync"
)

func Uuid() string {
	u := uuid.NewV4()
	return u.String()
}

var sfn *snowflake.Node
var once sync.Once

func Id() snowflake.ID {
	once.Do(func() {
		var err error
		sfn, err = snowflake.NewNode(global.CONFIG.System.Node)
		if err != nil {
			panic(err)
		}
	})
	return sfn.Generate()
}
