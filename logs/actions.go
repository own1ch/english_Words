package logs

import (
	"fmt"
	"os"
	"time"
)

func createFile(text string) {
	file, err := os.Create(initFileName())
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.WriteString(text + "\n")
}

func UpdateFile(text string) {
	file, err := os.OpenFile(initFileName()+".txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		createFile(text)
		return
	}
	defer file.Close()

	_, _ = file.WriteString(text + "\n")
}

func initFileName() string {
	currentTime := time.Now()
	return currentTime.Format("01-02-2006")
}

func main() {
	fmt.Println(parse())
}

func parse() [4]int {
	var res [4]int
	var k = 0
	var isEnglish = true
	text := "ая"
	for i := 0; i < len(text); i++ {
		l := len(text)
		fmt.Println(l)
		c := text[i]
		ascii := int(c)
		if ascii >= 224 && ascii <= 255 && !isEnglish {
			res[k] = i
			k++
			isEnglish = true
		} else if ascii >= 97 && ascii <= 122 && isEnglish {
			res[k] = i
			k++
			isEnglish = false
		}
	}
	return res
}
