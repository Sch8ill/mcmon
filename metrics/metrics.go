package metrics

import "github.com/prometheus/client_golang/prometheus"

var serverOnline = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "mcmon_online",
	Help: "wether the Minecraft Server is online",
}, []string{"server"})

var serverVersion = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "mcmon_version",
	Help: "version name of the Minecraft Server",
}, []string{"server", "version"})

var serverProtocolVersion = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "mcmon_protocol_version",
	Help: "protocol version of the Minecraft Server",
}, []string{"server"})

var serverOnlinePlayers = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "mcmon_players_online",
	Help: "number of current online players on the Minecraft Server",
}, []string{"server"})

var serverMaxPlayers = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "mcmon_players_max",
	Help: "max number of players allowed on the Minecraft Server",
}, []string{"server"})

var serverDescription = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "mcmon_description",
	Help: "description of the Minecraft Server",
}, []string{"server", "description"})

var serverLatency = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "mcmon_latency",
	Help: "latency of the Minecraft Server in ms",
}, []string{"server"})

func init() {
	registry.MustRegister(serverOnline, serverLatency, serverVersion, serverProtocolVersion, serverOnlinePlayers, serverMaxPlayers, serverDescription)
}

func ServerOnline(server string) {
	serverOnline.WithLabelValues(server).Set(float64(1))
}

func ServerOffline(server string) {
	serverOnline.WithLabelValues(server).Set(float64(0))
}

func ServerVersion(server string, version string) {
	serverVersion.WithLabelValues(server, version)
}

func ServerProtocolVersion(server string, protocolVersion int) {
	serverProtocolVersion.WithLabelValues(server).Set(float64(protocolVersion))
}

func ServerOnlinePlayers(server string, players int) {
	serverOnlinePlayers.WithLabelValues(server).Set(float64(players))
}

func ServerMaxPlayers(server string, players int) {
	serverMaxPlayers.WithLabelValues(server).Set(float64(players))
}

func ServerDescription(server string, description string) {
	serverDescription.WithLabelValues(server, description)
}

func ServerLatency(server string, ping int) {
	serverLatency.WithLabelValues(server).Set(float64(ping))
}
