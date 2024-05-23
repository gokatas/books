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
)

type Book struct {
	Title string
	Authors
	Year int
}

type Authors []string

func (as Authors) String() string {
	return strings.Join(as, ", ")
}

func printBooks(books []Book) {
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	format := "%v\t%v\t%v\n"
	fmt.Fprintf(tw, format, "Title", "Authors", "Year")
	fmt.Fprintf(tw, format, "-----", "-------", "----")
	for _, book := range books {
		fmt.Fprintf(tw, format, book.Title, book.Authors, book.Year)
	}
	tw.Flush()
}

func main() {
	books := []Book{
		{"The Lord of The Rings", Authors{"Tolkien"}, 1954},
		{"The Go Programming Language", Authors{"Kernighan", "Donovan"}, 2015},
		{"The Phoenix Project", Authors{"Kim", "Behr", "Spafford"}, 2013},
	}
	sort.Slice(books, func(i, j int) bool {
		return books[i].Year < books[j].Year
	})
	printBooks(books)
}
