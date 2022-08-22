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

    go user.GetSavedCommentsAndPosts(me, pch, cch)
    for {
        select {
            case posts, ok := <- pch:
                if !ok {
                    return
                }
                db.SaveAllPosts(posts)
            case cmts, ok := <- cch:
                if !ok {
                    return
                }
                db.SaveAllComments(cmts)
        }
    }
}
