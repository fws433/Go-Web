package main
import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type person1 struct{
	UserId int`db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string  `db:"email"`
}
var db1 *sqlx.DB
func init(){
	database,err:=sqlx.Open("mysql","root:fws666666@tcp(127.0.0.1:3306)/fws")
	if err!=nil{
		fmt.Println("open mysql failed,",err)
		return
	}
	db1=database
}
func main(){
	conn,err:=db1.Begin()
	if err!=nil{
		fmt.Println("begin failed:",err)
		return
	}
	r, err := conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ:", id)

	r, err = conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	id, err = r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ:", id)

	conn.Commit()
}


