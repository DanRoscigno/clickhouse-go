package examples

import (
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/stretchr/testify/require"
	"testing"
)

func multiHostVersion() (string, error) {
	var (
		conn, err = clickhouse.Open(&clickhouse.Options{
			Addr: []string{"127.0.0.1:9001", "127.0.0.1:9002", "127.0.0.1:9000"},
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

func multiHostRoundRobinVersion() (string, error) {
	var (
		conn, err = clickhouse.Open(&clickhouse.Options{
			Addr:             []string{"127.0.0.1:9001", "127.0.0.1:9002", "127.0.0.1:9000"},
			ConnOpenStrategy: clickhouse.ConnOpenRoundRobin,
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

func TestMultiHostConnect(t *testing.T) {
	_, err := multiHostVersion()
	require.NoError(t, err)
	_, err = multiHostRoundRobinVersion()
	require.NoError(t, err)
}
