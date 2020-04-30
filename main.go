package main
import "fmt"

//Sepeda is
type Sepeda struct{
	ban int
	gigi int
	warna string
	jarak float64
}

//WaktuTempuh is
func (s *Sepeda) WaktuTempuh(jarak float64) float64 {
	waktu := jarak * 2.5
	return waktu
}

func main(){

	arr := make([]Sepeda, 5, 10)

	arr[0] = Sepeda{2, 1, "Merah", 30}
	arr[1] = Sepeda{3, 2, "Hitam", 10}
	arr[2] = Sepeda{2, 3, "Putih", 25}
	arr[3] = Sepeda{4, 4, "Pink", 5}
	arr[4] = Sepeda{2, 5, "Biru", 10}


	for index, element := range arr{
			fmt.Println("Sepeda ke    :", index +1)
			fmt.Println("Jumlah ban   :", element.ban)
			fmt.Println("Gigi         :", element.gigi)
			fmt.Println("Warna        :", element.warna)
			fmt.Println("Jarak tempuh :", element.WaktuTempuh(element.jarak), "Menit")
			fmt.Println("-----------------------------------")
	}
}