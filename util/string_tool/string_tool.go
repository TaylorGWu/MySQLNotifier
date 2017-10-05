package string_tool

/*
	compare string
	if first string smaller than second string return -1,
	if equals than return 0, otherwise return 1.
*/
func CompareString(firstString string, secondString string) int {
	if firstString == secondString {
		return 0
	} else if firstString < secondString {
		return  -1
	} else  {
		return  1
	}
}
