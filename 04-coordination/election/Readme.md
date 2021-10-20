# Tugas Election

Implementasikan mekanisme pemilihan koordinator dengan menggunakan Algoritma Bullying. Beberapa asumsi dalam proyek ini antara lain :
1. Masing-masing proses memiliki ID yang berupa angka integer
2. Terdapat struktur data (array) yang menyimpan informasi mengenai semua proses yang terlibat dalam sistem. Informasi setiap proses meliputi IP host dan nomor Port.
3. Komunikasi dilakukan dengan metode RPC. Setiap proses berperan sebagai client dan server RPC sekaligus.
4. Terdapat beberapa method RPC yang perlu diimplementasikan, antara lain :
- method ```heartbeat()``` untuk mengetahui apakah koordinator masih hidup atau tidak
- method ```elect()``` untuk mengirimkan pesan election ke sebuah proses