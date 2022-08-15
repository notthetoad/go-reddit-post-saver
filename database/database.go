package mdb

import (
    "fmt"
    "database/sql"
    "github.com/go-sql-driver/mysql"
    "os"
    "log"
)

type Database struct {
    Db *sql.DB
}
//type tmp Database.Db

type Post struct {
    ID int64
    Title string
    Body string
    Url string
}

func (db *Database) InitDb() {
    var err error
    cfg := mysql.Config{
        User: os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net: "tcp",
        Addr: "127.0.0.1:3306",
        DBName: "post",
        AllowNativePasswords: true,
    }

    sqldb, err := sql.Open("mysql", cfg.FormatDSN())
    db.Db = sqldb
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
}


func (db *Database) QueryPostByID(id int64) (Post, error) {
    var pst Post
    row := db.Db.QueryRow("SELECT * FROM post WHERE id = ?", id)
    err := row.Scan(&pst.ID, &pst.Title, &pst.Body, &pst.Url); if err != nil {
        if err == sql.ErrNoRows {
            return pst, fmt.Errorf("postByID %d: no such post", id)
        }
        return pst, fmt.Errorf("postByID %d: %v", id, err)
    }
    return pst, nil
}

func (db *Database) QueryAllPosts() ([]Post, error) {
    var posts []Post
    rows, err := db.Db.Query("SELECT * FROM post;")
    if err != nil {
        return nil, fmt.Errorf("%v\n", err)
    }
    defer rows.Close()
    for rows.Next() {
        var pst Post
        if err := rows.Scan(&pst.ID, &pst.Title, &pst.Body, &pst.Url); err != nil {
            return nil, fmt.Errorf("%v\n", err)
        }
        posts = append(posts, pst)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("%v\n", err)
    }
    return posts, nil
}
