package main

import "calc_service/app/r"

func main() {
	if err := r.Run(); err != nil {
		panic(err)
	}
}
