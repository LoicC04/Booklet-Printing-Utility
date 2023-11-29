package main

import (
	"flag"
	"fmt"
	"log"
)

// main function.
// Computes the ordered pages to produce a booklet printing with A4 format, 2 pages per sheet and double-sided
func main() {
	nbPageArg := flag.Int("p", 4, "Number of pages")
	required := []string{"p"}
	flag.Parse()

	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	for _, req := range required {
		if !seen[req] {
			log.Fatalf("missing required -%s argument. See usage bookletprinting -h\n", req)
		}
	}
	if *nbPageArg <= 0 {
		log.Fatalf("Number of pages must be > 0")
	}

	log.Printf("You asked for a %v page document.\n", *nbPageArg)

	ordered := BookletOrderedPage(*nbPageArg)

	log.Printf("Printing a booklet ordered page : %v\n", ordered)
}

// BookletOrderedPage func.
// Computes the ordered pages
// nbPage is the number of pages of the document.
// returns a comma separated string with the ordered pages
func BookletOrderedPage(nbPage int) string {

	orderedPages := []int{}
	if nbPage%4 == 0 {
		orderedPages = bookletOrder(nbPage)
	} else {
		orderedPages = bookletOrder(nbPage + 4 - nbPage%4)
	}

	ordered := ""
	for i, v := range orderedPages {
		if i > 0 {
			ordered = fmt.Sprintf("%v,%v", ordered, v)
		} else {
			ordered = fmt.Sprintf("%v", v)
		}
	}

	return ordered
}

// bookletOrder internal func.
// Computes the ordered pages
// totalPages is the number of pages of the document.
// returns a comma separated string with the ordered pages (a multiple of 4)
func bookletOrder(totalPages int) []int {
	return bookletOrderDetailed(totalPages, totalPages, 1)
}

// bookletOrderDetailed internal func.
// Computes the ordered pages
// totalPages is the number of pages of the document.
// index is the current number of pages left to print
// idxSheet is the counter of sheets
// returns a comma separated string with the ordered pages
func bookletOrderDetailed(totalPages int, pages int, idxSheet int) []int {

	orderedPages := []int{pages, 2*idxSheet - 1, 2 * idxSheet, pages - 1}
	if idxSheet < totalPages/4 {
		orderedPages = append(orderedPages, bookletOrderDetailed(totalPages, pages-2, idxSheet+1)...)
	}

	return orderedPages
}
