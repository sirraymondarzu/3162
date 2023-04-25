package main

import (
	"context"
	"crypto/tls"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/sirraymondarzu/3162/internal/models"
	"github.com/alexedwards/scs/v2"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Share data across our handlers
type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
	reservation models.ReservationModel
	sessionManager *scs.SessionManager
}

func main() {
	// configure our server
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", os.Getenv("RCSYSTEM_DB_DSN"), "PostgreSQL DSN (Data Source Name)")
	flag.Parse()

	// get a database connection pool
	db, err := openDB(*dsn)
	if err != nil {
		log.Print(err)
		return
	}

	//create instances of error log and info log 
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	//set up a new session manager
	sessionManager := scs.New()
	sessionManager.Lifetime = 1 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.Secure = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode

	// share data across our handlers
	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
		reservation: models.ReservationModel{DB: db},
		sessionManager: sessionManager,
	}
	// cleanup the connection pool
	defer db.Close()
	// acquired a database connection pool
	infoLog.Printf("database connection pool established")
	// create and start a custom web server
	infoLog.Printf("starting server on %s", *addr)
	//configure TLS
	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}
	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		TLSConfig:    tlsConfig,
	}
	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	log.Fatal(err)
}
// The openDB() function returns a database connection pool or error
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	// create a context 
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	// test the DB connection
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}