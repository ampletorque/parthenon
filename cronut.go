package main

func main() {
	var crontabList [][]int
	var crontabPath string
	var crontabNextList [][]int

	crontabPath = cliParse()
	crontabList = crontabParse(crontabPath)
	crontabNextList = crontabGetNext(crontabList)
	toStdout(crontabNextList)
}

func cliParse() string {
	var crontabPath string

	return crontabPath
}

func toStdout(crontabNextList [][]int) {

	return
}
