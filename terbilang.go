package angka

import (
	"strconv"
	"strings"
)

// Definisi Array Angka dan Satuan
var arrAngka = [...]string{"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan"}
var arrSatuan = [...]string{"", "ribu", "juta", "milyar", "triliun", "quadriliun", "quintiliun", "sextiliun", "septiliun", "oktiliun", "noniliun", "desiliun",
	"undesiliun", "duodesiliun", "tredesiliun", "quattuordesiliun", "quindesiliun", "sexdesiliun", "septendesiliun", "oktodesiliun", "novemdesiliun", "vigintiliun"}

// ToTerbilang Function
func ToTerbilang(angka string) (string, error) {
	// Definisi Variabel Hasil Konversi Terbilang
	var resTerbilang string

	// Trim Inputan Angka
	// dan Cari Panjang Angka String
	strAngka := strings.TrimSpace(angka)
	lenAngka := len(strAngka) - 1

	// Loop Angka String dan Konversi
	for i := 0; i <= lenAngka; i++ {
		// Cari Posisi Digit
		posDigit := lenAngka - i
		grpDigit := posDigit % 3

		// Konversi Angka String ke Angka Int
		intAngka, err := strconv.Atoi(string(strAngka[i]))
		if err != nil {
			return "", err
		}

		// Konversi Angka ke Bilangan
		switch intAngka {
		case 1:
			switch grpDigit {
			case 2:
				// Proses Seratus
				resTerbilang += "seratus "

			case 1:
				// Konversi Angka String Selanjutnya ke Angka Int Selanjutnya
				nextIntAngka, err := strconv.Atoi(string(strAngka[i+1]))
				if err != nil {
					return "", err
				}

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
				posDigit--
				grpDigit--

			case 0:
				if i == lenAngka || lenAngka > 4 {
					// Proses Satu
					resTerbilang += arrAngka[intAngka] + " "
				} else {
					// Proses Seribu
					resTerbilang += "seribu "
					continue
				}
			}

		case 0:
			if i == lenAngka && lenAngka == 0 {
				// Proses Nol
				return "nol", nil
			}
			continue

		default:
			// Proses Angka
			resTerbilang += arrAngka[intAngka] + " "

			switch grpDigit {
			case 2:
				// Proses Ratusan
				resTerbilang += "ratus "

			case 1:
				// Proses Puluhan
				resTerbilang += "puluh "
			}
		}

		// Proses Satuan
		if grpDigit == 0 {
			resTerbilang += arrSatuan[posDigit/3] + " "
		}
	}

	// Trim Hasil Konversi dan Return
	return strings.TrimSpace(resTerbilang), nil
}
