package main

import (
	"fmt"
	"go-standard-library/media"
)

func main() {
	fmt.Println("My Favorite Movie")

	var myFav media.Catalogable = &media.Movie{}
	myFav.NewMovie("Top Gun", media.R, 43.2)
	fmt.Println(myFav)

	fmt.Printf("My favorite movie is %s.\n", myFav.GetTitle())
	fmt.Printf("It was rated %s.\n", myFav.GetRating())
	fmt.Printf("It made %f in the box office.\n", myFav.GetBoxOffice())

	//creates a new copy of the struct with the new title
	// myFav.Title = "Dumb and Dumber"

	myFav.SetTitle("Booger 2: Electric Boogerloo")
	fmt.Printf("My favorite movie is now %s\n", myFav.GetTitle())
}
