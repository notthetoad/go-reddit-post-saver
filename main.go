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

type Post struct {
    ID int64
    Title string
    Body string
    Url string
}

func main() {
    fmt.Println(mdb.Test())
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
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")

    posts, err1 := queryPosts(1)
    if err1 != nil {
        log.Fatal(err)
    }
    fmt.Printf("Posts found: %v\n", posts)

    post, err2 := postByID(2)
    if err2 != nil {
        log.Fatal(err)
    }
    fmt.Printf("Post found: %v\n", post)

    pstID, err3 := addPost(Post{
        Title: "Test",
        Body: "for adding data",
        Url: "tothedb.com",
    })
    if err3 != nil {
        log.Fatal(err3)
    }
    fmt.Printf("ID of added post: %v\n", pstID)
    
    newPost, err4 := mdb.QueryPostByID(1, db)
    if err4 != nil {
        log.Fatal(err4)
    }
    fmt.Printf("new post %v\n", newPost)

//    err1 := godotenv.Load("local.env")
//    if err1 != nil {
//        log.Fatal("Error: %s", err)
//    }    
//    
//    var id, username, password, secret string
//    id = os.Getenv("CLIENT_ID")
//    username = os.Getenv("USERNAME")
//    password = os.Getenv("PASSWORD")
//    secret = os.Getenv("CLIENT_SECRET")
//
//    credentials := reddit.Credentials{ID: id, Secret: secret, Username: username, Password: password}
//    client, _ := reddit.NewClient(credentials)
    
//    posts, _, _ := client.Subreddit.TopPosts(context.Background(), "golang", &reddit.ListPostOptions{
//        ListOptions: reddit.ListOptions{
//            Limit: 5,
//        },
//        Time: "all",
//    })
//    fmt.Printf("Received %d posts.\n", posts)
//    posts, _, _, err2 := client.User.Saved(context.Background(), &reddit.ListUserOverviewOptions{
//        ListOptions: reddit.ListOptions{},
//        Time: "all",
//    })
//    if err2 != nil {
//        fmt.Println(err)
//    }
//   fmt.Println(posts)
//   for _, val := range posts {
//        fmt.Println(val.Title)
//   }
}

func queryPosts(id int64) ([]Post, error) {
    var posts []Post

    rows, err := db.Query("SELECT * FROM post WHERE ID = ?", id)
    if err != nil {
        return nil, fmt.Errorf("queryPosts %q: %v", id, err)
    }
    defer rows.Close()

    for rows.Next() {
        var pst Post
        if err := rows.Scan(&pst.ID, &pst.Title, &pst.Body, &pst.Url); err != nil {
            return nil, fmt.Errorf("queryPosts %q: %v", id, err)
        }
        posts = append(posts, pst)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("queryPosts %q: %v", id, err)
    }
    return posts, nil   
}

func postByID(id int64) (Post, error) {
    var pst Post

    row := db.QueryRow("SELECT * FROM post WHERE id = ?", id)
    if err := row.Scan(&pst.ID, &pst.Title, &pst.Body, &pst.Url); err != nil {
        return pst, fmt.Errorf("postByID %d: no such post", id)
    }
    return pst, nil
}

func addPost(pst Post) (int64, error) {
    result, err := db.Exec("INSERT INTO post (title, body, url) VALUES (?, ?, ?)", pst.Title, pst.Body, pst.Url)
    if err != nil {
        return 0, fmt.Errorf("addPost: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("addPost: %v", err)
    }
    return id, nil
}
