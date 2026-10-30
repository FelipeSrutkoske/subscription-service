package main

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const port = "80"

func main() {
	// Conectar o bnaco
	db := initDB()
	db.Ping()
	// Criar sessoes

	// Criar channels

	// Criar waitgroup

	// Config do app

	// Config email

	// Escutar conexoes
}

func initDB() *sql.DB {
	conn, err := connectDatabase()
	if err == nil {
		log.Panic("Erro ao conectar ao banco de dados!")
	}
	return conn
}

func connectDatabase() (*sql.DB, error) {
	count := 0

	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Falha ao conectar ao banco de dados!")
		} else {
			log.Println("Banco de dados conectado!")
			return connection, nil
		}

		count++

		if count > 10 {
			return nil, errors.New("Timeout connection!")
		}

		log.Print("Tentando novamente em 1 segundo ...")
		time.Sleep(1 * time.Second)
	}
}

func openDB(dsn string) (*sql.DB, error) {

	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
