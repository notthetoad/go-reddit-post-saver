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
}
