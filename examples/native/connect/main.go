package main

import (
	"github.com/ClickHouse/clickhouse-go/v2"
	"log"
)

func version() (string, error) {
	var (
		conn, err = clickhouse.Open(&clickhouse.Options{
			Addr: []string{"127.0.0.1:9000"},
		})
	)
	if err != nil {
		return "", err
	}
	v, err := conn.ServerVersion()
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

func main() {
	version, err := version()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(version)
}
