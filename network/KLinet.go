package main



import (

	"bytes"

	"encoding/json"

	"fmt"

	"io"

	"log"

	"net/http"

	"strings"
	 
	"regexp"
	"container/list"
	

)



// Struktury do utworzenia obiektu JSON

type User struct {

	ID         string      `json:"ID"`

	Attributes []ParcelType `json:"ParcelType"`

}

type attribute struct {

	Name  string `json:"name"`

	Value string `json:"value"`

}





type ParcelType struct {

	ID           string `json:"id"`

	Destination  string `json:"destination"`

	Product_List string `json:"product_list"`

	Consignor    string `json:"consignor"`

	Localization string `json:"localization"`

	Track        string `json:"Track"`
	

}


type ParcelID struct {

	ID string `json:"ID"`

}

type Kurier struct {

	IDPaczki string `json:"ID"`

	IDKuriera string `json:"ID"`

}


func main() {

	// New list.
    values := list.New()
	printMenu()

	for {
	print("\n ====== Wykonaj operacje numer ======== ")	
	var wybor string
    fmt.Scanln(&wybor)
	print("\n")		
	if (wybor == "1"){
		print(" Obecne przesylki w rejestrze ")
		for temp := values.Front(); temp != nil; temp = temp.Next() {
			fmt.Println(temp.Value)
		}
	}else if (wybor == "2"){
		print("Podaj cel przesylki\n")
		var cel string
		fmt.Scanln(&cel)

		print("Podaj liste produktow\n")
		var Product_List string
		fmt.Scanln(&Product_List)

		print("Podaj nadawce przesylki\n")
		var Consignor string
		fmt.Scanln(&Consignor)

		var numerPrzesylki = RegisterParcel(cel , Product_List , Consignor)
		print("Nowo utworzona  przesylka ma numer = ",numerPrzesylki,"\n")
		values.PushFront(numerPrzesylki)
	}else if (wybor == "3"){
		print("Podaj numer przesylki na ktorej ma byc wykonana operacja\n")
		var numerPrzesylki string
		fmt.Scanln(&numerPrzesylki)
		var lokalizacjaPrzesylki = GetParcel(numerPrzesylki)
		print("Przesylka = ",lokalizacjaPrzesylki,"\n")
	}else if (wybor == "4"){
		print("Podaj numer przesylki na ktorej ma byc wykonana operacja\n")
		var numerPrzesylki string
		fmt.Scanln(&numerPrzesylki)
		var trasaPrzesylki = GetTrace(numerPrzesylki)
		print("Przesyka obecnie pokonala trase = ",trasaPrzesylki,"\n")
	}else if (wybor == "5"){
		print("Podaj numer przesylki na ktorej ma byc wykonana operacja\n")
		var numerPrzesylki string
		fmt.Scanln(&numerPrzesylki)
		var wynikSortownia2  = SortingO2(numerPrzesylki)
		print("Proba przeniesienia paczki do sortowni Org1MSP ", wynikSortownia2 ,"\n")
	}else if (wybor == "6"){
		print("Podaj numer przesylki na ktorej ma byc wykonana operacja\n")
		var numerPrzesylki string
		fmt.Scanln(&numerPrzesylki)
		var wynikSortownia1  = SortingO1(numerPrzesylki)
		print("Proba przeniesienia paczki do sortowni Org2MSP ", wynikSortownia1 ,"\n")
	}else if (wybor == "7"){
		print("Podaj numer przesylki na ktorej ma byc wykonana operacja\n")
		var numerPrzesylki string
		fmt.Scanln(&numerPrzesylki)
		var wynikBranch2  = BranchO1(numerPrzesylki)
		print("Proba przeniesienia paczki do odzialu  Org1MSP ", wynikBranch2 ,"\n")
	}else if (wybor == "8"){
		print("Podaj numer przesylki na ktorej ma byc wykonana operacja\n")
		var numerPrzesylki string
		fmt.Scanln(&numerPrzesylki)
		var wynikBranch1  = BranchO2(numerPrzesylki)
		print("Proba przeniesienia paczki do odzialu Org2MSP ", wynikBranch1 ,"\n")
	}else if (wybor == "9"){
		print("Podaj numer przesylki na ktorej ma byc wykonana operacja\n")
		var numerPrzesylki string
		fmt.Scanln(&numerPrzesylki)
		var numerKuriera string
		print("Podaj numer kuriera ktory ma dostarczyc przesylke")
		fmt.Scanln(&numerKuriera)
		var wynikKurier  = GiveToCourier(numerPrzesylki, numerKuriera)
		print("Proba dostarczenia przesylki ", wynikKurier ,"\n")
	}else if (wybor == "10"){
		print("Podaj numer przesylki na ktorej ma byc wykonana operacja\n")
		var numerPrzesylki string
		fmt.Scanln(&numerPrzesylki)
		var dostarczeniePrzesylki = Delivered(numerPrzesylki)
		print("Przesyka dotarla do celu = ",dostarczeniePrzesylki,"\n")
	}else if (wybor == "Exit"){
		break
	}else{
		printMenu()
	}

	}
	
		
	fmt.Println("Terminating the application...")

}

