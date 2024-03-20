package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	hobbies := [3]string{"Basketball", "Badminton", "Programming"}

	fmt.Println(hobbies)

	firstHobbies := hobbies[0]
	mixHobbies := hobbies[1:3]

	fmt.Println(firstHobbies, mixHobbies)

	fsHobbies1 := hobbies[0:2]
	fsHobbies2 := hobbies[:2]
	fmt.Println(fsHobbies1, fsHobbies2)

	fmt.Println(cap(fsHobbies1))
	//reslicing
	fsHobbies1 = fsHobbies1[1:3]
	fmt.Println(fsHobbies1)

	courseGoals := []string{"Learn Go!", "Learn all the basics"}
	fmt.Println(courseGoals)

	courseGoals[1] = "Learn all the details!"
	courseGoals = append(courseGoals, "Learn all the basics again!")

	fmt.Println(courseGoals)

	products := []Product{
		{"first-product", "A First Product", 12.8},
		{"second-product", "A second Product", 22.8},
	}

	fmt.Println(products)

	newProduct := Product{
		"third-product", "a third product", 15.99,
	}

	products = append(products, newProduct)

	fmt.Println(products)
}
