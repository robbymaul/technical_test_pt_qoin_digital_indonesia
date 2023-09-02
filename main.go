package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var jumlahPemain, jumlahDadu int

	fmt.Print("Masukkan jumlah player: ")
	fmt.Scan(&jumlahPemain)
	fmt.Print("Masukkan jumlah dadu: ")
	fmt.Scan(&jumlahDadu)

	fmt.Println("================")
	bermainDadu(jumlahPemain, jumlahDadu)
}

// function untuk bermain dadu
func bermainDadu(jumlahPemain, jumlahDadu int) {
	var putaran = make([][]int, jumlahPemain)
	var skor = make([]int, jumlahPemain)
	var ronde = 1

	for i := 0; i < jumlahPemain; i++ {
		putaranDadu := jumlahDadu
		putaran[i] = make([]int, jumlahDadu)

		for j := 0; j < putaranDadu; j++ {
			putaran[i][j] = rand.Intn(6) + 1
		}
	}

	fmt.Printf("Pemain = %d, Dadu = %d\n", jumlahPemain, jumlahDadu)
	fmt.Println("=================")
	putaranAwal(ronde, putaran)

	for {
		putaranDadu := putaran
		poin := skor
		kesempatan := 0

		for i := range putaranDadu {
			angka := []int{}
			if kesempatan != 0 {

				for i := 1; 1 <= kesempatan; i++ {
					angka = append(angka, 1)
					kesempatan -= 1
				}

				for j := range putaranDadu[i] {
					if putaranDadu[i][j] == 1 {
						kesempatan += 1
						continue
					}
					if putaranDadu[i][j] == 6 {
						poin[i] += 1
					} else {
						angka = append(angka, putaranDadu[i][j])
					}
				}
			} else {
				for j := range putaranDadu[i] {
					if putaranDadu[i][j] == 1 {
						kesempatan += 1
						continue
					}
					if putaranDadu[i][j] == 6 {
						poin[i] += 1
					} else {
						angka = append(angka, putaranDadu[i][j])
					}
				}
			}
			putaranDadu[i] = angka
		}

		angka := []int{}
		if kesempatan != 0 {
			for i := range putaranDadu[0] {
				if putaranDadu[0][i] == 1 {
					kesempatan += 1
					continue
				}
				if putaranDadu[0][i] == 6 {
					poin[0] += 1
				} else {
					angka = append(angka, putaranDadu[0][i])
				}
			}

			for i := 1; 1 <= kesempatan; i++ {
				angka = append(angka, 1)
				kesempatan -= 1
			}

			putaranDadu[0] = angka
		}
		skor = poin

		fmt.Println()
		evaluasi(ronde, putaranDadu, poin)

		empty := 0
		for _, v := range putaranDadu {
			if len(v) > 0 {
				empty++
			}
		}
		if empty <= 1 {
			break
		}

		for i := range putaran {
			putaranDadu[i] = make([]int, len(putaranDadu[i]))

			for j := range putaranDadu[i] {
				putaranDadu[i][j] = rand.Intn(6) + 1
			}
		}
		ronde++
		fmt.Printf("\nGiliran %d lempar dadu: \n", ronde)
		putaranSetelahEvaluasi(ronde, putaranDadu, poin)

	}

	pemenang := cariPemenang(skor)

	fmt.Printf("Game dimenangkan oleh pemain %d karena memiliki poin lebih banyak dari pemain lainnya", pemenang)

}

// fungsi untuk mencari pemenang dari putaran dadu yang telah selesai
func cariPemenang(arr []int) int {
	var pemenang int
	max := arr[0]
	found := false

	for i := 1; i < len(arr); i++ {
		num := arr[i]
		if num > max {
			max = num
			pemenang = i + 1
			found = true
		}
	}

	if found {
		return pemenang
	} else {
		return 1
	}

}

// fungsi untuk output hasil  pada putaran awal
func putaranAwal(ronde int, pemain [][]int) {
	fmt.Printf("Giliran %d lempar dadu:\n", ronde)
	for i, pemain := range pemain {
		fmt.Printf("Pemain #%d (%d): ", i+1, ronde-1)
		for j, putaran := range pemain {
			if j > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%d", putaran)
		}
		fmt.Println()
	}
}

// fungsi untuk output hasil putaran dadu setelah  setelah evaluasi
func putaranSetelahEvaluasi(ronde int, pemain [][]int, poin []int) {
	for i, pemain := range pemain {
		fmt.Printf("Pemain #%d (%d): ", i+1, poin[i])
		for j, putaran := range pemain {
			if j > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%d", putaran)
		}
		fmt.Println()
	}
}

// fungsi untuk hasil output pada saat evaluasi dan mencari poin
func evaluasi(ronde int, putaran [][]int, poin []int) {
	fmt.Println("Setelah evaluasi:")
	for i := range putaran {
		fmt.Printf("Pemain #%d (%d): ", i+1, poin[i])
		for j, putaran := range putaran[i] {
			if j > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%d", putaran)
		}
		fmt.Println()
	}
}
