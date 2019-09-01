package global

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/nikoksr/proji/internal/app/helper"
	"github.com/spf13/viper"
)

// Show shows detailed information about a global
func Show(globalType, globalID string) error {
	// Connect to database
	DBDir := helper.GetConfigDir() + "/db/"
	databaseName, ok := viper.Get("database.name").(string)

	if ok != true {
		return errors.New("could not read database name from config file")
	}

	db, err := sql.Open("sqlite3", DBDir+databaseName)
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	switch globalType {
	case "folder":
		err = showGlobalFolder(tx, globalID)
	case "file":
		err = showGlobalFile(tx, globalID)
	case "script":
		err = showGlobalFile(tx, globalID)
	default:
		err = fmt.Errorf("global type not valid")
	}

	if err != nil {
		return err
	}

	return tx.Commit()
}

// showGlobalFolder shows detailed information about a global folder
func showGlobalFolder(tx *sql.Tx, globalID string) error {
	stmt, err := tx.Prepare("SELECT target, template FROM global_folder WHERE global_folder_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	query, err := stmt.Query(globalID)
	if err != nil {
		return err
	}
	defer query.Close()
	var target, template string
	for query.Next() {
		query.Scan(&target, &template)
		fmt.Printf(" ID: %s | Target: %s | Template: %s\n", globalID, target, template)
	}
	fmt.Println()
	return nil
}

// showGlobalFile shows detailed information about a global file
func showGlobalFile(tx *sql.Tx, globalID string) error {
	stmt, err := tx.Prepare("SELECT target, template FROM global_file WHERE global_file_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	query, err := stmt.Query(globalID)
	if err != nil {
		return err
	}
	defer query.Close()
	var target, template string
	for query.Next() {
		query.Scan(&target, &template)
		fmt.Printf(" ID: %s | Target: %s | Template: %s\n", globalID, target, template)
	}
	fmt.Println()
	return nil
}

// showGlobalScript shows detailed information about a global script
func showGlobalScript(tx *sql.Tx, globalID string) error {
	stmt, err := tx.Prepare("SELECT name, run_as_sudo FROM global_script WHERE global_script_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	query, err := stmt.Query(globalID)
	if err != nil {
		return err
	}
	defer query.Close()
	var scriptName string
	var runAsSudo bool
	for query.Next() {
		query.Scan(&scriptName, &runAsSudo)
		fmt.Printf(" ID: %s | Script: %s | Sudo: %v\n", globalID, scriptName, runAsSudo)
	}
	fmt.Println()
	return nil
}