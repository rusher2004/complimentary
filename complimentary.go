// package complimentary reaffirms the user's self-worth by providing a complimentary message on initilization
package complimentary

import (
	"fmt"
	"math/rand/v2"
)

var compliments = []string{
	"You are a great person!",
	"Wow, your code is amazing!",
	"LGTM!",
	"You are handsome and everybody loves you!",
}

func init() {
	fmt.Println(rand.IntN(len(compliments)))
}
