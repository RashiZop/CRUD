package conn_sql

import(
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"reflect"
	//"fmt"
	errored "errors"
	"database/sql"
)

//test addEmp
func TestAddEmp(t *testing.T) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))	
	if err != nil {
	   t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
 
	defer db.Close()
 
	_ = mock.ExpectExec("INSERT INTO employee1 (Id, Name, Email, Role) VALUES (?,?,?,?)").WithArgs(1, "Rashi Singh", "rs@zopsmart.com", "intern").WillReturnResult(sqlmock.NewResult(1, 1))
	_ = mock.ExpectExec("INSERT INTO employee1 (Id, Name, Email, Role) VALUES (?,?,?,?)").WithArgs(2, "Anushi", "an@zopsmart.com","intern").WillReturnResult(sqlmock.NewResult(1, 1))

	testCases := []struct {
	   Id          int64
	   empOut      Employee
	   mockQuery   interface{}
	   expectError error
	   desc string
	}{
	   //Success case
	   {
		  Id:          1,
		  empOut:      Employee{1, "Rashi Singh", "rs@zopsmart.com", "intern"},
		  mockQuery:   mock.ExpectQuery("INSERT INTO employee1 (Id, Name, Email, Role) VALUES (?,?,?,?)").WithArgs(1, "Rashi Singh", "rs@zopsmart.com", "intern").WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Email", "Role"}).AddRow(1, "Rashi Singh", "rs@zopsmart.com", "intern")),
		  expectError: nil,
		  desc: 	   "success case",
		},
	   // Success case
	   {
		  Id:          2,
		  empOut:      Employee{2, "Anushi", "an@zopsmart.com", "intern"},
		  mockQuery:   mock.ExpectQuery("INSERT INTO employee1 (Id, Name, Email, Role) VALUES (?,?,?,?)").WithArgs(2, "Anushi", "an@zopsmart.com","intern").WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Email", "Role"}).AddRow(2, "Anushi", "an@zopsmart.com","intern")),
		  expectError: nil,
		  desc: 	   "success case",
		},
		// error case
		{
			Id:          -1,
			empOut:      Employee{-1, "Anushi", "an@zopsmart.com", "intern"},
			mockQuery:   mock.ExpectQuery("INSERT INTO employee1 (Id, Name, Email, Role) VALUES (?,?,?,?)").WithArgs(-1, "Anushi", "an@zopsmart.com","intern").WillReturnError(sql.ErrNoRows),
			expectError: nil,
			desc: 	   "error case",
		  },
	}
 
	for _, testCase := range testCases {
	   t.Run(testCase.desc, func(t *testing.T) {
		  _, err := addEmp(Employee{testCase.Id, testCase.empOut.Name, testCase.empOut.Email, testCase.empOut.Role}, db)
		  
		  if !reflect.DeepEqual(err, testCase.expectError) {
			 t.Errorf("expected: %v, got: %v", testCase.expectError, err)
		  }
	   })
	}
 }
 
//test empByID
func TestEmpByID(t *testing.T) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))	
	if err != nil {
	   t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
 
	defer db.Close()
 
	_ = sqlmock.NewRows([]string{"Id", "Name", "Email", "Role"}).AddRow(1, "Rashi Singh", "rs@zopsmart.com", "intern").AddRow(2, "Anushi", "an@zopsmart.com","intern")
 
	testCases := []struct {
	   Id          int64
	   empOut      *Employee
	   mockQuery   interface{}
	   expectError error
	   desc string
	}{
	   // Success case
	   {
		  Id:          1,
		  empOut:      &Employee{1, "Rashi Singh", "rs@zopsmart.com", "intern"},
		  mockQuery:   mock.ExpectQuery("Select * FROM employee1 WHERE Id = ?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Email", "Role"}).AddRow(1, "Rashi Singh", "rs@zopsmart.com", "intern")),
		  expectError: nil,
		  desc: 	   "success case",
		},
	   // Success case
	   {
		  Id:          2,
		  empOut:      &Employee{2, "Anushi", "an@zopsmart.com", "intern"},
		  mockQuery:   mock.ExpectQuery("Select * FROM employee1 WHERE Id = ?").WithArgs(2).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Email", "Role"}).AddRow(2, "Anushi", "an@zopsmart.com","intern")),
		  expectError: nil,
		  desc: 	   "success case",
		},
		// Error case
		{
			Id:          3,
			empOut:      nil,
			mockQuery:   mock.ExpectQuery("Select * FROM employee1 WHERE Id = ?").WithArgs(3).WillReturnError(sql.ErrNoRows),
			expectError: sql.ErrNoRows,
			desc: 	   "failure case",
		  },	
		  {
			Id:          -1,
			empOut:      nil,
			mockQuery:   mock.ExpectQuery("Select * FROM employee1 WHERE Id = ?").WithArgs(-1).WillReturnError(sql.ErrNoRows),
			expectError: sql.ErrNoRows,
			desc: 	   "failure case",
		  },
		  	 
	}
 
	for _, testCase := range testCases {
	   t.Run(testCase.desc, func(t *testing.T) {
		  _, err := empByID(testCase.Id, db)
	
		  if !reflect.DeepEqual(err, testCase.expectError) {
			 t.Errorf("expected: %v, got: %v", testCase.expectError, err)
		  }
	   })
	}
 }
 
