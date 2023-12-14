package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"

	"github.com/sch8ill/mcmon/config"
	"github.com/sch8ill/mcmon/metrics"
	"github.com/sch8ill/mcmon/monitor"
)

func main() {
	createLogger()

	app := createApp()
	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("error encountered")
	}
}

func run(ctx *cli.Context) error {
	config.SetConfig(ctx)
	log.Info().Str("servers", strings.Join(config.Servers, ", ")).Msg("Servers to be monitored")

	m := monitor.New(config.Servers, config.PingInterval)
	m.Start()
	defer m.Stop()

	if err := metrics.Listen(); err != nil {
		return fmt.Errorf("failed to start prometheus exporter: %w", err)
	}

	return nil
}

func createApp() *cli.App {
	return &cli.App{
		Name:      "mcmon",
		Usage:     "Prometheus exporter to monitor Minecraft Servers using SLP",
		Copyright: "Copyright (c) 2023 Sch8ill",
		Action:    run,
		Flags:     config.DeclareFlags(),
	}
}

func createLogger() {
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.DateTime,
	}
	log.Logger = log.Output(consoleWriter).Level(zerolog.DebugLevel).With().Timestamp().Logger()
}
