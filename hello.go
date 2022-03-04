package main

import "fmt"

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

	newMessage("Steven", "hi")
	newMessage("Steven", "bye")
	outputMessages()
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

func outputMessages() {
	for _, m := range messages {
		fmt.Println(m.toString())
	}
}
