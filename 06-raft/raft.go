package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var heartbeat_counter = 0
var leaderID = 0
var isLeader = false
var self Node

func generateRandomTimeoutCounter() {
	minTimer := 100
	maxTimer := 500
	heartbeat_counter = rand.Intn(maxTimer-minTimer) + minTimer
}

func heartbeatCountdown() {
	for heartbeat_counter > 0 {
		heartbeat_counter = heartbeat_counter - 1
		time.Sleep(1 * time.Millisecond)
	}
	// Implementasikan permintaan RPC RequestVote di sini

}

/*
Struct implementasi Raft
*/
type Raft struct {
}

/*
Fungsi untuk menerima heartbeat message dan append entry dari Leader
*/
func (t *Raft) AppendEntries(log *LogEntry, reply *int) {
	generateRandomTimeoutCounter()
	if log != nil {
		logs = append(logs, *log)
	}
	*reply = 200 // return okay
	return nil
}

/*
Fungsi untuk menerima permintaan vote dari Candidate leader
*/
func (t *Raft) RequestVote(args *Node, reply *int) {
	isLeader = false
	leaderID = *args.id
	for _, server := range members{
		if args == server{
			server.leader = true
		}else{
			server.leader = false
		}
	}
	generateRandomTimeoutCounter()
	*reply = 200
	return nil
}

/*
Fungsi untuk menerima request command dari client.
- Jika saya adalah Leader maka jalankan prosedur Raft
- Jika saya bukan Leader maka arahkan client untuk mengirim request ke Leader
*/
func (t *Raft) ClientRequest(args *LogEntry, reply *int) {
	if (isLeader){
		logs = append(logs, *args)
		for _, server := range members {
			if server != self {
				client, err := rpc.DialHTTP("tcp", server.address)
				handleError(err)
				var result int
				err = client.Call("Raft.AppendEntries", args, &result)
				handleError(err)
			}
		}	
	}else{
		for _, server:=range members {
			if server.id == leaderID {
				client, err := rpc.DialHTTP("tcp", server.address)
				handleError(err)
				var result int
				err = client.Call("Raft.ClientRequest", args, &result)
				handleError(err)
				break
			}
		}
	}
}

type Node struct {
	id      int16
	address string // address dengan format IP:Port
	port string
	leader  bool
}

var members = []Node{
	{1, "127.0.0.1:9000", "9000", false},
	{2, "127.0.0.1:9001", "9001", false},
	{3, "127.0.0.1:9003", "9003", false},
}

type LogEntry struct {
	index     int32
	term      int32
	command   string
	committed bool
}

var logs = []LogEntry{}

/*
Fungsi untuk mengirim heartbeat message
berupa RPC AppendEntries dengan isi kosong
*/
func sendHeartbeatToFollowers() {
	// Kirim heartbeat hanya selama saya menjadi leader
	for isLeader {
		for _, server := range members {
			fmt.Println("Send AppendEntries RPC to follower ", server.id)
			// Implementasikan Call RPC AppendEntries ke semua follower

		}
		time.Sleep(10 * time.Millisecond)
	}

}

func main() {
	port := os.Args[1]
	// Implementasi server RPC sebagai goroutines
	fmt.Println("Jalankan server di port ", port, " status leader ", isLeader)
	for _, server := range members {
		if server.port == port {
			self = server
		}
	}

	// Implementasi server RPC dalam sebuah thread
	raft := &Raft{}
	rpc.Register(raft)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", self.port)
	handleError(err)
	go http.Serve(listener, nil)
	
	// Generate random number untuk countdown timer
	generateRandomTimeoutCounter()

	// Jalankan countdown timer sebagai goroutines
	go heartbeatCountdown()
	go sendHeartbeatToFollowers()
	// Wait agar program tidak berhenti
	for{

	}
	<-(chan int)(nil)
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}
}
