package myMath

import "testing"

/*The 
$>go test 
command will look for any tests in 
any of the files in the current folder and run 
them. Tests are identified by starting a function 
with the word Test and taking one argument of type 
*testing.T. In our case since we're testing the 
Average function we name the test function TestAverage.*/
func TestDivide (t *testing.T) {
	var v int
	v = Divide(15, 3)
	if v != 5 {
		t.Error("Expected 5, got ", v)
	}
}