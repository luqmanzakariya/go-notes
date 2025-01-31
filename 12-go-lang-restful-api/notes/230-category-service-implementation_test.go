package main

/* === Category Service Implementation ===
1. Buat file go baru (category_service_impl)
2. Buat struct service implementation tersebut
3. Karena kita menggunakan database transactional (mySql) berarti kita butuh requestnya dalam bentuk transactional
4. Function harus di wrap dalam transactional dan apabila terjadi error maka data harus di rollback transaksinya
*/