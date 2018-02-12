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

// A getRulesID parse output string from IPTABLES and retrn slice of rules id.
// It also logs all deleted rules to file (filename param).
func getRulesID(s []byte, filename string) []int {
	reader := bytes.NewReader(s)
	scanner := bufio.NewScanner(reader)

	// Open file for logging
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	timestamp := time.Now().String()

	var ids []int

	// Go through command output and find DROP rules ids.
	for scanner.Scan() {
		s := scanner.Text()
		if strings.Contains(s, "DROP") {
			fmt.Printf("%s\n", s)

			// Write to logfile
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

// A removeIptablesIDS removes rules from IPTABLES based on provides rule ids.
func removeIptablesIDS(ids []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(ids)))
	for _, id := range ids {
		_, err := exec.Command("/sbin/iptables", "-D", "ARTILLERY", strconv.Itoa(id)).Output()
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Remove DROP rules from ARTILLERY table of IPTABLES and logs them to a file.
func main() {
	out, err := exec.Command("/sbin/iptables", "-L", "ARTILLERY", "-n", "--line-numbers").Output()
	if err != nil {
		log.Fatal(err)
	}

	ids := getRulesID(out, "artillery_cleanup.log")
	removeIptablesIDS(ids)
}
