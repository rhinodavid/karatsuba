package main

import "testing"

// adding
func TestAdd(t *testing.T) {
	nums1 := []string{"10", "35", "3567", "245673", "93478590"}
	nums2 := []string{"24", "256", "2567", "24532", "2345256"}
	sums := []string{"34", "291", "6134", "270205", "95823846"}
	for i, sum := range sums {
		addedSum := add(nums1[i], nums2[i])
		if sum != addedSum {
			t.Fatalf("Incorrect sum, expected %s plus %s to equal %s. Got %s", nums1[i], nums2[i], sum, addedSum)
		}
	}
}

// multiplying
func TestMultiply(t *testing.T) {
	num1 := "3141592653589793238462643383279502884197169399375105820974944592"
	num2 := "2718281828459045235360287471352662497757247093699959574966967627"
	ans := "8539734222673567065463550869546574495034888535765114961879601127067743044893204848617875072216249073013374895871952806582723184"
	if multiply(num1, num2) != ans {
		t.Fatal("Multiplication incorrect")
	}
}
