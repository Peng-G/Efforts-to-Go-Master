package go_basics

import (
	"bytes"
	"fmt"
	"io"
)

func Interfaces() {
	fmt.Println("Interfaces")
	fmt.Println("Use interfaces to define behavior, not data. Use structs to define data, not behavior.")

	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

	bwc := NewBufferWriterCloser()
	bwc.Write([]byte("Hello Go!"))
	bwc.Write([]byte("Only 8 bytes could be written at a time."))
	bwc.Close()

	var wc WriterCloser = NewBufferWriterCloser()
	wc.Write([]byte("Hello Go!"))
	wc.Close()

	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("wc does not implement io.Reader, conversion failed.")
	}

	// empty interface
	var myObj interface{} = NewBufferWriterCloser()
	// alligned with the interface type WriterCloser.
	if wc, ok := myObj.(WriterCloser); ok {
		wc.Write([]byte("Hello Go!"))
		wc.Close()
	}
	r, ok = myObj.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("wc does not implement io.Reader, conversion failed.")
	}

	// type switch
	var v interface{} = 0
	switch v.(type) {
	case int:
		fmt.Println("v is an int")
	case string:
		fmt.Println("v is a string")
	default:
		fmt.Println("v is another type")
	}
	
}

// Interfaces are named collections of method signatures.
// if the interface only realize one function, then name the interface with the func_name + 'er'.
type Writer interface {
	Write([]byte) (int, error)
}
// ConsoleWriter implements Writer
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
// this is called polymorphism, because the same method can be called on different types. We can have different types that implement the same interface.



type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

// compose interfaces together
type Closer interface {
	Close() error
}
// CloserWriter is a composed interface, which is a combination of two interfaces.
type WriterCloser interface {
	Writer
	Closer
}

// BufferWriterCloser implements WriterCloser with two functions.
type BufferWriterCloser struct {
	buffer *bytes.Buffer
}
// Write data to buffer.
func (bwc *BufferWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}
	v := make([]byte, 8)
	for bwc.buffer.Len() >= 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}
// Close implements Closer interface.
func (bwc *BufferWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		fmt.Println(string(data))
	}
	return nil
}

func NewBufferWriterCloser() *BufferWriterCloser {
	return &BufferWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}