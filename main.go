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
	batchLimit = kingpin.Flag("batch-limit", "Batch size per fetch").Default("50").Int64()

	// Input configuration arguments
	keyPattern = kingpin.Flag("key-pattern", "Regex pattern to find keys").Default("*").String()
	sampleTest = kingpin.Flag("test", "Pull once with the batch size for testing").Bool()

	// Output configuration arguments
	outputFilePath = kingpin.Flag("output-file-path", "Destination file path").Short('o').Default("data.csv").String()
	mergeKey = kingpin.Flag("merge-key", "Add key to output file, This will be the first element in the CSV").Bool()
	keySplitPattern = kingpin.Flag("key-split-pattern", "Search and replace by(,) in the key").String()
	valueSplitPattern = kingpin.Flag("value-split-pattern", "Search and replace by(,) in the value").String()

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
		},
		Input: config.InputConfig{
			KeyPattern: *keyPattern,
			SampleTest: *sampleTest,
			BatchLimit: *batchLimit,
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
