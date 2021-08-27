package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os/exec"
	"strings"
	"time"

	"github.com/nleeper/goment"
)

var count = 0

func add(file_name string) {
	com := "git add " + file_name
	args := strings.Split(com, " ")
	cmd := exec.Command(args[0], args[1:]...)
	_, _ = cmd.CombinedOutput()
}

func commit(data string) {
	s := `git commit -m "` + data + `" --date=` + data
	args := strings.Split(s, " ")
	cmd := exec.Command(args[0], args[1:]...)
	_, _ = cmd.CombinedOutput()
}

func push() {
	s := "git push -u origin main"
	args := strings.Split(s, " ")
	cmd := exec.Command(args[0], args[1:]...)
	_, _ = cmd.CombinedOutput()
}

// func pull() {
// 	s := "git pull origin main"
// 	args := strings.Split(s, " ")
// 	cmd := exec.Command(args[0], args[1:]...)
// 	_, _ = cmd.CombinedOutput()
// }

func main() {
	makeCommit()
	push()
}

func makeCommit() {
	count++
	if count <= 500 {
		rand.Seed(time.Now().UnixNano())
		x := 0 + rand.Intn(54-0+1)
		y := 0 + rand.Intn(6-0+1)
		g, err := goment.New()
		if err != nil {
			fmt.Println(err.Error())
		}
		Date := g.Subtract(1, "year").Add(1, "day").Add(x, "weeks").Add(y, "days").Format()
		data := []byte("{\"date\":\"" + Date + "\"}")
		err = ioutil.WriteFile("data.json", data, 0777)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			add("data.json")
			commit(Date)
			fmt.Println(Date)
		}
		makeCommit()
	}
}
