#  BigQuery SQL Driver for Golang
This is an implementation of the BigQuery Client as a database/sql/driver for easy integration and usage.


# Goals of project

This project is meant to be a basic database/sql driver implementation for Golang so that developers can easily use 
tools like Gorm, and *sql.DB functions, with Google's BigQuery database.

# Usage

Check out the example application in the `examples` directory, for a few examples.

As this is using the Google Cloud Go SDK, you will need to have your credentials available
via the GOOGLE_APPLICATION_CREDENTIALS environment variable point to your credential JSON file.

## Vanilla *sql.DB usage

Just like any other database/sql driver you'll need to import it 

```go
package main

import (
    "database/sql"
    _ "github.com/solcates/go-sql-bigquery"
    "log"
)

func main() {
    db, err := sql.Open("bigquery", 
        "bigquery://projectid/dataset")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close() 
    // Do Something with the DB

}
```

## Gorm Usage

For gorm

```go
package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/solcates/go-sql-bigquery/dialects/bigquery"
    "log"
)

func main() {
    db, err := gorm.Open("bigquery", 
        "bigquery://projectid/dataset")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close() 
    // Do Something with the DB

}
```


# Contribution

Pull Requests are welcome!  


# Current Support

* [x] driver.Conn implemented
* [x] driver.Querier implemented
* [x] driver.Pinger implemented
* [x] gorm Dialect - have only tested basic use cases
* [x] Prepared Statements - supported via a quick hack
* [ ] Parametiterized Queries
