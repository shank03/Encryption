/*
 * Copyright (c) 2021, Shashank Verma <shashank.verma2002@gmail.com>(shank03)
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 */

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// --------------------------------- USER INPUT ----------------------------------//

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var option string
	fmt.Println("[1] Encrypt")
	fmt.Println("[2] Decrypt")
	fmt.Print("Select an option: ")
	if scanner.Scan() {
		option = scanner.Text()
		for option != "1" && option != "2" {
			fmt.Println("Invalid option")
			fmt.Print("Select an option: ")
			if scanner.Scan() {
				option = scanner.Text()
			}
		}
		showOptions(option, scanner)
	}
}

func showOptions(mode string, scanner *bufio.Scanner) {
	var option string
	if mode == "1" {
		fmt.Println("\n[1] Encrypt text from file")
		fmt.Println("[2] Custom input (single line input)")
	}
	if mode == "2" {
		fmt.Println("\n[1] Decrypt text from file")
	}
	fmt.Print("Select an option: ")
	if scanner.Scan() {
		option = scanner.Text()
		if mode == "1" {
			for option != "1" && option != "2" {
				fmt.Println("Invalid option")
				fmt.Print("Select an option: ")
				if scanner.Scan() {
					option = scanner.Text()
				}
			}
		}
		if mode == "2" {
			for option != "1" {
				fmt.Println("Invalid option")
				fmt.Print("Select an option: ")
				if scanner.Scan() {
					option = scanner.Text()
				}
			}
		}

		if mode == "1" {
			var path string
			fmt.Print("Enter file name or path to save encrypted text:\n")
			if scanner.Scan() {
				path = scanner.Text()
				file, e := os.Create(path)
				for e != nil {
					fmt.Println("Error opening file/path: " + e.Error())
					fmt.Print("Enter file name or path to save encrypted text:\n")
					if scanner.Scan() {
						path = scanner.Text()
					}
					file, e = os.Create(path)
				}
				if option == "1" {
					var dat string
					fmt.Print("Enter file name to encrypt text from:\n")
					if scanner.Scan() {
						dat = scanner.Text()
						rFileDat, er := os.ReadFile(dat)
						for er != nil {
							fmt.Println("Error opening file/path: " + er.Error())
							fmt.Print("Enter file name to encrypt text from:\n")
							if scanner.Scan() {
								dat = scanner.Text()
							}
							rFileDat, er = os.ReadFile(dat)
						}

						encDat := encrypt(string(rFileDat), scanner)
						_, _ = file.WriteString(encDat)
						fmt.Println("Encrypted text stored in " + file.Name())
						_ = file.Close()
					}
				}
				if option == "2" {
					var input string
					fmt.Println("Enter input:")
					if scanner.Scan() {
						input = scanner.Text()
						for input == "" {
							fmt.Println("Enter input:")
							if scanner.Scan() {
								input = scanner.Text()
							}
						}

						encDat := encrypt(input, scanner)
						_, _ = file.WriteString(encDat)
						fmt.Println("Encrypted text stored in " + file.Name())
						_ = file.Close()
					}
				}
			}
		}
		if mode == "2" {
			var dat string
			fmt.Print("Enter file name to decrypt text from:\n")
			if scanner.Scan() {
				dat = scanner.Text()
				rFileDat, er := os.ReadFile(dat)
				for er != nil {
					fmt.Println("Error opening file/path: " + er.Error())
					fmt.Print("Enter file name to decrypt text from:\n")
					if scanner.Scan() {
						dat = scanner.Text()
					}
					rFileDat, er = os.ReadFile(dat)
				}

				decDat := decrypt(rFileDat, scanner)
				fmt.Println("Output:\n" + decDat)
			}
		}
	}
}

// --------------------------------- API ----------------------------------//

const lineSeparator = "--------------------------------------------------------------------- "

func formatBinaryString(bin string) string {
	if len(bin) == 64 {
		return bin
	}
	var out string
	for i := 0; i < 64-len(bin); i++ {
		out += "0"
	}
	return out + bin
}

