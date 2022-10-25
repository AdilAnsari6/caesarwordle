# Caesar's Wordle
Inspired by a class assignment showing that "arena" and "river" could both be solutions to a Caesar ciphertext such as "EVIRE"
## Run instructions
With the source code (caesar.go) and wordle.txt in the same folder:
```
go run caesar.go
```
## How it works
As a dictionary of valid words, the program uses the dictionary from Wordle.
To store words, each letter is stored as a 26-bit number, with the only high bit corresponding to that letter's position in the alphabet.
The program spawns 26 goroutines, each handling the generation every possible 5 letter ciphertext beginning with the letter the thread is assigned to.
```
Thread 0: all aXXXX ciphertexts
Thread 1: all bXXXX ciphertexts
.
.
Thread 25: all zXXXX ciphertexts
```
For each ciphertext, all caesar shifts are generated by rotating the high bit to the next position. These are then compared to the dictionary to see how many dictionary words a given word's Caesar shift cycle includes, or in practical terms, how many valid plaintexts our ciphertext can be deciphered into. Ciphertexts with the most valid plaintexts are then stored until the program has generated all possible ciphertexts (26<sup>5</sup>), upon which they are displayed.

My original implementation in C++ used threads, which are heavier than goroutines, and so the program had an estimated runtime of about 8 hours.
I switched to Go to make use of its faster multiprocessing. My [first Go implementation](https://github.com/AdilAnsari6/caesarwordle/commit/c4fb52ea8c6b451b8c0931e771a12042c43e99eb) involved looking up the next letter in the alphabet for each shift, which was 5-10x slower than using binary rotation. That took over an hour, it now takes about 5-10 minutes.
