package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func iterate( folder string, csvWriter *csv.Writer ) (size int64, folders int64, files int64 ) {
	fileList, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range fileList {
		if f.IsDir() {
			name := f.Name()
			if name == "." || name == ".." {
				continue
			}
			folders++
			fullpath := filepath.Join(folder, f.Name())
			s, fo, fi := iterate(fullpath, csvWriter)
			fmt.Printf("%s: Size:%v / Folders:%v / Files:%v\n", filepath.ToSlash(fullpath), s, fo, fi )
			csvWriter.Write([]string{
				filepath.ToSlash(fullpath),
				strconv.FormatInt(s, 10),
				strconv.FormatInt(fi, 10),
				strconv.FormatInt(fo, 10),
			})
			size += s
			folders += fo
			files += fi
		} else {
			size += f.Size()
			files++
		}
//		fmt.Println(f.Name())
	}
	return
}

func main() {
	dir := flag.String("dir", ".", "folder to start listing")
	csvfile := flag.String("csv", "", "output file")
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
	csvWriter.Write([]string{"folder", "size", "files", "subfolders"})

	path := filepath.ToSlash(filepath.Clean(*dir))

	s, fo, fi := iterate(path, csvWriter)
	fmt.Printf("%s: Size:%v / Folders:%v / Files:%v\n", path, s, fo, fi )
	csvWriter.Write([]string{
		path,
		strconv.FormatInt(s, 10),
		strconv.FormatInt(fi, 10),
		strconv.FormatInt(fo, 10),
	})
	csvWriter.Flush()
}
