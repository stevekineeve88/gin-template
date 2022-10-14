package managers

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"util/objects"
)

var connectionManager *ConnectionManager

func GetConnectionManager() ConnectionManager {
	if connectionManager == nil {
		port, err := strconv.ParseInt(os.Getenv("MYSQL_DB_PORT"), 10, 64)
		if err != nil {
			panic(err)
		}
		driverString := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s",
			os.Getenv("MYSQL_DB_USER"),
			os.Getenv("MYSQL_DB_PWD"),
			os.Getenv("MYSQL_DB_HOST"),
			port,
			os.Getenv("MYSQL_DB_NAME"),
		)
		driver, err := sql.Open("mysql", driverString)
		if err != nil {
			panic(err)
		}
		connectionManager = &ConnectionManager{
			db: driver,
		}
	}
	return *connectionManager
}

type ConnectionManager struct {
	db *sql.DB
}

func (manager *ConnectionManager) Select(query string) objects.DatabaseResult {
	res, err := manager.db.Query(query)
	if err != nil {
		return objects.DatabaseResult{
			Status:       false,
			Message:      err.Error(),
			Data:         nil,
			InsertId:     -1,
			AffectedRows: -1,
		}
	}

	return objects.DatabaseResult{
		Status:       true,
		Message:      "",
		Data:         res,
		InsertId:     -1,
		AffectedRows: -1,
	}
}

func (manager *ConnectionManager) Insert(query string) objects.DatabaseResult {
	res, err := manager.db.Exec(query)
	if err != nil {
		return objects.DatabaseResult{
			Status:       false,
			Message:      err.Error(),
			Data:         nil,
			InsertId:     -1,
			AffectedRows: -1,
		}
	}

	lastId, _ := res.LastInsertId()
	rowsAffected, _ := res.RowsAffected()
	return objects.DatabaseResult{
		Status:       true,
		Message:      "",
		Data:         nil,
		InsertId:     lastId,
		AffectedRows: rowsAffected,
	}
}

func (manager *ConnectionManager) Query(query string) objects.DatabaseResult {
	res, err := manager.db.Exec(query)
	if err != nil {
		return objects.DatabaseResult{
			Status:       false,
			Message:      err.Error(),
			Data:         nil,
			InsertId:     -1,
			AffectedRows: -1,
		}
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return objects.DatabaseResult{
			Status:       false,
			Message:      err.Error(),
			Data:         nil,
			InsertId:     -1,
			AffectedRows: -1,
		}
	}
	return objects.DatabaseResult{
		Status:       true,
		Message:      "",
		Data:         nil,
		InsertId:     -1,
		AffectedRows: rowsAffected,
	}
}
