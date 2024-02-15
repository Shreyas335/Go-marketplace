package main
import (
	"fmt"
	"os"
	"math/rand"
	"strconv"
	"strings"
)

var Names  = [...]string{"laptop", "phone", "water bottle", "peanuts", "hoodie", "charger", "paper", "pencil", "notebook", "water", "kettle"}

type Product struct {
    Name string
    Cost float64
	Quantity int
}

const FILE = "store.txt"

func main(){
	var input int 
	greeting()
	fmt.Scan(&input)
	for input != 3{
		switch input {
		case 1: 
			
			buyScreen()

		case 2:
			
			makeMarketplace()
		
		case 3:
			fmt.Println("Thanks")

		}
		clear()
		greeting()
		fmt.Scan(&input)

	}
}

func greeting(){
	fmt.Println("Welcome to the marketplace")
	fmt.Println("Press 1 to load the marketplace")
	fmt.Println("Press 2 to make a marketplace")
	fmt.Println("Press 3 to quit")
}

func makeMarketplace(){
	var numProducts uint
	fmt.Println("How many items do you want?")
	fmt.Scan(&numProducts)
	
	//random elements
	products := make([]Product, numProducts)
	for i  := 0; uint(i) < numProducts; i++{
		products[i] = Product{Names[rand.Int31n(int32((len(Names) - 1)))], rand.Float64() * 200.0, int(rand.Int31n(100))}

	}
	products = sortbyCostQuantityName(products, 1)

	//error checking
	err := os.WriteFile(FILE, write(products), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v", err)
		return
	}

	fmt.Println("Store generated")
}

func buyScreen(){
	content, err := os.ReadFile(FILE)
	//byte slice into string
	text := string(content)
	//error checking
	if err != nil {
		fmt.Println("File is empty. load first")
		return
	}

	if len(content) == 0 {
		fmt.Println("File is empty. load first")
		return 
	}
	products := make([]Product, 0, 10)

	fmt.Println("1) Sort by name")
	fmt.Println("2) Sort by price")
	fmt.Println("3) Sort by Stock")
	clear()
	index := 1
	tempName := ""
	var tempCost float64
	var tempQuant int64

	//making a product slice out of the text file
	for len(text) != 0{
		tempName = text[0:strings.Index(text, "|")]
		text = text[strings.Index(text, "|") + 1:]
		
		tempCost,_ = strconv.ParseFloat(text[0:strings.Index(text, "|")], 64)
		text = text[strings.Index(text, "|") + 1:]

		tempQuant,_ = strconv.ParseInt(text[0:strings.Index(text, "\n")],10, 64)
		text = text[strings.Index(text, "\n") + 1:]
		products = append(products, Product{tempName, tempCost, int(tempQuant)})
		index++
		
		
	}
	sort := -1
	fmt.Println("Sort by")
	fmt.Println("1) Cost")
	fmt.Println("2) Quantity")
	fmt.Println("3) Name")
	fmt.Scan(&sort)
	clear()
	
	//int matches sorting constraint
	products = sortbyCostQuantityName(products, sort)
	
	print(products)

	//product and quantity selection
	selection := 0
	quant := 0
	fmt.Println("Chose Product:")
	fmt.Scan(&selection)
	fmt.Println("Chose Quantity")
	fmt.Scan(&quant)
	fmt.Println("Bought!")
	products[selection - 1].Quantity = products[selection - 1].Quantity - quant

	//writing and error checking
	err2 := os.WriteFile(FILE, write(products), 0644)
	if err2 != nil {
		fmt.Printf("Error writing file: %v", err)
		return
	}


}

func clear(){
	for i := 0; i < 100; i++ {
		fmt.Print("\n")
	}
}

func write(products []Product) []byte{
	writeContent := ""

	for _,product := range products {
		writeContent += "" + product.Name + "|" + strconv.FormatFloat(product.Cost, 'f', 2, 64) + "|" + strconv.Itoa(product.Quantity) + "\n"
	}
	return []byte(writeContent)
}

func print(products []Product) {
	for i := 0; i < len(products); i++{
		fmt.Printf("%v) %v|%v|%v\n", i + 1, products[i].Name, products[i].Cost, products[i].Quantity)
	}
}

func sortbyCostQuantityName(slice []Product, sortType int) []Product{
	//insertion sort in ascending order

	for i := 0; i < len(slice); i++ {
		//sort by cost
		if sortType == 1{
			for j := i; j > 0 && slice[j-1].Cost > slice[j].Cost; j-- {
				slice[j], slice[j-1] = slice[j-1], slice[j]
			}

		//sort by quantity
		}else if sortType == 2{
			for j := i; j > 0 && slice[j-1].Quantity > slice[j].Quantity; j-- {
				slice[j], slice[j-1] = slice[j-1], slice[j]
			}

		//sort by name	
		}else if sortType == 3{
			
			for j := i; j > 0 && strings.Compare(slice[j - 1].Name, slice[j].Name) > 0; j-- {
				slice[j], slice[j-1] = slice[j-1], slice[j]
			}
		}
	}
	return slice
}
