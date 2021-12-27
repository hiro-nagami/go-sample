package database

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
    "app/utils"
)

func SetupDatabase() (*sql.DB, error){
    var host string = utils.MustGet("POSTGRES_HOST")
    var user string = utils.MustGet("POSTGRES_USER")
    var password string = utils.MustGet("POSTGRES_PASSWORD")
    var dbname string = utils.MustGet("POSTGRES_DB")

    var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)
    db, err := sql.Open("postgres", connectionString)

    if err != nil {
        fmt.Println(err)
    }
    
    // defer db.Close()
    
    return db, err
}