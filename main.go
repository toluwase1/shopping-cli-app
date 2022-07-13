package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

//map of phones with quantity
// switch case for selection via numbers

const conferenceTickets int = 50

var shopName = "Classic iPhone"
var remainingTickets uint = 50

var shopItems = map[int]map[string]string{
	1: {"name": "iphone 6", "qty": "10"},
	2: {"name": "iphone 7", "qty": "10"},
	3: {"name": "iphone 8", "qty": "10"},
	4: {"name": "iphone X", "qty": "10"},
}
var shopArray = map[int][]string{
	1: {"iphone 6", "10"},
	2: {"iphone 7", "10"},
	3: {"iphone 8", "10"},
	4: {"iphone 10", "10"},
}

var bookings = make([]UserData, 0)

type UserData struct {
	firstName      string
	lastName       string
	email          string
	phoneId        uint
	numberOfPhones uint
}

var wg = sync.WaitGroup{}

func main() {

	shopHomePage()

	for {
		firstName, lastName, email, phoneId, qty := getUserInput()
		isValidName, isValidEmail, isValidPhoneId, isValidQuantity, isQtyRequestValid := validateUserInput(firstName, lastName, email, phoneId, qty)
		fmt.Println(isQtyRequestValid)
		if isValidName && isValidEmail && isValidQuantity && isQtyRequestValid {

			orderProduct(qty, phoneId, firstName, lastName, email)

			//wg.Add(1)
			//go sendTicket(userTickets, firstName, lastName, email)
			individualPhone := shopItems[int(phoneId)]
			phoneName := individualPhone[string(phoneId)]
			generateReceipt(firstName, lastName, email, phoneName, int(qty))
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidPhoneId {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidQuantity {
				fmt.Println("number of tickets you entered is invalid")
			}
			if !isQtyRequestValid {
				fmt.Println("requested quantity larger than current inventory balance ")
			}
		}
	}
	//wg.Wait()
}

func shopHomePage() {
	fmt.Printf("Welcome to %v Store\n", shopName)
	fmt.Printf("We have total of %v phones and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint, uint) {
	var firstName string
	var lastName string
	var email string
	var phoneId uint
	var quantity uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter the phone ID you would like to purchase: ")
	fmt.Scan(&phoneId)

	fmt.Println("Enter number of phones: ")
	fmt.Scan(&quantity)

	return firstName, lastName, email, phoneId, quantity
}

func inventory(shopItems map[int]map[string]string, productId int, qty int) map[int]map[string]string {
	innerValueOfMap, ok := shopItems[productId]
	if !ok {
		fmt.Println("error")
		return map[int]map[string]string{}
	}
	bal, _ := strconv.Atoi(innerValueOfMap["qty"])
	bal = bal - qty
	innerValueOfMap["qty"] = strconv.Itoa(bal)
	//shopItems[productId] = innerValueOfMap

	return shopItems

}
func productBalance(shopItems map[int]map[string]string) {

}

func orderProduct(quantity uint, phoneId uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - quantity
	inventoryBalance := inventory(shopItems, int(phoneId), int(quantity))
	var userData = UserData{
		firstName:      firstName,
		lastName:       lastName,
		email:          email,
		phoneId:        phoneId,
		numberOfPhones: quantity,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of current orders made %v\n", bookings)

	fmt.Printf("Thank you %v %v for ordering %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, quantity, email)
	fmt.Printf("Check the list of remaining products %v as follows %v\n", shopName, inventoryBalance)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}

func validateUserInput(firstName string, lastName string, email string, phoneId uint, qty uint) (bool, bool, bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := qty > 0 && qty <= remainingTickets
	isValidPhoneId := phoneId > 0 && phoneId <= 4
	individualPhone := shopItems[int(phoneId)]
	balance, _ := strconv.Atoi(individualPhone["qty"])
	isQtyRequestValid := balance >= int(qty)
	return isValidName, isValidEmail, isValidPhoneId, isValidTicketNumber, isQtyRequestValid
}

func generateReceipt(firstName string, lastName string, email string, product string, qty int) {
	file, err := os.Create("receipt.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	_, err2 := file.WriteString("firstName : " + firstName + "\n" + "lastName : " + lastName + "\n" +
		"email : " + email + "\n" + "product ordered: " + product + "\n" + "quantity ordered: " + string(rune(qty)))

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("receipt completely generated")
}
