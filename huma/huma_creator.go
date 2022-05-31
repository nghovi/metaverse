package huma

import "fmt"

type ConsoleCreator struct {
	Name     string
	Sex      int8
	Birthday string
}

func CreateHumaFromConsole() Huma {
	humaCreator := ConsoleCreator{}
	fmt.Println("Enter huma name:")
	fmt.Scan(&humaCreator.Name)
	fmt.Println("Huma's birthday:")
	fmt.Scan(&humaCreator.Birthday)
	fmt.Println("Huma's sex: (1: Male, 0: Gay, -1: Female)")
	fmt.Scan(&humaCreator.Sex)
	return humaCreator.Create()
}

func (consoleCreator ConsoleCreator) Create() Huma {
	return Huma{
		Name:     consoleCreator.Name,
		Sex:      consoleCreator.Sex,
		Birthday: consoleCreator.Birthday,
	}
}

type HumaCreator interface {
	Create() Huma
}
