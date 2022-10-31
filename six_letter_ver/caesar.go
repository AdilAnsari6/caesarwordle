// Finds what Caesar ciphered word has the most valid dictionary entries it could be deciphered as.
// by Adil Ansari
package main
import (
"fmt"
"log"
"os"
"bufio"
"sync"
"math"
"math/bits"
"time"
)

var dict1 []uint
var dict2 []uint
var dict3 []uint
var dict4 []uint
var dict5 []uint
var dict6 []uint
var alphabet string = "abcdefghijklmnopqrstuvwxyz"
var wordstested int = 0
var most int = 0
var oldmost int = 0
var bestwords []string;
var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	start := time.Now()

	infile, err := os.Open("words.txt")
	var word string = ""
    if err != nil {
        log.Fatal(err)
    }
	defer infile.Close()
	
	// Each letter is a stored as a binary number with the digit corresponding to its position in the alphabet turned on
	scanner := bufio.NewScanner(infile)
	for scanner.Scan() {
		word = scanner.Text()
		dict1 = append(dict1, uint(math.Pow(2, float64(word[0] - 97))))
		dict2 = append(dict2, uint(math.Pow(2, float64(word[1] - 97))))
		dict3 = append(dict3, uint(math.Pow(2, float64(word[2] - 97))))
		dict4 = append(dict4, uint(math.Pow(2, float64(word[3] - 97))))
		dict5 = append(dict5, uint(math.Pow(2, float64(word[4] - 97))))
		dict6 = append(dict6, uint(math.Pow(2, float64(word[5] - 97))))
	}

	// 26 routines, one for each first letter of the words (aXXXXX, bXXXXX, etc)
	for i:=0; i<26; i++ {
		go generator(i)
		wg.Add(1)
	}
	wg.Wait()	
					
	fmt.Printf("\nThe best words are: ")
	
	for _,bestword := range bestwords {
		fmt.Printf("%s, ", bestword)
	}
	fmt.Printf(" with %d dictionary entries in their Caesar cycle.", most)
	elapsed := time.Since(start)
	fmt.Printf("\nTotal time: %s", elapsed)
}

func generator(h int) {
	defer wg.Done()
	var caesar1 [26]uint
	var caesar2 [26]uint
	var caesar3 [26]uint
	var caesar4 [26]uint
	var caesar5 [26]uint
	var caesar6 [26]uint
	var count int = 0
	var currbest string = ""
	var currbestshift [26]string
	var prevwin bool
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			for k := 0; k < 26; k++ {
				for l := 0; l < 26; l++ {
					for m := 0; m < 26; m++ {
						for n := 0; n < 26; n++ {
							// Using binary rotation to simulate a Caesar shift, wrapping after the 26th digit
							caesar1[n] = bits.RotateLeft(uint(math.Pow(2,float64(h))), n)
							if caesar1[n] > 33554432 {
								caesar1[n] = bits.RotateLeft(caesar1[n], -26)
							}
							
							caesar2[n] = bits.RotateLeft(uint(math.Pow(2,float64(i))), n)
							if caesar2[n] > 33554432 {
								caesar2[n] = bits.RotateLeft(caesar2[n], -26)
							}
							
							caesar3[n] = bits.RotateLeft(uint(math.Pow(2,float64(j))), n)
							if caesar3[n] > 33554432 {
								caesar3[n] = bits.RotateLeft(caesar3[n], -26)
							}	
							caesar4[n] = bits.RotateLeft(uint(math.Pow(2,float64(k))), n)
							if caesar4[n] > 33554432 {
								caesar4[n] = bits.RotateLeft(caesar4[n], -26)
							}
							
							caesar5[n] = bits.RotateLeft(uint(math.Pow(2,float64(l))), n)
							if caesar5[n] > 33554432 {
								caesar5[n] = bits.RotateLeft(caesar5[n], -26)
							}
							
							caesar6[n] = bits.RotateLeft(uint(math.Pow(2,float64(m))), n)
							if caesar6[n] > 33554432 {
								caesar6[n] = bits.RotateLeft(caesar6[n], -26)
							}
						}
						// Checking to see if any shift in a word's cycle matches a dictionary entry
						count = 0
						for a:=0; a<len(caesar1); a++ {
							for b:=0; b<len(dict1); b++ {
								if caesar1[a] == dict1[b] && caesar2[a] == dict2[b] && caesar3[a] == dict3[b] && caesar4[a] == dict4[b] && caesar5[a] == dict5[b] && caesar6[a] == dict6[b] {
									count++
								}
							}
						}
						mutex.Lock()
						wordstested++
						// If this word isn't the best, only print out every one hundred thousand words tested
						if count < most {
							if wordstested % 100000 == 0 {
								fmt.Printf("\nROUTINE %s\nWords tested: %d (%.2f%% done)\nCurrent best words: ", string(alphabet[h]), wordstested, float64(wordstested)/3089157.76)
								for ind,bestword := range bestwords {
									if ind >= 5 {
										fmt.Printf("and %d more.", len(bestwords)-5)
										break
									} else {
										fmt.Printf("%s, ", bestword)
									}
								}
								fmt.Printf("\nValid shifts: %d", most)
							}
							mutex.Unlock()
							continue
						// If the word is the new best, update the high score and clear all previous best words
						} else if count > most {
							most = count
							bestwords = bestwords[:0]
						}
						// Check to see if this best word is just a Caesar shift of another best word
						prevwin = false
						currbest = string(alphabet[h]) + string(alphabet[i]) + string(alphabet[j]) + string(alphabet[k]) + string(alphabet[l]) + string(alphabet[m])
						for n := 0; n<26; n++ {
							currbestshift[n] = ""
							currbestshift[n] += string(alphabet[(int(currbest[0])-97+n)%26])
							currbestshift[n] += string(alphabet[(int(currbest[1])-97+n)%26])
							currbestshift[n] += string(alphabet[(int(currbest[2])-97+n)%26])
							currbestshift[n] += string(alphabet[(int(currbest[3])-97+n)%26])
							currbestshift[n] += string(alphabet[(int(currbest[4])-97+n)%26])
							currbestshift[n] += string(alphabet[(int(currbest[5])-97+n)%26])
							for _, bestword := range bestwords {
								if currbestshift[n] == bestword {
									prevwin = true
									break
								}
							}
							if prevwin {
								break
							}
						}
						// If this word is not a Caesar shift of another winner, add it to the list and print out the current best words
						if !prevwin {
							bestwords = append(bestwords, currbest)
							fmt.Printf("\nROUTINE %s\nWords tested: %d (%.2f%% done)\nCurrent best words: ", string(alphabet[h]), wordstested, float64(wordstested)/3089157.76)
							for ind,bestword := range bestwords {
								if ind >= 5 {
									fmt.Printf("and %d more.", len(bestwords)-5)
									break
								} else {
									fmt.Printf("%s, ", bestword)
								}
							}
							fmt.Printf("\nValid shifts: %d", most)
						}
						mutex.Unlock()
					}	
				}	
			}
		}
	}
}