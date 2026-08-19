package database

import "fmt"

func ErrorCreateDB(nameDb string, err error) error {
	return fmt.Errorf("Error create database %s, error: %w", nameDb, err)
}
