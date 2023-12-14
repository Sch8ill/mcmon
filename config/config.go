package config

import (
	"time"

	"github.com/urfave/cli/v2"
)

const (
	DefaultPingTimeout    = time.Second * 5
	DefaultPingInterval   = time.Second * 30
	DefaultMetricsAddress = ":9100"

	ServersFlag        = "server"
	PingTimeoutFlag    = "timeout"
	PingIntervalFlag   = "interval"
	MetricsAddressFlag = "metrics-address"
)

var (
	DefaultServers = cli.NewStringSlice("localhost:25565")

	Servers        []string
	PingTimeout    time.Duration
	PingInterval   time.Duration
	MetricsAddress string
)

func DeclareFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringSliceFlag{
			Name:  ServersFlag,
			Usage: "address of the server to be monitored",
			Value: DefaultServers,
		},
		&cli.DurationFlag{
			Name:  PingTimeoutFlag,
			Usage: "max timeout for SLP requests",
			Value: DefaultPingTimeout,
		},
		&cli.DurationFlag{
			Name:  PingIntervalFlag,
			Usage: "interval the servers should checked in",
			Value: DefaultPingInterval,
		},
		&cli.StringFlag{
			Name:  MetricsAddressFlag,
			Usage: "address the Prometheus metrics exporter listens on",
			Value: DefaultMetricsAddress,
		},
	}
}

func SetConfig(ctx *cli.Context) {
	Servers = ctx.StringSlice(ServersFlag)
	PingTimeout = ctx.Duration(PingTimeoutFlag)
	PingInterval = ctx.Duration(PingIntervalFlag)
	MetricsAddress = ctx.String(MetricsAddressFlag)
}
