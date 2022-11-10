package flags

import (
	"context"
	"flag"
)

type Data struct {
	SERVER1 string
	SERVER2 string
	PORT1   string
	PORT2   string
}

type DataType string

// Parse user input's flags
func ParseUserFlags(ctx *context.Context) error {
	userData := GetData(ctx)

	flag.StringVar(&userData.SERVER1, "s1", "127.0.0.1", "server 1 IP address")
	flag.StringVar(&userData.SERVER2, "s2", "127.0.0.1", "server 2 IP address")
	flag.StringVar(&userData.PORT1, "port1", "54321", "server 1 port")
	flag.StringVar(&userData.PORT2, "port2", "54322", "server 2 port")
	//TODO: add debug key
	flag.Parse()

	return nil
}

func GetData(ctx *context.Context) *Data {
	var dataCtx DataType = "dataType"
	userData := (*ctx).Value(dataCtx).(*Data)
	return userData
}
