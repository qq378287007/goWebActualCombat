package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("index.html")

	buf := bufio.NewReader(f)
	var rep = []string{"<meta name=\"testkey\" content=\"", arg1, "\" /> "}
	var result = ""
	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		if strings.Contains(string(a), "baidu-site-verification") {
			result += strings.Join(rep, "") + "\n"
		} else {
			result += string(a) + "\n"
		}
	}

	fw, err := os.OpenFile("index.html", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	w := bufio.NewWriter(fw)
	w.WriteString(result)
	if err != nil {
		panic(err)
	}
	w.Flush()

}
