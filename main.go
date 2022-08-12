package main

import (
    "fmt"
    "os"
    "log"
    "context"
    "github.com/joho/godotenv"
    "github.com/vartanbeno/go-reddit/v2/reddit"
)

func main() {
    err := godotenv.Load("local.env")
    if err != nil {
        log.Fatal("Error: %s", err)
    }    
    
    var id, username, password, secret string
    id = os.Getenv("CLIENT_ID")
    username = os.Getenv("USERNAME")
    password = os.Getenv("PASSWORD")
    secret = os.Getenv("CLIENT_SECRET")

    credentials := reddit.Credentials{ID: id, Secret: secret, Username: username, Password: password}
    client, _ := reddit.NewClient(credentials)
    
//    posts, _, _ := client.Subreddit.TopPosts(context.Background(), "golang", &reddit.ListPostOptions{
//        ListOptions: reddit.ListOptions{
//            Limit: 5,
//        },
//        Time: "all",
//    })
//    fmt.Printf("Received %d posts.\n", posts)
    posts, _, _, err := client.User.Saved(context.Background(), &reddit.ListUserOverviewOptions{
        ListOptions: reddit.ListOptions{},
        Time: "all",
    })
    if err != nil {
        fmt.Println(err)
    }
   fmt.Println(posts)
   for _, val := range posts {
        fmt.Println(val.Title)
   }
}
