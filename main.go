package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	core "core/func"
)

func main() {
	args := os.Args
	var newData []string

	// check that there are two arguments
	if len(args) < 3 {
		fmt.Println("Error: There must be two file names")
	} else if len(args) > 3 {
		fmt.Println("Error: Too many arguments")
	} else {
		fileNameS := args[1]
		fileNameO := args[2]
		// check that the two arguments ends with .txt
		arg1 := strings.HasSuffix(fileNameS, ".txt")
		arg2 := strings.HasSuffix(fileNameO, ".txt")

		if !arg1 {
			fmt.Println("Error: sample file is not txt")
		}
		if !arg2 {
			fmt.Println("Error: output file is not txt")
		}

		if arg1 && arg2 {

			// Read file data
			sample, error := os.ReadFile(fileNameS)
			if error != nil {
				fmt.Println("Error:", error)

				//check if empty
			} else if len(sample) == 0 {
				fmt.Println("Error: The content is empty")

			} else {

				data := string(sample)
				// from string to an array of strings
				newData = core.SplitWhiteSpaces(data)

				for i := 0; i < len(newData); i++ {
					if newData[i] == "(hex)" {
						if i > 0 {
							//convert from hex to decimal version
							decNum, error := strconv.ParseInt(newData[i-1], 16, 64)
							if error != nil {
								fmt.Println("Error", error)
							}
							//replace the word and convert to string but in decimal form
							newData[i-1] = strconv.FormatInt(decNum, 10)
							//remove current word which is (hex)
							newData = append(newData[:i], newData[i+1:]...)
							i-- // Decrement i to recheck the current index after removal
						} else {
							fmt.Println("Error: There is no word before (hex)")
						}

					} else if newData[i] == "(bin)" {
						if i > 0 {
							//convert from bin to decimal version
							decNum, error := strconv.ParseInt(newData[i-1], 2, 64)
							if error != nil {
								fmt.Println("Error", error)
							}
							newData[i-1] = strconv.FormatInt(decNum, 10)
							newData = append(newData[:i], newData[i+1:]...)
							i-- // Decrement i to recheck the current index after removal
						} else {
							fmt.Println("Error: There is no word before (bin)")
						}

					} else if newData[i] == "(up)" {
						if i > 0 {
							newData[i-1] = core.ToUpper(newData[i-1])
							newData = append(newData[:i], newData[i+1:]...)
							i-- // Decrement i to recheck the current index after removal
						} else {
							fmt.Println("Error: There is no word before (up)")
							break
						}
 
					} else if newData[i] == "(low)" {
						if i > 0 {
							newData[i-1] = core.ToLower(newData[i-1])
							newData = append(newData[:i], newData[i+1:]...)
							i-- // Decrement i to recheck the current index after removal
						} else {
							fmt.Println("Error: There is no word before (low)")
							break
						}

					} else if newData[i] == "(cap)" {
						if i > 0 {
							newData[i-1] = core.ToCap(newData[i-1])
							newData = append(newData[:i], newData[i+1:]...)
							i-- // Decrement i to recheck the current index after removal
						} else {
							fmt.Println("Error: There is no word before (cap)")
							break
						}

						//cap, low, up with number
					} else if newData[i] == "(cap," {
						if i+1 >= len(newData) {
							fmt.Println("Error: there is no number after cap")
							break
						} else {
							//get the number and remove unwanted characters
							tempN := strings.Trim(newData[i+1], "(), ")

							//convert the tempN into an int
							number, err := strconv.Atoi(tempN)

							if err != nil {
								fmt.Println("Error: ", err)

								//error if number is negative
							} else if number < 0 {
								fmt.Println("Error: the number should be positive")
								break
							} else {

								for j := 1; j <= number; j++ {
									//check if index out of range
									if i-j < 0 {
										fmt.Println("Error: Index out of range, please change the number")
										break
									} else {
										//adjust the word and replace it
										newData[i-j] = core.ToCap(newData[i-j])
									}

								}

								newData = append(newData[:i], newData[i+2:]...)
								i-- // Decrement i to recheck the current index after removal
							}
						}

						//repeat for the low and up
					} else if newData[i] == "(low," {
						if i+1 >= len(newData) {
							fmt.Println("Error: there is no number after low")
							break
						}
						tempN := strings.Trim(newData[i+1], "(), ")
						number, err := strconv.Atoi(tempN)

						if err != nil {
							fmt.Println("Error: ", err)
						} else if number < 0 {
							fmt.Println("Error: the number should be positive")
						} else {
							for j := 1; j <= number; j++ {
								if i-j < 0 {
									fmt.Println("Error: Index out of range, please change the number")
									break
								} else {
									newData[i-j] = core.ToLower(newData[i-j])
								}

							}
							newData = append(newData[:i], newData[i+2:]...)
							i-- // Decrement i to recheck the current index after removal
						}

					} else if newData[i] == "(up," {
						if i+1 >= len(newData) {
							fmt.Println("Error: there is no number after up")
							break
						}
						tempN := strings.Trim(newData[i+1], "(), ")
						number, err := strconv.Atoi(tempN)

						if err != nil {
							fmt.Println("Error: ", err)
						} else if number < 0 {
							fmt.Println("Error: the number should be positive")
							break
						} else {
							for j := 1; j <= number; j++ {
								if i-j < 0 {
									fmt.Println("Error: Index out of range, please change the number")
									break
								} else {
									newData[i-j] = core.ToUpper(newData[i-j])
								}

							}
							newData = append(newData[:i], newData[i+2:]...)
							i-- // Decrement i to recheck the current index after removal
						}

					}
				}
				//checking for vowels after an 'a' and adjusting it
				newData = core.Vowels(newData)

				//check for punctuations
				newData = core.Punctuation(newData)

				//check the apostrophes
				output := core.Apostrophe(newData)

				_, err := os.Stat(args[2])

				if os.IsNotExist(err) {
					fmt.Println("Error: The output file does not exist")
				} else {
					file, err := os.Create(args[2])

					if err != nil {
						fmt.Println("Failed to create file:", err)
					} else {
						_, err = file.WriteString(output)

						if err != nil {
							fmt.Println("Failed to write to file:", err)
						}
						file.Close()
					}
				}
			}
		}
	}
}
