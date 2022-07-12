package examples

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/stretchr/testify/require"
	"testing"
)

func createCreateDrop() error {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"127.0.0.1:9000"},
	})
	if err != nil {
		return err
	}
	defer func() {
		conn.Exec(context.Background(), "DROP TABLE test_table")
	}()
	err = conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS test_table (
			Col1 UInt8,
			Col2 String
		) engine=Memory
	`)
	if err != nil {
		return err
	}
	return conn.Exec(context.Background(), "INSERT INTO test_table VALUES (1, 'test-1')")
}

func TestExec(t *testing.T) {
	require.NoError(t, createCreateDrop())
}