func printMenu(){
	print("=============Menu=============\n")
	print("1. Wypisz obecne przesyki w rejestrze\n")
	print("2. Dodaj nowa przesylke do rejestru\n")
	print("3. Pobierz doane o przesylce\n")
	print("4. Wypisz trase przesylki\n")
	print("5. Wyslij przesylke do sortowni firmy Org1MSP\n")
	print("6. Wyslij przesylke do sortowni firmy Org2MSP\n")
	print("7. Wyslij paczke do odzialu Org1MSP\n")
	print("8. Wyslij paczke do odzialu Org2MSP\n")
	print("9. Przekaz paczke kurierowi do dostarczenia\n")
	print("10. Zmien status przesylki na dostarczony\n")
	print("Exit. Wyjdz z rejestru\n")
}



func Delivered(numerPrzesylki string) string{
	//hx := hex.EncodeToString([]byte(numerPrzesylki))

	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	numerPrzesylki = re.ReplaceAllString(numerPrzesylki, "")

	jsonValue, _ := json.Marshal(numerPrzesylki)

	// Wyslanie danych do serwera i obsluga odpowiedzi -- Rjestracja przesyki

	response, err := http.Post("http://127.0.0.1:8080/Delivered", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

	} else {

		data, _ := io.ReadAll(response.Body)

		body := strings.Split(string(data), "\n")

		fmt.Println(body[0])
		
		return body[0]
		
}
return string("Blad")
}

func GiveToCourier(numerPrzesylki string, numerKuriera string) string{
	
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	numerPrzesylki = re.ReplaceAllString(numerPrzesylki, "")
	//UserObject := Kurier{numerPrzesylki, numerKuriera}
	//UserObject := Kurier{"fwfawf", "wafawf"}
	UserObject := ParcelType{numerPrzesylki,numerKuriera,"null"  ,"null","null","null"}
	jsonValue, _ := json.Marshal(UserObject)

	log.Printf(string(jsonValue))

	//print(UserObject.IDKuriera,"      ", UserObject.IDPaczki, "Koniec \n")
	// Wyslanie danych do serwera i obsluga odpowiedzi -- Rjestracja przesyki

	response, err := http.Post("http://127.0.0.1:8080/GiveToCourier", "application/json", bytes.NewBuffer(jsonValue))



	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

	} else {

		data, _ := io.ReadAll(response.Body)

		body := strings.Split(string(data), "\n")

		fmt.Println(body[0])
		
		return body[0]
		
}
return string("NULL")
}


