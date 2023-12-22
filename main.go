package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// Struct untuk menyimpan data kategori
type Kategori struct {
	Nama  string `json:"nama-kategori"`
	Kunci string `json:"kata kunci"`
	Harga int    `json:"harga-barang"`
}

// Struct untuk menyimpan data belanjaan
type ShoppingItem struct {
	Name     string  `json:"name"`     // untuk nama barang yang akan di input
	Quantity int     `json:"quantity"` // jumlah barang
	Category string  `json:"kategori"` // kategori barang
	Price    float64 `json:"price"`    // harga barang
}

// Struct untuk menyimpan data aplikasi
type ShoppingApp struct {
	Categories   []Kategori      `json:"kategori"`
	ShoppingList []ShoppingItem `json:"Daftar Belanja"`
}


// FUNGSI CLEAR, UNTUK MEMBERSIHKAN SEBUAH TERMINAL SETELAH APLIKASI DI GUNAKAN
func clearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// MENAMPILKAN SEBUAH MENU SETELAH RUNNING FILE main.go
func showMenu() {
	clearScreen()
	fmt.Println("┌────────────────────────────────────────────────┐")
	fmt.Println("|             APLIKASI DAFTAR BELANJA            │")
	fmt.Println("├────────────────────────────────────────────────┤")
	fmt.Println("|   PILIHAN MENU                                 |")
	fmt.Println("│   1. Tambah Kategori                           │")
	fmt.Println("│   2. Cari Barang                               │")
	fmt.Println("│   3. Hapus Barang                              │")
	fmt.Println("│   4. Tambah Belanjaan                          │")
	fmt.Println("│   5. Cari Belanjaan                            │")
	fmt.Println("│   6. Hapus Belanjaan                           │")
	fmt.Println("│   7. Hitung Total Belanjaan                    │")
	fmt.Println("│   8. Keluar                                    │")
	fmt.Println("|                                                |")
	fmt.Println("└────────────────────────────────────────────────┘")
}

// KODE UNTUK MEMILIH PILIHAN YANG TERSEDIA SESUAI FUNGSI MULAI DARI 1-8
func main() {
	app := loadAppData()

	for {
		showMenu()
		var choice int
		fmt.Print("Pilih Menu [1-8]: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			tambahKategori(&app)
		case 2:
			cariKategori(&app)
		case 3:
			hapusKategori(&app)
		case 4:
			tambahBelanjaan(&app)
		case 5:
			cariBelanjaan(&app)
		case 6:
			hapusBelanjaan(&app)
		case 7:
			hitungBelanjaan(&app)
		case 8:
			// 8. KELUAR
			clearScreen()
			fmt.Printf("┌────────────────────────────────────────────────┐\n")
			fmt.Printf("│               SAMPAI JUMPA                     │\n")
			fmt.Printf("└────────────────────────────────────────────────┘\n\n")
			return
			default:
			// TERJADI JIKA TIDAK MEMILIH ANGKA YANG DI SEDIAKAN
			fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
		}
		for {
			// FUNGSI YANG MUNCUL SETELAH SELESAI MEMILIH
			fmt.Print("Apakah Anda ingin membuka menu lagi? (y/n): ")
			resp := getUserInput()
			// JIKA MEMILIH "Y" MAKA PROGRAM DI ULANG KEMBALI KE BAGIAN MENU
			if strings.ToLower(resp) == "y" {
				break
				// JIKA MEMILIH "N" MAKA PROGRAM BERHENTI
			} else if strings.ToLower(resp) == "n" {
				clearScreen()
				fmt.Printf("┌────────────────────────────────────────────────┐\n")
				fmt.Printf("│               SAMPAI JUMPA                     │\n")
				fmt.Printf("└────────────────────────────────────────────────┘\n\n")
				return
			} else {
				clearScreen()
				fmt.Println("Pilihan tidak valid. Harap masukkan 'y' atau 'n'.")
			}
		}
		clearScreen()
	}
}

// MEMILIKI HUBUNGAN DENGAN USER INPUT DI MANA AGAR BISA MENEKAN Y / N
func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// 1. PILIHAN UNTUK MENAMBAHKAN KATEGORI
func tambahKategori(app *ShoppingApp) {
	var category Kategori
	clearScreen()
	fmt.Printf("┌────────────────────────────────────────────────┐\n")
	fmt.Printf("│               TAMBAH KATEGORI                  │\n")
	fmt.Printf("└────────────────────────────────────────────────┘\n\n")
	fmt.Print("Nama Kategori: ")//nama benda/barang
	fmt.Scanln(&category.Nama)
	fmt.Print("Satuan Kata Kunci: ")//kategori benda 
	fmt.Scanln(&category.Kunci)
	fmt.Print("Harga Barang: ")
	fmt.Scanln(&category.Harga)

	app.Categories = append(app.Categories, category)
	saveAppData(*app)
	fmt.Println("")
	fmt.Println("Kategori berhasil ditambahkan!")
	fmt.Println("")
}

