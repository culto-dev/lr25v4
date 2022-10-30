package main

import (
    "lr25v4_back/database"
    "lr25v4_back/server"
)

func main() {
    database.Init()
    server.StartServer()
}
