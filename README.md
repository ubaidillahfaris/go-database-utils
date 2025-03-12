# go-database-utils 
🇮🇩 Bahasa Indonesia | 🇬🇬 English | 🇨🇳 中文
## Installation
Untuk menginstal package ini, gunakan perintah berikut<br>
To install this package, use the following command<br>
要安装此软件包，请使用以下命令<br>
  

```
go get github.com/ubaidillahfaris/go-database-utils@v1.0.0  
```
## Environment Variables
Pastikan Anda memiliki file .env dengan konfigurasi berikut <br>
Make sure you have a .env file with the following configuration:<br>
确保您具有以下配置的 .env 文件：<br>
  

```
DB_CONNECTION=  
DB_HOST=  
DB_PORT=  
DB_DATABASE=  
DB_USERNAME=  
DB_PASSWORD=  
```
# Usage
Contoh penggunaan untuk mengambil data dari tabel parfums dengan filter<br>
Example usage: Retrieve data from the parfums table with filters<nr>
示例用法：从 parfums 表中获取符合筛选条件的数据。<br>


### package models
```
package models

type Parfum struct {
	ID     int     `json:"id"`
	Kode   string  `json:"kode"`
	Nama   string  `json:"nama"`
	Harga  float64 `json:"harga"`
	Volume float64 `json:"volume"`
}

```


```
package main  

import (  
    "fmt"  
    "github.com/ubaidillahfaris/go-database-utils/database"  
    "github.com/ubaidillahfaris/go-database-utils/models"  
)  

func main() {  
    var parfums []models.Parfum  

    err := database.DB("parfums").  
        Where("harga > ?", 50000).  
        Where("volume > ?", 50).  
        All(&parfums)  

    if err != nil {  
        fmt.Println("Error:", err)  
        return  
    }  

    for _, parfum := range parfums {  
        fmt.Println("ID:", parfum.ID)  
        fmt.Println("Kode:", parfum.Kode)  
        fmt.Println("Nama:", parfum.Nama)  
        fmt.Println("Harga:", parfum.Harga)  
        fmt.Println("Volume:", parfum.Volume)  
        fmt.Println("---")  
    }  
}  
```

Kode ini memungkinkan penggunaan beberapa koneksi database sekaligus, seperti MySQL dan PostgreSQL.
Kita dapat menentukan koneksi yang diinginkan dengan parameter kedua pada fungsi 
```
database.DB(nama_tabel, skema)
```
<br>

This code allows multiple database connections simultaneously, such as MySQL and PostgreSQL.
We can specify the desired connection using the second parameter in the 
```
database.DB(table_name, scheme)
```
<br>

此代码允许同时使用多个数据库连接，例如 MySQL 和 PostgreSQL。
我们可以通过 
```
database.DB(table_name, scheme) 
```
函数的第二个参数指定所需的连接。

Menggunakan MySQL untuk mengakses tabel users. <br>
Using MySQL to access the users table. <br>
使用 MySQL 访问 users 表。<br>

```
func main() {
	users := database.DB("users", "mysql")
	print(users)
}

func main() {
	users := database.DB("users", "pgsql")
	print(users)
}

```


## Features
Mudah digunakan dengan metode chaining seperti ORM Laravel  
Dukungan berbagai database seperti MySQL, PostgreSQL, SQLite  
Query builder sederhana untuk operasi database umum seperti Where, All, First, dll <br>

Easy to use with method chaining like Laravel ORM.
Supports multiple databases such as MySQL, PostgreSQL, and SQLite. Simple query builder for common database operations such as Where, All, First, etc. <br>

易于使用，支持类似 Laravel ORM 的方法链式调用。
支持多种数据库，如 MySQL、PostgreSQL、SQLite。
简单的查询构造器，适用于常见数据库操作，如 Where、All、First 等。<br>




## License

[MIT](https://choosealicense.com/licenses/mit/)
