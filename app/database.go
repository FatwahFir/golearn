package app

import (
	"database/sql"
	"time"

	"github.com/FatwahFir/golearn/helpers"
)

func NewDb() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golearn_migration")
	helpers.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db

	//UP
	//migrate -database "mysql://root@tcp(localhost:3306)/golearn_migration" -path db/migrations up

	//DOWN
	//migrate -database "mysql://root@tcp(localhost:3306)/golearn_migration" -path db/migrations down

	//UP SPECIFIC VERSION
	//migrate -database "mysql://root@tcp(localhost:3306)/golearn_migration" -path db/migrations UP (1..N)

	//DOWN SPECIFIC VERSION
	//migrate -database "mysql://root@tcp(localhost:3306)/golearn_migration" -path db/migrations DOWN (1..N)

	//CHECK VERSION
	//migrate -database "mysql://root@tcp(localhost:3306)/golearn_migration" -path db/migrations version

	//FORCE MIGRATION
	//migrate -database "mysql://root@tcp(localhost:3306)/golearn_migration" -path db/migrations force 000001(version)

}
