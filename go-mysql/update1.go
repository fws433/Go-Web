package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person3 struct{
	UserId int `db;"user_id"`
	Username string `db:"username"`
	Sex string `db:"sex"`
	Email string `db:"email"`

}
var Db3 *sqlx.DB
func init(){
	database,err:=sqlx.Open("mysql","root:fws666666@tcp(127.0.0.1:3306)/fws")
	if err!=nil{
		fmt.Println("open mysql failed,",err)
		return
	}
	Db3=database
}
func main(){
	res,err:=Db3.Exec("update person set username=? where user_id=?","stu003",2)
	if err!=nil{
		fmt.Println("exec failed,",err)
		return
	}
	row,err:=res.RowsAffected()
	if err!=nil{
		fmt.Println("rows failed,",err)

	}
	fmt.Println("update succ;",row)
}
