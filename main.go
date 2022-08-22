package main

import (
    "fmt"

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
    go func(pch chan []*reddit.Post, cch chan []*reddit.Comment) {
        user.GetSavedCommentsAndPosts(me, pch, cch)
        for p := range pch {
            fmt.Println(p)
        } 
    }(pch, cch)
    //fmt.Println("posts: ", fmt.Println(<-pch))
    //fmt.Println("comments: ", fmt.Println(<-pch))

    //db.SaveAllPosts(posts)
    //db.SaveAllComments(cmts)
}
