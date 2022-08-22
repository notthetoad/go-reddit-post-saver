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
            case posts := <- pch:
                db.SaveAllPosts(posts)
            case cmts := <- cch:
                db.SaveAllComments(cmts)
        }
    }
    close(pch)
    close(cch)
}
