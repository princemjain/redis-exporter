
build	:
	GO111MODULE=on go build -ldflags "-X main.versionInfo=`cat version.txt` -X main.versionDate=`date -u +%Y-%m-%d.%H:%M:%S`"  -o ./redis-exporter main.go

package: build
		tar -czf redis-exporter-`cat version.txt`.tar.gz redis-exporter
