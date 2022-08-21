package main

import (
    "fmt"

    "example.com/user"
    "example.com/user/mdb"
)

func main() {
    db := mdb.Database{}
    db.InitDb()

    me := user.SignIn()    

    posts, cmts := user.GetSavedCommentsAndPosts(me)
    fmt.Println("posts: ", len(posts))
    fmt.Println("comments: ", len(cmts))

    db.SaveAllPosts(posts)
    db.SaveAllComments(cmts)
}
