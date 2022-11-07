package main

import "fmt"

func main() {
	//if else
	/*
		x := 21
		if x%2 == 0 {
			fmt.Printf("%d is an even number\n", x)
		} else {
			fmt.Printf("%d is an odd number\n", x)
		}
	*/

	if x := 21; x%2 == 0 { // the scope of the variable 'x' is limited only to the if-else block
		fmt.Printf("%d is an even number\n", x)
	} else {
		fmt.Printf("%d is an odd number\n", x)
	}

	//switch case
	/*
		Rank by the score
		score 0 - 3 => "Terrible"
		score 4 - 7 => "Not Bad"
		score 8 - 9 => "Very Good"
		score 10 => "Excellent"
		otherwise => "Invalid score"
	*/
	//score := 6
	/*
		switch score := 6; score {
		case 0:
			fmt.Println("Terrible")
		case 1:
			fmt.Println("Terrible")
		case 2:
			fmt.Println("Terrible")
		case 3:
			fmt.Println("Terrible")
		case 4:
			fmt.Println("Not Bad")
		case 5:
			fmt.Println("Not Bad")
		case 6:
			fmt.Println("Not Bad")
		case 7:
			fmt.Println("Not Bad")
		case 8:
			fmt.Println("Very Good")
		case 9:
			fmt.Println("Very Good")
		case 10:
			fmt.Println("Excellent")
		default:
			fmt.Println("Invalid score")
		}
	*/

	switch score := 6; score {
	case 0, 1, 2, 3:
		fmt.Println("Terrible")
	case 4, 5, 6, 7:
		fmt.Println("Not Bad")
	case 8, 9:
		fmt.Println("Very Good")
	case 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Invalid score")
	}

	//mimicing if- else if
	switch x := 22; {
	case x%2 == 0:
		fmt.Printf("%d is an even number\n", x)
	case x%2 != 0:
		fmt.Printf("%d is an odd number\n", x)
	}

	//fallthrough
	switch x := 2; x {
	case 0:
		fmt.Println("x is zero")
		fallthrough
	case 1:
		fmt.Println("x is <= 1")
		fallthrough
	case 2:
		fmt.Println("x is <= 2")
		fallthrough
	case 3:
		fmt.Println("x is <= 3")
		fallthrough
	case 4:
		fmt.Println("x is <= 4")
		fallthrough
	case 5:
		fmt.Println("x is <= 5")
		fallthrough
	case 6:
		fmt.Println("x is <= 6")
		//fallthrough
	case 7:
		fmt.Println("x is <= 7")
		fallthrough
	case 8:
		fmt.Println("x is <= 8")
	}

	//fallthrough - use case
	switch plan := "pro"; plan {
	case "super":
		fmt.Println("All super features")
		fallthrough
	case "premium":
		fmt.Println("All premium features")
		fallthrough
	case "pro":
		fmt.Println("All pro features")
		fallthrough
	case "free":
		fmt.Println("All free features")
	}

	//for construct
	fmt.Println("for (v 1.0)")
	for no := 1; no <= 10; no++ {
		fmt.Println(no)
	}

	fmt.Println("for (v 2.0) (as a while loop)")
	x := 1
	for x < 100 {
		x += x
	}
	fmt.Printf("x = %d\n", x)

	fmt.Println("for (v 3.0) (as a infinite loop)")
	no := 1
	for {
		no += no
		if no > 100 {
			break
		}
	}
	fmt.Printf("no = %d\n", no)

	fmt.Println("for (using continue)")
	for no := 1; no <= 10; no++ {
		if no%2 == 0 {
			continue
		}
		fmt.Println(no)
	}

	fmt.Println("Using labels")
OUTER_LOOP:
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
			if i == j {
				fmt.Println("==================")
				continue OUTER_LOOP
			}
		}
	}
}
