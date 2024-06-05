# Manajemen Pendapatan dari YouTube AdSense

## Description

YouTube AdSense adalah program periklanan yang memungkinkan pemilik saluran YouTube untuk menghasilkan pendapatan dari iklan yang ditayangkan di video mereka. Untuk memaksimalkan pendapatan dari YouTube AdSense, diperlukan pemahaman yang baik tentang kinerja iklan, jumlah tayangan, dan metrik lainnya. Aplikasi ini dibangun untuk membantu pemilik saluran YouTube dalam mengelola dan menganalisis pendapatan dari YouTube AdSense.

## Features

### Fitur Utama
- Penyimpanan Data Akun: Setiap akun memiliki informasi seperti nama, channel, email, saldo, status YouTube (silver, platinum, gold, dll.), dan daftar judul video (durasi, banyaknya ditonton video tersebut).
- Pencarian Data Akun: Pencarian data akun berdasarkan channel atau status YouTube.
- Operasi CRUD: Pengubahan (edit) dan penghapusan data akun dan video dari akun tertentu.
- Tampilan Data Akun: Menampilkan data akun channel beserta daftar video yang terurut berdasarkan tanggal publish, berisi informasi channel, saldo, dan daftar judul video.

### Fitur Tambahan
- Perhitungan Monetisasi: Pengguna dapat memasukkan channel dan judul video yang ditonton. Setiap judul yang ditonton akan menambah saldo sebesar:
  - 1000 jika channel tersebut berstatus silver,
  - 2000 jika platinum,
  - 5000 jika gold.
  
    Jumlah saldo merupakan total yang diperoleh pada setiap video yang ditonton.

## Technologies Used
- **Programming Language:** Go
- **Build Tool:** `go build cmd/main.go`