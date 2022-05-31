package planet

import "fmt"

type Planet struct {
	Name string
	Area int64
}

func (planet Planet) SelfIntroduce() {
	fmt.Println("Hi, I'm a planet. My name is ", planet.Name, "with area of", planet.Area, "km2")
}
