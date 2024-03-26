package database

// Database connection pooling base example
// https://shahidyousuf.com/database-connection-pooling-example-in-golang

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var Pool *ConnectionPool

type ConnectionPool struct {
	// the underlying connection pool
	pool *sql.DB

	// the maximum number of connections in the pool
	maxConnections int

	// the current number of connections in the pool
	numConnections int

	// the mutex to synchronize access to the connection pool
	mutex *sync.Mutex
}

func NewConnectionPool(dsn string, maxConnections int) (*ConnectionPool, error) {
	// create a new connection pool
	pool, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// set the maximum number of connections in the pool
	pool.SetMaxOpenConns(maxConnections)

	// create a new ConnectionPool instance
	p := &ConnectionPool{
		pool:           pool,
		maxConnections: maxConnections,
		numConnections: 0,
		mutex:          &sync.Mutex{},
	}

	return p, nil
}

func (p *ConnectionPool) GetConnection() (*sql.DB, error) {
	// acquire the mutex lock
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// check if the pool is full
	if p.numConnections == p.maxConnections {
		return nil, fmt.Errorf("connection pool is full")
	}

	// increment the number of connections in the pool
	p.numConnections++
	println(p.numConnections)

	// return a connection from the underlying pool
	return p.pool, nil
}

func (p *ConnectionPool) ReleaseConnection(conn *sql.DB) {
	// acquire the mutex lock
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// decrement the number of connections in the pool
	p.numConnections--
}

func InitDatabase() {
	// create a new connection pool
	port := os.Getenv("DB_PORT")
	maxConnections, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTION"))
	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	connPool, err := NewConnectionPool(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database), maxConnections)

	if err != nil {
		log.Fatal(err)
	}

	defer connPool.pool.Close()

	Pool = connPool
}
