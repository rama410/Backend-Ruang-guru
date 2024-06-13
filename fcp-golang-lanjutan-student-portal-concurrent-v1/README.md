# Pemrograman Backend Lanjutan

## Student Portal Concurrent

### Description

Kita akan menggunakan Final Project dari Dasar Pemrograman Backend, yaitu Student Portal 3 di sini. Jadi jika kamu belum mengerjakan, silahkan diselesaikan dahulu. Jika sudah, maka bisa kamu gunakan kode kamu dari Final Project Dasar Pemrograman Backend ke sini.

Ada beberapa tambahan fitur yang akan kamu implementasikan di sini:

1. Import Student: Registrasi banyak data mahasiswa sekaligus.
2. Submit Assignment: Mahasiswa bisa submit tugas.
3. Update Welcome Message on Login: Menambahkan program study di welcome message pada function Login.
4. Maximum Login Attempts: Melakukan pengecekan percobaan login dengan ID yang sama, jika sudah melebihi batas maksimum, ID tersebut akan diblokir dari proses login.

#### Import Student

Kamu diberikan sebuah function `ReadStudentsFromCSV()` yang menerima satu parameter yaitu `filename` dan hasil dari function tersebut dalam bentuk slice `[]Student` yang disimpan di variable `students`

Data untuk import sudah diberikan dalam 3 file, masing-masing berisi 1.000 data yaitu `students1.csv`, `students2.csv`, `students3.csv` dan diterima di function `ImportStudents` sebagai slice `filenames`.

Disarankan untuk menggunakan teknik `concurrency` yang sudah kamu pelajari di proyek ini.

Silahkan kamu lanjutkan function `ImportStudents` agar data bisa di-registrasi dengan cepat dan tepat. Dimana kalian diharapkan melakukan import data yang berada di slice `filenames` menggunakan `ReadStudentsFromCSV()` dan melakukan registrasi menggunakan Fungsi `Register` yang sudah kamu miliki.

### Submit Assignment

Setiap siswa dapat mengirimkan tugas ke portal. Hanya saja dikarenakan masing-masing tugas merupakan file yang cukup besar, maka kamu akan diharapkan untuk menggunakan teknik `job queue` untuk meng-handle pengiriman tugas dari siswa.

`job queue` secara sederhana adalah ketika kita membuat jumlah `goroutine` yang terbatas sehingga misal kita memiliki 1000 tugas yang harus dikirimkan dan jumlah `goroutine` yang kita buat hanya 10, maka 10 `goroutine` tersebut akan meng-handle 10 tugas pertama, ketika salah satu `goroutine` sudah selesai maka `goroutine` tersebut akan meng-handle tugas selanjutnya dalam antrian. Ini ditujukan agar proses yang berat tidak membebani CPU dan memori.

Kamu ditugaskan untuk membuat 3 goroutine saja untuk meng-handle pengiriman tugas dari siswa, masing-masing goroutine wajib memanggil ke `SubmitAssignmentLongProcess()` untuk pemrosesan pengiriman tugas.

### Update Welcome Message on Login

Pada saat login, kita tambahkan di Welcome Message, program study dari mahasiswa yang berhasil login.

Sebelum update:
```bash
=== Login ===
ID: A12345
Name: Aditira
Login berhasil: Aditira
Press any key to continue...
```

Setelah update:
```bash
=== Login ===
ID: A12345
Name: Aditira
Login berhasil: Selamat datang Aditira! Kamu terdaftar di program studi: Teknik Informatika.
Press any key to continue...
```

Perhatikan baris `Login berhasil`, ada perubahan di sana yaitu kata "Selamat datang" serta tanda seru dan kalimat `Kamu terdaftar di program studi: Teknik Informatika`.

### Maximum Login Attempts

Untuk menjaga keamanan data mahasiswa, maka proses login yang gagal kita batasi maksimum 3 kali saja. Setelah 3 kali gagal, ID akan diblokir dan proses login dengan ID tersebut akan ditolak.

Apabila sebelum 3 kali proses login, ada 1 login yang berhasil, maka kita akan reset percobaan gagal ID tersebut menjadi 0.

Berikut langkah-langkah yang perlu dilakukan:
- Kamu perlu menambahkan sebuah `map[string]int` dengan nama `failedLoginAttempts` pada `InMemoryStudentManager`.
- Kamu juga perlu menginisialisasi `failedLoginAttempts` di function `NewInMemoryStudentManager`.
- Ubah kode di function `Login` menggunakan `failedLoginAttempts` untuk menyimpan ID dan berapa kali proses login sudah dilakukan untuk ID tersebut.
- Jika proses login untuk sebuah ID sudah lebih dari 3 kali, maka berikan pesan error `Login gagal: Batas maksimum login terlampaui`.
- Jika proses login berhasil, maka reset data ID tersebut di `failedLoginAttempts` jika ada.

### Constraints

- Isi dari semua csv file dipastikan unik, tidak ada ID student yang sama.
- Setiap student pasti memiliki program studi yang valid.

### Test Case Examples

#### Test Case 1

**Input**:

```go
sm := NewInMemoryStudentManager()
err := sm.ImportStudents([]string{"students1.csv", "students2.csv", "students3.csv"})
```

**Expected Output / Behavior**:

Tidak ada error, i.e. err should be nil.

**Explanation**:

