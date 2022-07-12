package examples

import (
	"crypto/tls"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/stretchr/testify/require"
	"testing"
)

func sslVersion() (string, error) {
	var (
		conn, err = clickhouse.Open(&clickhouse.Options{
			Addr: []string{"play.clickhouse.com:9440"},
			TLS:  &tls.Config{},
			Auth: clickhouse.Auth{
				Username: "explorer",
			},
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

func TestSSL(t *testing.T) {
	_, err := sslVersion()
	require.NoError(t, err)
}
