package main

import (
    "fmt"
    "os"
    "log"
    "context"
    "database/sql"
    "github.com/joho/godotenv"
    "github.com/vartanbeno/go-reddit/v2/reddit"
    "github.com/go-sql-driver/mysql"
)

var db *sql.DB

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
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")

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
