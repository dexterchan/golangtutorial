package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sort"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
)

//"golang.org/x/crypto/scrypt"
type person struct {
	First string
	Last  string
	Age   int
}

type Person struct {
	person
	Height      float64 `json:"height"`
	rightHanded bool
}

func testMarshal() string {
	pp1 := Person{

		person: person{
			"Cobi",
			"Bryant",
			40,
		},
		Height:      220,
		rightHanded: true, //Field must be capitalized to be marshalled
	}
	pp2 := Person{
		person: person{
			First: "Lebon",
			Last:  "James",
			Age:   35,
		},
		Height:      240,
		rightHanded: true, //Field must be capitalized to be marshalled
	}

	if bs, err := json.Marshal([]Person{pp1, pp2}); err == nil {
		fmt.Println(string(bs))
		return string(bs)
	} else {
		fmt.Println("error:", err)
		return "nothing"
	}

}

func testUnMarshal(input string) {
	bs := []byte(input)
	people := make([]Person, 0, 10)
	if err := json.Unmarshal(bs, &people); err == nil {
		fmt.Println("Unmarshal:", people)
	} else {
		fmt.Println("error:", err)
	}
}
func main() {
	//testMarshal()

	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	p1 = p1
	ret := testMarshal()
	testUnMarshal(ret)

	writeInterface()

	sortfunc()
	sortCustom()

	bcryptfunc()
	scruptfunc()
}

func writeInterface() {
	fmt.Fprintln(os.Stdout, "Hello, world")
	io.WriteString(os.Stdout, "Hello, world\n")
}

func sortfunc() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	sort.Ints(xi)
	fmt.Println(xi)
	sort.Strings(xs)
	fmt.Println(xs)
}

//to String function
func (p person) String() string {
	return fmt.Sprintf("%s %s :%d", p.First, p.Last, p.Age)
}

//Implements sort.Interface for []person
type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// Swap swaps the elements with indexes i and j.
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func sortCustom() {
	p1 := person{"James", "B", 32}
	p2 := person{"Moneypenny", "P", 27}
	p3 := person{"Q", "Q", 64}
	p4 := person{"M", "M", 56}

	people := []person{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
}

func sortCustom2() {
	p1 := Person{person: person{"James", "B", 32}, Height: 200, rightHanded: true}
	p2 := Person{person: person{"Moneypenny", "B", 32}, Height: 200, rightHanded: true}
	p3 := Person{person: person{"Q", "Q", 64}, Height: 170, rightHanded: true}
	p4 := Person{person: person{"M", "M", 56}, Height: 150, rightHanded: true}

	people := []Person{p1, p2, p3, p4}

	fmt.Println(people)
	//sort.Sort(ByAge(pp))
	fmt.Println(people)
}

func bcryptfunc() {
	pwd := `password123`
	b, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	if err == nil {
		fmt.Println("encrypted:", string(b))
	} else {
		fmt.Println(err)
	}
	inputpwd := "abcd"
	hashedPassword := b
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(inputpwd))

	fmt.Println("compare:", err)
}

func scruptfunc() {
	pwd := `password123`
	salt := []byte("testingscript")
	dk, err := scrypt.Key([]byte(pwd), salt, 1<<15, 8, 1, 32)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Scrypt", base64.StdEncoding.EncodeToString(dk))
}
