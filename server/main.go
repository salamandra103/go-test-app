package main

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"time"
)

type User struct {
	Id     int64
	Name   string
	Emails []string
}

type Story struct {
	Id       int64
	Title    string
	AuthorId int64
	Author   *User `pg:"rel:has-one"`
}

type Log struct {
	tableName struct{} `pg:"logs,partition_by:RANGE(log_time)"`

	Id        int       `pg:"id,pk"`
	LogString string    `pg:"log_string"`
	LogTime   time.Time `pg:"log_time,pk"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "test"
	password = "1234"
	dbname   = "postgres"
)

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %v>", u.Id, u.Name, u.Emails)
}
func (s Story) String() string {
	return fmt.Sprintf("Story<%d %s %s>", s.Id, s.Title, s.Author)
}

func main() {
	exampleDbModel()
	startServer()
}

func exampleDbModel() {
	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	//	"password=%s dbname=%s sslmode=disable",
	//	host, port, user, password, dbname)
	//
	//db, err := sql.Open("postgres", psqlInfo)
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer db.Close()
	//
	//err = db.Ping()
	//if err != nil {
	//	panic(err)
	//}

	db := pg.Connect(&pg.Options{
		User:     "test",
		Password: "1234",
		Database: "postgres",
	})

	//opt, err := pg.ParseURL("postgres://test:1234@127.0.0.1:5432/postgres")
	//if err != nil {
	//	panic(err)
	//}
	//db := pg.Connect(opt)

	err := db.Ping(db.Context())
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = createSchema(db)

	//if err != nil {
	//	panic(err)
	//}
	//
	//user1 := &User{
	//	Name:   "user",
	//	Emails: []string{"email1", "email2"},
	//	Id:     200,
	//}
	//
	//_, err = db.Model(user1).Insert()
	//
	//if err != nil {
	//	panic(err)
	//}

}

func startServer() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hell123oo! <3")
	})

	e.GET("/createModel", func(c echo.Context) error {
		exampleDbModel()
		return c.String(http.StatusOK, "dsadsa")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}

func createSchema(db *pg.DB) error {
	//models := []interface{}{
	//	(*User)(nil),
	//	(*Story)(nil),
	//}
	//
	//for _, model := range models {
	//	err := db.Model(model).CreateTable(&orm.CreateTableOptions{Temp: true})
	//
	//	if err != nil {
	//		return err
	//	}
	//}

	err := db.Model(&Log{}).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})

	if err != nil {
		return err
	}

	return nil
}
