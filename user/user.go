package user

import (
    "log"
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

func GetSavedCommentsAndPosts(client *reddit.Client, pch chan []*reddit.Post, cch chan []*reddit.Comment) {
    //var allPosts []*reddit.Post
    //var allCmts []*reddit.Comment
    posts, cmts, resp, _ := client.User.Saved(context.Background(), &reddit.ListUserOverviewOptions{
        ListOptions: reddit.ListOptions{
            Limit: 5,
        },
        Time: "all",
    })
    //allPosts = append(allPosts, posts...)
    pch <- posts  
    //allCmts = append(allCmts, cmts...)
    cch <- cmts
    for resp.After != "" {
        posts, cmts, resp, _ = client.User.Saved(context.Background(), &reddit.ListUserOverviewOptions{
            ListOptions: reddit.ListOptions{
                Limit: 100,
                After: resp.After,
            },
            Time: "all",
        })
        //allPosts = append(allPosts, posts...)
        pch <- posts
        //allCmts = append(allCmts, cmts...)
        cch <- cmts
    } 
    close(pch)
    close(cch)
}

