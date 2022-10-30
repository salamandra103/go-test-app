package postgres

import "github.com/go-pg/pg/v10"

func ConnectToDb() *pg.DB {
	//opt, err := pg.ParseURL("postgres://postgres:1234@postgres_container:5432/postgres?sslmode=disable")
	opt, err := pg.ParseURL("postgres://postgres:1234@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	db := pg.Connect(opt)
	return db
}
