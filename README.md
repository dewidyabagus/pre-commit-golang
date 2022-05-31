# pre-commit for go repository
Melakukan pengecekan pustaka kode untuk memastikan penformatan, kerentanan, permasalahan dan kualitas kode yang digunakan secara otomatis sebelum anda membuat komitmen dengan perubahan kode yang sudah dibuat.

# Pemasangan
Sebelum menggunakan pengait untuk melakukan tugasnya anda harus memasang terlebih dahulu paket ```pre-commit```.

### Pengguna Mac OS
```console
brew install pre-commit
```

### Pengguna Windows dan Linux
Sebelum memasang paket, pastikan Anda sudah memasang bahasa pemrograman ```python``` dikarenakan pemasangan menggunakan paket manajer dan berjalan diatas bahasa tersebut. Jika Anda belum memasangnya dapat mengunjungi link [download and install python](https://www.python.org/downloads/release/python-397/) untuk windows dan secara default untuk linux sudah dibekali python dalam instalasinya.

```console
python -m pip install pre-commit
```

### Pengecekan hasil pemasangan
Untuk memastikan pemasangan Anda berhasil jalankan perintah ini pada terminal:
```console
pre-commit --version

output: pre-commit 2.19.0
```

# Menambahkan Pada Pustaka Kode
Sebelum proses berjalan secara otomatis ada beberapa hal yang perlu dilakukan:

- Menambahkan pada pustaka kode.
   ```console
   pre-commit install
   ```
- Menambahkan file ```.pre-commit-config.yaml``` untuk daftar proses yang akan dilakukan sebelum komitmen perubahan kode dibuat.
- Atau dapat membuat dengan konfigurasi default ```pre-commit sample-config```.

# Menjalankan Pengcekan
Untuk menjalankan pengecekan tanpa harus melakukan commit
```console
pre-commit run --all-files
```

# Sumber Referensi
- https://pre-commit.com/index.html
- https://github.com/pre-commit/pre-commit-hooks
- https://github.com/Bahjat/pre-commit-golang
