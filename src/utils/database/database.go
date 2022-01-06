package database

import (
	"app/ent"
	"app/utils"
	"fmt"
	_ "github.com/lib/pq"
)

type Database interface {
	GetEntClient()
}

func GetEntClient() (*ent.Client, error) {
	var host string = utils.MustGet("POSTGRES_HOST")
	var user string = utils.MustGet("POSTGRES_USER")
	var password string = utils.MustGet("POSTGRES_PASSWORD")
	var dbname string = utils.MustGet("POSTGRES_DB")

	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)
	client, err := ent.Open("postgres", connectionString)

	if err != nil {
		fmt.Println(err)
	}

	return client, err
}
