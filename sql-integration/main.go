package main

import (
	"fmt"
	"sql-integration/model"
	"sql-integration/repository"
)

func main(){
	db:=repository.DB()
	var users []model.User
	db.Raw("Select * from users").Scan(&users)
	// if err!=nil{
	// 	fmt.Println(err)
	// }
	// defer rows.Close()
	for val,index := range users {
		fmt.Println(val)
		fmt.Println(index)
	}
	
	
	fmt.Println(users)
}