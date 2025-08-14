package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

/**
题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
    要求 ：编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
CREATE TABLE employees (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    department VARCHAR(50) NOT NULL,
    salary DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

题目2：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
    要求 ： 定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
CREATE TABLE books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    author VARCHAR(50) NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);
*/

type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

type Book struct {
	ID     int    `db:"id"`
	Title  string `db:"title"`
	Author string `db:"author"`
	Price  string `db:"price"`
}

func getEmployeesByDepartment(db *sqlx.DB, dept string) ([]Employee, error) {
	var employees []Employee
	err := db.Select(&employees,
		"SELECT id, name, department, salary FROM employees WHERE department = ?",
		dept)
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func getMaxSalaryEmployee(db *sqlx.DB) ([]Employee, error) {
	var employees []Employee
	err := db.Select(&employees,
		"SELECT id, name, department, salary FROM employees where salary = (select max(salary) from employees)")
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func getExpensiveBooks(db *sqlx.DB, minPrice float64) ([]Book, error) {
	var books []Book
	query := `
		SELECT id, title, author, price
		FROM books
		WHERE price > ?
		ORDER BY price DESC
	`
	err := db.Select(&books, query, minPrice)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func main() {
	// 连接数据库
	db, err := sqlx.Connect("mysql", "root:root@tcp(127.0.0.1:3306)/go")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//1、查询技术部员工
	techEmps, err := getEmployeesByDepartment(db, "技术部")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("技术部员工:")
	for _, emp := range techEmps {
		fmt.Printf("%d: %s, 薪资: %.2f\n", emp.ID, emp.Name, emp.Salary)
	}
	//2、查询薪资最多的员工
	techEmpss, err := getMaxSalaryEmployee(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("薪资最多的员工:")
	for _, emp := range techEmpss {
		fmt.Printf("%d: %s\n", emp.ID, emp.Name, emp.Salary)
	}

	expensiveBooks, err := getExpensiveBooks(db, 50.0)
	if err != nil {
		log.Fatalf("查询书籍失败: %v", err)
	}
	fmt.Println("价格大于50元的书籍:")
	for _, book := range expensiveBooks {
		fmt.Printf("ID: %d, 书名: %s, 作者: %s, 价格: ￥%.2f\n",
			book.ID, book.Title, book.Author, book.Price)
	}
}
