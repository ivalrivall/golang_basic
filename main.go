package main

import "fmt"

func main() {
	// Go memiliki aturan cukup ketat dalam hal penulisan variabel. Ketika deklarasi, tipe data yg digunakan harus dituliskan juga. Istilah dari metode deklarasi variabel ini adalah manifest typing.
	var firstName string = "john"
	var lastName string
	lastName = "wick"
	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Printf("halo john wick!\n")
	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Println("halo", firstName, lastName+"!")

	// Deklarasi menggunakan := hanya bisa dilakukan di dalam blok fungsi, misalnya dalam blok fungsi main()
	// Deklarasi := disebut juga dengan konsep type inference, yaitu metode deklarasi variabel yang tipe data-nya diketahui secara otomatis dari data/nilai variabel.
	middleName := "wick"
	var x = "john"
	fmt.Printf("halo %s!\n", middleName)
	fmt.Println("halo", x)

	// Multi Variable
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	var fourth, fifth, sixth string = "empat", "lima", "enam"
	seventh, eight, ninth := "tujuh", "delapan", "sembilan"
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"
	fmt.Print(first, second, third)
	fmt.Print(fourth, fifth, sixth)
	fmt.Println(seventh, eight, ninth)
	fmt.Println(one, isFriday, twoPointTwo, say)

	// Underscore (_) adalah reserved variable yang bisa dimanfaatkan untuk menampung nilai yang tidak dipakai. Bisa dibilang variabel ini merupakan keranjang sampah.
	_ = "belajar Golang"
	_ = "Golang itu mudah"
	name, _ := "john", "wick"
	fmt.Println(name)

	// Fungsi new() digunakan untuk membuat variabel pointer dengan tipe data tertentu. Nilai data default-nya akan menyesuaikan tipe datanya.
	nameX := new(string)
	fmt.Println(nameX)  // 0x20818a220
	fmt.Println(*nameX) // ""

	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644
	// String format %d pada fmt.Printf() digunakan untuk memformat data numerik non-desimal.
	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	var decimalNumber = 2.62
	/*
		String format %f digunakan untuk memformat data numerik desimal menjadi string.
		Digit desimal yang akan dihasilkan adalah 6 digit.
		Pada contoh di atas, hasil format variabel decimalNumber adalah 2.620000.
		Jumlah digit yang muncul bisa dikontrol menggunakan %.nf, tinggal ganti n dengan angka yang diinginkan.
		Contoh: %.3f maka akan menghasilkan 3 digit desimal, %.10f maka akan menghasilkan 10 digit desimal.
	*/
	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	var exist bool = true
	// Gunakan %t untuk memformat data bool menggunakan fungsi fmt.Printf().
	fmt.Printf("exist? %t \n", exist)

	var message = `Nama saya "John Wick".
	Salam kenal.
	Mari belajar "Golang".`
	fmt.Println(message)
}
