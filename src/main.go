package main

func main() {

	// For loop with single condition
	for i := 0; i <= 10; i++ {
		println(i)
	}

	// For loop with multiple conditions
	for i, j := 0, 10; i <= 10 && j >= 0; i, j = i+1, j-1 {
		println(i, j)
	}

	// For while loop
	i := 0
	for i <= 10 {
		println(i)
		i++
	}

	// For loop with break
	for i := 0; i <= 10; i++ {
		if i == 5 {
			break
		}
		println(i)
	}

	// For loop with continue
	for i := 0; i <= 10; i++ {
		if i == 5 {
			continue
		}
		println(i)

	}

	// For loop with labels
Loop:
	for i := 0; i <= 10; i++ {
		for j := 0; j <= 10; j++ {
			if i == 5 && j == 5 {
				break Loop
			}
			println(i, j)
		}
	}

	// For loop with range
	for i, j := range []int{1, 2, 3, 4, 5} {
		println(i, j)
	}

	// For loop with range and _
	for _, j := range []int{1, 2, 3, 4, 5} {
		println(j)
	}

	// for forever loop
	for {
		println("forever")
	}
}
