package main

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)

var messages []Message

func main() {
	//slice()

	//mapr()

	// Pointers are rarely used with Go's built-in types,
	// but they are useful when paired with structs
	// x := 42
	// fmt.Println(x)
    // pointer(&x)
	// fmt.Println(x)

	// daysTravelled := 28
	// avg := averageSpeedToMars(daysTravelled)
	// fmt.Println("Average speed is " + fmt.Sprintf("%.2f", avg))

	// newMessage("Steven", "hi")
	// newMessage("Steven", "bye")
	// outputMessages()

	// deleteMessageBySender("Steven")
	// outputMessages()

	// newMessage("Steven", "hi")
	// newMessage("Steven", "hi")
	// deleteMessageByMessageText("hi")
	// outputMessages()

	// a := []int{1, 2, 3, 4, 5}
	// result, err := joesImperativeChallenge(a, 3)

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result)
	// }

	theNumber := 3
	fmt.Println(guessTheNumber(theNumber))
}

func guessTheNumber(theNumber int) string {
	rand.Seed(time.Now().UnixNano())

	for {
			guess := rand.Intn(100) + 1

			switch {
			case guess == theNumber:
				return fmt.Sprintf("%v is the correct number", guess)
			case guess < theNumber:
				fmt.Printf("%v is too low\n", guess)
			default:
				fmt.Printf("%v is too high\n", guess)
			}
		}
}

func joesImperativeChallenge(nums []int, partitionSize int) ([]int, error) {
	partitioned, err := partition(nums, partitionSize)

	if err != nil {
		return nil, err
	} else {
		reduced := reduce(partitioned)
		return reduced, nil
	}
}

func partition(slice []int, partitionSize int) ([][]int, error) {
	if len(slice) < partitionSize {
		return nil, errors.New("Slice length must be at least as large as partition size")
	}

	a := make([][]int, 0)
	l := len(slice)
	counter := 0

	for {
		if counter + partitionSize <= l {
			a = append(a, slice[counter:partitionSize + counter])
		} else {
			break
		}

		counter++
	}

	return a, nil
}

func reduce(nums [][]int) []int {
	var accum int
	var x []int

	for _, v := range nums {
		for _, n := range v {
			accum += n
		}

		x = append(x, accum)
		accum = 0
	}

	return x
}

func slice() {
	s := make([]string, 0)
	s = append(s, "hello")
	s = append(s, "world")
	s = append(s, "!")

	x := len(s) - 2

	for i, item := range s {
		if i < x {
			print(item + " ")
		} else {
			print(item)
		}
	}
}

func mapr() {
	m := make(map[string]string)
	m["steven"] = "kent"
	m["kitty"] = "boo"

	for key, value := range m {
		println(key + " " + value)
	}
}

func pointer(xPtr *int) {
	*xPtr += 58
}

func averageSpeedToMars(daysTravelled int) float64 {
	distanceMars := 56_000_000
	return float64(distanceMars / (daysTravelled * 24))
}

type Message struct {
	sender string
	text   string
}

func (m *Message) toString() string {
  return m.text + " from " + m.sender
}

func newMessage(sender string, text string) Message {
	msg := Message{sender, text}
	addMessage(msg)
	return msg
}

func addMessage(msg Message) {
	if messages == nil {
		messages = make([]Message, 0)
	}

	messages = append(messages, msg)
}

func deleteMessageBySender(sender string) {
	tempMessages := make([]Message, 0)

	for _, msg := range messages {
		if msg.sender != sender {
			tempMessages = append(tempMessages, msg)
		}
	}

	messages = tempMessages
}

func deleteMessageByMessageText(messageText string) {
	tempMessages := make([]Message, 0)

	for _, msg := range messages {
		if msg.text != messageText {
			tempMessages = append(tempMessages, msg)
		}
	}

	messages = tempMessages
}

func outputMessages() {
	if len(messages) == 0 {
		fmt.Println("There are no messages")
	} else {
		for _, m := range messages {
			fmt.Println(m.toString())
		}
	}
}
