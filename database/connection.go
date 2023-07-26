package database

import(
	"fmt"
	"database/sql"
    _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "postgres"
    dbname   = "TestDb"
)

func Connect() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
    if err != nil {
        return nil, err
    }

    return db, nil
}

func Close(db *sql.DB){
	db.Close()
}
 
func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}
