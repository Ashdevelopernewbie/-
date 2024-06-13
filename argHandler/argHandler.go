package arghandler

func ArgHandler(argument string) string {
	if argument == "" || argument != "-c" {
		return "invalid argument"
	}
	return "checking size of the file"
}
