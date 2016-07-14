package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

type Node interface{} // CharData or *Element

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func NewTree(dec *xml.Decoder) (Node, error) {
	var stack []*Element // stack of elements
	var result *Element
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			newElem := &Element{tok.Name, tok.Attr, []Node{}}
			if len(stack) > 0 {
				elem := stack[len(stack)-1]
				elem.Children = append(elem.Children, newElem)
			}
			stack = append(stack, newElem)
			if result == nil {
				result = newElem
			}
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if len(stack) == 0 {
				continue
			}
			elem := stack[len(stack)-1]
			elem.Children = append(elem.Children, CharData(tok))
		}
	}
	return result, nil
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	tree, err := NewTree(dec)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", tree)
}
