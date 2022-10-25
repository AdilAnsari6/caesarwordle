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
var dict1 []float64
var dict2 []float64
var dict3 []float64
var dict4 []float64
var dict5 []float64
var alphabet string = "abcdefghijklmnopqrstuvwxyz"
var wg sync.WaitGroup
func main() {
	start := time.Now()
	infile, err := os.Open("wordle.txt")
	var word string = ""
    if err != nil {
        log.Fatal(err)
    }
	defer infile.Close()
	scanner := bufio.NewScanner(infile)
	for scanner.Scan() {
		word = scanner.Text()
		dict1 = append(dict1, math.Pow(2, float64(word[0]) - 97))
		dict2 = append(dict2, math.Pow(2, float64(word[1]) - 97))
		dict3 = append(dict3, math.Pow(2, float64(word[2]) - 97))
		dict4 = append(dict4, math.Pow(2, float64(word[3]) - 97))
		dict5 = append(dict5, math.Pow(2, float64(word[4]) - 97))
	}

	for i:=0; i<26; i++ {
		go generator(i)
		wg.Add(1)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Total time: %s", elapsed)
}

func generator(i int) {
	defer wg.Done()
	var caesar1 [26]float64
	var caesar2 [26]float64
	var caesar3 [26]float64
	var caesar4 [26]float64
	var caesar5 [26]float64
	var most int = 0
	var wordstested int = 0
	var bestword string = ""
	var count int = 0
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
					for n := 0; n<26; n++ {
						caesar1[n] = float64(bits.RotateLeft(uint(math.Pow(2,float64(i))), n))
						caesar2[n] = float64(bits.RotateLeft(uint(math.Pow(2,float64(j))), n))
						caesar3[n] = float64(bits.RotateLeft(uint(math.Pow(2,float64(k))), n))
						caesar4[n] = float64(bits.RotateLeft(uint(math.Pow(2,float64(l))), n))
						caesar5[n] = float64(bits.RotateLeft(uint(math.Pow(2,float64(m))), n))
					}
					count = 0
					for a:=0; a<len(caesar1); a++ {
						for b:=0; b<len(dict1); b++ {
							if caesar1[a] == dict1[b] && caesar2[a] == dict2[b] && caesar3[a] == dict3[b] && caesar4[a] == dict4[b] && caesar5[a] == dict5[b] {
								count++
							}
						}
					}
					wordstested++
					if count <= most {
						if wordstested % 25000 == 0 {
							fmt.Printf("\nTHREAD %s\nWords tested: %d\nCurrent best word: %s\nValid shifts: %d", string(alphabet[i]), wordstested, bestword, most)
						}
						continue
					}
					most = count
					bestword = string(alphabet[int(math.Logb(caesar1[0]))]) + string(alphabet[int(math.Logb(caesar2[0]))]) + string(alphabet[int(math.Logb(caesar3[0]))]) + string(alphabet[int(math.Logb(caesar4[0]))]) + string(alphabet[int(math.Logb(caesar5[0]))])
					fmt.Printf("\nTHREAD %s\nWords tested: %d\nCurrent best word: %s\nValid shifts: %d", string(alphabet[i]), wordstested, bestword, most)
				}	
			}	
		}
	}
	fmt.Printf("\nBest word is: %s with %d valid rotated words.\n", bestword, most)
}