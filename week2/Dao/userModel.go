package Dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var DB *sql.DB

func InitDb() error {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/test"
	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		return err
	}

	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	return nil
}

type User struct {
	Id   int64          `db:"id"`
	Name sql.NullString `db:"string"`
	Age  int            `db:"age"`
}

func IsertData()(int64,error) {
	result, err := DB.Exec("insert into user(name, age) values(?, ?)", "tom", 18)
	if err != nil {
		return 0,err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0,err
	}
	return id,nil
}

func QueryData(id int) (User,error) {
		var user User
		row := DB.QueryRow( "select id, name, age from user where id=?", id)
		err := row.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			if err == sql.ErrNoRows {
				return user,errors.Wrap(err,"scan no rows")
			} else {
				return user,err
			}
		}

		return user,nil
}
