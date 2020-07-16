package main

import (
	//"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person1 struct{
	UserId int `db:"user_id"`
	Username string `db:"username"`
	Sex string `db;"sex"`
	Email string `db:"email"`
}
var Db1 *sqlx.DB
func init(){
	database,err:=sqlx.Open("mysql","root:fws666666@tcp(127.0.0.1:3306)/fws")
	if err!=nil{
		fmt.Println("open mysql failed,",err)
		return
	}
	Db1=database
}
func main(){
 var person []Person1
 err:=Db1.Select(&person,"select user_id,username,sex,email from person where user_id=?",2)
 if err!=nil{
 	fmt.Println("exec failed,",err)
	 return
 }
 fmt.Println("select succ:",person)
}
