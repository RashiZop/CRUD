package handler

import(
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	_"github.com/go-sql-driver/mysql"
	"strconv"
	"golangprog/rest_api-mysql/store"
	"golangprog/rest_api-mysql/driver"
)

//get
func EmpByID(w http.ResponseWriter, r *http.Request){
    
	db:=driver.ConnectToSQL()
	defer db.Close()

	params:=mux.Vars(r)
	id, _:=strconv.Atoi(params["id"])
	log.Print(id)
	use, err:=store.EmpByID(id, db)
	if(err!=nil){
		w.Write([]byte("Employee doesn't exist"))
	}else{
		res, _:=json.Marshal(use)
		w.Write(res)
	}
}

//post
func AddEmp(w http.ResponseWriter, r *http.Request) {  
	
	db:=driver.ConnectToSQL()
	defer db.Close()

	var emp store.Employee
	emp.Id, _=strconv.Atoi(r.PostFormValue("Id"))
	emp.Name=r.PostFormValue("Name")
	emp.Email=r.PostFormValue("Email")
	emp.Role=r.PostFormValue("Role")

	_,err:=store.AddEmp(emp, db)
	if(err!=nil){
		w.Write([]byte("Failed to add new emp"))
	}else{
		w.Write([]byte("Successfully added"))
	}
}

//put
func Employeeupdate(w http.ResponseWriter, r *http.Request) {

	db:=driver.ConnectToSQL()
	defer db.Close()
	
	params := mux.Vars(r)  
	id, _:=strconv.Atoi(params["id"])
	
	use, err:=store.EmpByID(id, db)
	use.Name=r.PostFormValue("Name")
	use.Email=r.PostFormValue("Email")
	use.Role=r.PostFormValue("Role")

	err=store.Employeeupdate(use, db)
	if(err!=nil){
		w.Write([]byte("Failed to update emp"))
	}else{
		w.Write([]byte("Successfully updated"))
	}
  }


//delete
func DelEmp(w http.ResponseWriter, r *http.Request) {

	db:=driver.ConnectToSQL()
	defer db.Close()

	params := mux.Vars(r)  
	id,_:=strconv.Atoi(params["id"])
	err:=store.DelEmp(id, db)
 	if err != nil {
		w.Write([]byte("Failed to delete emp"))
	}else{
		w.Write([]byte("Successfully deleted"))
	}

}
