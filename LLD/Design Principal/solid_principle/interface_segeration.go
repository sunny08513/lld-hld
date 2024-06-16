package solidprinciple

import "container/list"
//Suppose we have a Document interface with methods for reading, writing, and printing:


type Document interface {
	Read() string
	Write(content string)
	Print() string
}

type TextDocument struct {
	content string
	lruList *list.List
}

func (d *TextDocument) Read() string {
	return d.content
}

func (d *TextDocument) Write(content string) {
	d.content = content
}

func (d *TextDocument) Print() string {
	return "Printing: " + d.content
}

/*If we have a ReadOnlyDocument that should only support reading,
the current interface forces us to implement unnecessary methods:*/

type ReadOnlyDocument struct {
	content string
}

func (d *ReadOnlyDocument) Read() string {
	return d.content
}

func (d *ReadOnlyDocument) Write(content string) {
	// Not supported
}

func (d *ReadOnlyDocument) Print() string {
	// Not supported
	return ""
}

/*To follow the Interface Segregation Principle,
we can break the Document interface into smaller, more focused interfaces:*/

type Reader interface {
	Read() string
}

type Writer interface {
	Write(content string)
}

type Printer interface {
	Print() string
}

// Implement Reader, Writer, and Printer for TextDocument

/* Implement only Reader for ReadOnlyDocument
   Now, our ReadOnlyDocument only depends on the Reader interface, which aligns with the ISP.*/
