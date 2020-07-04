package importer

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const tempFolder = "./.temp"

// Run will import sql model in too sqlite db
// 1. uznip mwb, and take xml file
// 2. run over xml file, and extract all diagrams
// 3. save them to sqlite database
func Run(path string, name string, version string) {
	log.Println("Import start ...")

	files, err := unzip(path, tempFolder)
	if err != nil {
		log.Fatal(err)
	}

	removeTempFolder()

	fmt.Printf("name : %q , version %q!", name, version)
	fmt.Println("Unzipped:\n" + strings.Join(files, "\n"))

}

// Unzip will decompress a zip archive, moving all files and folders
// within the zip file (parameter 1) to an output directory (parameter 2).
func unzip(src string, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}

func removeTempFolder() {
	os.RemoveAll(tempFolder)
	os.MkdirAll(tempFolder, 1)
	log.Println("temp folder removed.")
}
