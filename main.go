package main

import (
	"main.go/packages/server"
)

func main() {
	s := server.NewServer("8080")
	s.ListenAndServe()

}

/*
func ReadProviders(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
*/
