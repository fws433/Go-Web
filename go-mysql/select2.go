package main
import (
	//"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person2 struct{
	UserId int `db:"user_id"`
	Username string `db:"username"`
	Sex string `db;"sex"`
	Email string `db:"email"`
}
var Db2 *sqlx.DB
func init(){
	database,err:=sqlx.Open("mysql","root:fws666666@tcp(127.0.0.1:3306)/fws")
	if err!=nil{
		fmt.Println("open mysql failed,",err)
		return
	}
	Db2=database
}
func main(){
	var person1 []Person2
	//err:=Db2.Select(&person,"select user_id,username,sex,email from person where user_id=?",2)
	//查询所有 的记录
	rows,err:=Db2.Query("select user_id,username,sex,email from person")
	if err!=nil{
		fmt.Println("exec failed,",err)
		return
	}
	for rows.Next(){
		person:=Person2{}
		err:=rows.Scan(&person.UserId,&person.Username,&person.Sex,&person.Email)
		if err!=nil{
			log.Fatal(err)
		}
		person1=append(person1,person)
	}
	//fmt.Println("select succ:",person1)
	rows.Close()
	fmt.Println("select success:",person1)
}
