package creational_patterns

import "fmt"

/* Provide an interface
for creating families of related or dependent objects
without specifying their concrete classes
*/

type Bear interface{}
type StoneBear struct{}
type IceBear struct{}

type Wolf interface{}
type StoneWolf struct{}
type IceWolf struct{}

type NaberiusFactory interface {
	CreateWolf() Wolf
	CreateBear() Bear
}

type ForestFactory struct{}

func (forest ForestFactory) CreateWolf() Wolf {
	fmt.Println("Creating a stone wolf...")
	return StoneWolf{}
}

func (forest ForestFactory) CreateBear() Bear {
	fmt.Println("Creating a stone bear...")
	return StoneBear{}
}

type TundraFactory struct{}

func (tundra TundraFactory) CreateWolf() Wolf {
	fmt.Println("Creating an ice wolf...")
	return IceWolf{}
}

func (tundra TundraFactory) CreateBear() Bear {
	fmt.Println("Creating an ice wolf...")
	return IceBear{}
}

func createAllMonster(naberius NaberiusFactory) {
	naberius.CreateWolf()
	naberius.CreateBear()
}

func AbstractFactoryExample() {
	fmt.Println("1 - Build Naberius monsters in forest")
	createAllMonster(ForestFactory{})

	fmt.Println("2 - Build Naberius monsters in tundra")
	createAllMonster(TundraFactory{})
}
