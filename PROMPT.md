# Aufgabe 2 – Prompt

Implementiere das Composite-Konstruktionsprinzip in Go am Beispiel einer Verzeichnishierarchie:

- `Doc`: Interface mit `Size() int`
- `TxtDoc`: Struct mit `size int`. `Size()` gibt die gespeicherte Größe zurück.
- `Folder`: Struct mit `docs []Doc`. Mit `AddDoc(d Doc)` können Docs hinzugefügt werden. `Size()` berechnet die Summe der Größen aller enthaltenen Docs. Verwende Pointer-Receiver für die Methoden von `Folder`.

Ein `Folder` kann `TxtDoc`-Instanzen und weitere `Folder`-Instanzen enthalten. Die Größenberechnung erfolgt polymorph über `Doc` ohne Typabfragen.

Erstelle in `main()` einen Folder mit zwei TxtDocs der Größen 10 und 50 und einem Unterordner mit einem TxtDoc der Größe 100. Gib die Gesamtgröße aus (erwartetes Ergebnis: 160).

Verwende nur die Go-Standardbibliothek und überprüfe die Implementierung mit `go test ./...`.