func BranchO1(numerPrzesylki string) string{
	//hx := hex.EncodeToString([]byte(numerPrzesylki))

	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	numerPrzesylki = re.ReplaceAllString(numerPrzesylki, "")

	jsonValue, _ := json.Marshal(numerPrzesylki)

	// Wyslanie danych do serwera i obsluga odpowiedzi -- Rjestracja przesyki

	response, err := http.Post("http://127.0.0.1:8080/BranchO1", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

	} else {

		data, _ := io.ReadAll(response.Body)

		body := strings.Split(string(data), "\n")

		fmt.Println(body[0])
		
		return body[0]
		
}
return string("Blad")
}
func BranchO2(numerPrzesylki string) string{
	//hx := hex.EncodeToString([]byte(numerPrzesylki))

	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	numerPrzesylki = re.ReplaceAllString(numerPrzesylki, "")

	jsonValue, _ := json.Marshal(numerPrzesylki)

	// Wyslanie danych do serwera i obsluga odpowiedzi -- Rjestracja przesyki

	response, err := http.Post("http://127.0.0.1:8080/BranchO2", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

	} else {

		data, _ := io.ReadAll(response.Body)

		body := strings.Split(string(data), "\n")

		fmt.Println(body[0])
		
		return body[0]
		
}
return string("Blad")
}


func SortingO2(numerPrzesylki string) string{
	//hx := hex.EncodeToString([]byte(numerPrzesylki))

	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	numerPrzesylki = re.ReplaceAllString(numerPrzesylki, "")

	jsonValue, _ := json.Marshal(numerPrzesylki)

	// Wyslanie danych do serwera i obsluga odpowiedzi -- Rjestracja przesyki

	response, err := http.Post("http://127.0.0.1:8080/SortingO2", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

	} else {

		data, _ := io.ReadAll(response.Body)

		body := strings.Split(string(data), "\n")

		fmt.Println(body[0])
		
		return body[0]
		
}
return string("Blad")
}

func SortingO1(numerPrzesylki string) string{
	//hx := hex.EncodeToString([]byte(numerPrzesylki))

	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	numerPrzesylki = re.ReplaceAllString(numerPrzesylki, "")

	jsonValue, _ := json.Marshal(numerPrzesylki)

	// Wyslanie danych do serwera i obsluga odpowiedzi -- Rjestracja przesyki

	response, err := http.Post("http://127.0.0.1:8080/SortingO1", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

	} else {

		data, _ := io.ReadAll(response.Body)

		body := strings.Split(string(data), "\n")

		fmt.Println(body[0])
		
		return body[0]
		
}
return string("Blad")
}



func GetTrace(numerPrzesylki string) string{
	//hx := hex.EncodeToString([]byte(numerPrzesylki))

	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	numerPrzesylki = re.ReplaceAllString(numerPrzesylki, "")


	
	jsonValue, _ := json.Marshal(numerPrzesylki)




	// Wyslanie danych do serwera i obsluga odpowiedzi -- Rjestracja przesyki

	response, err := http.Post("http://127.0.0.1:8080/GetTrace", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

	} else {

		data, _ := io.ReadAll(response.Body)

		body := strings.Split(string(data), "\n")

		fmt.Println(body[0])
		
		return body[0]
		
}
return string("NULL")
}

func RegisterParcel(Destination string, Product_List string, Consignor string) string{
	//var attrib []ParcelType

	//attrib = append(attrib, ParcelType{"id",string(Destination),string(Product_List),string(Consignor),"localization","Track"})

	UserObject := ParcelType{"id",Destination, Product_List ,Consignor,"localization","Track"}

	jsonValue, _ := json.Marshal(UserObject)

	log.Printf(string(jsonValue))


	// Wyslanie danych do serwera i obsluga odpowiedzi -- Rjestracja przesyki

	response, err := http.Post("http://127.0.0.1:8080/RegisterParcel", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

	} else {

		data, _ := io.ReadAll(response.Body)

		body := strings.Split(string(data), "\n")

		fmt.Println(body[0])
		
		return body[0]
		
}
return string("NULL")
}

func GetParcel(numerPrzesylki string) string{
	//hx := hex.EncodeToString([]byte(numerPrzesylki))

	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	numerPrzesylki = re.ReplaceAllString(numerPrzesylki, "")


	
	jsonValue, _ := json.Marshal(numerPrzesylki)




	// Wyslanie danych do serwera i obsluga odpowiedzi -- Rjestracja przesyki

	response, err := http.Post("http://127.0.0.1:8080/GetParcel", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

	} else {

		data, _ := io.ReadAll(response.Body)

		body := strings.Split(string(data), "\n")

		fmt.Println(body[0])
		
		return body[0]
		
}
return string("NULL")
}