package main

import (
    "example.com/user"
    "example.com/user/mdb"
    "github.com/vartanbeno/go-reddit/v2/reddit"
)

func main() {
    db := mdb.Database{}
    db.InitDb()

    pch := make(chan []*reddit.Post)
    cch := make(chan []*reddit.Comment)

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
