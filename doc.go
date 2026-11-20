package main

// Doc ist das gemeinsame Interface für TxtDoc und Folder.
type Doc interface {
	Size() int
}

// TxtDoc ist ein Blatt der Verzeichnishierarchie.
type TxtDoc struct {
	size int
}

func (d TxtDoc) Size() int {
	return d.size
}

// Folder enthält Docs, die wiederum Folder sein können.
type Folder struct {
	docs []Doc
}

func (f *Folder) AddDoc(d Doc) {
	f.docs = append(f.docs, d)
}

func (f *Folder) Size() int {
	size := 0
	for _, doc := range f.docs {
		size += doc.Size()
	}
	return size
}
