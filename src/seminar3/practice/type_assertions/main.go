package main

import "fmt"

type Writer interface {
	Write([]byte) (int, error)
}

type FileWriter struct{}

func (fw FileWriter) Write(data []byte) (int, error) {
	// Write to file
	return len(data), nil
}

type NetworkWriter struct{}

func (nw NetworkWriter) Write(data []byte) (int, error) {
	// Send over network
	return len(data), nil
}

type LogWriter struct{}

func (lw LogWriter) Write(data []byte) (int, error) {
	fmt.Printf("LOG: %s\n", string(data))
	return len(data), nil
}

func process(w Writer) {
	if lw, ok := w.(LogWriter); ok {
		fmt.Println("It's a LogWriter! Let's add a timestamp:")
		_, err := lw.Write([]byte("with timestamp"))
		if err != nil {
			return
		}
	} else {
		fmt.Println("It's some other kind of Writer")
		_, err := w.Write([]byte("normal write"))
		if err != nil {
			return
		}
	}
}

type Speaker interface {
	Speak() string
}

type Dog struct{ Name string }

func (d Dog) Speak() string {
	return "Woof! I'm " + d.Name
}

func MakeItTalk(s Speaker) {
	fmt.Println(s.Speak())
}

func TypeSwitch(x interface{}) {
	switch v := x.(type) {
	case int:
		fmt.Printf("Twice %d is %d\n", v, v*2)
	case string:
		fmt.Printf("%q is %d bytes long\n", v, len(v))
	case Dog:
		fmt.Printf("Dog %s says: %s\n", v.Name, v.Speak())
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	var w Writer = FileWriter{}
	w = NetworkWriter{}
	w = LogWriter{}

	process(FileWriter{})
	process(w)

	TypeSwitch(21)
	TypeSwitch("hello")
	TypeSwitch(Dog{Name: "Rex"})
	TypeSwitch(3.14)

	var w2 Writer // nil значение интерфейса
	fmt.Println(w2 == nil)

	var lw *LogWriter = nil // nil - конкретное значение
	w2 = lw                 // в w2 содержится (nil, *LogWriter)
	fmt.Println(w2 == nil)
}
