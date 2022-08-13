package mdb

import (
    "fmt"
    "database/sql"
    //"github.com/go-sql-driver/mysql"
)

type Post struct {
    ID      int64
    Title   string
    Body    string
    Url     string
}

//func (db *sql.DB) PostByID(id int64) (Post, error) {
//    var pst Post
//    row := db.Query("SELECT * FROM post WHERE id = ?", id)
//    err := row.Scan(&pst.ID, &pst.Title, &pst.Body, &pst.Url); if err != nil {
//        if err == sql.ErrNoRows {
//            return pst, fmt.Errorf("postByID %d: no such post", id)
//        }
//        return pst, fmt.Errorf("postByID %d: %v", id, err)
//    }
//    return pst, nil
//}

func Test() string {
    return "Hello World"
}

func QueryPostByID(id int64, db *sql.DB) (Post, error) {
    var pst Post
    row := db.QueryRow("SELECT * FROM post WHERE id = ?", id)
    err := row.Scan(&pst.ID, &pst.Title, &pst.Body, &pst.Url); if err != nil {
        if err == sql.ErrNoRows {
            return pst, fmt.Errorf("postByID %d: no such post", id)
        }
        return pst, fmt.Errorf("postByID %d: %v", id, err)
    }
    return pst, nil
}
