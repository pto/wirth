package main

func Example() {
	main()
	// Output:
	// .5
	// .25
	// .125
	// .0625
	// .03125
	// .015625
	// .0078125
	// .00390625
	// .001953125
	// .0009765625
}

func ExamplePower_parm1() {
	Power(1)
	// Output:
	// .5
}

func ExamplePower_parm2() {
	Power(2)
	// Output:
	// .5
	// .25
}

func ExamplePower_parm5() {
	Power(5)
	// Output:
	// .5
	// .25
	// .125
	// .0625
	// .03125
}
func ExamplePower_parm11() {
	Power(11)
	// Output:
	// .5
	// .25
	// .125
	// .0625
	// .03125
	// .015625
	// .0078125
	// .00390625
	// .001953125
	// .0009765625
	// .00048828125
}
