package mysql

import (
	"heaven/storage"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"fmt"
	"strings"
	"github.com/jianfengye/collection"
)

const DB_TABLE  = "s_user_miner"
var fields = []string{"app_id", "client_ip", "client_time", "server_time", "sdk_version", "resolution", "os",
	"page", "event_id", "arg1", "arg2", "arg3", "args"}

type mysqlStorage struct {
	isOpened bool
	db *sql.DB
	insertStmt string
}

func NewStorage() storage.Storage  {
	db, err := sql.Open("mysql", "root:12345678@tcp(localhost:3306)/miner?charset=utf8")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	processedFields, _ := collection.NewStrCollection(fields).Map(func(item interface{}, key int) interface{} {
		return item.(string) + "=?"
	}).ToStrings()
	q := fmt.Sprintf("INSERT s_user_miner SET %s", strings.Join(processedFields, ", "))
	log.Println("insert query is :", q)

	return &mysqlStorage{
		true,
		db,
		q,
	}
}

func (st *mysqlStorage) Open()  {
	if ! st.isOpened {
		// TODO open db
	}
}


func (st *mysqlStorage) Write(record map[string]string)  {
	stmt, err := st.db.Prepare(st.insertStmt)
	if err != nil {
		log.Println(err)
		return
	}
	if len(record) != len(fields) {
		log.Println("Record does not match with filds")
		return
	}
	vals := make([]interface{}, len(fields))
	for idx, v := range fields {
		vals[idx] = record[v]
	}
	_, err = stmt.Exec(vals...)
	if err != nil {
		log.Println("Insert failed", err)
		return
	}
}