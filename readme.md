# Nama : Muhammad Arif Sulaksono

A.
1. [System Design](https://drive.google.com/file/d/1R4qQaLSKhEtfTGDl2wnqvIrCnwPq8KcW/view?usp=sharing)
2. Hal dapat dilakukan untuk mempertimbangkan performa kinerja sistem dan tetap dapat diakses oleh banyak orang:
- Cache : Menyimpan data yang sering diakses, seperti profil user, data kursi dan jadwal film untuk meningkatkan kecepatan response time dan mengurangi beban server.
- Reservation system : mengirimkan notifikasi kepada pemesan untuk segera melakukan pembayaran. jika dalam beberapa waktu (10-15 menit) tidak menyelesaikan transaksi, maka reservasi akan otomatis dibatalkan dan kursi akan kembali tersedia untuk dibeli.  
- Database : memaksimal query database seperti membuat indexing.
- Server : menggunakan teknologi load balancing untuk mendistribusikan traffic ke beberapa server, sehingga performa sistem tetap optimal meskipun banyak orang yang mengaksesnya secara bersamaan.

B. [Database Design](https://drive.google.com/file/d/10V7YNghY0c4EwLBUrHZ9hneWQkyDTGUr/view?usp=sharing)

C. Please clone this project and follow the instruction:
## Import database
- Download [Database](https://drive.google.com/file/d/11PxoPOJR2I-tqOwt-yBGlfbYIz6FN7IC/view?usp=sharing), and import to your local.

## Step
- Clone this repository with command :
```git clone https://github.com/marifsulaksono/eticketing-test-mkp.git```
- Configure .env file with following .env.example.
- Run project. ```go run .```

### Contact me
[WhatsApp](wa.me/6285331828197)
[Email](mailto:marifsulaksono@gmail.com)