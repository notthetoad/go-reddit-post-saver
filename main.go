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

    posts := user.GetSavedPosts(me)
    for i, val := range posts {
        fmt.Println(i, val.Title, "https://reddit.com"+val.Permalink, val.URL)
    }
}
