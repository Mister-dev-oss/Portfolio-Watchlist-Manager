package data

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite", "portfolio.db")
	if err != nil {
		log.Fatalf("Errore apertura DB: %v", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatalf("Errore connessione DB: %v", err)
	}

	// Abilita le foreign key in SQLite
	_, err = DB.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		log.Fatalf("Errore nell'abilitazione foreign keys: %v", err)
	}

	log.Println("Connessione al DB aperta con successo e foreign keys abilitate")
}