// 2. PILIHAN UNTUK MENCARI KATEGORI
func cariKategori(app *ShoppingApp) {
	var searchNama string
	clearScreen()
	fmt.Printf("┌────────────────────────────────────────────────┐\n")
	fmt.Printf("│               CARI KATEGORI                    │\n")
	fmt.Printf("└────────────────────────────────────────────────┘\n\n")
	fmt.Print("Masukkan nama kategori: ")
	fmt.Scanln(&searchNama)

	// Slice untuk menyimpan kategori yang cocok dengan kriteria pencarian
	var matchingCategories []Kategori

	for _, cat := range app.Categories {
		if strings.Contains(strings.ToLower(cat.Nama), strings.ToLower(searchNama)) {
			// Menambahkan kategori yang cocok ke dalam slice
			matchingCategories = append(matchingCategories, cat)
		}
	}

	if len(matchingCategories) > 0 {
		fmt.Println("")
		fmt.Println("Kategori ditemukan!")

		// Menampilkan semua kategori yang cocok
		for _, matchedCat := range matchingCategories {
			fmt.Println("")
			fmt.Printf("Nama Kategori: %s\n", matchedCat.Nama)
			fmt.Printf("Kata Kunci: %s\n", matchedCat.Kunci)
			fmt.Printf("Harga Barang: %d\n", matchedCat.Harga)
			fmt.Println("")
		}
	} else {
		fmt.Println("Kategori tidak ditemukan.")
	}
}

