package examples

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func sslCustomCertsVersion() (string, error) {
	t := &tls.Config{}
	caCert, err := ioutil.ReadFile("play.clickhouse.pem")
	if err != nil {
		return "", err
	}
	caCertPool := x509.NewCertPool()
	successful := caCertPool.AppendCertsFromPEM(caCert)
	if !successful {
		return "", err
	}
	t.RootCAs = caCertPool

	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"play.clickhouse.com:9440"},
		TLS:  t,
		Auth: clickhouse.Auth{
			Username: "explorer",
		},
	})
	if err != nil {
		return "", err
	}
	v, err := conn.ServerVersion()
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

func TestSSLCustomCerts(t *testing.T) {
	_, err := sslCustomCertsVersion()
	require.NoError(t, err)
}
