package models

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

var database *sql.DB

func InitDatabase(databaseName string) {

    var err error

    database, err = sql.Open("sqlite3", databaseName)

    if err == nil {
        err = database.Ping()
    }

    if err != nil {
        log.Panic(err)
    }

}
