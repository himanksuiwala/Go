package main

import "fmt"

func main() {
	fmt.Println("Maps in GOlang")

	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["RB"] = "RubyOnRails"
	languages["PY"] = "Python"

	fmt.Println(languages)
	fmt.Println(languages["JS"])
	delete(languages, "RB")

	fmt.Println(languages)

	for key, val := range languages {
		fmt.Printf("Key %v has value %v\n", key, val)

	}

}
