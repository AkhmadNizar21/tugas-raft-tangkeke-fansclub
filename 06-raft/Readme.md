# Tugas Raft

Implementasikan protokol Raft untuk dua operasi :
1. Pemilihan Leader
2. Replikasi Log

## Pemilihan Leader
Pada awalnya seluruh node berstatus follower. Masing-masing node melakukan generate angka random untuk coutndown timer election. Jika coundown timer election habis dan belum ada RPC AppendEntries (heartbeat) dari Leader, maka Follower berubah status menjadi Candidate.

Catatan penting : 
1. Setiap follower memiliki countdown timer sendiri
2. Countdown timer akan direset dan digenerate ulang setelah menerima heartbeat

## Replikasi Log
Yang perlu disimpan di dalam log adalah informasi mengenai command, index, term dari entry tersebut dan penanda apakah entry sudah committed. 

