package db

import (
	"database/sql"
	"fmt"
	"Api/configs"

	_ "github.com/lib/pq"
)

func OpenConnection()(*sql.DB, error){
	conf:= configs.GetDB()


	sc:= fmt.Sprint(
		"host=", conf.Host,
		" port=", conf.Port,
		" user=", conf.User,
		" password", conf.Password,
		" dbname=", conf.Database,
		" sslmode=disable",
	)

	conn, err:= sql.Open("postgres", sc)

	if err!= nil {
		panic(err)

	}

	err = conn.Ping()

	return conn, err

}