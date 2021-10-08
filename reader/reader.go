package reader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadSubs() {
	f, err := os.Open("./mock/subs/EN.srt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		line := scanner.Text()

		if strings.Contains(line, " --> ") {
			fmt.Println(line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
