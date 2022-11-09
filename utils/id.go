package utils

import (
	"aixinge/global"
	"aixinge/utils/snowflake"
	uuid "github.com/satori/go.uuid"
)

func Uuid() string {
	u := uuid.NewV4()
	return u.String()
}

func Id() snowflake.ID {
	sfn, err := snowflake.NewNode(global.CONFIG.System.Node)
	if err != nil {
		panic(err)
	}
	return sfn.Generate()
}
