package main

import (
	"fmt"
	"time"

	"github.com/coreos/go-systemd/daemon"
)

func main() {
	time.Sleep(time.Second * 3)
	fmt.Println(daemon.SdNotify(false, daemon.SdNotifyReady))
}
