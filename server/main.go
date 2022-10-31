package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	_ "github.com/lib/pq"
	"github.com/olliefr/docker-gs-ping/controller"
	"github.com/olliefr/docker-gs-ping/db"
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

var (
	articlesController controller.ArticlesController = controller.NewArticlesController()
)

func main() {
	db.ConnectToDb()
	r := gin.Default()

	articlesRoutes := r.Group("api/articles")
	{
		articlesRoutes.GET("/", articlesController.GetArticles)
		articlesRoutes.POST("/create", articlesController.SetArticles)
		articlesRoutes.DELETE("/delete", articlesController.DeleteArticles)
	}

	r.Run()
}

func exampleDbModel(db *pg.DB) {
	err := createSchema(db)

	if err != nil {
		panic(err)
	}
}

func createSchema(db *pg.DB) error {
	err := db.Model(&Log{}).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})

	if err != nil {
		return err
	}

	return nil
}
