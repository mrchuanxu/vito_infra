package util

import (
	"context"
	"encoding/json"
	"regexp"

	"google.golang.org/grpc/metadata"
)

func Convert(s, v interface{}) {
	if v == nil {
		panic("val is nil")
	}
	if s == nil {
		panic("source is nil")
	}
	b, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(b, v); err != nil {
		panic(err)
	}
}

func SetDBCodeCtx(ctx context.Context, dbCode string) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return metadata.NewIncomingContext(ctx, metadata.MD{})
	}
	newMD := md.Copy()
	newMD.Set("db_code", dbCode)
	return metadata.NewIncomingContext(ctx, newMD)
}

func GetNumberByStr(str string) string {
	reg := regexp.MustCompile(`[0-9]\d{0,16}`)
	return reg.FindString(str)
}
