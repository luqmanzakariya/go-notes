package belajargolanggoroutine

/* === Pengenalan Channel  ===
- Channel adalah tempat komunikasi secara synchronous yang bisa dilakukan oleh goroutine
- Di Channel terdapat pengirim dan penerima, biasanya pengirim dan penerima adalah goroutine yang
	berbeda
- Saat melakukan pengiriman data ke Channel, goroutine akan ter-block, sampai ada yang menerima
	data tersebut
- Maka dari itu, channel disebut sebagai alat komunikasi synchronous (blocking)
- Channel cocok sekali sebagai alternatif seperti mekanisme async await yang terdapat di beberapa
	bahasa pemrograman lain
*/

/* === Karakteristik Channel  ===
- Secara default channel hanya bisa menampung satu data, jika kita ingin menambahkan data lagi,
	harus menunggu data yang ada di channel diambil
- Channel hanya bisa menerima satu jenis data
- Channel bisa diambil lebih dari satu goroutine
- Channel harus di close jika tidak digunakan, atau bisa menyebabkan memory leak
*/
