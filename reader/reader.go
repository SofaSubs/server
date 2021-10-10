package reader

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/SofaSubs/server/model"
)

func GetSubs() (map[int64]model.Sub, error) {
	subsMap := map[int64]model.Sub{}

	fO, errO := os.Open("./mock/subs/EN.srt")
	if errO != nil {
		return nil, errO
	}
	defer fO.Close()

	fT, errT := os.Open("./mock/subs/ES.srt")
	if errT != nil {
		return nil, errO
	}
	defer fT.Close()

	subs, err := readSubs(fO, subsMap, false)

	if err != nil {
		return nil, err
	}

	if subs, err := readSubs(fT, subs, true); err != nil {
		return nil, err
	} else {
		return subs, nil
	}
}

func readSubs(f *os.File, subsMap map[int64]model.Sub, tranlateMode bool) (map[int64]model.Sub, error) {
	scanner := bufio.NewScanner(f)

	subLine := false
	sub := ""
	second := ""
	for scanner.Scan() {

		line := scanner.Text()

		if line == "" {
			if s, err := subTimeFormatToSeconds(second); err == nil {
				if tranlateMode {
					if mapValue, exist := subsMap[s]; exist {
						mapValue.Translate = sub
						subsMap[s] = mapValue
					}
				} else {
					subsMap[s] = model.Sub{
						Second:    s,
						Original:  sub,
						Translate: "",
					}
				}
			}
			subLine = false
			sub = ""
			second = ""
		} else {
			if subLine {
				sub += line + " \n"
			}

			if strings.Contains(line, " --> ") {
				subLine = true
				second = strings.Split(line, " --> ")[0]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return subsMap, nil
}

//sub time format => 00:00:00,00
//TODO: refactor repeat code
func subTimeFormatToSeconds(time string) (int64, error) {
	var seconds int64
	times := strings.Split(time, ":")

	secondsH, err := strconv.ParseInt(times[0], 10, 64)
	if err != nil {
		return 0, err
	}
	seconds = ((secondsH * 60) * 60)

	secondsM, err := strconv.ParseInt(times[1], 10, 64)
	if err != nil {
		return 0, err
	}
	seconds = seconds + (secondsM * 60)

	secondsS, err := strconv.ParseInt(strings.Split(times[2], ",")[0], 10, 64)
	if err != nil {
		return 0, err
	}
	seconds = seconds + (secondsS)

	return seconds, nil
}
