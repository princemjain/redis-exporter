package config

import "fmt"

type RedisExporterConfig struct {
	Redis  RedisConfig
	Input  InputConfig
	Output OutputConfig
}

type RedisConfig struct {
	HostName   string
	Port       int
	Password   string
	BatchLimit int64
}

type InputConfig struct {
	KeyPattern string
}

type OutputConfig struct {
	FilePath          string
	MergeKey          bool
	KeySplitPattern   string
	ValueSplitPattern string
}

func (re RedisExporterConfig) GetRedisConfig() RedisConfig {
	return RedisConfig{
		HostName: re.Redis.HostName,
		Port:     re.Redis.Port,
		Password: re.Redis.Password,
	}
}

func (re RedisExporterConfig) GetOutputFilePath() string {
	return re.Output.FilePath
}

func (rc RedisConfig) GetRedisURI() string {
	return fmt.Sprintf("%s:%d", rc.HostName, rc.Port)
}
