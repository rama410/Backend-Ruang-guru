# Artificial Intelligence menggunakan Golang

## Final Project AI-Powered Smart Home Energy Management System

### Description

Kamu akan mengembangkan Sistem Manajemen Energi Rumah Pintar menggunakan Golang dan [model AI Tapas](https://huggingface.co/google/tapas-base-finetuned-wtq) dari Huggingface Model Hub. Sistem ini akan memprediksi dan mengelola penggunaan energi dalam lingkungan rumah pintar. Aplikasi ini akan menerima data tentang penggunaan energi rumah dan memberikan wawasan dan rekomendasi tentang cara mengoptimalkan konsumsi energi.

Fitur:

- Prediksi Konsumsi Energi: Sistem ini akan memprediksi konsumsi energi rumah berdasarkan data historis.

- Rekomendasi Penghematan Energi: Sistem ini akan memberikan rekomendasi tentang cara menghemat energi berdasarkan konsumsi energi yang diprediksi.

Data input dalam bentuk format CSV dengan kolom berikut:

- Date: Tanggal data penggunaan energi.
- Time: Waktu data penggunaan energi.
- Appliance: Nama alat.
- Energy_Consumption: Konsumsi energi alat dalam kWh.
- Room: Ruang tempat alat berada.
- Status: Status alat (On/Off).

Contoh:

```txt
Date,Time,Appliance,Energy_Consumption,Room,Status
2022-01-01,00:00,Refrigerator,1.2,Kitchen,On
2022-01-01,01:00,Refrigerator,1.2,Kitchen,On
...
2022-01-01,08:00,TV,0.8,Living Room,Off
2022-01-01,09:00,TV,0.8,Living Room,On
2022-01-01,10:00,TV,0.8,Living Room,On
...
```

Untuk contoh, kalian bisa menggunakan file yang telah disiapkan `data-series.csv`.

#### Penggunaan Model AI:

Model AI Tapas `tapas-base-finetuned-wtq` akan digunakan untuk memahami data tabel dan membuat prediksi tentang konsumsi energi masa depan. Model ini akan menerima file CSV sebagai input dan menghasilkan prediksi untuk total konsumsi energi hari berikutnya.

Buatlah interface untuk aplikasi ini, bisa berupa aplikasi CLI maupun Web Application. Silahkan dikembangkan sehingga mirip dengan chatbot dimana user bisa bertanya mengenai data-data yang ada di file input.

Silahkan menggunakan model AI lainnya dari Hugging Face Hub untuk membuat aplikasi ini lebih menarik, misal-nya dengan menambahkan model AI `openai-community/gpt2` agar bisa memberikan rekomendasi tentang alat apa yang bisa digunakan lebih sedikit untuk menghemat energi.

### Constraints

Function `CsvToSlice` dan `ConnectAIModel` sudah diberikan dan wajib kalian gunakan. Silahkan membuat function-function lain yang kalian perlukan.

### Test Case Examples

#### Test Case CsvToSlice

**Input**:

```txt
"Name,Age\nJohn,30\nDoe,40"
```

**Expected Output / Behavior**:

{
    "Name": ["John", "Doe"],
    "Age": ["30", "40"]
}

**Explanation**:

Fungsi CsvToSlice menerima string dari file CSV sebagai input dan mengembalikan `map` di mana `key`-nya adalah header kolom dan `value`nya adalah data untuk setiap kolom. Dalam hal ini, string CSV input memiliki dua kolom "Name" dan "Age", dan dua baris data "John, 30" dan "Doe, 40". Fungsi ini harus mengembalikan map dengan dua data. Data pertama harus memiliki key "Name" dan value ["John", "Doe"], dan key kedua harus memiliki key "Age" dan value ["30", "40"].

#### Test Case 2

**Input**:

```txt
Payload: {
    "table": {
        "Name": ["John", "Doe"],
        "Age": ["30", "40"]
    },
    "query": "What is the age of John?"
}
```

**Expected Output / Behavior**:

```txt
{
    "answer": "30",
    "coordinates": [[0, 1]],
    "cells": ["30"],
    "aggregator": ""
}
```

**Explanation**:

Fungsi ConnectAIModel menerima payload dan Huggingface Token sebagai input dan mengembalikan struktur Response. Payload adalah struktur yang berisi `Table` dan `Query`. `Tabel` adalah sebuah map di mana `key`-nya adalah header kolom dan `value`-nya adalah irisan yang berisi data untuk setiap kolom. `Query` adalah string yang mewakili pertanyaan tentang data di tabel. Dalam hal ini, querynya adalah "Berapa umur John?". Fungsi ini harus mengembalikan struktur Response dengan jawaban "30", koordinat [[0, 1]], sel ["30"], dan aggregator.

Happy Coding!