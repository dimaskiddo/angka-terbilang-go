package angka

import (
	"strconv"
	"strings"
)

// Definisi Array Angka dan Units
var arrAngka = [...]string{"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan"}
var arrUnits = [...]string{"", "ribu", "juta", "milyar", "triliun", "quadriliun", "quintiliun", "sextiliun", "septiliun", "oktiliun", "noniliun", "desiliun",
	"undesiliun", "duodesiliun", "tredesiliun", "quattuordesiliun", "quindesiliun", "sexdesiliun", "septendesiliun", "oktodesiliun", "novemdesiliun", "vigintiliun"}

// ToTerbilang Function
func ToTerbilang(angka string) (string, error) {
	// Definisi Variabel Hasil Konversi Terbilang
	var resTerbilang string

	// Trim Inputan Angka
	// dan Cari Panjang Angka String
	strAngka := strings.TrimSpace(angka)
	lenAngka := len(strAngka)

	// Loop Angka String dan Konversi
	for i := 0; i < lenAngka; i++ {
		if i == (lenAngka-1) && (lenAngka-1) == 0 {
			// Proses Nol
			return "nol", nil
		}

		// Konversi Angka String ke Angka Int
		intAngka, err := strconv.Atoi(string(strAngka[i]))
		if err != nil {
			return "", err
		}

		// Konversi Angka ke Bilangan
		if arrAngka[intAngka] != "" {
			resTerbilang += arrAngka[intAngka] + " "
		} else {
			resTerbilang += arrAngka[intAngka]
		}

		// Cari Posisi Digit
		posDigit := (lenAngka - 1) - i
		grpDigit := posDigit % 3

		if intAngka != 0 {
			switch grpDigit {
			case 2:
				// Proses Ratusan
				if intAngka > 1 {
					resTerbilang += "ratus "
				} else if intAngka == 1 {
					// Cara Bar-Bar Hapus Konversi "Satu"
					// Dan Append Bentuk Baru :D
					resTerbilang = strings.TrimSuffix(resTerbilang, arrAngka[intAngka]+" ")
					resTerbilang += "seratus "
				}
			case 1:
				// Proses Puluhan
				if intAngka > 1 {
					resTerbilang += "puluh "
				} else if intAngka == 1 {
					// Konversi Angka String Selanjutnya ke Angka Int Selanjutnya
					nextIntAngka, err := strconv.Atoi(string(strAngka[i+1]))
					if err != nil {
						return "", err
					}

					// Cara Bar-Bar Hapus Konversi "Satu"
					// Dan Append Bentuk Baru :D
					resTerbilang = strings.TrimSuffix(resTerbilang, arrAngka[intAngka]+" ")
					switch nextIntAngka {
					case 0:
						// Proses Sepuluh
						resTerbilang += "sepuluh "
					case 1:
						// Proses Sebelas
						resTerbilang += "sebelas "
					default:
						// Proses Belasan
						resTerbilang += arrAngka[nextIntAngka] + " belas "
					}

					// Skip Angka Selanjutnya
					i++

					// Cari Ulang Posisi Digit
					posDigit = (lenAngka - 1) - i
					grpDigit = posDigit % 3
				}
			}
		}

		if grpDigit == 0 {
			grpUnits := posDigit / 3
			if i == 0 && intAngka == 1 && grpUnits <= 1 {
				// Cara Bar-Bar Hapus Konversi "Satu"
				// Dan Append Bentuk Baru :D
				resTerbilang = strings.TrimSuffix(resTerbilang, arrAngka[intAngka]+" ")
				resTerbilang += "se" + arrUnits[grpUnits] + " "
			} else {
				resTerbilang += arrUnits[grpUnits] + " "
			}
		}
	}

	// Trim Hasil Konversi dan Return Hasil
	resTerbilang = strings.TrimSpace(resTerbilang)
	return resTerbilang, nil
}
