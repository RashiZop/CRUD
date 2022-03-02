package main

import(
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	_"github.com/go-sql-driver/mysql"
	"golangprog/rest_api-mysql/CRUD/handler"
)

func main(){
	
	fmt.Println("REST API-MuxRouter")
	router:=mux.NewRouter()
	
	router.HandleFunc("/employees/{id}", handler.EmpByID).Methods("GET")
	router.HandleFunc("/employees", handler.AddEmp).Methods("POST")
	router.HandleFunc("/employees/{id}", handler.Employeeupdate).Methods("PUT")
	router.HandleFunc("/employees/{id}", handler.DelEmp).Methods("DELETE")

	
	http.ListenAndServe(":8000", router)
	
}