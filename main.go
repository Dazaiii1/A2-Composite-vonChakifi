package main

import "fmt"

func main() {
	sub := &Folder{}
	sub.AddDoc(TxtDoc{size: 100})

	root := &Folder{}
	root.AddDoc(TxtDoc{size: 10})
	root.AddDoc(TxtDoc{size: 50})
	root.AddDoc(sub)

	fmt.Println("Größe des Unterordners:", sub.Size())
	fmt.Println("Gesamtgröße:", root.Size())
}
