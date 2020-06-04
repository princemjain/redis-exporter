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
	hostName   = kingpin.Flag("hostname", "Redis master server hostname(any one)").Default("127.0.0.1").Short('h').String()
	port       = kingpin.Flag("port", "Redis server port").Short('p').Default("6379").Int()
	password   = kingpin.Flag("password", "Redis server password").String()
	batchLimit = kingpin.Flag("batch-limit", "Number of records limit to fetch").Default("50").Int64()

	// Input configuration arguments
	keyPattern = kingpin.Flag("key-pattern", "Regex pattern to find keys").Default("*").String()

	// Output configuration arguments
	outputFilePath = kingpin.Flag("output-file-path", "Destination file path").Short('o').Default("data.csv").String()
	mergeKey = kingpin.Flag("merge-key", "Add record key to output file").Bool()
	keySplitPattern = kingpin.Flag("key-split-pattern", "Before split the key format").String()
	valueSplitPattern = kingpin.Flag("value-split-pattern", "Before split the value format(ex: delimiter `:` )").String()

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
