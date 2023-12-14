package monitor

import (
	"fmt"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/sch8ill/mclib/server"
	"github.com/sch8ill/mclib/slp"

	"github.com/sch8ill/mcmon/config"
	"github.com/sch8ill/mcmon/metrics"
)

type Monitor struct {
	servers   []string
	interval  time.Duration
	stopCh    chan struct{}
	waitGroup sync.WaitGroup
}

func New(servers []string, interval time.Duration) *Monitor {
	return &Monitor{
		servers:  servers,
		interval: interval,
	}
}

func (m *Monitor) Start() {
	log.Info().Msg("Starting server monitor")
	m.waitGroup.Add(1)
	go m.run()
}

func (m *Monitor) Stop() {
	close(m.stopCh)
	m.waitGroup.Wait()
}

func (m *Monitor) run() {
	defer m.waitGroup.Done()

	for {
		select {
		case <-m.stopCh:
			return

		default:
			for _, srv := range m.servers {
				go m.monitor(srv)
			}
			time.Sleep(m.interval)
		}
	}
}

func (m *Monitor) monitor(srv string) {
	if err := m.ping(srv); err != nil {
		log.Warn().Err(err).Msg("failed to ping server")
		metrics.ServerOffline(srv)
	}
}

func (m *Monitor) ping(srv string) error {
	s, err := server.New(srv, server.WithTimeout(config.PingTimeout))
	if err != nil {
		return fmt.Errorf("failed to create server instance: %w", err)
	}

	res, err := s.StatusPing()
	if err != nil {
		return fmt.Errorf("failed to ping server: %w", err)
	}

	m.submitMetrics(srv, res)
	return nil
}

func (m *Monitor) submitMetrics(srv string, res *slp.Response) {
	metrics.ServerOnline(srv)
	metrics.ServerVersion(srv, res.Version.Name)
	metrics.ServerProtocolVersion(srv, res.Version.Protocol)
	metrics.ServerOnlinePlayers(srv, res.Players.Online)
	metrics.ServerMaxPlayers(srv, res.Players.Max)
	metrics.ServerDescription(srv, res.Description.String())
	metrics.ServerLatency(srv, res.Latency)
}
