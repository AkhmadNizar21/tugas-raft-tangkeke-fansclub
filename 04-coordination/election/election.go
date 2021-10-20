package main

type Node struct {
	id      int16
	address string // address dengan format IP:Port
}

var list_node = []Node{
	{1, "127.0.0.1:9000"},
	{2, "127.0.0.1:9001"},
	{3, "127.0.0.1:9003"},
}

type Election struct {
}

// Fungsi untuk mengecek secara periodik apakah leader masih hidup atau tidak
func (t *Election) heartbeat() error {
	return nil
}

// Fungsi untuk melakukan election
func (t *Election) elect() error {
	return nil
}

func main() {
	// Implementasi RPC client dan server dalam satu program

	// Server bisa dihandle dengan menggunakan goroutines

}