Fungsi `ImportStudents` akan membaca isi dari semua file CSV dan mendaftarkan data student tersebut menggunakan fungsi `Register`. Kalau tidak ada error, maka `err` akan berisi `nil`, dan jumlah mahasiswa di `manager` setelah import akan bertambah sesuai dengan jumlah dari dari file CSV yaitu sebanyak 3000+3 data bawaan awal. Proses import sendiri harus berlangsung cepat, kurang dari 100ms (100 milidetik).

#### Test Case 2

**Input**:

```go
sm := NewInMemoryStudentManager()
sm.SubmitAssignments(10)
```

**Expected Output / Behavior**:

```bash
=== Submit Assignment ===
Enter the number of assignments you want to submit: 10
Worker 3: Processing assignment 1
Worker 1: Processing assignment 2
Worker 2: Processing assignment 3
Worker 2: Finished assignment 3
Worker 3: Finished assignment 1
Worker 3: Processing assignment 5
Worker 1: Finished assignment 2
Worker 1: Processing assignment 6
Worker 2: Processing assignment 4
Worker 2: Finished assignment 4
Worker 2: Processing assignment 7
Worker 3: Finished assignment 5
Worker 3: Processing assignment 8
Worker 1: Finished assignment 6
Worker 1: Processing assignment 9
Worker 1: Finished assignment 9
Worker 1: Processing assignment 10
Worker 2: Finished assignment 7
Worker 3: Finished assignment 8
Worker 1: Finished assignment 10
Submitting 10 assignments took 122.333986ms
Press any key to continue...
```

**Explanation**:

Fungsi `SubmitAssignments` akan berjalan sebanyak jumlah assignment yang diberikan di user input melalui terminal. Disimulasikan bahwa setiap assignment perlu diproses selama 3 detik. Perhatikan bahwa output di atas hanyalah contoh, kemungkinan besar urutan dari Finished Assignment akan acak.

#### Test Case 3

**Input**:

```go
sm := NewInMemoryStudentManager()
message, err := sm.Login("A12345", "Aditira")
```

**Expected Output / Behavior**:

```bash
=== Login ===
ID: A12345
Name: Aditira
Login berhasil: Selamat datang Aditira! Kamu terdaftar di program studi: Teknik Informatika.
Press any key to continue...
```

**Explanation**:

Sudah cukup jelas, lihat bagian Description.

#### Test Case 4

**Input**:

```go
sm := NewInMemoryStudentManager()
message, err := sm.Login("A12345", "no_name") // Output: Login gagal: data mahasiswa tidak ditemukan
message, err := sm.Login("A12345", "no_name") // Output: Login gagal: data mahasiswa tidak ditemukan
message, err := sm.Login("A12345", "no_name") // Output: Login gagal: data mahasiswa tidak ditemukan
message, err := sm.Login("A12345", "no_name") // Output: Login gagal: Batas maksimum login terlampaui
```

**Expected Output / Behavior**:

Pada login pertama sampai ketiga:
```bash
=== Login ===
ID: A12345
Name: no_name
Error: Login gagal: data mahasiswa tidak ditemukan

Press any key to continue...
```

Pada login keempat dan seterusnya:
```bash
=== Login ===
ID: A12345
Name: no_name
Error: Login gagal: Batas maksimum login terlampaui

Press any key to continue...
```

**Explanation**:

Pada saat login pertama sampai ketiga, dengan ID yang sama, namun dengan nama yang salah, sistem masih memperbolehkan login. Namun saat login keempat dan seterusnya, ID tersebut tidak bisa login.

#### Test Case 5

**Input**:

```go
sm := NewInMemoryStudentManager()
message, err := sm.Login("A12345", "no_name") // Output: Login gagal: data mahasiswa tidak ditemukan
message, err := sm.Login("A12345", "no_name") // Output: Login gagal: data mahasiswa tidak ditemukan
message, err := sm.Login("A12345", "Aditira") // Output: Login berhasil: Selamat datang Aditira! Kamu terdaftar di program studi: Teknik Informatika.
message, err := sm.Login("A12345", "no_name") // Output: Login gagal: data mahasiswa tidak ditemukan
message, err := sm.Login("A12345", "no_name") // Output: Login gagal: data mahasiswa tidak ditemukan
```

**Expected Output / Behavior**:

Pada login pertama sampai kedua:
```bash
=== Login ===
ID: A12345
Name: no_name
Error: Login gagal: data mahasiswa tidak ditemukan

Press any key to continue...
```

Pada login ketiga:
```bash
=== Login ===
ID: A12345
Name: Aditira
Login berhasil: Selamat datang Aditira! Kamu terdaftar di program studi: Teknik Informatika.
Press any key to continue...
```

Pada login keempat dan kelima:
```bash
=== Login ===
ID: A12345
Name: no_name
Error: Login gagal: data mahasiswa tidak ditemukan

Press any key to continue...
```

**Explanation**:

Pada saat login pertama dan kedua, login dilakukan dengan ID yang sama, namun dengan nama yang salah. Pada login ketiga, login dilakukan dengan ID dan nama yang benar dan mahasiswa berhasil login. Kemudian pada login keempat dan kelima, ID tersebut melakukan login namun dengan nama yang salah, di sini tidak terjadi pemblokiran login, karena pada saat login ketiga dengan ID dan nama yang benar, sudah kita reset menjadi 0, sehingga percobaan berikutnya dihitung dari 1, bukan dari 3.