package main

import (
	"fmt"

	"github.com/tkXreplica/cinema/movie"
	"github.com/tkXreplica/cinema/ticket"
)

func main() {
	movieName := movie.FindMovieName("tt0111161")
	fmt.Println(movieName)
	movie.Review(movieName, 9.3)
	ticket.BuyTicket(movieName)
}
