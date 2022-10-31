package db

import "github.com/go-pg/pg/v10"

const (
	host     = "localhost"
	port     = 5432
	user     = "test"
	password = "1234"
	dbname   = "db"
)

func ConnectToDb() *pg.DB {
	//opt, err := pg.ParseURL("postgres://postgres:1234@postgres_container:5432/db?sslmode=disable")
	opt, err := pg.ParseURL("postgres://postgres:1234@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	db := pg.Connect(opt)
	return db
}
