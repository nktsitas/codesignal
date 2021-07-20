package main

func allLongestStrings(inputArray []string) []string {
	maxLength := 0

	toReturn := make([]string, 0)

	for _, iterString := range inputArray {
		if len(iterString) > maxLength {
			toReturn = make([]string, 0)

			toReturn = append(toReturn, iterString)

			maxLength = len(iterString)
		} else if len(iterString) == maxLength {
			toReturn = append(toReturn, iterString)
		}
	}

	return toReturn
}
