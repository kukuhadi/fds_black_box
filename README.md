# Fraud Detection - Skimming Card

### Package yang dibutuhkan 

Package - package ini harus diinstal di mesin source code yang digunakan untuk membangun aplikasi FDS - Skimming Card. 

```sh
> github.com/keltia/leftpad
> gopkg.in/resty.v1
> github.com/mediocregopher/radix.v2/redis
> github.com/kelvins/geocoder/structs
> github.com/gin-gonic/gin
> github.com/go-sql-driver/mysql
```

### Initialize Aplikasi
Setelah seluruh kebutuhan diatas selesai, maka periksa kembali parameter yang akan digunakan di dalam operasional aplikasi. 

```sh
  os.Setenv("hostDatabase", "[IP SERVER DATABASE]")
  os.Setenv("usernameDBMysql", "[USERNAME DB]")
  os.Setenv("passwordDBMysql", "[PASSWORD DB]")
  os.Setenv("schemaDatabase", "[SCHEMA DB - ATM LOCATION]")
  os.Setenv("schemaDatabase_Seq", "[SCHEMA DB - SEQ ACTIVITY]")
```
