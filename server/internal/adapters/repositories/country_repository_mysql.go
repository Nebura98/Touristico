package repositories

import (
	"context"
	"database/sql"

	dbconfig "main/internal/adapters/database"
	query "main/internal/adapters/database/letstour"

	_ "github.com/go-sql-driver/mysql"
)

var dbConnection *sql.DB

func init() {
	dbConnection, _ = dbconfig.Pool.GetConnection()
}

func PostNewReview() {
	ctx := context.Background()

	q := query.New(dbConnection)
	r := query.PostReviewParams{Title: "Volcan Poas", Score: 5, Description: "El mejor que haya visitado"}
	_, err := q.PostReview(ctx, r)

	if err != nil {
		panic(err.Error())
	}
}
