package golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	script := "INSERT INTO customer(id,name) VALUES ('joko','JOKO')"
	_, error := db.ExecContext(ctx, script)
	if error != nil {
		panic(error)
	}
	fmt.Println("SUCCESS INSERT DATA")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	script := "SELECT id,name  FROM customer"
	rows, error := db.QueryContext(ctx, script)
	if error != nil {
		panic(error)
	}
	defer rows.Close()
	fmt.Println("SUCCESS INSERT DATA")

	for rows.Next() {
		var id, name string
		error := rows.Scan(&id, &name)
		if error != nil {
			panic(error)
		}
		fmt.Println("Id", id)
		fmt.Println("Name", name)
	}
}
func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	script := "SELECT id,name,email,balance,rating,birth_date,married,created_at  FROM customer"
	rows, error := db.QueryContext(ctx, script)
	if error != nil {
		panic(error)
	}
	defer rows.Close()
	fmt.Println("SUCCESS INSERT DATA")

	for rows.Next() {
		var id, name, email string
		var balance int32
		var rating float64
		var birthDate, createdAt time.Time
		var married bool
		error := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if error != nil {
			panic(error)
		}
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		fmt.Println("Email:", email)
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		fmt.Println("BirthDate:", birthDate)
		fmt.Println("Married:", married)
		fmt.Println("CreatedAt:", createdAt)
	}
}
func TestQuerySqlComplexNull(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	script := "SELECT id,name,email,balance,rating,birth_date,married,created_at  FROM customer"
	rows, error := db.QueryContext(ctx, script)
	if error != nil {
		panic(error)
	}
	defer rows.Close()
	fmt.Println("SUCCESS INSERT DATA")

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var createdAt time.Time
		var birthDate sql.NullTime
		var married bool
		error := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if error != nil {
			panic(error)
		}
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("Email:", email.String)
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		if birthDate.Valid {
			fmt.Println("BirthDate:", birthDate.Time)
		}
		fmt.Println("Married:", married)
		fmt.Println("CreatedAt:", createdAt)
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	username := "admin'; #"
	password := "salah"
	script := "SELECT username  FROM user WHERE username = '" + username + "' AND password='" + password + "' LIMIT 1"
	fmt.Println(script)
	rows, error := db.QueryContext(ctx, script)
	if error != nil {
		panic(error)
	}
	defer rows.Close()
	if rows.Next() {
		var username string
		error := rows.Scan(&username)
		if error != nil {
			panic(error)
		}
		fmt.Println("SUCCESS LOGIN")
	} else {
		fmt.Println("LOGIN FAILED")
	}
}
func TestSqlInjectionSave(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	username := "admin'; #"
	password := "salah"
	script := "SELECT username  FROM user WHERE username = ? AND password= ? LIMIT 1"
	fmt.Println(script)
	rows, error := db.QueryContext(ctx, script, username, password)
	if error != nil {
		panic(error)
	}
	defer rows.Close()
	if rows.Next() {
		var username string
		error := rows.Scan(&username)
		if error != nil {
			panic(error)
		}
		fmt.Println("SUCCESS LOGIN")
	} else {
		fmt.Println("LOGIN FAILED")
	}
}

func TestExecSqlParamter(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	username := "eko"
	password := "eko"
	script := "INSERT INTO user(username,password) VALUES (?,?)"
	_, error := db.ExecContext(ctx, script, username, password)
	if error != nil {
		panic(error)
	}
	fmt.Println("SUCCESS INSERT DATA")
}

func TestLastInsertID(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	email := "eko@gamil.com"
	comment := "test comments"
	script := "INSERT INTO comments(email,comment	) VALUES (?,?)"
	result, error := db.ExecContext(ctx, script, email, comment)
	if error != nil {
		panic(error)
	}
	insertID, error := result.LastInsertId()
	if error != nil {
		panic(error)
	}

	fmt.Println("SUCCESS INSERT DATA new comment", insertID)

}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	statement, error := db.PrepareContext(ctx, "INSERT INTO comments(email,comment	) VALUES (?,?)")
	if error != nil {
		panic(error)
	}
	for i := 0; i <= 10; i++ {
		email := "eko" + strconv.Itoa(i) + "@gamil.com"
		comment := "test comments ke" + strconv.Itoa(i)
		result, error := statement.ExecContext(ctx, email, comment)
		if error != nil {
			panic(error)
		}
		lastInsertID, _ := result.LastInsertId()
		fmt.Println("INSERT COMMENT", lastInsertID)
	}
	defer statement.Close()
}

func TestDatabaseTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, error := db.Begin()
	if error != nil {
		panic(error)
	}
	script := "INSERT INTO comments(email,comment) VALUES (?,?)"
	for i := 0; i <= 10; i++ {
		email := "eko" + strconv.Itoa(i) + "@gamil.com"
		comment := "test comments ke" + strconv.Itoa(i)
		result, error := tx.ExecContext(ctx, script, email, comment)
		if error != nil {
			panic(error)
		}
		lastInsertID, _ := result.LastInsertId()
		fmt.Println("INSERT COMMENT", lastInsertID)
	}
	//do transaction

	//err := tx.Commit()
	err := tx.Rollback()

	if err != nil {
		panic(error)
	}
}
