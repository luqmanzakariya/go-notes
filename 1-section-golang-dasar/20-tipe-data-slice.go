package main

import "fmt"

func slice() {
	// === Tipe Data Slice ===
	// Tipe data slice adalah potongan dari data array
	// Slice mirip dengan Array yang membedakan adalah ukuran Slice bisa berubah
	// Slice dan Array selalu terkoneksi, dimana Slice adalah data yang mengakses sebagian atau seluruh data di array

	// === Detail Tipe Data Slice ===
	// Tipe data Slice memiliki 3 data, yaitu pointer, length, dan capacity
	// Pointer adalah penunjuk data pertama di array para slice
	// Length adalah panjang dari slice, dan
	// Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity

	// === Membuat slice dari Array ===
	// array[low:high] >> membuat slice dari array dimulai index low sampai index sebelum high
	// array[low:] >> membuat slice dari array dimulai index low sampai index akhir di array
	// array[:high] >> membuat slice dari array dimulai index 0 sampai index sebelum high
	// array[:] >> membuat slice dari array dimulai index 0 sampai index akhir di array

	// # Data slice mengacu ke array, jadi jika data array diubah maka data slice akan berubah
	var months = [...]string{
		"Jan",
		"Feb",
		"Mar",
		"Apr",
		"May",
		"Jun",
		"Jul",
		"Aug",
		"Sep",
		"Oct",
		"Nov",
		"Dec",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(cap(slice1))

	months[5] = "Diubah"
	fmt.Println(slice1)

	// Function Slice
	// len(slice) >> Untuk pendapatkan panjang
	// cap(slice) >> Untuk mendapatkan kapasitas
	// append(slice, data) >> Membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas sudah penuh, maka akan membuat array baru
	// make([]TypeData, length, capacity) >> Membuat slice baru
	// copy(destination, source) >> Menyalin slice dari source ke destination

	days := [...]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	daysSlice1 := days[5:]
	daysSlice1[0] = "NewSat"
	daysSlice1[1] = "NewSun"
	fmt.Println(days)

	// Append
	fmt.Println("=== append")

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "XXX")
	fmt.Println(slice3)
	slice3[1] = "NoDec"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	// Kode Program Make Slice
	fmt.Println("=== Make Slice")
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Eko"
	newSlice[1] = "Kurniawan"
	fmt.Println(newSlice)

	fmt.Println("=== Copy Slice")
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// # Hati hati saat membuat array
	// Saat membuat array, kita harus berhati hati, jika salah, maka yang kita buat bukanlah Array, melainkan Slice
	// iniArray := [...]int{1,2,3,4,5} >> Harus dimasukkan jumlah kapasitas
	// iniSlice := []int{1,2,3,4,5} >> Tidak perlu dimasukkan jumlah kapasitas

}
