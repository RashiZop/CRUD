package store

import(
	_ "github.com/go-sql-driver/mysql"
	//"time"
	//"context"
	//"log"
	"fmt"
	"errors"
	"database/sql"
	"log"

)

type Employee struct{
	Id int `json:"Id"`
	Name string `json:"Name"`
	Email string `json:"Email"`
	Role string `json:"Role`
}

//readById
func EmpByID(Id int, db *sql.DB)(*Employee, error){
	var emp Employee
	log.Print(Id)
	
	if(Id<0){
		return nil, sql.ErrNoRows
	}

	rows, err:=db.Query("Select * FROM employee1 WHERE Id = ?", Id)
	if(err!=nil){
		log.Print(err)
		return &emp, err
	}
	
	defer rows.Close()
	
	for rows.Next(){
		log.Print("1")
		
		if err:=rows.Scan(&emp.Id, &emp.Name, &emp.Email, &emp.Role); err!=nil{
			return &emp, err
		}

	}

	log.Print(emp)

	return &emp, nil
	
}

//insert
func AddEmp(emp Employee, db *sql.DB)(Employee, error){
	
	_, err:=db.Exec("INSERT INTO employee1 (Id, Name, Email, Role) VALUES (?,?,?,?)",emp.Id, emp.Name, emp.Email, emp.Role)
	if err!=nil{
		return emp, nil
	}
	return emp, nil
}

func Employeeupdate(emp *Employee, db *sql.DB) error {
	_, err := db.Exec("UPDATE employee1 SET Name = ?, Email=?, Role=? WHERE ID = ?", emp.Name, emp.Email, emp.Role, emp.Id)
	if err != nil {
	   return errors.New("update failed")
	}
	return nil
 }

//deleteById
func DelEmp(cond1 int, db *sql.DB)(error){
	
	if(cond1<0){
		return sql.ErrNoRows
	}
	rows, err:=db.Query("delete from employee1 Where Id = ?", cond1)
	if err!=nil{
		return err
	}

	defer rows.Close()

	return nil
}

/*
//create table
func table(db *sql.DB) error {  
    query := "CREATE TABLE IF NOT EXISTS employee1(Id int primary key auto_increment, Name text, Email text, Role text)"
    ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelfunc()
    res, err := db.ExecContext(ctx, query)
    if err != nil {
        log.Printf("Error %s when creating table", err)
        return err
    }
    rows, err := res.RowsAffected()
    if err != nil {
        log.Printf("Error %s when getting rows affected", err)
        return err
    }
    log.Printf("Rows affected when creating table: %d", rows)
    return nil
}

*/