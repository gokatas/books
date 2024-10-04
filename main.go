// Books sorts and prints a collection of books.
// For more see
//   - https://pkg.go.dev/sort#pkg-examples
//   - https://github.com/adonovan/gopl.io/blob/master/ch7/sorting
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
	"time"
)

type Book struct {
	Title string
	Authors
	Age int
}

type Authors []string

func (as Authors) String() string {
	return strings.Join(as, ", ")
}

func age(yearBorn int) (yearsOld int) {
	yearNow := time.Now().Year()
	return yearNow - yearBorn
}

func printBooks(books []Book) {
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	format := "%v\t%v\t%v\n"
	fmt.Fprintf(tw, format, "Age", "Title", "Authors")
	fmt.Fprintf(tw, format, "---", "-----", "-------")
	for _, book := range books {
		fmt.Fprintf(tw, format, book.Age, book.Title, book.Authors)
	}
	tw.Flush()
}

func main() {
	books := []Book{
		{"The Lord of The Rings", Authors{"Tolkien"}, age(1954)},
		{"The Go Programming Language", Authors{"Kernighan", "Donovan"}, age(2015)},
		{"The Phoenix Project", Authors{"Kim", "Behr", "Spafford"}, age(2013)},
	}
	sort.Slice(books, func(i, j int) bool {
		return books[i].Age < books[j].Age
	})
	printBooks(books)
}
