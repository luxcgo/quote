package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"

	_ "github.com/boyter/go-string"
)

func main() {
	const needle = "This is not the web page you are looking for"

	res := (StringPermutation("xor"))
	for _, v := range res {
		// s, _ := GetURLContent("https://github.com/iloveaa")
		// if !strings.Contains(s, needle) {
		// 	log.Println(v)
		// }
		fmt.Println(v)

	}
}

func StringPermutation(str string) []string {
	permutations := []string{}
	if str == "" {
		return permutations
	}
	chars := []rune(str)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	visited := make([]bool, len(chars))
	stringPermutationDfs(chars, visited, []rune{}, &permutations)
	return permutations
}

func stringPermutationDfs(chars []rune, visited []bool, permutation []rune, permutations *[]string) {
	if len(permutation) == len(chars) {
		*permutations = append(*permutations, string(permutation))
	}

	for i, char := range chars {
		if visited[i] {
			continue
		}

		if i > 0 && char == chars[i-1] && !visited[i-1] {
			continue
		}

		visited[i] = true
		permutation = append(permutation, char)
		stringPermutationDfs(chars, visited, permutation, permutations)
		permutation = permutation[:len(permutation)-1]

		visited[i] = false
	}
}

func GetURLContent(url string) (string, error) {
	webUserAgent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", webUserAgent)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	content := string(respBody)
	return content, nil
}
