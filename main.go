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

       for i, post := range posts {
            err := db.SaveSinglePost(post)
            if err != nil {
                fmt.Errorf("Error inserting post: %v", err)
            }
            fmt.Println("post", i)
        }
        fmt.Println("posts added to db")

        for i, cmt := range cmts {
            err := db.SaveSingleComment(cmt)
            if err != nil {
                fmt.Errorf("Error inserting comment: %v", err)
            }
            fmt.Println("comment", i)
        }
        fmt.Println("comments added to db")
}
