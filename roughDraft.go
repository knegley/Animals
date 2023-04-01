package main

import (
	"errors"
	"fmt"
	"strings"
)

type Animal interface {
	Eat() string
	Move() string
	Speak() string
	GetName() string
	SetName(string)
	SetFood(string)
	SetMovement(string)
	SetSound(string)
}

type Cow struct {
	Name     string
	Food     string
	Movement string
	Sound    string
}

type Bird struct {
	Name     string
	Food     string
	Movement string
	Sound    string
}

type Snake struct {
	Name     string
	Food     string
	Movement string
	Sound    string
}

//Cow getters and Setters

func (c *Cow) Eat() string {
	return c.Food
}

func (c *Cow) Move() string {
	return c.Movement
}

func (c *Cow) Speak() string {
	return c.Sound
}

func (c *Cow) GetName() string {
	return c.Name
}

func (c *Cow) SetName(name string) {
	c.Name = name
}

func (c *Cow) SetMovement(movement string) {
	c.Movement = movement
}

func (c *Cow) SetSound(sound string) {
	c.Sound = sound
}

func (c *Cow) SetFood(food string) {
	c.Food = food
}

//// Bird Getters and Setters

func (b *Bird) Eat() string {
	return b.Food
}

func (b *Bird) Move() string {
	return b.Movement
}

func (b *Bird) Speak() string {
	return b.Sound
}

func (b *Bird) GetName() string {
	return b.Name
}

func (b *Bird) SetName(name string) {
	b.Name = name
}

func (b *Bird) SetMovement(movement string) {
	b.Movement = movement
}

func (b *Bird) SetSound(sound string) {
	b.Sound = sound
}

func (b *Bird) SetFood(food string) {
	b.Food = food
}

/// Snake getters and setters

func (s *Snake) Eat() string {
	return s.Food
}

func (s *Snake) Move() string {
	return s.Movement
}

func (s *Snake) Speak() string {
	return s.Sound
}

func (s *Snake) GetName() string {
	return s.Name
}

func (s *Snake) SetName(name string) {
	s.Name = name
}

func (s *Snake) SetMovement(movement string) {
	s.Movement = movement
}

func (s *Snake) SetSound(sound string) {
	s.Sound = sound
}

func (s *Snake) SetFood(food string) {
	s.Food = food
}

// / fmt Helper
func prompt(input *string) {

	fmt.Print("> ")
	if _, err := fmt.Scan(input); err != nil {
		panic(err)
	}

}

type AnimalFactory interface {
	MakeAnimal() Animal
}

type CowFactory struct{}
type SnakeFactory struct{}
type BirdFactory struct{}

func (c *CowFactory) MakeAnimal() Animal {
	return &Cow{Name: "", Food: "grass", Movement: "walking", Sound: "moo"}
}

func (s *SnakeFactory) MakeAnimal() Animal {
	return &Snake{Name: "", Food: "mice", Movement: "slither", Sound: "hiss"}
}

func (b *BirdFactory) MakeAnimal() Animal {
	return &Bird{Name: "", Food: "worms", Movement: "fly", Sound: "peep"}
}

func GetAnimalFactoryMethod(animal string) (AnimalFactory, error) {
	switch strings.ToLower(animal) {
	case "snake":
		return &SnakeFactory{}, nil
	case "cow":
		return &CowFactory{}, nil
	case "bird":
		return &BirdFactory{}, nil
	default:

		return nil, errors.New("wrong animal selected")
	}

}

func Commands(command string, name string) error {

	if len(animals) == 0 {
		return errors.New("no animals to select")
	}
Loop:
	for _, creature := range animals {

		if (*creature).GetName() != name {
			continue
		}

		switch strings.ToLower(command) {
		case "eat":
			fmt.Println((*creature).GetName()+" eats", (*creature).Eat())
			break Loop
		case "move":
			fmt.Println((*creature).GetName()+" makes movement by", (*creature).Move())
			break Loop
		case "speak":
			fmt.Println((*creature).GetName()+" makes sound by", (*creature).Speak())
			break Loop
		default:
			return errors.New("wrong animal command selected")
		}

	}
	return nil
}

var animals []*Animal

func main() {

	var animal, command, name, input string
	var creature *Animal

Program:
	for {

		fmt.Println("Enter newanimal or query to get started. Type exit to quit")

		prompt(&input)

		if input == "exit" {
			break Program
		} else if input == "newanimal" {
			goto AnimalLoop
		} else if input == "query" {
			goto CommandLoop
		} else {
			continue
		}

	AnimalLoop:
		for {

			fmt.Println("Please Enter Cow, Bird, or Snake. Type exit to quit")
			prompt(&animal)

			if animal == "exit" {
				break Program
			}

			aFactory, err := GetAnimalFactoryMethod(strings.ToLower(animal))

			if err != nil {
				fmt.Println(err)
				continue AnimalLoop
			}

			fmt.Println("Please enter name for animal. Type exit to quit")
			prompt(&name)

			if name == "exit" {
				break Program
			}

			factoryCreature := aFactory.MakeAnimal()
			creature = &factoryCreature

			(*creature).SetName(name)
			animals = append(animals, creature)
			// break AnimalLoop
			continue Program

			// switch c := strings.ToLower(animal); c {
			// case "cow":
			// 	// creature = &Cow{}
			// 	// creature.SetFood("grass")
			// 	// creature.SetMovement("walk")
			// 	// creature.SetSound("moo")
			// 	// fmt.Println("created It")
			// 	aFactory, err := GetAnimalFactoryMethod(c)
			// 	if err != nil {
			// 		panic(err)
			// 	}
			// 	creature = aFactory.MakeAnimal()
			// 	fmt.Println(creature)

			// 	break AnimalLoop
			// case "bird":
			// 	creature = &Bird{}
			// 	creature.SetFood("worms")
			// 	creature.SetMovement("fly")
			// 	creature.SetSound("peep")
			// 	fmt.Println("created It")
			// 	break AnimalLoop
			// case "snake":
			// 	creature = &Snake{}
			// 	creature.SetFood("mice")
			// 	creature.SetMovement("slither")
			// 	creature.SetSound("hss")
			// 	fmt.Println("created It")
			// 	break AnimalLoop
			// default:
			// 	fmt.Println("Invalid Animal Entry")
			// 	continue AnimalLoop

			// }

		}

	CommandLoop:
		for {

			fmt.Println("Please enter name for animal. Type exit to quit")
			prompt(&name)

			if name == "exit" {
				break Program
			}

			fmt.Println("Please Select Animal Command: Eat, Move, Speak. Type exit to quit")
			prompt(&command)

			if command == "exit" {
				break Program
			}

			if err := Commands(command, name); err != nil {
				fmt.Println(err)
				continue CommandLoop
			}

			// switch strings.ToLower(command) {
			// case "eat":
			// 	fmt.Println((*creature).GetName()+" eats", (*creature).Eat())
			// 	break CommandLoop
			// case "move":
			// 	fmt.Println((*creature).GetName()+" makes movement by", (*creature).Move())
			// 	break CommandLoop
			// case "speak":
			// 	fmt.Println((*creature).GetName()+" makes sound by", (*creature).Speak())
			// 	break CommandLoop
			// default:
			// 	fmt.Println("Invalid Animal Command")
			// 	continue CommandLoop
			// }

			break CommandLoop

		}
	}

}
