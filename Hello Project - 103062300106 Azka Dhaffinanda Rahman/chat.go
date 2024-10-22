package main

import (
    "fmt"
)

const (
    maxAkun      = 100
    maxPesan     = 100
    maxGrup      = 10
    maxAnggota   = 10
    maxPesanGrup = 100
)

type Akun struct {
    Username string
    Password string
}

type Pesan struct {
    Pengirim string
    Penerima string
    IsiPesan string
}

type Grup struct {
    Nama    string
    Anggota [maxAnggota]string
    Jumlah  int
}

type PesanGrup struct {
    NamaGrup string
    Pengirim string
    IsiPesan string
}

var akunList [maxAkun]Akun
var pesanList [maxPesan]Pesan
var grupList [maxGrup]Grup
var pesanGrupList [maxPesanGrup]PesanGrup
var opt int
var back, konfirmasi string
var jumlahAkun, jumlahPesan, jumlahGrup, jumlahPesanGrup int

func main() {
    start()
}

func start() {
    var input string

    fmt.Println("===================")
    fmt.Println("=====TelU Chat=====")
    fmt.Println("===================")
    fmt.Println("                 ")
    fmt.Println("---------------------")
    fmt.Println("Tekan x Untuk Mulai")
    fmt.Println("---------------------")
    fmt.Scan(&input)
    menu()
}

func menu() {
    fmt.Println("                 ")
    fmt.Println("Menu Utama")
    fmt.Println("                 ")
    fmt.Println("1. Login Akun")
    fmt.Println("2. Buat Akun")
    fmt.Println("3. Mode Admin")
    fmt.Println("4. Keluar")

    fmt.Print("Pilih Opsi = ")
    fmt.Scan(&opt)

    if opt == 1 {
        login()
    } else if opt == 2 {
        regis()
    } else if opt == 3 {
        admin()
    } else if opt == 4 {
        fmt.Println("Keluar dari aplikasi.")
        return
    } else {
        fmt.Println("Opsi tidak valid.")
        menu()
    }
}

func login() {
    var username, password string

    fmt.Print("Username: ")
    fmt.Scan(&username)
    fmt.Print("Password: ")
    fmt.Scan(&password)

    for i := 0; i < jumlahAkun; i++ {
        if akunList[i].Username == username && akunList[i].Password == password {
            fmt.Println("Login berhasil!")
            fmt.Println("Lanjut Ke Akun (x)")
            fmt.Scan(&opt)
            loggedInMenu(username)
            return
        }
    }

    fmt.Println("Username Atau Password Salah.")
    fmt.Println("Kembali Ke Menu (x)")
    fmt.Scan(&back)
    menu()
}

func regis() {
    if jumlahAkun >= maxAkun {
        fmt.Println("Jumlah akun sudah maksimal.")
        fmt.Println("Kembali Ke Menu (x)")
        fmt.Scan(&back)
        menu()
        return
    }

    var username, password string

    fmt.Print("Username: ")
    fmt.Scan(&username)
    fmt.Print("Password: ")
    fmt.Scan(&password)

    fmt.Print("Konfirmasi Pembuatan Akun? (Y/N): ")
    fmt.Scan(&konfirmasi)

    if konfirmasi != "Y" && konfirmasi != "y" {
        fmt.Println("Pembuatan Akun Batal")
        fmt.Println("Kembali Ke Menu (x)")
        fmt.Scan(&back)
        menu()
    } else {
        akunList[jumlahAkun] = Akun{Username: username, Password: password}
        jumlahAkun++
        fmt.Println("Akun berhasil dibuat!")
        fmt.Println("Kembali Ke Menu (x)")
        fmt.Scan(&back)
        menu()
    }
}

func admin() {
    var useradmin, passadmin string

    fmt.Print("Username Admin: ")
    fmt.Scan(&useradmin)
    fmt.Print("Password Admin: ")
    fmt.Scan(&passadmin)

    if useradmin == "Manggarai" && passadmin == "2024" {
        adminControl()
    } else {
        fmt.Println("Login Invalid")
        fmt.Println("Kembali Ke Menu (x)")
        fmt.Scan(&back)
        menu()
    }
}

func adminControl() {
    fmt.Println("                    ")
    fmt.Println("Menu Admin")
    fmt.Println("                 ")
    fmt.Println("1. List Username Dan Password")
    fmt.Println("2. Mengurutkan Data Secara Alfabet")
    fmt.Println("3. Mencari User")
    fmt.Println("4. Kembali Ke Menu")

    fmt.Print("Pilih Opsi = ")
    fmt.Scan(&opt)

    if opt == 1 {
        adminControl1()
    } else if opt == 2 {
        adminControl2()
    } else if opt == 3 {
        adminControl3()
    } else if opt == 4 {
        menu()
    } else {
        fmt.Println("Opsi tidak valid.")
        adminControl()
    }
}

