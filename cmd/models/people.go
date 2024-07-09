// All of the stuff for dealing with people in and out of the db
package people 

import (
	"fmt"
  "database/sql"
  //"log"
)


type Person struct {
  FName string
  LName string
}

type People = []Person


func getAll(db *sql.DB) {
 // rows, err := db.Query(`SELECT fname, lname FROM people`)


  //log.Print("error: ", err)
  //return nil 
}

func Test() {
  fmt.Println("A better place")
}
