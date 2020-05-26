package main

import (
	"fmt"
)


type dataDump struct {
	dataLD map[string]string
	output string
	root bool
}

// -- Setters --
func (d *dataDump) SetDataLD(data map[string]string) {
	d.dataLD = data
}

func (d *dataDump) SetOutput(output string) {
	d.output = output
}

func (d *dataDump) SetRoot(root bool) {
	d.root = root
}
// -- Setters --

func main() {
	x := make(map[string]string)
	x["/root/root.txt"] = "Password123"


	model := dataDump{}
	model.SetDataLD(x)
	model.SetOutput("/tmp/out.tar")
	model.SetRoot(false)

	fmt.Println(model)
}

// The full program will stay private and not pushed here, due to the potential of scammed that way
// I will keep you updated on the program by sending you videos
// After everything has been completed, and our transaction has been fullfilled. I will push everything here.
