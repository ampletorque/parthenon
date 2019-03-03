package main

func main() {
	var crontabRaw [][]string
	var crontabParsedList []map[string]string
	var crontabPath string
	var crontabNextList []map[string]string

	crontabPath = cliParse()
	crontabRaw = crontabRead(crontabPath)
	crontabParsedList = crontabParse(crontabRaw)
	crontabNextList = crontabGetNext(crontabParsedList)
	toStdout(crontabNextList)
}

func cliParse() string {
	var crontabPath string
	// check if arg and save, else stdin
	crontabPath = "/Users/andrewdavid/go/src/github.com/ampletorque/parthenon/crontab_1_parthenon"

	return crontabPath
}

func toStdout(crontabNextList []map[string]string) {

	return
}
