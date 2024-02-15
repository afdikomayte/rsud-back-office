package database

import (
	"testing"

	"github.com/afdikomayte/rsud-back-office/internal/config"
	"github.com/stretchr/testify/require"
)

func init() {
	filename := "../../cmd/config.yaml"

	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}
}

func TestConnectionMysql(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, err := ConnectMysql(config.Cfg.DB)
		require.Nil(t, err)
		require.NotNil(t, db)
	})
}
