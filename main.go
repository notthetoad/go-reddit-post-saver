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

    db.SaveAllComments(cmts)
    //if errp := db.SaveAllPosts(posts); errp != nil {
    //    for i, p := range posts {
    //        fmt.Println(i, p.FullID)
    //    }
    //}
    for _, p := range posts {
        db.SaveSinglePost(p)
    }
}
