package main

import (
	"github.com/R0bson/sqlboiler/drivers"
	"github.com/R0bson/sqlboiler/drivers/sqlboiler-mssql/driver"
)

func main() {
	drivers.DriverMain(&driver.MSSQLDriver{})
}
