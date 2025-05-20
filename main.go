package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
	//"github.com/jackc/pgx/v5"
)

type User struct {
	Id               uint32
	FirstName        string
	LastName         string
	NikName          string
	Age              uint8
	Status           uint8
	Level            uint8
	Raiting          uint
	RegistrationTime time.Time
}

// Users functions
// =============================================================================
func (u *User) getAllInfo() string {
	outputString := fmt.Sprintf("id: %d\nFirstname: %s\n", u.Id, u.FirstName) +
		fmt.Sprintf("Lastname: %s\nNikname: %s\n", u.LastName, u.NikName) +
		fmt.Sprintf("Age: %d\nStatus: %d\n", u.Age, u.Status) +
		fmt.Sprintf("Level: %d\nRaiting: %d\n", u.Level, u.Raiting) +
		fmt.Sprintf("Registration time: %s", u.RegistrationTime.Format("2006-01-02\t15:04:05"))
	return outputString
}

func (u *User) setNewFirstName(newFirstName string) {
	u.FirstName = newFirstName
}

func usersPage(w http.ResponseWriter, r *http.Request) {
	admin := User{1, "Pavel", "Gasparyan", "Piligrim", 40, 8, 0, 0, time.Now()}
	admin.setNewFirstName("Pavel Grigorevich")
	fmt.Fprintf(w, admin.getAllInfo())
}

//=============================================================================

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to IT-dojo!</h1>")
	//fmt.Fprintf(w, `<h2>Main text</h2>
	//    							<h3>Choose a programming language</h3>`)
	admin := User{1, "Pavel", "Gasparyan", "Piligrim", 40, 8, 0, 0, time.Now()}
	tmpl := template.Must(template.ParseFiles("templates/homePage.html"))
	tmpl.Execute(w, admin)
}

func contactsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts/", contactsPage)
	http.HandleFunc("/users/", usersPage)
	http.ListenAndServe(":8080", nil)
}

func main() {
	/*
		conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
		defer conn.Close(context.Background())

		var name string
		var weight int64
		err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(name, weight)
	*/
	handleRequest()
}
