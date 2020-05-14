package main

import (
	"crypto/sha256"
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/goph/emperror"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func calcSha256(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", emperror.Wrapf(err, "cannot open file %s", filename)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", emperror.Wrapf(err, "cannot read file %s", filename)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func iterate(folder string, csvWriter *csv.Writer, checksum bool) (size int64, folders int64, files int64, cs string, err error) {
	var fileList []os.FileInfo
	fileList, err = ioutil.ReadDir(folder)
	if err != nil {
		err = emperror.Wrapf(err, "cannot read %s", folder)
		return
	}

	folderHash := sha256.New()
	hasError := false
	for _, f := range fileList {
		fullpath := filepath.Join(folder, f.Name())
		if !f.IsDir() {
			fsize := f.Size()
			size += fsize
			files++
			if checksum {
				error := ""
				sha256, err := calcSha256(fullpath)
				if err != nil {
					error = fmt.Sprintf("%v", err)
					hasError = true
				}
				if !hasError {
					io.Copy(folderHash, strings.NewReader(sha256))
				}
				fmt.Printf("%s: Size:%v / Checksum:%v\n", filepath.ToSlash(fullpath), fsize, sha256)
				data := []string{
					"file",
					filepath.ToSlash(fullpath),
					strconv.FormatInt(fsize, 10),
					"",
					"",
					sha256,
					error,
				}
				csvWriter.Write(data)
			}
		} else {
			name := f.Name()
			if name == "." || name == ".." {
				continue
			}
			folders++
			s, fo, fi, csf, err := iterate(fullpath, csvWriter, checksum)
			t := "folder"
			if err != nil {
				t = fmt.Sprintf("unreadable folder: %v", err)
			}
			io.Copy(folderHash, strings.NewReader(csf))
			fmt.Printf("%s: Size:%v / Folders:%v / Files:%v \n", filepath.ToSlash(fullpath), s, fo, fi)
			data := []string{
				t,
				filepath.ToSlash(fullpath),
				strconv.FormatInt(s, 10),
				strconv.FormatInt(fi, 10),
				strconv.FormatInt(fo, 10),
			}
			if checksum {
				data = append(data, csf, "")
			}
			csvWriter.Write(data)
			size += s
			folders += fo
			files += fi
		}
		//		fmt.Println(f.Name())
	}
	cs = fmt.Sprintf("%x", folderHash.Sum(nil))

	return
}

func main() {
	dir := flag.String("dir", ".", "folder to start listing")
	csvfile := flag.String("csv", "", "output file")
	checksum := flag.Bool("checksum", false, "calculate checksums for all files")
	flag.Parse()

	if *csvfile == "" {
		flag.PrintDefaults()
		return
	}

	writer, err := os.Create(*csvfile)
	if err != nil {
		log.Fatalf("cannot open %s: %v", *csvfile, err)
	}
	defer writer.Close()
	csvWriter := csv.NewWriter(writer)
	cols := []string{"type", "name", "size", "files", "subfolders"}
	if *checksum {
		cols = append(cols, "checksum", "error")
	}
	csvWriter.Write(cols)

	path := filepath.ToSlash(filepath.Clean(*dir))

	s, fo, fi, cs, _ := iterate(path, csvWriter, *checksum)
	fmt.Printf("%s: Size:%v / Folders:%v / Files:%v\n", path, s, fo, fi)
	data := []string{
		"folder",
		path,
		strconv.FormatInt(s, 10),
		strconv.FormatInt(fi, 10),
		strconv.FormatInt(fo, 10),
	}
	if *checksum {
		data = append(data, cs, "")
	}
	csvWriter.Write(data)
	csvWriter.Flush()
}
