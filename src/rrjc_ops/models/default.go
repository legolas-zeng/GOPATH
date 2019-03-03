package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type Message struct {
	Msg		string
	Status  int
}