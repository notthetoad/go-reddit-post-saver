package mdb

import (
    "fmt"
    "os"
    "log"
    "database/sql"
    "github.com/go-sql-driver/mysql"
    "github.com/vartanbeno/go-reddit/v2/reddit"
)

type Database struct {
    Db *sql.DB
}

func (db *Database) InitDb() {
    var err error
    cfg := mysql.Config{
        User: os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net: "tcp",
        Addr: "127.0.0.1:3306",
        DBName: "goreddit",
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


//func (db *Database) QueryPostByID(id int64) (Post, error) {
//    var pst Post
//    row := db.Db.QueryRow("SELECT * FROM post WHERE id = ?", id)
//    err := row.Scan(&pst.ID, &pst.Title, &pst.Body, &pst.Url); if err != nil {
//        if err == sql.ErrNoRows {
//            return pst, fmt.Errorf("postByID %d: no such post", id)
//        }
//        return pst, fmt.Errorf("postByID %d: %v", id, err)
//    }
//    return pst, nil
//}
//
//func (db *Database) QueryAllPosts() ([]Post, error) {
//    var posts []Post
//    rows, err := db.Db.Query("SELECT * FROM post;")
//    if err != nil {
//        return nil, fmt.Errorf("%v\n", err)
//    }
//    defer rows.Close()
//    for rows.Next() {
//        var pst Post
//        if err := rows.Scan(&pst.ID, &pst.Title, &pst.Body, &pst.Url); err != nil {
//            return nil, fmt.Errorf("%v\n", err)
//        }
//        posts = append(posts, pst)
//    }
//    if err := rows.Err(); err != nil {
//        return nil, fmt.Errorf("%v\n", err)
//    }
//    return posts, nil

func (db *Database) SaveSinglePost(p *reddit.Post) error {
    _, err := db.Db.Exec("INSERT IGNORE INTO reddit_post (post_id, permalink, url, title, selftext, subreddit) VALUES (?, ?, ?, ?, ?, ?);", &p.FullID, &p.Permalink, &p.URL, &p.Title, &p.Body, &p.SubredditName)
    if err != nil {
        return err
    }
    return nil
}

func (db *Database) SaveSingleComment(cmt *reddit.Comment) error {
    _, err := db.Db.Exec("INSERT IGNORE INTO reddit_comment (comment_id, permalink, selftext, subreddit) VALUES (?, ?, ?, ?);", &cmt.FullID, &cmt.Permalink, &cmt.Body, &cmt.SubredditName) 
    if err != nil {
        return err
    }
    return nil
}
