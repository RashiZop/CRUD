package main

import(
	_ "github.com/go-sql-driver/mysql"
	"golangprog/CRUD/store"
	//"time"
	//"context"
	"database/sql"
	"log"
	"fmt"
)
func main(){

	
	db, err:=sql.Open("mysql","root:password@tcp(127.0.0.1:3306)/test")
	if err!=nil{
		panic(err.Error())
	}else{
		fmt.Println("Connected")
	}

	// //create
	// err = store.table(db)
    // if err != nil {
    //     log.Printf("Create table failed with error %s", err)
    //     return
    // }

	//inserting new elements
	_, err=store.AddEmp(store.Employee{
		Id: 5,
		Name: "Kashish Gupta",
		Email: "kg@zopsmart.com",
		Role: "intern",
	}, db)
	if err!=nil{
		log.Fatal(err)
	}

	_, err=store.AddEmp(store.Employee{
		Id: 1,
		Name: "Rashi Singh",
		Email: "rs@zopsmart.com",
		Role: "intern",
	}, db)
	if err!=nil{
		log.Fatal(err)
	}
	
	_, err=store.AddEmp(store.Employee{
		Id: 2,
		Name: "Anushi",
		Email: "an@zopsmart.com",
		Role: "intern",
	}, db)
	if err!=nil{
		log.Fatal(err)
	}

	//employee by Id
	fmt.Println(store.EmpByID(1, db))

	//delete
	store.DelEmp(5,db)

	//update
	fmt.Println(store.Employeeupdate(store.Employee{
		Id: 2,
		Name: "Anushi",
		Email: "an@zopsmart.com",
		Role: "intern",
	}, db))

	
}