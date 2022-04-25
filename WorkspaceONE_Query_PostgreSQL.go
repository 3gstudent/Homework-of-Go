package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	_ "github.com/lib/pq"
)


func connectDB() *sql.DB{
	fmt.Println("[*] Get password of the db")
	b, err := ioutil.ReadFile("/usr/local/horizon/conf/db.pwd")
    if err != nil {
        fmt.Print(err)
    }

	password := string(b)
	fmt.Println("[+] " + password)

	var host     = "localhost"
	var port int = 5432
	var user     = "horizon"
	var dbname   = "saas"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	fmt.Println("[*] psqlInfo:" + psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("[+] Successfully connected!")
	return db
}


func queryUser(db *sql.DB){
	fmt.Println("[*] Querying User")
	var strUsername,strEmail,idUser,createdDate,domain string
	var searchStr     = `SELECT "strUsername","strEmail","idUser","createdDate","domain" FROM "Users"`

	fmt.Println("    " + searchStr)
	rows,err:=db.Query(searchStr)
	if err!= nil{
		panic(err)
	}
	defer rows.Close()
	for rows.Next(){
		err:= rows.Scan(&strUsername,&strEmail,&idUser,&createdDate,&domain)
		if err!= nil{
			//fmt.Println(err)
		}
		fmt.Println(" - idUser     : " + idUser)		
		fmt.Println("   strUsername: " + strUsername)
		fmt.Println("   strEmail   : " + strEmail)
		fmt.Println("   createdDate: " + createdDate)
		fmt.Println("   domain     : " + domain)
	}
	err = rows.Err()
	if err!= nil{
		panic(err)
	}
}


func main()  {
	db:=connectDB()
	queryUser(db)
}