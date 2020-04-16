# Clean Code

Clean code biasanya diartikan sebagai kode yang mudah dipahami dan mudah diubah.

## Why?

- Project biasanya dikerjakan bersama dengan tim
- Kode yang kotor akan sangat sulit untuk diubah
- Project jangka panjang membutuhkan maintenance code yang baik

## Contoh code yang tidak "Clean"

```go
    func SearchForWork() {
        i := 5
        last := 30 // Last created on days
        toddhms := DateTime.now

        search(i, toddhms)
    }
```

## Contoh code yang "Clean"

```go
    func SearchForWork() {
        workInWeeks := 5
        lastCreatedDays := 30
        today := DateTime.now

        getWorks(workInWeeks, lastCreatedDays, today)
    }
```

## Ciri-ciri

- Mudah dicari

    ```go
        i := 5
    ```

    ```go
        workInWeeks := 5
    ```

- Kadang tidak perlu komentar

    ```go
        last := 30 // Last created on days
    ```

    ```go
        lastCreatedDays := 30
    ```

- Mudah disebut

    ```go
        toddhms := DateTime.now
    ```

    ```go
        today := DateTime.now
    ```

- Mudah dipahami

    ```go
        search(i, toddhms)
    ```

    ```go
        getWorks(workInWeeks, lastCreatedDays, today)
    ```

- Kadang perlu komentar (contohnya pada pemanfaatan regex)

- Mematuhi konvensi, standar, dan peraturan

## KIS (Keep It Simple)

- Fungsi dan class harus kecil​

- Fungsi dan calss dibuat untuk melakukan hanya 1 tugas​

- Jangan gunakan terlalu banyak argumen dalam fungsi​

## Ada 3 pilihan konsep yang biasa digunakan

- DRY (Don’t Repeat Yourself)​

- WET (Write Everything Twice)​

- AHA (Avoid Hasty Abstraction)​

## Saran Formating

- Lebar baris 80 - 120 karakter​

- Satu class 300 - 500 baris​

- Code yang berhubungan usahakan berdekatan​

- Fungsi yang memanggil fungsi lainnya usahakan didekatkan​

- Deklarasi variabel sedekat mungkin dengan penggunaannya​

- Perhatikan indentasi​

- Prettier / Formatter​

# DO IT NOW!
