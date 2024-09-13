package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/SabariGanesh-K/21BPS1209_Backend.git/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries;
var testDB  *sql.DB;
func TestMain(m *testing.M) {
	// config,Configerr:= util.LoadConfig("../../")
	// if Configerr!= nil {
	// 	log.Fatal("error loading config",Configerr)
	// }
	var err error
	testDB,err= sql.Open("postgres","postgresql://root:secret@localhost:5432/filemanager?sslmode=disable")
	log.Print("hello")
	if err!=nil {
		log.Fatal("Cannot connect to DB: ",err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}