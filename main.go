package main

import (
    "fmt"
    "sync"

    "example.com/mdb"
    "example.com/user"
    "github.com/vartanbeno/go-reddit/v2/reddit"
)

func main() {
    db := mdb.Database{}
    db.InitDb()

    me := user.SignIn()    

    var wg sync.WaitGroup

    posts, cmts := user.GetSavedCommentsAndPosts(me)
    fmt.Println("posts: ", len(posts))
    fmt.Println("comments: ", len(cmts))

    wg.Add(2) 
    go func(posts []*reddit.Post) {
        for i, post := range posts {
            err := db.SaveSinglePost(post)
            if err != nil {
                fmt.Errorf("Error inserting post: %v", err)
            }
            fmt.Println("post", i)
        }
        fmt.Println("posts added to db")
        wg.Done()
    }(posts)

    go func(cmts []*reddit.Comment) {
        for i, cmt := range cmts {
            err := db.SaveSingleComment(cmt)
            if err != nil {
                fmt.Errorf("Error inserting comment: %v", err)
            }
            fmt.Println("comment", i)
        }
        fmt.Println("comments added to db")
        wg.Done()
    }(cmts)

    wg.Wait()
}
