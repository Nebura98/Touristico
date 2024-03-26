package repositories

import (
	"context"

	dbconfig "main/internal/adapters/database"
	query "main/internal/adapters/database/letstour"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbConnection, _ = dbconfig.Pool.GetConnection()
}

func GetUser(email string) query.GetUserRow {
	ctx := context.Background()

	q := query.New(dbConnection)
	user, err := q.GetUser(ctx, email)

	if err != nil {
		panic(err.Error())
	}

	return user
}