func adminControl1() {
    fmt.Println("Daftar Akun:")
    for i := 0; i < jumlahAkun; i++ {
        fmt.Printf("%d. Username: %s, Password: %s\n", i+1, akunList[i].Username, akunList[i].Password)
    }
    fmt.Println("Kembali Ke Menu Admin (x)")
    fmt.Scan(&back)
    adminControl()
}

func adminControl2() {
    insertionSort(akunList[:jumlahAkun])
    fmt.Println("Data telah diurutkan secara alfabet:")
    for i := 0; i < jumlahAkun; i++ {
        fmt.Printf("%d. Username: %s, Password: %s\n", i+1, akunList[i].Username, akunList[i].Password)
    }
    fmt.Println("Kembali Ke Menu Admin (x)")
    fmt.Scan(&back)
    adminControl()
}

func adminControl3() {
    insertionSort(akunList[:jumlahAkun])

    var username string
    fmt.Print("Masukkan username yang ingin dicari: ")
    fmt.Scan(&username)

    index := binarySearch(akunList[:jumlahAkun], username)
    if index != -1 {
        fmt.Printf("Akun dengan username '%s' ditemukan.\n", akunList[index].Username)
    } else {
        fmt.Printf("Akun dengan username '%s' tidak ditemukan.\n", username)
    }
    fmt.Println("Kembali Ke Menu Admin (x)")
    fmt.Scan(&back)
    adminControl()
}

func insertionSort(akunList []Akun) {
    for i := 1; i < len(akunList); i++ {
        key := akunList[i]
        j := i - 1

        for j >= 0 && akunList[j].Username > key.Username {
            akunList[j+1] = akunList[j]
            j = j - 1
        }
        akunList[j+1] = key
    }
}

