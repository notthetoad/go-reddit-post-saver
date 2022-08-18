package main

import (
    "example.com/mdb"
    "example.com/user"
    "fmt"
    "sync"

    //"github.com/vartanbeno/go-reddit/v2/reddit"
)

func main() {
    db := mdb.Database{}
    db.InitDb()

    me := user.SignIn()    

    posts, cmts := user.GetSavedCommentsAndPosts(me)
    fmt.Println("posts: ", len(posts))
    fmt.Println("comments: ", len(cmts))

    //db.Cache(posts)
    for _, post := range posts {
        err := db.SaveSinglePost(post)
        if err != nil {
            fmt.Errorf("Error inserting post: %v", err)
        }
    }
    fmt.Println("posts added to db")

    for _, cmt := range cmts {
        err := db.SaveSingleComment(cmt)
        if err != nil {
            fmt.Errorf("Error inserting comment: %v", err)
        }
    }
    fmt.Println("comments added to db")
}
