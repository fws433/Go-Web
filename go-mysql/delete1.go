package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type person struct{
	UserId int`db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string  `db:"email"`
}
var db *sqlx.DB
func init(){
	database,err:=sqlx.Open("mysql","root:fws666666@tcp(127.0.0.1:3306)/fws")
	if err!=nil{
		fmt.Println("open mysql failed,",err)
		return
	}
	db=database
}
func main(){
	res,err:=db.Exec("delete from person where user_id=?",2)
	if err!=nil{
		fmt.Println("exec failed,",err)
		return
	}
	row,err:=res.RowsAffected()
	if err!=nil{
		fmt.Println("rows failed,",err)

	}
	fmt.Println("delete succ:",row)
}
