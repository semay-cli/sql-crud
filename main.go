package main

import "github.com/semay-cli/sql-crud/manager"

// go build -tags netgo -ldflags '-s -w' -o /home/bushu/go/bin/scrud
func main() {
	manager.Execute()
}
