package main

import (
    "example.com/mdb"
    "example.com/user"
    "fmt"
)

func main() {
    db := mdb.Database{}
    db.InitDb()

    me := user.SignIn()    

    posts, cmts := user.GetSavedCommentsAndPosts(me)
    fmt.Println("posts: ", len(posts))
    fmt.Println("comments: ", len(cmts))

    for _, post := range posts {
        err := db.SaveSinglePost(post)
        if err != nil {
            fmt.Errorf("Error: %v", err)
        }
    }
    fmt.Println("posts added to db")
}
