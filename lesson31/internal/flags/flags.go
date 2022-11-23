package flags

import (
	"context"
	"flag"
)

type Data struct {
	PORT  string
	PEERS string
}

type DataType string

// ParseUserFlags parses user input's flags
func ParseUserFlags(ctx *context.Context) error {
	userData := GetData(ctx)

	flag.StringVar(&userData.PORT, "port", "54321", "server's port")
	flag.StringVar(&userData.PEERS, "peers", "127.0.0.1:54322,127.0.0.1:54323", "balancer upstream addresses in <IP:PORT>[,<IP:PORT>] format")
	//TODO: add debug key
	flag.Parse()

	return nil
}

func GetData(ctx *context.Context) *Data {
	var dataCtx DataType = "dataType"
	userData := (*ctx).Value(dataCtx).(*Data)
	return userData
}
