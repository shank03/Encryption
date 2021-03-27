package main

import (
	"bufio"
	"fmt"
	"os"
)

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

        var key string
        fmt.Print("Enter key:\n")
        if scanner.Scan() {
            key = scanner.Text()
            for key == "" {
                fmt.Print("Enter key:\n")
                if scanner.Scan() {
                    key = scanner.Text()
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

						encDat := encrypt(string(rFileDat), key)
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

						encDat := encrypt(input, key)
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

				decDat := decrypt(string(rFileDat), key)
				fmt.Println("Output:\n" + decDat)
			}
		}
	}
}
