//unions help compiler vectorize
#include "immintrin.h"
#include <iostream>
#include <fstream>
#include <vector>
#include <chrono>

using namespace std;
using namespace std::chrono;

struct Word { 
	int8_t ltr1;
	int8_t ltr2;
	int8_t ltr3;
	int8_t ltr4;
	int8_t ltr5;
};
union WordDelta
{
	uint8_t ltr[4];
	uint32_t word;
};
vector<int8_t> ltr1;
vector<WordDelta> dictionary;
int8_t tmpltr2;
int8_t tmpltr3;
int8_t tmpltr4;
int8_t tmpltr5;

int main()
{
	auto start = high_resolution_clock::now();

	string dictIn;
	fstream infile;
	WordDelta dict = {0};
	infile.open("wordle.txt",ios::in);
	while(!infile.eof()) {
		infile >> dictIn;
		
		ltr1.push_back(dictIn[0]);

		tmpltr2 = dictIn[1] - dictIn[0];
		if(tmpltr2 < 0) tmpltr2 += 26;
		dict.ltr[0] = tmpltr2;

		tmpltr3 = dictIn[2] - dictIn[0];
		if(tmpltr3 < 0) tmpltr3 += 26;
		dict.ltr[1] = tmpltr3;

		tmpltr4 = dictIn[3] - dictIn[0];
		if(tmpltr4 < 0) tmpltr4 += 26;
		dict.ltr[2] = tmpltr4;

		tmpltr5 = dictIn[4] - dictIn[0];
		if(tmpltr5 < 0) tmpltr5 += 26;
		dict.ltr[3] = tmpltr5;
		
		dictionary.push_back(dict);

		if(infile.eof())
			break;
	}
	infile.close();

	vector<Word> bestWords = {};
	int count = 0;
	int most = 0;
	const int size = ltr1.size();
	for(int i = 0; i < size; i++) {
		count = 0;
		for(int j = i+1; j < size; j++) {
			if(	dictionary[i].word == dictionary[j].word ) {
				count++;
			}
		}
		if(count > most){
			most = count;
			bestWords.clear();
			bestWords.push_back({ltr1[i],(signed char)dictionary[i].ltr[0],(signed char)dictionary[i].ltr[1],(signed char)dictionary[i].ltr[2],(signed char)dictionary[i].ltr[3]});
		} else if(count == most){
			bestWords.push_back({ltr1[i],(signed char)dictionary[i].ltr[0],(signed char)dictionary[i].ltr[1],(signed char)dictionary[i].ltr[2],(signed char)dictionary[i].ltr[3]});
		}
	}
	
	auto stop = high_resolution_clock::now();
	auto duration = duration_cast<milliseconds>(stop - start);
	cout << duration.count() << "ms" << endl;

	string bestOut;
	for(Word bestWord : bestWords) {
		bestOut = (char)(bestWord.ltr1);

		bestWord.ltr2 += bestWord.ltr1;
		if(bestWord.ltr2 > 90) bestWord.ltr2 -= 26;
		bestOut += (char)bestWord.ltr2;

		bestWord.ltr3 += bestWord.ltr1;
		if(bestWord.ltr3 > 90) bestWord.ltr3 -= 26;
		bestOut += (char)bestWord.ltr3;

		bestWord.ltr4 += bestWord.ltr1;
		if(bestWord.ltr4 > 90) bestWord.ltr4 -= 26;
		bestOut += (char)bestWord.ltr4;

		bestWord.ltr5 += bestWord.ltr1;
		if(bestWord.ltr5 > 90) bestWord.ltr5 -= 26;
		bestOut += (char)bestWord.ltr5;
		
		cout << "Best word is: " << bestOut << " with " << most << " valid rotated words.\n";
	}
	
	return 0;
}