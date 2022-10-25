package main
import (
"fmt"
"log"
"os"
"bufio"
"sync"
)
var dictionary []string
var hof []string
var wg sync.WaitGroup
func main() {
	infile, err := os.Open("wordle.txt")
    if err != nil {
        log.Fatal(err)
    }
	defer infile.Close()
	scanner := bufio.NewScanner(infile)
	for scanner.Scan() {
		dictionary = append(dictionary, scanner.Text())
	}

	for i:=0; i<26; i++ {
		go generator(i)
		wg.Add(1)
	}
	wg.Wait()
}

func generator(i int) {
	defer wg.Done()
	var caesar [26]string
	var alphabet string = "abcdefghijklmnopqrstuvwxyz"
	var most int = 0
	var wordstested int = 0
	var bestword string = ""
	var count int = 0
	var prevwin bool = false
	for j := 0; j < 26; j++ {
		for k := 0; k < 26; k++ {
			if i == j && j == k {
				wordstested += 676
				continue
			}
			for l := 0; l < 26; l++ {
				if j == k && k == l {
					wordstested += 26
					continue
				}
				for m := 0; m < 26; m++ {
					if k == l && l == m {
						wordstested++
						continue
					}
					prevwin = false
					for n := 0; n<26; n++ {
						caesar[n] = ""
						caesar[n] += string(alphabet[(i+n)%26])
						caesar[n] += string(alphabet[(j+n)%26])
						caesar[n] += string(alphabet[(k+n)%26])
						caesar[n] += string(alphabet[(l+n)%26])
						caesar[n] += string(alphabet[(m+n)%26])
						for _, winner := range hof {
							if caesar[n] == winner {
								prevwin = true
								break
							}
						}
						if prevwin == true {
							break
						}
					}
					if prevwin == true {
						wordstested++
						continue
					}
					count = 0
					for _, shift := range caesar {
						if shift == bestword {
							break
						}
						for _, word := range dictionary {
							if word == shift {
								count++
							}
						}
					}
					if count < most {
						wordstested++
						continue
					}
					wordstested++
					most = count
					bestword = caesar[0]
					hof = append(hof, bestword)
					fmt.Printf("\nTHREAD %s\nWords tested: %d\nCurrent best word: %s\nValid shifts: %d", string(alphabet[i]), wordstested, bestword, most)
				}	
			}	
		}
	}
	fmt.Printf("\nBest word is: %s with %d valid rotated words.\n", bestword, most)
}