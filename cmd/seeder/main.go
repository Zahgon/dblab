package main

import (

	// mysql driver.
	_ "github.com/go-sql-driver/mysql"

	// postgres driver.
	_ "github.com/lib/pq"

	// sqlite driver.
	_ "modernc.org/sqlite"
)

func main() {
	handleArgs()
}

func handleArgs() { _ = "STUB: not implemented"; return }
