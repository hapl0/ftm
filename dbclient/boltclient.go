package dbclient

import "github.com/asdine/storm"

db, err := storm.Open("my.db")
defer db.Close()

