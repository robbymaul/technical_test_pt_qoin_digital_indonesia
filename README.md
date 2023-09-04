# technical_test_pt_qoin_digital_indonesia

# ERD RUMAH MAKAN
![erd_rumah_makan drawio (2)](https://github.com/robbymaul/technical_test_pt_qoin_digital_indonesia/assets/104151418/a6dea00d-43f6-4e64-91c0-ff8215e9a4b1)

# Requirement
1. Admin dapat membuat menu pada aplikasi dan membuat user cashier
2. Chasier dapat membuat orderan sesaui menu yang ditambahkan pada cart
3. Chasier dapat membuat transaksi dari orderan yang telah dibuat dan mencetak struk sesaui dengan transaksi yang baru saja dibuat
4. laporan penghasilan mingguan dan bulanan bisa didapatkan dari tabel report yang didalamnya terdapat kumpulan transaksi yang telah dibuat jika ingin mendapatkan laporan selama seminggu dapat didapatkan dari created at pada setiap laporan dibuat, dan jika 30 hari bisa kita dapatkan dari relasi tabel transaksi yang mengandung tabel bulan.

# Tech Stack
Aplikasi ini lebih baik menggunakan basis mobile app karena mobile app sangat effesiensi pada pada aplikasi rumah makan membutuhkan seperti akses ke perangkat lain akan lebih mudah. cukup mobile kita bisa menjalankan aplikasi sehingga menghemat biaya dari perangkat yang akan digunakan.

1. Front End
untuk frontend saya menggunakan beberapa framework seperti react native dan flutter karena kedua technologi tersebut dapat sekaligus membuat aplikasi android dan ios dengan kode yang sama sehingga menghemat waktu dan biaya.

2. Back End
Untuk backend saya akan menggunakan golang karena memiliki pengelolaan memori yang baik, dukungan terhap multiplatform. golang merupakan bahasa yang mudah untuk dipelajari sehingga dapat mempercepat dalam pengembangan aplikasi.

3. DBMS
untuk database saya akan menggunakan sql karena aplikasi yang akan dibangun secara terstruktur dan definisi yang jelas. sehingga dapat mengurangi kesalahan pada saat pencatatan setiap data.

# Pemainan Dadu
untuk menjalankan aplikasi ketikan pada terminal :
```
go run main.go
```

``` masukan inputan pemain```

Masukkan jumlah player: 3
Masukkan jumlah dadu: 3
================      
Pemain = 3, Dadu = 3  
================= 

```permainan akan berjalan dan menentukan pemenangnya```

Giliran 1 lempar dadu: <br >
Pemain #1 (0): 1, 6, 1 <br >
Pemain #2 (0): 6, 4, 5 <br >
Pemain #3 (0): 1, 5,  <br >

Setelah evaluasi: <br >
Pemain #1 (1): 1 <br >
Pemain #2 (1): 1, 1, 4, 5 <br >
Pemain #3 (0): 5, 3 <br >

Giliran 2 lempar dadu: <br >
Pemain #1 (1): 5 <br >
Pemain #2 (1): 6, 1, 6, 6 <br >
Pemain #3 (0): 1, 1 <br >

Setelah evaluasi:
Pemain #1 (1): 5, 1, 1 <br >
Pemain #2 (4): <br >
Pemain #3 (0): 1 <br >

Giliran 3 lempar dadu: <br >
Pemain #1 (1): 2, 1, 1 <br >
Pemain #2 (4): <br >
Pemain #3 (0): 2 <br >

Setelah evaluasi: <br >
Pemain #1 (1): 2 <br >
Pemain #2 (4): 1, 1 <br >
Pemain #3 (0): 2 <br >

Giliran 4 lempar dadu: <br >
Pemain #1 (1): 5 <br >
Pemain #2 (4): 3, 5 <br >
Pemain #3 (0): 6 <br >

Setelah evaluasi: <br >
Pemain #1 (1): 5 <br >
Pemain #2 (4): 3, 5 <br >
Pemain #3 (1): <br >

Giliran 5 lempar dadu: <br >
Pemain #1 (1): 3 <br >
Pemain #2 (4): 2, 5 <br >
Pemain #3 (1): <br >

Setelah evaluasi:
Pemain #1 (1): 3 <br >
Pemain #2 (4): 2, 5 <br >
Pemain #3 (1): <br >

Giliran 6 lempar dadu: <br >
Pemain #1 (1): 3 <br >
Pemain #2 (4): 3, 2 <br >
Pemain #3 (1): <br >

Setelah evaluasi: <br >
Pemain #1 (1): 3 <br >
Pemain #2 (4): 3, 2 <br >
Pemain #3 (1): <br >

Giliran 7 lempar dadu: <br >
Pemain #1 (1): 4 <br >
Pemain #2 (4): 3, 5 <br >
Pemain #3 (1): <br >

Setelah evaluasi: <br >
Pemain #1 (1): 4 <br >
Pemain #2 (4): 3, 5 <br >
Pemain #3 (1): <br >

Giliran 8 lempar dadu: <br >
Pemain #1 (1): 5 <br >
Pemain #2 (4): 2, 1 <br >
Pemain #3 (1): <br >

Setelah evaluasi: <br >
Pemain #1 (1): 5 <br >
Pemain #2 (4): 2 <br >
Pemain #3 (1): 1 <br >

Giliran 9 lempar dadu: <br >
Pemain #1 (1): 6 <br >
Pemain #2 (4): 5 <br >
Pemain #3 (1): 1 <br >

Setelah evaluasi: <br >
Pemain #1 (2): 1 <br >
Pemain #2 (4): 5 <br >
Pemain #3 (1): <br >

Giliran 10 lempar dadu: <br >
Pemain #1 (2): 5 <br >
Pemain #2 (4): 2 <br >
Pemain #3 (1): <br >

Setelah evaluasi:
Pemain #1 (2): 5 <br >
Pemain #2 (4): 2 <br >
Pemain #3 (1): <br >

Giliran 11 lempar dadu <br >
Pemain #1 (2): 3 <br >
Pemain #2 (4): 5 <br >
Pemain #3 (1): <br >

Setelah evaluasi:
Pemain #1 (2): 3 <br >
Pemain #2 (4): 5 <br >
Pemain #3 (1): <br >

Giliran 12 lempar dadu: <br >
Pemain #1 (2): 1 <br >
Pemain #2 (4): 6 <br >
Pemain #3 (1): <br >

Setelah evaluasi: <br >
Pemain #1 (2): <br >
Pemain #2 (5): 1 <br >
Pemain #3 (1): <br >

> Game dimenangkan oleh pemain 2 karena memiliki poin lebih banyak dari pemain lainnya
