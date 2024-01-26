package main

import "fmt"

func main() {
	websites := map[string]string{
		"Amazon Web Services": "https://www.aws.com",
		"Google":              "https://www.google.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])

	websites["LinkedIn"] = "https://linkedin.com"

	fmt.Println(websites)

	// deleting keys
	delete(websites, "Google")

	fmt.Println(websites)
}
