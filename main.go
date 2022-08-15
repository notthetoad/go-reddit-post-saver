package main

import (
    "fmt"
    "os"
    "log"
    //"context"
    "database/sql"
    //"github.com/joho/godotenv"
    //"github.com/vartanbeno/go-reddit/v2/reddit"
    "github.com/go-sql-driver/mysql"
    "example.com/mdb"
)

var db *sql.DB

//type Post struct {
//    ID int64
//    Title string
//    Body string
//    Url string
//}

func main() {
    cfg := mysql.Config{
        User: os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net: "tcp",
        Addr: "127.0.0.1:3306",
        DBName: "post",
        AllowNativePasswords: true,
    }

    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    mdbsql := mdb.Database{
        Db: db,
    }
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("mdb: ", mdbsql.Test())

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")

    post, err1 := mdbsql.QueryPostByID(1)
    if err1 != nil {
        log.Fatal(err1)
    }
    fmt.Printf("post found: %v\n", post)

    posts, err2 := mdbsql.QueryAllPosts()
    if err2 != nil {
        log.Fatal(err2)
    }
    fmt.Printf("Query all posts: %v\n", posts)
    pmap := make(map[int]mdb.Post)
    for i, val := range posts {
        fmt.Println(i, val)
        pmap[i] = val
    }
    fmt.Println(pmap)
    for key, val := range pmap {
        fmt.Println(key, val.Title) 
    }
}
