package main
import (
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "fmt"
    //"time"
)
func main() {
    db, err := sql.Open("mysql", "root:123456@/test?charset=utf8") //format:[user:password@tcp(localhost:5555)/dbname?charset=utf8]
    //db, err := sql.Open("mysql", "mysql@/test?charset=utf8")
    checkErr(err)
    stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
    checkErr(err)
    res, err := stmt.Exec("astaxie", "dev-department", "2012-12-09")
    checkErr(err)
    id, err := res.LastInsertId()
    checkErr(err)
    fmt.Println(id)

    //update data
    stmt, err = db.Prepare("update userinfo set username=? where uid=?")
    checkErr(err)
    res, err = stmt.Exec("astaxieupdate", id)
    checkErr(err)
    affect, err := res.RowsAffected()
    checkErr(err)
    fmt.Println(affect)

    //query data
    rows, err := db.Query("SELECT * FROM userinfo")
    checkErr(err)
    for rows.Next() {
        var uid int
        var username string
        var department string
        var created string
        err = rows.Scan(&uid, &username, &department, &created)
        checkErr(err)
        fmt.Println(uid)
        fmt.Println(username)
        fmt.Println(department)
        fmt.Println(created)
    }

    //delete data
    stmt, err = db.Prepare("delete from userinfo where uid=?")
    checkErr(err)
    res, err = stmt.Exec(id)
    checkErr(err)
    affect, err = res.RowsAffected()
    checkErr(err)
    fmt.Println(affect)
    db.Close()
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
