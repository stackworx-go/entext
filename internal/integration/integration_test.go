package integration

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/stackworx-go/entext"
	"github.com/stackworx-go/entext/internal/integration/ent"
	"github.com/stackworx-go/entext/internal/integration/ent/enttest"
	"github.com/stackworx-go/entext/internal/integration/ent/hook"
	"github.com/stackworx-go/entext/internal/integration/ent/migrate"
	"github.com/stackworx-go/entext/internal/integration/ent/user"
	"github.com/stretchr/testify/require"

	_ "github.com/jackc/pgx/v4/stdlib"
	iz "github.com/matryer/is"
)

func Test(t *testing.T) {
	is := iz.New(t)
	db, err := sql.Open("pgx", connectionString("postgres"))
	is.NoErr(err)
	for _, tt := range tests {
		name := runtime.FuncForPC(reflect.ValueOf(tt).Pointer()).Name()
		t.Run(name[strings.LastIndex(name, ".")+1:], func(t *testing.T) {
			_, err = db.ExecContext(context.Background(), "drop database if exists testing")
			is.NoErr(err)
			_, err = db.ExecContext(context.Background(), "create database testing")
			is.NoErr(err)

			testingDb, err := sql.Open("pgx", connectionString("testing"))
			is.NoErr(err)
			drv := entsql.OpenDB(dialect.Postgres, testingDb)
			defer drv.Close()
			tt(t, enttest.NewClient(t, enttest.WithOptions(ent.Driver(drv)), opts), drv)
		})
	}
}

var (
	opts = enttest.WithMigrateOptions(
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	tests = [...]func(*testing.T, *ent.Client, *entsql.Driver){
		Readonly,
		Scenario2,
		AuditMixin,
		Scenario4,
		Scenario5,
	}
)

func Readonly(t *testing.T, mutableClient *ent.Client, driver *entsql.Driver) {
	client := ent.NewReadonlyClient(ent.Driver(driver))
	client.Use(hook.Reject(ent.OpCreate | ent.OpUpdate | ent.OpUpdateOne | ent.OpDelete | ent.OpDeleteOne))

	u, err := mutableClient.User.Create().SetUsername("Ciaran").Save(context.Background())
	require.NoError(t, err)

	count := client.User.Query().CountX(context.Background())
	require.Equal(t, count, 1)

	mutableUserClient := client.User.(*ent.UserClient)
	_, err = mutableUserClient.Delete().Exec(context.Background())
	require.Errorf(t, err, "Op(31) operation is not allowed")
	client.User.Query().AllX(context.Background())
	err = mutableUserClient.DeleteOne(&ent.User{}).Exec(context.Background())
	require.Errorf(t, err, "Op(31) operation is not allowed")
	err = mutableUserClient.DeleteOneID(1).Exec(context.Background())
	require.Errorf(t, err, "Op(31) operation is not allowed")
	err = mutableUserClient.Update().Exec(context.Background())
	require.Errorf(t, err, "Op(31) operation is not allowed")
	err = mutableUserClient.UpdateOne(&ent.User{}).Exec(context.Background())
	require.Errorf(t, err, "Op(31) operation is not allowed")
	err = mutableUserClient.UpdateOneID(1).Exec(context.Background())
	require.Errorf(t, err, "Op(31) operation is not allowed")

	tx, err := client.BeginTx(context.Background(), nil)
	require.NoError(t, err)
	tx.User.GetX(context.Background(), u.ID)
	err = tx.Commit()
	require.NoError(t, err)
}

func Scenario2(t *testing.T, client *ent.Client, driver *entsql.Driver) {
	is := iz.New(t)

	ctx := entext.WithUserID(context.Background(), 1)
	u, err := client.User.Create().SetUsername("ciaran").Save(ctx)
	is.NoErr(err)
	_, err = client.UserStatus.Create().SetUser(u).Save(ctx)
	is.NoErr(err)
}

func AuditMixin(t *testing.T, client *ent.Client, driver *entsql.Driver) {
	is := iz.New(t)

	ctx := entext.WithUserID(context.Background(), 1)
	u, err := client.User.Create().SetUsername("ciaran").Save(ctx)
	is.NoErr(err)
	_, err = client.UserStatus.Create().SetUser(u).Save(ctx)
	is.NoErr(err)

	_, err = client.UserStatus.UpdateUserStatus(ctx, u, client.UserStatus.Create())
	is.NoErr(err)
}

func Scenario4(t *testing.T, client *ent.Client, driver *entsql.Driver) {
	is := iz.New(t)

	ctx := entext.WithUserID(context.Background(), 1)
	u, err := client.User.Create().SetUsername("ciaran").Save(ctx)
	is.NoErr(err)
	_, err = client.UserStatus.Create().SetUser(u).Save(ctx)
	is.NoErr(err)

	err = client.UserStatus.DeactivateUserStatus(ctx, u)
	is.NoErr(err)
}

func Scenario5(t *testing.T, client *ent.Client, driver *entsql.Driver) {
	is := iz.New(t)

	ctx := entext.WithUserID(context.Background(), 1)
	u, err := client.User.Create().SetUsername("ciaran").Save(ctx)
	is.NoErr(err)
	_, err = client.UserStatus.Create().SetUser(u).Save(ctx)
	is.NoErr(err)

	u, err = client.User.Query().WithActiveStatuses().Where(user.ID(u.ID)).Only(ctx)
	is.NoErr(err)
	is.Equal(len(u.Edges.Statuses), 1)
}

func connectionString(dbname string) string {
	ci := os.Getenv("CI")
	if ci == "true" {
		return fmt.Sprintf("host=%s port=%d user=postgres dbname=%s password=password sslmode=disable",
			"localhost", 5432, dbname)
	}
	return fmt.Sprintf("host=localhost port=%d user=postgres dbname=%s password=password sslmode=disable", 5432, dbname)
}

/*
func drop(t *testing.T, client *ent.Client) {
	t.Log("drop data from database")
	ctx := context.Background()
	client.User.Delete().ExecX(ctx)
}
*/
