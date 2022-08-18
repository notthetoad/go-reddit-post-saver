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

    db.Cache(posts)
    for _, post := range posts {
        tmp, err := db.SaveSinglePost(post)
        if err != nil {
            fmt.Errorf("Error: %v", err)
        }
        fmt.Println(tmp)
    }
}
