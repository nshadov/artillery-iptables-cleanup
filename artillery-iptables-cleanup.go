package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"
)

func getRulesID(s []byte, filename string) []int {
	reader := bytes.NewReader(s)
	scanner := bufio.NewScanner(reader)

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	timestamp := time.Now().String()

	var ids []int

	for scanner.Scan() {
		s := scanner.Text()
		if strings.Contains(s, "DROP") {
			fmt.Printf("%s\n", s)

			if _, err := f.Write([]byte(timestamp + ": " + s + "\n")); err != nil {
				log.Fatal(err)
			}

			i := strings.Split(s, " ")[0]
			x, _ := strconv.Atoi(i)
			ids = append(ids, x)
		}
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	return ids
}

func removeIptablesIDS(ids []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(ids)))
	for _, id := range ids {
		_, err := exec.Command("/sbin/iptables", "-D", "ARTILLERY", strconv.Itoa(id)).Output()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	out, err := exec.Command("/sbin/iptables", "-L", "ARTILLERY", "-n", "--line-numbers").Output()
	// out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	ids := getRulesID(out, "artillery_cleanup.log")
	removeIptablesIDS(ids)
}