/*
 //test create
 func TestTable(t *testing.T) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))	
	if err != nil {
	   t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
 
	defer db.Close()
 
	_ = sqlmock.NewRows([]string{"Id", "Name", "Email", "Role"}).AddRow(1, "Rashi Singh", "rs@zopsmart.com", "intern").AddRow(2, "Anushi", "an@zopsmart.com","intern")

	testCases := []struct {
	   Id          int64
	   empOut      Employee
	   mockQuery   interface{}
	   expectError error
	   desc string
	}{
	   // Success case
	   {
		  Id:          1,
		  empOut:      Employee{1, "Rashi Singh", "rs@zopsmart.com", "intern"},
		  mockQuery:   mock.ExpectExec("CREATE TABLE IF NOT EXISTS employee1(Id int primary key auto_increment, Name text, Email text, Role text)").WillReturnResult(sqlmock.NewResult(1, 0)),
		  expectError: nil,
		  desc: 	   "success case",
		},
	   // Success case
	   {
		  Id:          2,
		  empOut:      Employee{2, "Anushi", "an@zopsmart.com", "intern"},
		  mockQuery:   mock.ExpectExec("CREATE TABLE IF NOT EXISTS employee1(Id int primary key auto_increment, Name text, Email text, Role text)").WillReturnResult(sqlmock.NewResult(1, 1)),
		  expectError: nil,
		  desc: 	   "success case",
		},
		// Error case
		{
			Id:          -1,
			empOut:      Employee{-1, "", "", ""},
			mockQuery:   mock.ExpectExec("CREATE TABLE IF NOT EXISTS employee1(Id int primary key auto_increment, Name text, Email text, Role text)").WillReturnError(sql.ErrNoRows),
			expectError: sql.ErrNoRows,
			desc: 	   "failure case",
		  },
	}

	query:="CREATE TABLE IF NOT EXISTS employee1(Id int primary key auto_increment, Name text, Email text, Role text)"
	prep:=mock.ExpectPrepare(query)

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			prep.ExpectExec().WillReturnResult(sqlmock.NewResult(0, 1))

			err = table(db)
			if !reflect.DeepEqual(err, testCase.expectError) {
				t.Errorf("Create() error = %v, wantErr %v", err, testCase.expectError)
				return
			 }
		   
		})
	 }
 }
*/

 //test delete
 func TestDelEmp(t *testing.T){
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))	
	if err != nil {
	   t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	_ = sqlmock.NewRows([]string{"Id", "Name", "Email", "Role"}).AddRow(1, "Rashi Singh", "rs@zopsmart.com", "intern").AddRow(2, "Anushi", "an@zopsmart.com","intern")
 
	testCases := []struct {
	   Id          int64
	   empOut      Employee
	   mockQuery   interface{}
	   expectError error
	   desc string
	}{
	   // Success case
	   {
		  Id:          1,
		  empOut:      Employee{1, "Rashi Singh", "rs@zopsmart.com", "intern"},
		  mockQuery:   mock.ExpectQuery("delete from employee1 Where Id = ?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Email", "Role"}).AddRow(1, "Rashi Singh", "rs@zopsmart.com", "intern")),
		  expectError: nil,
		  desc: 	   "success case",
		},
	   // Success case
	   {
		  Id:          2,
		  empOut:      Employee{2, "Anushi", "an@zopsmart.com", "intern"},
		  mockQuery:   mock.ExpectQuery("delete from employee1 Where Id = ?").WithArgs(2).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Email", "Role"}).AddRow(2, "Anushi", "an@zopsmart.com","intern")),
		  expectError: nil,
		  desc: 	   "success case",
		},
		// Error case
		{
			Id:          3,
			empOut:      Employee{3, "", "", ""},
			mockQuery:   mock.ExpectQuery("delete from employee1 Where Id = ?").WithArgs(3).WillReturnError(sql.ErrNoRows),
			expectError: sql.ErrNoRows,
			desc: 	   "failure case",
		  },
		  // Error case
		{
			Id:          -1,
			empOut:      Employee{-1, "", "", ""},
			mockQuery:   mock.ExpectQuery("delete from employee1 Where Id = ?").WithArgs(-1).WillReturnError(sql.ErrNoRows),
			expectError: sql.ErrNoRows,
			desc: 	   "failure case",
		  },
	}



	query := "delete from employee1 Where Id = ?"
	prep := mock.ExpectPrepare(query)

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			prep.ExpectExec().WithArgs(testCase.Id).WillReturnResult(sqlmock.NewResult(0, 1))

			err = delEmp(testCase.Id,db)
			if !reflect.DeepEqual(err, testCase.expectError) {
				t.Errorf("delete() error = %v, wantErr %v", err, testCase.expectError)
				return
			 }
		   
		})
	 }



 }
 

 func TestEmployeeupdate(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
	   t.Error(err)
	}
	updaterr := errored.New("update failed")
 
	tests := []struct {
	   desc      string
	   expecterr error
	   input_emp Employee
	   mockCall  *sqlmock.ExpectedExec
	}{
	   {
		  desc:      "update succes",
		  expecterr: nil,
		  input_emp: Employee{1, "Rashi Singh", "rs@gmail.com", "Intern"},
		  mockCall:  mock.ExpectExec("UPDATE employee SET Name = ?, Email=?, Role=? WHERE ID = ?").WithArgs("Rashi Singh", "rs@gmail.com", "Intern", 1).WillReturnResult(sqlmock.NewResult(1, 1)),
	   },
	   {
		  desc:      "update fail",
		  expecterr: updaterr,
		  input_emp: Employee{3, "", "", ""},
		  mockCall:  mock.ExpectExec("UPDATE employee SET Name=?,Email=?,Role=? WHERE ID = ?").WithArgs("", "", "", 3).WillReturnError(updaterr),
	   },
	}
	for _, tc := range tests {
	   err := Employeeupdate(tc.input_emp, db)
 
	   if !reflect.DeepEqual(err, tc.expecterr) {
		  t.Errorf("Expected: %v, Got: %v", tc.expecterr, err)
	   }
 
	}
 
 }