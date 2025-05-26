package main

import "fmt"

//https://leetcode.com/problems/powx-n/description/
func main() {

	//how to calculate power x of power y
	//run loop
	//2nd way is to get the binary digits

	x := 3
	y := 6

	fmt.Println(calculatePower(x, y))
	fmt.Println(calculatePower(2.00000, -2))
}

func calculatePower(x int, power int) int {
	fmt.Println("calcualte value of ", x, power)
	if power == 0 {
		return 1
	}
	res := 1
	for power > 0 {
		fmt.Println("current value of power:", power, res)
		//step 1 : if x & 1 or 0, will give last element just multipley if its one
		if (power & 1) == 1 {
			res = res * x
		}
		//step 2 - every time loop runs keep increasing it like to 3, 9, 27
		x = x * x
		//step 3 is right shift with 1, that means dividing with 10 means, remocing the last digit
		power = power >> 1
	}
	return res
}

func calculatePowerNegativeAndFloat(x float64, n int) float64 {
	res := 1.0
	if n < 0 {
		//we are looking at negative power
		// x power -n = 1/x power n
		x = 1 / x
		n = -n

	}

	for n > 0 {
		//step 1 calculate the digit
		if n&1 == 1 {
			res = res * x
		}
		// step 2 multiply the x with itself
		x = x * x
		// step 3 is right shift to remove the last binary digit
		n = n >> 1
	}
	return res

}
