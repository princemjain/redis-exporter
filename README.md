# Redis Exporter

Exports all data from redis cluster to CSV

## Installation

To install on OS X
```
brew tap princemjain/homebrew-redis-exporter
brew install redis-exporter
```  
To upgrade
```
brew update
brew upgrade redis-exporter
```

## Usage
``` sh
$ redis-exporter --help
  usage: redis-exporter [<flags>]
  
  Flags:
        --help                  Show context-sensitive help (also try --help-long and --help-man).
    -h, --hostname="127.0.0.1"  Redis master server hostname(any one)
    -p, --port=6379             Redis server port
        --password=PASSWORD     Redis server password
        --batch-limit=50        Batch size per fetch
        --key-pattern="*"       Regex pattern to find keys
        --test                  Pull once with the batch size for testing
    -o, --output-file-path="data.csv"
                                Destination file path
        --merge-key             Add key to output file, This will be the first element in the CSV
        --key-split-pattern=KEY-SPLIT-PATTERN
                                Search and replace by(,) in the key
        --value-split-pattern=VALUE-SPLIT-PATTERN
                                Search and replace by(,) in the value
    -v, --version
```