// 3. PILIHAN UNTUK MENGHAPUS KATEGORI
func hapusKategori(app *ShoppingApp) {
	clearScreen()
	fmt.Printf("┌────────────────────────────────────────────────┐\n")
	fmt.Printf("│               HAPUS KATEGORI                   │\n")
	fmt.Printf("└────────────────────────────────────────────────┘\n\n")
	fmt.Println("Daftar Kategori Tersedia:")
	for i, cat := range app.Categories {
		fmt.Printf("%d. %s\n", i+1, cat.Nama)//%d untuk int ,%s untuk string
	}

	var choice int
	fmt.Print("Pilih Kategori yang akan dihapus [1-", len(app.Categories), "]: ")
	fmt.Scanln(&choice)

	if choice >= 1 && choice <= len(app.Categories) {
		app.Categories = append(app.Categories[:choice-1], app.Categories[choice:]...)
		saveAppData(*app)
		fmt.Println("Kategori berhasil dihapus!")
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}


// 4. PILIHAN UNTUK MENAMBAH BELANJAAN
func tambahBelanjaan(app *ShoppingApp) {
	var item ShoppingItem
	clearScreen()
	fmt.Printf("┌────────────────────────────────────────────────┐\n")
	fmt.Printf("│               TAMBAH BELANJAAN                 │\n")
	fmt.Printf("└────────────────────────────────────────────────┘\n\n")
	fmt.Println("Daftar Kategori Tersedia:")
	for i, cat := range app.Categories {
		fmt.Printf("%d. %s\n", i+1, cat.Nama)
	}

	for {
		var categoryChoice int
		fmt.Println("")
		fmt.Print("Pilih Kategori Barang [1-", len(app.Categories), "]: ")
		fmt.Scanln(&categoryChoice)

		if categoryChoice >= 1 && categoryChoice <= len(app.Categories) {
			item.Category = app.Categories[categoryChoice-1].Nama
			break
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}

	fmt.Print("Nama Barang: ")
	fmt.Scanln(&item.Name)
	fmt.Print("Masukan jumlah barang: ")
	fmt.Scanln(&item.Quantity)
	fmt.Print("Harga Barang: ")
	fmt.Scanln(&item.Price)

	app.ShoppingList = append(app.ShoppingList, item)
	saveAppData(*app)
	fmt.Println("Daftar belanja berhasil ditambahkan!")
}

// 5. PILIHAN UNTUK MENCARI BELANJAAN
func cariBelanjaan(app *ShoppingApp) {
	var searchItem string
	clearScreen()
	fmt.Printf("┌────────────────────────────────────────────────┐\n")
	fmt.Printf("│               CARI BELANJAAN                   │\n")
	fmt.Printf("└────────────────────────────────────────────────┘\n\n")
	fmt.Print("Masukkan nama barang: ")
	fmt.Scanln(&searchItem)

	// Slice untuk menyimpan belanjaan yang cocok dengan kriteria pencarian
	var matchingItems []ShoppingItem

	for _, item := range app.ShoppingList {
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(searchItem)) {
			// Menambahkan barang yang cocok ke dalam slice
			matchingItems = append(matchingItems, item)
		}
	}

	if len(matchingItems) > 0 {
		fmt.Println("")
		fmt.Println("Belanjaan ditemukan!")

		// Menampilkan semua belanjaan yang cocok
		for _, matchedItem := range matchingItems {
			fmt.Println("")
			fmt.Printf("Nama Barang: %s\n", matchedItem.Name)
			fmt.Printf("Jumlah Barang: %d\n", matchedItem.Quantity)
			fmt.Printf("Kategori Barang: %s\n", matchedItem.Category)
			fmt.Printf("Harga Barang: %.2f\n", matchedItem.Price)
			fmt.Println("")
		}
	} else {
		fmt.Println("Belanjaan tidak ditemukan.")
	}
}
// 6. PILIHAN UNTUK MENGHAPUS BELANJAAN
func hapusBelanjaan(app *ShoppingApp) {
	clearScreen()
	fmt.Printf("┌────────────────────────────────────────────────┐\n")
	fmt.Printf("│               HAPUS BELANJAAN                  │\n")
	fmt.Printf("└────────────────────────────────────────────────┘\n\n")
	fmt.Println("Daftar Belanja Tersedia:")
	for i, item := range app.ShoppingList {
		fmt.Printf("%d. %s\n", i+1, item.Name)
	}

	var choice int
	fmt.Println("")
	fmt.Print("Pilih Barang yang akan dihapus [1-", len(app.ShoppingList), "]: ")
	fmt.Scanln(&choice)

	if choice >= 1 && choice <= len(app.ShoppingList) {
		app.ShoppingList = append(app.ShoppingList[:choice-1], app.ShoppingList[choice:]...)
		saveAppData(*app)
		fmt.Println("Barang berhasil dihapus!")
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

// 7. PILIHAN UNTUK MENGHITUNG TOTAL BELANJAAN
func hitungBelanjaan(app *ShoppingApp) {
	clearScreen()
	fmt.Printf("┌────────────────────────────────────────────────┐\n")
	fmt.Printf("│               HITUNGAN DAFTAR BELANJAAN        │\n")
	fmt.Printf("└────────────────────────────────────────────────┘\n\n")
	fmt.Println("Berikut adalah daftar belanjaan:")
	var totalPrice float64

	for i, item := range app.ShoppingList {
		fmt.Printf("%d. %s - %d %s - Rp %.2f\n", i+1, item.Name, item.Quantity, getCategoryAlias(app, item.Category), item.Price)
		totalPrice += float64(item.Quantity) * item.Price
	}

	fmt.Printf("Total harga belanjaan adalah Rp %.2f\n", totalPrice)
	fmt.Println("")
}

// FUNGSI UNTUK MENDAPATKAN KATA KUNCI KATEGORI BERDASARKAN NAMA KATEGORI
func getCategoryAlias(app *ShoppingApp, categoryName string) string {
	for _, cat := range app.Categories {
		if cat.Nama == categoryName {
			return cat.Kunci
		}
	}
	return ""
}

// Fungsi untuk menyimpan data aplikasi ke file JSON
func saveAppData(app ShoppingApp) {
	data, err := json.MarshalIndent(app, "", "  ")
	if err != nil {
		fmt.Println("Gagal menyimpan data aplikasi:", err)
		return
	}

	err = ioutil.WriteFile("data-kategori.json", data, 0644)
	if err != nil {
		fmt.Println("Gagal menyimpan data aplikasi:", err)
	}
}

// Fungsi untuk memuat data aplikasi dari file JSON
func loadAppData() ShoppingApp {
	data, err := ioutil.ReadFile("data-kategori.json")
	if err != nil {
		return ShoppingApp{}
	}

	var app ShoppingApp
	err = json.Unmarshal(data, &app)
	if err != nil {
		fmt.Println("Gagal memuat data aplikasi:", err)
		return ShoppingApp{}
	}

	return app
}