package bug

import (
	"ariga.io/atlas/sql/migrate"
	"context"
	"entgo.io/ent/dialect/sql/schema"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"github.com/lagzi/EntSymbolBug/ent"
	"github.com/lagzi/EntSymbolBug/ent/enttest"
)

func TestBugPostgres(t *testing.T) {
	for version, port := range map[string]int{"12": 5432} {
		t.Run(version, func(t *testing.T) {
			client := enttest.Open(t, dialect.Postgres, fmt.Sprintf("host=localhost port=%d user=postgres dbname=test password=pass sslmode=disable", port), enttest.WithMigrateOptions(schema.WithUniversalID(), schema.WithAtlas(true)))
			defer client.Close()
			test(t, client)
		})
	}
}

func test(t *testing.T, client *ent.Client) {
	ctx := context.Background()

	dir, err := migrate.NewLocalDir("migrations")
	require.NoError(t, err)

	// Diff should be empty on happy flow
	err = client.Schema.Diff(ctx, schema.WithDir(dir), schema.WithUniversalID(), schema.WithAtlas(true))
	require.NoError(t, err)
}
