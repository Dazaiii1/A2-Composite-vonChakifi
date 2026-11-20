package main

import "testing"

func TestTxtDocSize(t *testing.T) {
	var doc Doc = TxtDoc{size: 10}
	if got := doc.Size(); got != 10 {
		t.Errorf("Size() = %d, erwartet 10", got)
	}
}

func TestEmptyFolderSize(t *testing.T) {
	var doc Doc = &Folder{}
	if got := doc.Size(); got != 0 {
		t.Errorf("Size() = %d, erwartet 0", got)
	}
}

func TestFolderSize(t *testing.T) {
	sub := &Folder{}
	sub.AddDoc(TxtDoc{size: 100})

	root := &Folder{}
	root.AddDoc(TxtDoc{size: 10})
	root.AddDoc(TxtDoc{size: 50})
	root.AddDoc(sub)

	if got := sub.Size(); got != 100 {
		t.Errorf("sub.Size() = %d, erwartet 100", got)
	}
	if got := root.Size(); got != 160 {
		t.Errorf("root.Size() = %d, erwartet 160", got)
	}
}

func TestNestedFolderSize(t *testing.T) {
	inner := &Folder{}
	inner.AddDoc(TxtDoc{size: 30})

	sub := &Folder{}
	sub.AddDoc(TxtDoc{size: 20})
	sub.AddDoc(inner)

	root := &Folder{}
	root.AddDoc(TxtDoc{size: 10})
	root.AddDoc(sub)

	if got := root.Size(); got != 60 {
		t.Errorf("root.Size() = %d, erwartet 60", got)
	}

	// Änderungen an einem Unterordner wirken sich auf die Gesamtgröße aus.
	inner.AddDoc(TxtDoc{size: 40})
	if got := root.Size(); got != 100 {
		t.Errorf("root.Size() = %d, erwartet 100", got)
	}
}
