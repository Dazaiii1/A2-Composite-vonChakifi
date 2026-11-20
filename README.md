# Aufgabe 2 – Composite-Konstruktionsprinzip

Implementierung des Composite-Konstruktionsprinzips in Go am Beispiel einer Verzeichnishierarchie (Folder-Docs).

Der zur Implementierung verwendete Prompt steht in [PROMPT.md](PROMPT.md).

## Beispiel

```text
root (Folder)
├── TxtDoc (10)
├── TxtDoc (50)
└── sub (Folder)
    └── TxtDoc (100)
```

`TxtDoc` und `*Folder` erfüllen das Interface `Doc`. `Folder.Size()` summiert die Größen der enthaltenen Docs über das Interface. Die indirekte Rekursion ergibt sich aus der verschachtelten Baumstruktur. Das Beispiel setzt eine zyklenfreie Verzeichnishierarchie voraus und greift nicht auf das Dateisystem zu.

## Ausführen

Voraussetzung: Go 1.22 oder neuer. Keine externen Abhängigkeiten.

```sh
go run .
```

Ausgabe:

```text
Größe des Unterordners: 100
Gesamtgröße: 160
```

## Tests und Formatierung

```sh
go test ./...
go fmt ./...
```
