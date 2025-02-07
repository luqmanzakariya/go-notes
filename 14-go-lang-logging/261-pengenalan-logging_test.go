package main

/* === Pengenalan Logging ===
- Log file adalah file yang berisikan informasi kejadian dari sebuah sistem
- Biasanya dalam log file, terdapat informasi waktu kejadian dan pesan kejadian
- Logging adalah aksi menambah informasi log ke log file
- Logging sudah menjadi standar industri untuk menampilkan informasi yang terjadi di aplikasi yang
	kita buat
- Logging bukan hanya untuk menampilkan informasi, kadang digunakan untuk proses debugging ketika
	masalah terjadi di aplikasi kita
*/

/* === Diagram Logging ===
Aplikasi >> Mengirim Log Event >> Log File
*/

/* === Ekosistem Logging ===
Aplikasi >> Mengirim Log Event >> Log File >> Log Aggregator >> Log Database >> Log Management
*/

/* === Contoh Ekosistem Logging ===
- Log database yang sering digunakan: Elastic Search
- Log aggregator yang sering digunakan: Log Stash, Fluentd, Filebeat
- Log management: Kibana (melihat detail log), graylog
*/
