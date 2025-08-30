package main

func main() {

	type Human struct {
		Age    string
		Gender string
		Name   string
	}

	type Action struct {
		Human
		Run  string
		Talk string
		Feel string
	}
}
