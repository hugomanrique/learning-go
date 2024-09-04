package main

import "fmt"

func main() {
	structs()
	// var matriz = [...]int{10,20,30,40} // when no know the length of array, use this

	// for i:=0; i<len(matriz); i++ {
	// 	matriz[i] = matriz[i] * 2
	// }

	// for index,value := range matriz {
	// 	fmt.Println(index,value)
	// }
	// fmt.Println(matriz)

	// var matriz2 = [3][3]int{{1,2,3},{4,5,6},{7,8,9}}
	// for _,value := range matriz2 {
	// 	for _,v := range value {
	// 		fmt.Println(v)
	// 	}
	// }
}

func slices() {
	// diasSemana := []string{"Domingo", "Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado"}
	// fmt.Println(diasSemana)

	// diaSlice := diasSemana[0:5]
	// fmt.Println(diaSlice)
	// diaSlice = append(diaSlice, "Viernes")

	// fmt.Println(len(diaSlice))
	// fmt.Println(cap(diaSlice))
	// fmt.Println(diaSlice)

	// diaSlice = append(diaSlice[:2], diaSlice[3:]...)
	// fmt.Println(diaSlice)

	nombres := make([]string, 5, 10)
	nombres[0] = "Juan"
	nombres[1] = "Pedro"
	nombres[2] = "Carlos"
	nombres[3] = "Luis"
	nombres[4] = "Maria"
	fmt.Println(nombres)
}

func maps(){
	var colors = map[string]string{
		"red":   "ff0000",
		"green": "00ff00",
		"blue":  "0000ff",
	}
	colors["black"]	= "000000"
	fmt.Println(colors)
	if valor, ok := colors["white"]; ok {
		fmt.Println(valor)
	}else{
		fmt.Println("not found")
	}

	delete(colors, "black")
	fmt.Println(colors)

	for key,value := range colors {
		fmt.Println(key,value)
	}
}

type Person struct {
	Name string
	Age int
	Email string
}

func structs(){
	
}