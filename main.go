package main

import (
	"database/sql"
	"fmt"
	"os"
	"os/user"
	"strconv"
)

func getgroup() string {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	gidInt, _ := strconv.Atoi(currentUser.Gid)
	group, _ := user.LookupGroupId(strconv.Itoa(gidInt))
	if group == nil {
		return "err: unknown group"
	}
	return group.Name
}

func checkdb() bool {
	dbPath := os.Getenv("QUOTA_DB_PATH")
	var dbStatus bool
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		fmt.Printf("Database file does not exist: %s\n", dbPath)
		dbStatus = false
	} else if err != nil {
		fmt.Printf("Error checking file: %v\n", err)
		dbStatus = false
	} else {
		fmt.Printf("File exists: %s\n", dbPath)
		dbStatus = true
	}
	return dbStatus
}

func readdb(group string) (bool, error) {
	var status string
	// see if the db exists first
	if checkdb() == true {
		db, _ := sql.Open("sqlite3", os.Getenv("QUOTA_DB_PATH"))
		if err := db.QueryRow("SELECT status FROM quotas where group_name = ?", group).Scan(&status); err != nil {
			status = err
			if err == sql.ErrNoRows {
				return false, fmt.Errorf("status %d: unknown status", group)
			}
			return false, fmt.Errorf("status %d: %v", group, err)
		}
	}
}

func main() {
	fmt.Println(readdb(getgroup()))
}
