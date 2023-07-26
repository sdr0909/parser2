package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var protocols = map[string][]int{
	"http":   {80, 8080, 3128, 8888, 8000, 808, 3127, 3129, 8081, 81, 1080, 888, 9999, 8118, 8008},
	"https":  {443, 8443, 3128, 8080, 8888, 4433, 2053, 4443, 8444, 9999, 4125, 1443, 8181, 4430, 4434},
	"socks4": {1080, 8080, 3128, 9050, 9150, 1081, 8118, 4145, 9999, 8888, 8119, 9000, 1082, 8115, 8089},
}

func main() {
	inFile, err := os.Open("ruholy.txt")
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	outFile, err := os.Create("ruholysorted.txt")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	scanner := bufio.NewScanner(inFile)
	writer := bufio.NewWriter(outFile)

	for scanner.Scan() {
		line := scanner.Text()

		words := strings.Split(line, " ")
		if len(words) != 5 {
			continue
		}

		port, err := strconv.Atoi(words[2])
		if err != nil {
			continue
		}

		ip := words[3]

		for protocol, ports := range protocols {
			for _, p := range ports {
				if p == port {
					output := fmt.Sprintf("%s://%s:%d\n", protocol, ip, port)
					writer.WriteString(output)
				}
			}
		}
	}

	writer.Flush()
}
