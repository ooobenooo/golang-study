package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // 加载package,但在代码中没有显示调用
)

func main() {
	// 创建数据库链接
	db, err := sql.Open("mysql", "demo:123456@tcp(localhost:3306)/demo?charset=utf8")
	checkErr(err)
	defer db.Close()

	// 预编译
	stmt, err := db.Prepare("INSERT INTO T_USER (NAME, AGE) values (?, ?)")
	checkErr(err)

	// 传入参数
	res, err := stmt.Exec("ben", 38)
	checkErr(err)

	// 获取自动递增的ID
	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("id:", id)

	stmt, err = db.Prepare("UPDATE T_USER SET AGE = ? WHERE ID = ?")
	checkErr(err)
	res, err = stmt.Exec(40, id)
	checkErr(err)

	rowCount, err := res.RowsAffected() // 影响行数
	checkErr(err)
	fmt.Println("row affacted: ", rowCount)

	stmt, err = db.Prepare("SELECT * FROM T_USER")
	checkErr(err)
	rows, err := stmt.Query()
	checkErr(err)

	columns, _ := rows.Columns()
	for i, v := range columns {
		fmt.Printf("%d: %s\n", i, v)
	}

	for rows.Next() {
		var id int
		var name string
		var age int
		// scan是通过变量的指针作为参数传入，修改指针的内存数据从而达到赋值操作
		err = rows.Scan(&id, &name, &age)
		checkErr(err)
		fmt.Println("id: ", id)
		fmt.Println("name: ", name)
		fmt.Println("age: ", age)
	}

	stmt, err = db.Prepare("delete from T_USER where ID = ?")
	checkErr(err)
	res, err = stmt.Exec(id)
	checkErr(err)
	rowCount, err = res.RowsAffected()
	checkErr(err)
	fmt.Println("delete rows: ", rowCount)

	fmt.Println("OK")
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
