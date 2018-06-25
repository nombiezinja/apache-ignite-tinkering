package main

import (
	"crypto/tls"
	"net"
	"time"

	"github.com/amsokol/ignite-go-client/binary/v1"
	u "github.com/nombiezinja/chstub/utils"
)

func main() {

	// ctx := context.Background()

	// connect
	c, err := ignite.Connect(ignite.ConnInfo{
		Network: "tcp",
		Host:    "localhost",
		Port:    10800,
		Major:   1,
		Minor:   1,
		Patch:   0,
		// Username: "ignite",
		// Password: "ignite",
		Dialer: net.Dialer{
			Timeout: 10 * time.Second,
		},
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	})
	u.FailOnError(err, "Failed to connect to server")

	defer c.Close()
}
