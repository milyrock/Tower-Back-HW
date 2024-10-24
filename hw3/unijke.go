package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func sfi(text []string, num_fields int, num_chars int, iflag bool) []string {
	changed := []string{}
	for _, line := range text {
		if iflag {
			line = strings.ToLower(line)
		}

		words := strings.Split(line, " ")
		words = words[num_fields:]
		new_line := strings.Join(words, " ")
		new_line = new_line[min(len(new_line), num_chars):]
		changed = append(changed, new_line)
	}
	return changed
}

func noFlags(changed []string) []string {
	a := []string{}
	if len(changed) == 0 {
		return a
	}
	last := 0
	for i := 1; i < len(changed); i++ {
		if changed[i] != changed[last] {
			a = append(a, changed[last])
			last = i
		}
	}
	a = append(a, changed[last])
	return a
}

func cFlag(changed []string) []string {
	a := []string{}
	last := 0
	counter := 1           // одна строка такого типа всегда есть
	if len(changed) == 0 { // такое возможно только если изначальный пустой, тк и так и так запускаем обработку нум филдс и нум чарс и и
		return a
	}
	for i := 1; i < len(changed); i++ {
		if changed[i] != changed[last] {
			a = append(a, strconv.Itoa(counter)+" "+changed[last])
			counter = 1
			last = i
		} else {
			counter += 1
		}
	}
	a = append(a, strconv.Itoa(counter)+" "+changed[last]) //хз поч нельзя стринг(каунтер) чтото внутри происходит видимо
	return a
}

func dFlag(changed []string) []string {
	a := []string{}
	last := 0
	counter := 1
	if len(changed) == 0 {
		return a
	}
	for i := 1; i < len(changed); i++ {
		if changed[i] != changed[last] {
			if counter > 1 {
				a = append(a, changed[last])
			}
			counter = 1
			last = i
		} else {
			counter += 1
		}
	}

	if counter > 1 {
		a = append(a, changed[last])
	}
	return a
}

func uFlag(changed []string) []string {
	a := []string{}
	last := 0
	counter := 1
	if len(changed) == 0 {
		return a
	}
	for i := 1; i < len(changed); i++ {
		if changed[i] != changed[last] {
			if counter == 1 {
				a = append(a, changed[last])
			}
			counter = 1
			last = i
		} else {
			counter += 1
		}
	}

	if counter == 1 {
		a = append(a, changed[last])
	}
	return a
}

func main() {

	var count = flag.Bool("c", false, "Counting reapeated strings")
	var down = flag.Bool("d", false, "Returning repeated strings")
	var up = flag.Bool("u", false, "Returing unique strings")

	var field = flag.Int("f", 0, "Not counting first n fields in a row")
	var chars = flag.Int("s", 0, "Not counting first n chars in a row")
	var size = flag.Bool("i", false, "Size does not matter")
	flag.Parse()

	if (*count && *up) || (*count && *down) || (*down && *up) {
		log.Fatal("You can't use those flags together, please reconsider your input and run this program again")
	}

	file_name_in := os.Stdin
	if flag.NArg() > 0 {
		var err error
		file_name_in, err = os.Open(flag.Arg(0))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer file_name_in.Close()
	}

	text := []string{}
	reader := bufio.NewReader(file_name_in)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimRight(line, "\n")
		text = append(text, line)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}

	}

	changed := sfi(text, *field, *chars, *size)
	if *count {
		changed = cFlag(changed)
	} else if *down {
		changed = dFlag(changed)
	} else if *up {
		changed = uFlag(changed)
	} else {
		changed = noFlags(changed)
	}
	if flag.NArg() > 1 {
		var file_name_out *os.File
		file_name_out, err := os.Create(flag.Arg(1))
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		defer file_name_out.Close()

		for _, line := range changed {
			fmt.Fprintln(file_name_out, line)
		}
	} else {
		file_name_out := os.Stdout
		for _, line := range changed {
			fmt.Fprintln(file_name_out, line)
		}
	}
}
