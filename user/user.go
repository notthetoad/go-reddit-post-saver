package user

import (
    "log"
    "fmt"
    "os"
    "context"
    "github.com/vartanbeno/go-reddit/v2/reddit"
    "github.com/joho/godotenv"
)

func SignIn() *reddit.Client {
    err := godotenv.Load("local.env")
    if err != nil {
        log.Fatal(err)
    }
    var id, username, password, secret string
    id = os.Getenv("CLIENT_ID")
    username = os.Getenv("USERNAME")
    password = os.Getenv("PASSWORD")
    secret = os.Getenv("CLIENT_SECRET")

    credentials := reddit.Credentials{ID: id, Secret: secret, Username: username, Password: password}
    client, _ := reddit.NewClient(credentials)
    return client
}

func GetSavedPosts(client *reddit.Client) []*reddit.Post{
    posts, _, _, err := client.User.Saved(context.Background(), &reddit.ListUserOverviewOptions{
        ListOptions: reddit.ListOptions{},
        Time: "all",
    })
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(posts)
    return posts
}

