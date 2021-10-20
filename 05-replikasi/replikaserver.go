package main

import (
	"fmt"
	"os"
)

type Row struct {
	// Definisikan struktur data yang akan disimpan
	// Ini hanya contoh saja
	id   string
	nama string
	ipk  float32
}

type Node struct {
	id      int16
	address string // address dengan format IP:Port
}

var data = []Row{}

var list_node = []Node{
	{1, "127.0.0.1:9000"},
	{2, "127.0.0.1:9001"},
	{3, "127.0.0.1:9003"},
}

func main() {
	serverID := os.Args[1]
	fmt.Println(serverID)

	// Implementasi server RPC dalam sebuah thread

	// Implementasi client RPC
}
