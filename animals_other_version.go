package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
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

	var isFound bool = false

	if len(animals) == 0 {
		return errors.New("no animals to select")
	}

Loop:
	for _, creature := range animals {

		if (*creature).GetName() != name {
			continue
		}

		isFound = true

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

	if !isFound {
		return errors.New("animal not found")
	}

	return nil
}

var animals []*Animal

func main() {

	var animal, command, name string
	// var input string
	var userInput []string
	var creature *Animal

Program:
	for {

		fmt.Println("Enter newanimal or query to get started. Type exit to quit")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue Program
		}
		userInput = strings.Fields(scanner.Text())

		// prompt(&input)

		if userInput[0] == "exit" {
			break Program
		} else if len(userInput) != 3 {
			fmt.Println("3 input strings are required")
			continue Program
		} else if userInput[0] == "newanimal" {
			animal = userInput[1]
			name = userInput[2]
			goto AnimalLoop
		} else if userInput[0] == "query" {
			name = userInput[1]
			command = userInput[2]
			goto CommandLoop
		} else {
			continue
		}

	AnimalLoop:
		for {

			// fmt.Println("Please Enter Cow, Bird, or Snake. Type exit to quit")
			// prompt(&animal)

			if animal == "exit" {
				break Program
			}

			aFactory, err := GetAnimalFactoryMethod(strings.ToLower(animal))

			if err != nil {
				fmt.Println(err)
				// continue AnimalLoop
				continue Program
			}

			// fmt.Println("Please enter name for animal. Type exit to quit")
			// prompt(&name)

			if name == "exit" {
				break Program
			}

			factoryCreature := aFactory.MakeAnimal()
			creature = &factoryCreature

			fmt.Println("Created Animal!")

			(*creature).SetName(name)
			animals = append(animals, creature)

			continue Program

		}

	CommandLoop:
		for {

			// fmt.Println("Please enter name for animal. Type exit to quit")
			// prompt(&name)

			if name == "exit" {
				break Program
			}

			// fmt.Println("Please Select Animal Command: Eat, Move, Speak. Type exit to quit")
			// prompt(&command)

			if command == "exit" {
				break Program
			}

			if err := Commands(command, name); err != nil {
				fmt.Println(err)
			}

			break CommandLoop

		}
	}

}
