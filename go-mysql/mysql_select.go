package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Tag struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func main() {
    // Open up our database connection.
    db, err := sql.Open("mysql", "root:pass1@tcp(127.0.0.1:3306)/tuts")

    // if there is an error opening the connection, handle it
    if err != nil {
        log.Print(err.Error())
    }
    defer db.Close()

    // Execute the query
    results, err := db.Query("SELECT id, name FROM tags")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    for results.Next() {
        var tag Tag
        // for each row, scan the result into our tag composite object
        err = results.Scan(&tag.ID, &tag.Name)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
                // and then print out the tag's Name attribute
        log.Printf(tag.Name)
    }

}


//Note - If you retrieve 3 fields from the database and Scan only has 2 parameters, it will fail. They need to match up exactly.




//---->Querying a Single Row
//--->Say we wanted to query a single row this time and had an ID and again wanted to populate our struct. We could do that like so:


/*

var tag Tag

err = db.QueryRow("SELECT id, name FROM tags where id = ?", 2).Scan(&tag.ID, &tag.Name)
if err != nil {
    panic(err.Error()) 
}

log.Println(tag.ID)
log.Println(tag.Name)


*/