func compress(input string, m *map[rune][]uint64) {
	for i := 0; i < len(input); i++ {
		_, found := (*m)[rune(input[i])]
		if found {
			(*m)[rune(input[i])] = append((*m)[rune(input[i])], []uint64{uint64(i + 1)}...)
		} else {
			var pos []uint64
			pos = append(pos, []uint64{uint64(i + 1)}...)
			(*m)[rune(input[i])] = pos
		}
	}
}

func binaryToUint(bin string) uint64 {
	var out uint64 = 0
	for i := 0; i < len(bin); i++ {
		if bin[len(bin)-1-i] == '1' {
			out += uint64(math.Pow(2, float64(i)))
		}
	}
	return out
}

func encrypt(text string, scanner *bufio.Scanner) string {
	var outStr string

	var k string
	fmt.Print("Enter key to encrypt:\n")
	if scanner.Scan() {
		k = scanner.Text()
		for k == "" {
			fmt.Print("Enter key to encrypt:\n")
			if scanner.Scan() {
				k = scanner.Text()
			}
		}
	}
	var keyMask uint64 = 0
	for i := 0; i < len(k); i++ {
		keyMask += uint64(k[i])
	}

	for _, line := range strings.Split(text, "\n") {
		if line == "" {
			continue
		}
		dataMap := make(map[rune][]uint64)
		compress(line, &dataMap)

		var out = uint64(len(line)) ^ keyMask
		outStr += lineSeparator + "\n" + formatBinaryString(strconv.FormatUint(out, 2)) + " \n"

		for key, pos := range dataMap {
			if key == '\n' || key == '\r' {
				outStr += "nl "
			} else if key == '\t' {
				outStr += "tb "
			} else if key == ' ' {
				outStr += "sp "
			} else {
				outStr += string(key) + " "
			}

			var ps, c uint64 = 0, 0
			for i := 0; i < len(pos); i++ {
				ps |= pos[i] << (8 * c)
				c++
				if c == 4 {
					ps |= 4 << 32
					outStr += formatBinaryString(strconv.FormatUint(ps^keyMask, 2)) + " "
					c, ps = 0, 0
				}
			}
			if ps != 0 {
				ps |= c << 32
			}
			outStr += formatBinaryString(strconv.FormatUint(ps^keyMask, 2)) + " \n"
		}
	}
	return outStr + lineSeparator + "\n"
}

func decrypt(encData []byte, scanner *bufio.Scanner) string {
	var outStr string

	var k string
	fmt.Print("Enter key to decrypt:\n")
	if scanner.Scan() {
		k = scanner.Text()
		for k == "" {
			fmt.Print("Enter key to decrypt:\n")
			if scanner.Scan() {
				k = scanner.Text()
			}
		}
	}
	var keyMask uint64 = 0
	for i := 0; i < len(k); i++ {
		keyMask += uint64(k[i])
	}

	var intoData = false
	var data []rune = nil
	var char rune
	for _, line := range strings.Split(string(encData), "\n") {
		if line == lineSeparator {
			intoData = true
			if data != nil {
				for i := range data {
					outStr += string(data[i])
				}
			}
			continue
		}
		if intoData {
			datLength := binaryToUint(strings.Split(line, " ")[0]) ^ keyMask
			data = nil
			data = make([]rune, datLength)
			intoData = false
		} else {
			tokens := strings.Split(line, " ")
			for t, tok := range tokens {
				if tok == "" {
					continue
				}
				if t == 0 {
					if len(tok) == 1 {
						char = rune(tok[0])
					}
					if len(tok) == 2 {
						if tok == "nl" {
							char = '\n'
						}
						if tok == "sp" {
							char = ' '
						}
						if tok == "tb" {
							char = '\t'
						}
					}
				} else {
					if len(tok) == 64 {
						o := binaryToUint(tok) ^ keyMask
						numPos := o >> 32
						for p := uint64(0); p < numPos; p++ {
							pos := (o >> (8 * p)) & 0xFF
							if data != nil {
								data[pos-1] = char
							}
						}
					}
				}
			}
		}
	}
	return outStr
}
