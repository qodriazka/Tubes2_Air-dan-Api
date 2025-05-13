# Little Alchemy 2 Recipe Finder
Aplikasi web ini digunakan untuk memudahkan pengguna mencari resep-resep dalam permainan Little Alchemy 2 berdasarkan elemen target yang dimasukkan ketika pertama kali mengakses aplikasi web. Aplikasi web ini memiliki tiga pilihan algoritma pencarian resep target elemen tersebut yaitu BFS (Breadth First Search), DFS (Depth First Search), atau Bidirectional untuk menemukan cara mengkombinasikan elemen-elemen dasar hingga mencapai elemen target.

## Penjelasan Algoritma Pencarian
BFS adalah algoritma pencarian graf yang menjelajahi graf dengan cara lapis demi lapis, mulai dari simpul akar dan mengunjungi semua tetangga pada kedalaman d sebelum berpindah ke kedalaman d+1. 
DFS (Depth-First Search) adalah algoritma pencarian graf yang mengeksplorasi kedalaman suatu jalur terlebih dahulu sebelum melakukan backtracking yaitu kembali ke simpul sebelumnya untuk mengeksplorasi jalur lain yang belum dikunjungi.
Bidirectional search menjalankan dua BFS sekaligus—satu dari sumber dan satu dari tujuan—hingga mereka bertemu di tengah.

## Requirement Program
- Golang
- Next.js

## Cara Menjalankan Program
- Clone repositori ini dengan cara mengetik ini di terminal (atau jika ada masalah maka dari Release download Assets -> Source Code (zip) -> Ekstrak zipnya):
```shell
git clone https://github.com/qodriazka/Tubes2_Air-dan-Api.git
```

- Navigasi ke repositori yang sudah di-clone, lalu arahkan ke folder source codenya dan ke backend:
```shell
cd src/backend
```

- Menjalankan program backend terlebih dahulu:
```shell
go run main.go
```

- Membuka terminal baru dan tidak menutup terminal sebelumnya yang menjalankan backend, lalu mengarahkan ke source code dan folder frontend:
```shell
cd src/frontend
```

- Melakukan instalasi pada direktori tersebut:
```shell
npm install mermaid
npm install
```

- Setelah itu, cara menjalankan aplikasi web ini yaitu:
```shell
npm run dev
```

<div id="Author">
  <strong>
    <h3>Author</h3>
    <table align="center">
      <tr>
        <td>NIM</td>
        <td>Nama</td>
      </tr>
      <tr>
        <td>10122078</td>
        <td>Ghaisan Zaki Pratama</td>
      </tr>
      <tr>
        <td>13523010</td>
        <td>Qodri Azkarayan</td>
      </tr>
      <tr>
        <td>13523016</td>
        <td>Clarissa Nethania Tambunan</td>
    </table>
  </strong>
</div>