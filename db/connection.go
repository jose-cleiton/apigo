package db

import (
	"Api/configs"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func OpenConnection(*sql.DB, error){
	conf:= configs.GetDB()


	sc:= fmt.Sprint(
		"host=%s",
		" port=%s",
		" user=%s",
		" password=%s",
		" dbname=%s",
		" sslmode=disable",

		conf.Host,
		conf.Port,
		conf.User,
		conf.Password,
		conf.Database,

		
	)

	conn, err:= sql.Open("postgre", sc)

	if err!= nil {
		panic(err)

	}

	err = conn.Ping()

	return conn, err

}