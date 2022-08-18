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
    fmt.Println(db.Cached)
    fmt.Println(db.Cached[2])
    //db.SaveSinglePost(posts[len(posts)-1])
}
