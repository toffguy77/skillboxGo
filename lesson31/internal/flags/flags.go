package flags

import (
	"context"
	"flag"
)

type Data struct {
	PORT  string
	PEERS string
	DB    string
}

type DataType string

// ParseUserFlags parses user input's flags
func ParseUserFlags(ctx *context.Context) error {
	userData := GetData(ctx)

	flag.StringVar(&userData.PORT, "port", "54321", "server's port")
	flag.StringVar(&userData.PEERS, "peers", "server-instance1:54321,server-instance2:54321", "balancer upstream addresses in <IP:PORT>[,<IP:PORT>] format")
	flag.StringVar(&userData.DB, "db", "mongodb:27017", "container and port for database connections")
	//TODO: add debug key
	//TODO: add IP:HOST validation

	flag.Parse()

	return nil
}

func GetData(ctx *context.Context) *Data {
	var dataCtx DataType = "dataType"
	userData := (*ctx).Value(dataCtx).(*Data)
	return userData
}
