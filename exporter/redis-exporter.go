package exporter

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/princemjain/redis-exporter/config"
	"os"
	"regexp"
)

func GenerateCSV(redisExporterConfig *config.RedisExporterConfig) {
	client := initRedis(redisExporterConfig.GetRedisConfig())
	if redisExporterConfig.Input.SampleTest {
		runSampleTest(client, redisExporterConfig)
	}

	i := func(ctx context.Context, c *redis.Client) error {
		iter := c.Scan(c.Context(), 0, redisExporterConfig.Input.KeyPattern, redisExporterConfig.Input.BatchLimit).Iterator()
		var values [][]string
		for iter.Next(c.Context()) {
			key := iter.Val()
			value := client.SMembers(c.Context(), key).Val()
			values = append(values, buildOutputFormat(key, value, redisExporterConfig.Output))
		}
		writeToFile(redisExporterConfig, values)
		return nil
	}
	client.ForEachMaster(client.Context(), i)
}

func initRedis(redisConfig config.RedisConfig) *redis.ClusterClient {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{redisConfig.GetRedisURI()},
		Password: redisConfig.Password,
	})

	_, err := client.Ping(client.Context()).Result()
	if err != nil {
		fmt.Printf("[Redis] Error connecting to server %s: %s\n", redisConfig.GetRedisURI(), err)
		os.Exit(1)
	}
	fmt.Printf("[Redis] connected to Redis server: %s\n", redisConfig.GetRedisURI())
	return client
}

func buildOutputFormat(key string, values []string, outputConfig config.OutputConfig) []string {
	var data []string
	if outputConfig.MergeKey {
		if outputConfig.KeySplitPattern != "" {
			keys := RegSplit(key, outputConfig.KeySplitPattern)
			data = append(data, keys...)
		} else {
			data = append(data, key)
		}
	}

	if outputConfig.ValueSplitPattern != "" {
		values = RegSplit(values[0], outputConfig.ValueSplitPattern)
	}
	data = append(data, values...)
	return data
}

func RegSplit(text string, delimiter string) []string {
	reg := regexp.MustCompile(delimiter)
	indexes := reg.FindAllStringIndex(text, -1)
	lastStart := 0
	result := make([]string, len(indexes)+1)
	for i, element := range indexes {
		result[i] = text[lastStart:element[0]]
		lastStart = element[1]
	}
	result[len(indexes)] = text[lastStart:]
	return result
}

func writeToFile(redisExporterConfig *config.RedisExporterConfig, data [][]string) {
	file, err := os.OpenFile(redisExporterConfig.GetOutputFilePath(), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("[File] Error(%s) while opening: %s\n", err, redisExporterConfig.GetOutputFilePath())
		os.Exit(1)
	}
	defer file.Close()

	err = csv.NewWriter(file).WriteAll(data)
	if err != nil {
		fmt.Printf("[File] Error(%s) while writing : %s\n", err, redisExporterConfig.GetOutputFilePath())
		os.Exit(1)
	}
	fmt.Printf("Writing data....\n")
}

func runSampleTest(client *redis.ClusterClient, redisExporterConfig *config.RedisExporterConfig) {
	keys, _ := client.Scan(client.Context(), 0, redisExporterConfig.Input.KeyPattern, redisExporterConfig.Input.BatchLimit).Val()
	var values [][]string
	for _, key := range keys {
		value := client.SMembers(client.Context(), key).Val()
		values = append(values, buildOutputFormat(key, value, redisExporterConfig.Output))
	}
	writeToFile(redisExporterConfig, values)
	os.Exit(0)
}
