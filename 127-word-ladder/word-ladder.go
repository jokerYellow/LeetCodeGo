package _27_word_ladder

func ladderLength(beginWord string, endWord string, wordList []string) int {
	reached := map[string]struct{}{}
	reached[beginWord] = struct{}{}
	wordList = append(wordList, beginWord)
	distance := 1
	for {
		if _, ok := reached[endWord]; ok {

		}
	}
	return distance
}
