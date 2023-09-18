package main

import "sp-3-proxy/structs"

func main() {
	client := structs.NewClient(structs.NewProtectionRealDatabaseProxy("./dataSource.txt"))
	client.Start()
}
