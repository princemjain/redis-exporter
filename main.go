package main

import (
	"fmt"
	"github.com/princemjain/redis-exporter/config"
	"github.com/princemjain/redis-exporter/exporter"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	// Redis configuration arguments
	hostName   = kingpin.Flag("hostname", "").Default("127.0.0.1").Short('h').String()
	port       = kingpin.Flag("port", "").Short('p').Default("6379").Int()
	password   = kingpin.Flag("password", "").Default("").String()
	batchLimit = kingpin.Flag("batch-limit", "").Default("50").Int64()

	// Input configuration arguments
	keyPattern = kingpin.Flag("key-pattern", "").String()

	// Output configuration arguments
	outputFilePath = kingpin.Flag("output-file-path", "").Short('o').Default("data.csv").String()
	mergeKey = kingpin.Flag("merge-key", "").Bool()
	keySplitPattern = kingpin.Flag("key-split-pattern", "").Default("").String()
	valueSplitPattern = kingpin.Flag("value-split-pattern", "").Default("").String()

	// redis-exporter configuration arguments
	version = kingpin.Flag("version", "").Short('v') .Bool()

	// make provides the version details for the release executable
	versionInfo string
	versionDate string
)

func main() {
	// Load configuration file into redisExporterConfig
	kingpin.Parse()
	redisExporterConfig := loadConfiguration()
	if *version {
		fmt.Printf("%s - %s", versionInfo, versionDate)
		os.Exit(0)
	}

	exporter.GenerateCSV(redisExporterConfig)
}

func loadConfiguration() *config.RedisExporterConfig {
	redisExporterConfig := config.RedisExporterConfig{
		Redis: config.RedisConfig{
			HostName: *hostName,
			Port: *port,
			Password: *password,
			BatchLimit: *batchLimit,
		},
		Input: config.InputConfig{
			KeyPattern: *keyPattern,
		},
		Output: config.OutputConfig{
			FilePath: *outputFilePath,
			MergeKey: *mergeKey,
			KeySplitPattern: *keySplitPattern,
			ValueSplitPattern: *valueSplitPattern,
		},
	}
	return &redisExporterConfig
}
