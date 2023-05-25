# Golang Dependency Injection
Dependency Injection merupakan sebuah teknik dimana sebuah object menerima object lain yang membutuhkannya (Dependencies) 
ketika pembuatan object itu sendiri.<br>
Biasanya object yang menerima Dependencies disebut client, proses mengirim Dependencies ke object tersebut biasanya disebut inject
### Function Sebagai Constructor 
Constructor yaitu sebuah function yang digunakan ketika sebuah object dibuat<br>
Di Golang, biasanya kita juga bisa membuat sebuah function untuk membuat object, dan mirip seperti Constructor tugasnya, yaitu membuat object baru<br>
Biasanya kita akan membuat object dengan memanggil function Constructor tersebut, lalu mengirim Dependencies yang dibutuhkan pada function Constructor tersebut. <br>
Cara ini mudah dilakukan ketika kode program aplikasi kita tidak terlalu besar <br>
Namun saat kode program aplikasi kita semakin besar, akan semakin sulit melakukan hal ini, terutama kita harus tahu urutan object mana yang harus dibuat terlebih dahulu <br>
Oleh karena itu, proses Dependency Injection sebenarnya bisa dipermudah dengan memanfaatkan library
