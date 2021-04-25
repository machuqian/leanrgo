package main

import (
	"database/sql"
	"fmt"
	"geektime/week2/Dao"
	"github.com/pkg/errors"

)



func main(){
	err := Dao.InitDb()


	if err != nil {
		fmt.Printf("init db failed, err:%v\n", err)
		return
	}

	var user Dao.User
	user,err= Dao.QueryData(2)

	if err!=nil{
		if(sql.ErrNoRows==errors.Cause(err)){
			//
			//fmt.Println(err)
		}else{
			fmt.Println(err)
		}
		return
	}

	fmt.Println(user)


}
