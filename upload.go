package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var uploadDir string

func setUploadDir(dir string) {
	uploadDir = dir
	http.HandleFunc("/upload", handleUpload)
}

func handleUpload(w http.ResponseWriter, req *http.Request) {
	var names = []string{}

	w.WriteHeader(200)

	r, err := req.MultipartReader()
	if err == nil {
		f, err := r.ReadForm(20 * 1024 * 1024)
		if err == nil {
			for _, v := range f.File {
				//fmt.Printf("File:%s\n", k)
				for i := 0; i < len(v); i++ {
					var filename string
					for n := 1; true; n++ {
						filename = filepath.Join(uploadDir, fmt.Sprintf("%v-%v", n, v[i].Filename))
						_, err := os.Stat(filename)
						if err != nil {
							break
						}
					}
					of1, _ := os.Create(filename)
					if1, _ := v[i].Open()
					for n, _ := io.Copy(of1, if1); n > 0; n, _ = io.Copy(of1, if1) {
					}
					names = append(names, v[i].Filename)
					fmt.Printf("Upload:%s\n", filepath.Base(filename))
					console.Append(fmt.Sprintf("Upload: %s\n", filepath.Base(filename)))
					of1.Close()
					if1.Close()
				}
			}
		} else {
			log.Println("ReadForm:" + err.Error())
			//resp.Write([]byte("Error"))
			//return
		}
	}
	t := template.New("")
	_, err = t.Parse(tmplUpload)
	if err != nil {
		panic(err)
	}
	t.Execute(w, names)
}
