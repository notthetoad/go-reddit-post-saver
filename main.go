package main

import (
    "example.com/mdb"
    "example.com/user"
    "fmt"
)

func main() {
    db := mdb.Database{}
    db.InitDb()

    //posts, err := db.QueryAllPosts()
    //if err != nil {
    //    panic(err)
    //}
    //fmt.Printf("posts: %v\n", posts)

    me := user.SignIn()    
    fmt.Println(me)

    posts := user.GetSavedPosts(me)
    for _, val := range posts {
       fmt.Println(val.Title)
    }
}
