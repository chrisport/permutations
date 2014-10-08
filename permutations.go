package main

func permutations(input string) []string {
	if len(input) == 1 {
		return []string{input}
	}

	var merge = func(ins []rune, c rune) (result []string) {
		for i := 0; i <= len(ins); i++ {
			result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
		}
		return
	}

	runes := []rune(input)
	subPermutations := permutations(string(runes[0:len(input) - 1]))

	result := []string{}
	for _, s := range subPermutations {
		result = append(result, merge([]rune(s), runes[len(input)-1])...)
	}

	return result
}
