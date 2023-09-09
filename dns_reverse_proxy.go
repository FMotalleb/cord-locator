//go:build !test

package main

import (
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/FMotalleb/cord-locator/lib/config"
	"github.com/FMotalleb/cord-locator/lib/utils"
	"github.com/fsnotify/fsnotify"
	"github.com/miekg/dns"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var (
	// DNSConfig is the configuration data of the instance
	DNSConfig config.Config
)

func main() {
	log.Info().Msg("Starting DNS Server")
	address := DNSConfig.Global.Address
	udpServer := &dns.Server{Addr: address, Net: "udp"}
	tcpServer := &dns.Server{Addr: address, Net: "tcp"}
	dns.HandleFunc(".", handle)

	go func() {
		if err := udpServer.ListenAndServe(); err != nil {
			log.Fatal().Msg(err.Error())
		}
	}()
	go func() {
		if err := tcpServer.ListenAndServe(); err != nil {
			log.Fatal().Msg(err.Error())
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	_ = udpServer.Shutdown()
	_ = tcpServer.Shutdown()
}

func init() {
	consoleLogger := zerolog.ConsoleWriter{Out: os.Stdout}

	log.Logger = zerolog.New(consoleLogger).With().Timestamp().Logger()

	logLevel, hasConfigFile := os.LookupEnv("LOG_LEVEL")

	if hasConfigFile {
		levelValue, err := zerolog.ParseLevel(logLevel)
		if err == nil {
			log.Info().Msgf("log level set to %s", levelValue)
			zerolog.SetGlobalLevel(levelValue)
		}
	}

	logFilePath, hasLogFilePath := os.LookupEnv("LOG_FILE")

	if hasLogFilePath {
		fileLogger, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
		if err == nil {
			log.Info().Msgf("log file set to %s", logFilePath)
			multiLogger := zerolog.MultiLevelWriter(consoleLogger, fileLogger)
			log.Logger = zerolog.New(multiLogger).With().Timestamp().Logger()
		} else {
			log.Fatal().Msgf("cannot set log file to `%s` error: `%s`", logFilePath, err)
		}
	}

	configPath, hasConfigFile := os.LookupEnv("CONFIG_FILE")
	if !hasConfigFile {
		log.Warn().Msg("`CONFIG_FILE` was missing from environment table default value is `./config.yaml`")
		configPath = "config.yaml"
	}
	file, err := os.OpenFile(configPath, os.O_RDONLY, 0664)
	if err != nil {
		log.Fatal().Msgf("config file does not found please set `CONFIG_FILE` environment, error: %v", err)
	}
	_ = file.Close()
	log.Info().Msgf("reading from config file at: %s", configPath)
	viper.SetConfigFile(configPath)

	refreshConfig()
	watchConfig, _ := os.LookupEnv("WATCH_CONFIG_FILE")
	boolVal, _ := strconv.ParseBool(watchConfig)
	if boolVal {
		log.Info().Msg("watching config file for changes")
		viper.WatchConfig()
		viper.OnConfigChange(resetDNSConfiguration)
	}
}
func handle(w dns.ResponseWriter, req *dns.Msg) {
	utils.HandleRequest(DNSConfig, w, req)
}
func refreshConfig() {
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Msgf("%v", err)
		return
	}
	err := viper.Unmarshal(&DNSConfig)
	if err != nil {
		log.Warn().Msgf("config changed but it has an error: %v", err)
	}
	if !DNSConfig.Validate() {
		panic("config validation failed")
	}
}

func resetDNSConfiguration(event fsnotify.Event) {
	if event.Op == fsnotify.Write {
		refreshConfig()
		log.Info().Msg("Dns Config refreshed. Keep in mind that serving port will not change until you reset dns server")
		dns.HandleRemove(".")
		dns.HandleFunc(".", handle)
	}
}
