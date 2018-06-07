package filemgmt

import (
	"fmt"
	"io"
	"os"
)

// CopyFile kopieert source naar destination
func CopyFile(source string, destination string) error {
	inputFile, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("Fout tijdens openen van source: %s", err)
	}
	outputFile, err := os.Create(destination)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("Kon destination niet openen: %s", err)
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("Schrijven naar destination niet gelukt: %s", err)
	}

	return nil
}

// MoveFile verplaatst een bestand van source naar destination test
func MoveFile(source string, destination string) error {
	err := CopyFile(source, destination)
	if err != nil {
		return err
	}

	err = os.Remove(source)
	if err != nil {
		return fmt.Errorf("Fout tijdens verwijderen van bestand: %s", err)
	}

	return nil
}
