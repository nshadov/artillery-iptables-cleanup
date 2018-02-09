package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

func getRulesID(s []byte) []int {
	reader := bytes.NewReader(s)
	scanner := bufio.NewScanner(reader)

	var ids []int

	for scanner.Scan() {
		s := scanner.Text()
		if strings.Contains(s, "DROP") {
			fmt.Printf("%s\n", s)
			i := strings.Split(s, " ")[0]
			x, _ := strconv.Atoi(i)
			ids = append(ids, x)
		}
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
	ids := getRulesID(out)
	removeIptablesIDS(ids)
}
