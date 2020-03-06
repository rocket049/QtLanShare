package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"

	qrcode "github.com/skip2/go-qrcode"
)

func showAddr() {
	ifs, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for n, if1 := range ifs {
		addrs, err := if1.Addrs()
		if err != nil {
			panic(err)
		}
		for i, addr := range addrs {
			if strings.HasPrefix(addr.String(), "127.") {
				continue
			} else if strings.Contains(addr.String(), ":") {
				continue
			} else {
				vs := strings.Split(addr.String(), "/")

				dir1, err := os.UserCacheDir()
				if err != nil {
					panic(err)
				}
				png := filepath.Join(dir1, fmt.Sprintf("fileserver-%d-%d.png", n, i))
				fmt.Println(png)
				var addr string
				if strings.Contains(vs[0], ":") {
					addr = fmt.Sprintf("http://[%s]:6868/index", vs[0])
				} else {
					addr = fmt.Sprintf("http://%s:6868/index", vs[0])
				}
				fmt.Printf("Access URL: %s\n", addr)
				console.Append(fmt.Sprintf("Access URL: %s\n", addr))
				qrcode.WriteFile(addr, qrcode.Highest, 400, png)
				showImg(window, png, addr)
			}
		}
	}
}
