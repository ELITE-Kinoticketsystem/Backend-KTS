package main

import (
	"fmt"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
)

func main() {
	dbConnection, _ := managers.InitializeDB()
	fmt.Print(dbConnection)
}
