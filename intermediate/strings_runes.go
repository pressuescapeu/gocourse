package intermediate

import "fmt"

func strings_runes() {
	// strings are immutable
	message := "string \n"
	another_message := "string \tsee that's tab"
	// \r means it goes to the start of the string and starts rewriting
	// so the following would be: in the way mskdjskdjsdger and longe
	another_message_r := "something but I go longer and longer \rin the way mskdjskdjsd"

	// can use backticks for raw string literals
	// raw means that \n or \t idk won't work as intended
	// and would work as literally \n rather than newline character
	raw := `raw string \n`

	fmt.Println(message)
	fmt.Println(another_message_r)
	fmt.Println(another_message)
	fmt.Println(raw)
	// strings have length bc they're basically array of runes
	fmt.Println(len(another_message_r))
	// also indexing through
	// in this case, it returns an uint8 type, so here s = 115 in ASCII
	fmt.Println(message[0])

	// string concatenation
	first_message := "first message"
	second_message := "second message"
	fmt.Println(first_message + second_message)
	// SINGLE CODE ONLY FOR RUNES
	var ch rune = 's'
	rch := 'щ'
	// will print out numbers
	fmt.Println(ch)
	fmt.Println(rch)

	// but in the formatted print
	fmt.Printf("%c\n", rch) // prints out regular
	converted_string := string(ch)
	fmt.Printf("type of the variable is %T\n", converted_string)

	const RUSSISH = "я покушал"
	fmt.Println(RUSSISH)

	for _, runeValue := range RUSSISH {
		fmt.Printf("%c\n", runeValue)
	}

	// also support for smileys
}
