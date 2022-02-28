package conn_sql
//package main
//hgchg
import(
	_ "github.com/go-sql-driver/mysql"
	//"time"
	//"context"
	"database/sql"
	//"log"
	//"fmt"
	"errors"
)

type Employee struct{
	Id int64
	Name string
	Email string
	Role string
}

//readById
func empByID(Id int64, db *sql.DB)(Employee, error){
	var emp Employee
	
	if(Id<0){
		return emp, sql.ErrNoRows
	}

	rows, err:=db.Query("Select * FROM employee1 WHERE Id = ?", Id)
	if(err!=nil){
		return emp, err
	}
	
	defer rows.Close()
	
	for rows.Next(){
		
		if err:=rows.Scan(&emp.Id, &emp.Name, &emp.Email, &emp.Role); err!=nil{
			return emp, err
		}

	}
	return emp, nil
	
}

//insert
func addEmp(emp Employee, db *sql.DB)(Employee, error){
	
	_, err:=db.Exec("INSERT INTO employee1 (Id, Name, Email, Role) VALUES (?,?,?,?)",emp.Id, emp.Name, emp.Email, emp.Role)
	if err!=nil{
		return emp, nil
	}
	return emp, nil
}

func Employeeupdate(emp Employee, db *sql.DB) error {
	_, err := db.Exec("UPDATE employee SET Name = ?, Email=?, Role=? WHERE ID = ?",
	   &emp.Name, &emp.Email, &emp.Role, &emp.Id)
	if err != nil {
	   return errors.New("update failed")
	}
	return nil
 }

//deleteById
func delEmp(cond1 int64, db *sql.DB)(error){
	
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

func main(){

	
	db, err:=sql.Open("mysql","root:password@tcp(127.0.0.1:3306)/test")
	if err!=nil{
		panic(err.Error())
	}else{
		fmt.Println("Connected")
	}

	//create
	err = table(db)
    if err != nil {
        log.Printf("Create table failed with error %s", err)
        return
    }

	//inserting new elements
	_, err=addEmp(Employee{
		Id: 5,
		Name: "Kashish Gupta",
		Email: "kg@zopsmart.com",
		Role: "intern",
	}, db)
	if err!=nil{
		log.Fatal(err)
	}

	_, err=addEmp(Employee{
		Id: 1,
		Name: "Rashi Singh",
		Email: "rs@zopsmart.com",
		Role: "intern",
	}, db)
	if err!=nil{
		log.Fatal(err)
	}
	
	_, err=addEmp(Employee{
		Id: 2,
		Name: "Anushi",
		Email: "an@zopsmart.com",
		Role: "intern",
	}, db)
	if err!=nil{
		log.Fatal(err)
	}

	//employee by Id
	fmt.Println(empByID(1, db))

	//delete
	delEmp(5,db)

	//update
	fmt.Println(updateEmp("rs2@zopsmart.com", 1, db))

	
}
*/