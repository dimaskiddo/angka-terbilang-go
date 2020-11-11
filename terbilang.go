package angka

import (
	"errors"
	"strconv"
)

// Definisi Array Angka dan Satuan
var arrAngka = [...]string{"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan"}
var arrSatuan = [...]string{"", "ribu", "juta", "milyar", "triliun", "quadriliun", "quintiliun", "sextiliun", "septiliun", "oktiliun", "noniliun", "desiliun",
	"undesiliun", "duodesiliun", "tredesiliun", "quattuordesiliun", "quindesiliun", "sexdesiliun", "septendesiliun", "oktodesiliun", "novemdesiliun", "vigintiliun"}

// ToTerbilang Function
func ToTerbilang(strAngka string) (string, error) {
	// Jika Inputan Bukan Angka Maka Return Error
	if _, err := strconv.ParseFloat(strAngka, 10); err != nil {
		return "", errors.New("Error, input is not a valid number!")
	}

	// Cari Panjang Angka String
	lenAngka := len(strAngka) - 1

	// Konversi Angka String ke Nol
	intAngka, err := strconv.Atoi(string(strAngka[0]))
	if err != nil {
		return "", err
	}

	// Jika Panjang Angka Nol dan Angka Pertama adalah Nol Maka Proses Nol
	if lenAngka == 0 && intAngka == 0 {
		return "nol", nil
	}

	// Jika Angka Over dari Satuan Maka Return Error
	if (lenAngka / 3) > len(arrSatuan) {
		return "", errors.New("Error, number is to big!")
	}

	// Definisi Variabel Hasil Konversi Terbilang
	resTerbilang := ""

	// Set Counter Nol
	cntZero := 0

	// Loop Angka String dan Konversi
	for i := 0; i <= lenAngka; i++ {
		// Set Variable Sementara Hasil Konversi
		tmpTerbilang := ""

		// Cari Posisi Digit
		posDigit := lenAngka - i
		grpDigit := posDigit % 3

		// Konversi Angka String ke Angka Int
		intAngka, err = strconv.Atoi(string(strAngka[i]))
		if err != nil {
			return "", err
		}

		// Konversi Angka ke Bilangan
		switch intAngka {
		case 1:
			switch grpDigit {
			case 2:
				// Proses Seratus
				tmpTerbilang += "seratus"

			case 1:
				// Konversi Angka String Selanjutnya ke Angka Int Selanjutnya
				nextIntAngka, err := strconv.Atoi(string(strAngka[i+1]))
				if err != nil {
					return "", err
				}

				switch nextIntAngka {
				case 1:
					// Proses Sebelas
					tmpTerbilang += "sebelas"

				case 0:
					// Proses Sepuluh
					tmpTerbilang += "sepuluh"

				default:
					// Proses Belasan
					tmpTerbilang += arrAngka[nextIntAngka] + " belas"
				}

				// Skip Angka Selanjutnya
				i++

				// Cari Ulang Posisi Digit
				posDigit--
				grpDigit--

			case 0:
				if (intAngka == 1 && posDigit == 3) && (cntZero == 2 || lenAngka == 3) {
					// Tambah Spasi
					if resTerbilang != "" {
						resTerbilang += " "
					}

					// Proses Seribu
					resTerbilang += "seribu"

					// Reset Penghitung Nol
					cntZero = 0
					continue
				} else {
					// Proses Satu
					tmpTerbilang += arrAngka[intAngka]
				}
			}

		case 0:
			// Hitung Nol
			cntZero++
			break

		default:
			// Proses Angka
			tmpTerbilang += arrAngka[intAngka]

			switch grpDigit {
			case 2:
				// Proses Ratusan
				tmpTerbilang += " ratus"

			case 1:
				// Proses Puluhan
				tmpTerbilang += " puluh"
			}
		}

		// Prepand Spasi
		if tmpTerbilang != "" {
			// Tambah Spasi
			if resTerbilang != "" {
				resTerbilang += " " + tmpTerbilang
			} else {
				resTerbilang += tmpTerbilang
			}
		}

		// Cari Posisi Satuan
		posSatuan := posDigit / 3

		// Konversi Satuan
		if grpDigit == 0 && posSatuan != 0 {
			if cntZero != 3 {
				// Proses Satuan
				resTerbilang += " " + arrSatuan[posSatuan]
			}

			// Reset Pneghitung Nol
			cntZero = 0
		}
	}

	// Trim Hasil Konversi dan Return
	return resTerbilang, nil
}