func binarySearch(akunList []Akun, username string) int {
    low, high := 0, len(akunList)-1
    for low <= high {
        mid := (low + high) / 2
        if akunList[mid].Username == username {
            return mid
        } else if akunList[mid].Username < username {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return -1
}

func loggedInMenu(username string) {
    fmt.Println("                 ")
    fmt.Println("Menu Akun")
    fmt.Println("                 ")
    fmt.Println("1. Kirim Pesan")
    fmt.Println("2. Lihat Pesan")
    fmt.Println("3. Buat Grup")
    fmt.Println("4. Tambah Anggota Grup")
    fmt.Println("5. Kirim Pesan ke Grup")
    fmt.Println("6. Lihat Pesan di Grup")
    fmt.Println("7. Keluar")

    fmt.Print("Pilih Opsi = ")
    fmt.Scan(&opt)

    if opt == 1 {
        kirimPesan(username)
    } else if opt == 2 {
        lihatPesan(username)
    } else if opt == 3 {
        buatGrup(username)
    } else if opt == 4 {
        tambahAnggotaGrup(username)
    } else if opt == 5 {
        kirimPesanGrup(username)
    } else if opt == 6 {
        lihatPesanGrup(username)
    } else if opt == 7 {
        fmt.Println("Keluar dari akun.")
        menu()
    } else {
        fmt.Println("Opsi tidak valid.")
        fmt.Println("Kembali Ke Menu (x)")
        loggedInMenu(username)
    }
}

func kirimPesan(pengirim string) {
    if jumlahPesan >= maxPesan {
        fmt.Println("Jumlah pesan sudah maksimal.")
        fmt.Println("Kembali Ke Menu Akun (x)")
        fmt.Scan(&back)
        loggedInMenu(pengirim)
        return
    }

    var penerima, isiPesan string

    fmt.Print("Masukkan username penerima: ")
    fmt.Scan(&penerima)
    fmt.Print("Masukkan pesan: ")
    fmt.Scan(&isiPesan)

    // Validasi apakah penerima ada dalam daftar akun
    found := false
    for i := 0; i < jumlahAkun; i++ {
        if akunList[i].Username == penerima {
            found = true
            i = jumlahAkun
        }
    }

    if !found {
        fmt.Println("Penerima tidak ditemukan.")
        fmt.Println("Kembali Ke Menu Akun (x)")
        fmt.Scan(&back)
        loggedInMenu(pengirim)
        return
    }

    pesanList[jumlahPesan] = Pesan{Pengirim: pengirim, Penerima: penerima, IsiPesan: isiPesan}
    jumlahPesan++

    fmt.Println("Pesan berhasil dikirim!")
    fmt.Println("Kembali Ke Menu Akun (x)")
    fmt.Scan(&back)
    loggedInMenu(pengirim)
}

func lihatPesan(username string) {
    fmt.Println("Pesan yang diterima:")

    found := false
    for i := 0; i < jumlahPesan; i++ {
        if pesanList[i].Penerima == username {
            fmt.Printf("Dari: %s\nPesan: %s\n", pesanList[i].Pengirim, pesanList[i].IsiPesan)
            found = true
        }
    }

    if !found {
        fmt.Println("Tidak ada pesan yang diterima.")
    }

    fmt.Println("Kembali Ke Menu Akun (x)")
    fmt.Scan(&back)
    loggedInMenu(username)
}

func buatGrup(pembuat string) {
    if jumlahGrup >= maxGrup {
        fmt.Println("Jumlah grup sudah maksimal.")
        fmt.Println("Kembali Ke Menu Akun (x)")
        fmt.Scan(&back)
        loggedInMenu(pembuat)
        return
    }

    var namaGrup string

    fmt.Print("Masukkan nama grup: ")
    fmt.Scan(&namaGrup)

    grupList[jumlahGrup] = Grup{Nama: namaGrup, Anggota: [maxAnggota]string{pembuat}, Jumlah: 1}
    jumlahGrup++

    fmt.Println("Grup berhasil dibuat!")
    fmt.Println("Kembali Ke Menu Akun (x)")
    fmt.Scan(&back)
    loggedInMenu(pembuat)
}

func tambahAnggotaGrup(username string) {
    var namaGrup, anggotaBaru string

    fmt.Print("Masukkan nama grup: ")
    fmt.Scan(&namaGrup)
    fmt.Print("Masukkan username anggota baru: ")
    fmt.Scan(&anggotaBaru)

    // Validasi apakah anggota baru ada dalam daftar akun
    found := false
    for i := 0; i < jumlahAkun; i++ {
        if akunList[i].Username == anggotaBaru {
            found = true
            i = jumlahAkun
        }
    }

    if !found {
        fmt.Println("Anggota tidak ditemukan.")
        fmt.Println("Kembali Ke Menu Akun (x)")
        fmt.Scan(&back)
        loggedInMenu(username)
        return
    }

    // Tambahkan anggota baru ke grup
    for i := 0; i < jumlahGrup; i++ {
        if grupList[i].Nama == namaGrup {
            if grupList[i].Jumlah >= maxAnggota {
                fmt.Println("Jumlah anggota grup sudah maksimal.")
                fmt.Println("Kembali Ke Menu Akun (x)")
                fmt.Scan(&back)
                loggedInMenu(username)
                return
            }
            grupList[i].Anggota[grupList[i].Jumlah] = anggotaBaru
            grupList[i].Jumlah++
            fmt.Println("Anggota berhasil ditambahkan ke grup!")
            fmt.Println("Kembali Ke Menu Akun (x)")
            fmt.Scan(&back)
            loggedInMenu(username)
            return
        }
    }

    fmt.Println("Grup tidak ditemukan.")
    fmt.Println("Kembali Ke Menu Akun (x)")
    fmt.Scan(&back)
    loggedInMenu(username)
}

func kirimPesanGrup(pengirim string) {
    if jumlahPesanGrup >= maxPesanGrup {
        fmt.Println("Jumlah pesan grup sudah maksimal.")
        fmt.Println("Kembali Ke Menu Akun (x)")
        fmt.Scan(&back)
        loggedInMenu(pengirim)
        return
    }

    var namaGrup, isiPesan string

    fmt.Print("Masukkan nama grup: ")
    fmt.Scan(&namaGrup)
    fmt.Print("Masukkan pesan: ")
    fmt.Scan(&isiPesan)

    // Validasi apakah pengirim adalah anggota grup
    for i := 0; i < jumlahGrup; i++ {
        if grupList[i].Nama == namaGrup {
            isMember := false
            for j := 0; j < grupList[i].Jumlah; j++ {
                if grupList[i].Anggota[j] == pengirim {
                    isMember = true
                    i = jumlahAkun
                }
            }

            if !isMember {
                fmt.Println("Anda bukan anggota grup ini.")
                fmt.Println("Kembali Ke Menu Akun (x)")
                fmt.Scan(&back)
                loggedInMenu(pengirim)
                return
            }

            pesanGrupList[jumlahPesanGrup] = PesanGrup{NamaGrup: namaGrup, Pengirim: pengirim, IsiPesan: isiPesan}
            jumlahPesanGrup++

            fmt.Println("Pesan berhasil dikirim ke grup!")
            fmt.Println("Kembali Ke Menu Akun (x)")
            fmt.Scan(&back)
            loggedInMenu(pengirim)
            return
        }
    }

    fmt.Println("Grup tidak ditemukan.")
    fmt.Println("Kembali Ke Menu Akun (x)")
    fmt.Scan(&back)
    loggedInMenu(pengirim)
}

func lihatPesanGrup(username string) {
    var namaGrup string

    fmt.Print("Masukkan nama grup: ")
    fmt.Scan(&namaGrup)

    fmt.Printf("Pesan di grup %s:\n", namaGrup)

    found := false
    for i := 0; i < jumlahPesanGrup; i++ {
        if pesanGrupList[i].NamaGrup == namaGrup {
            fmt.Printf("Dari: %s\nPesan: %s\n", pesanGrupList[i].Pengirim, pesanGrupList[i].IsiPesan)
            found = true
        }
    }

    if !found {
        fmt.Println("Tidak ada pesan di grup ini.")
    }

    fmt.Println("Kembali Ke Menu Akun (x)")
    fmt.Scan(&back)
    loggedInMenu(username)
}
