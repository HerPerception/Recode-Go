package myDir

import "strings"

func PrintHistory(s string) string {
	maxHistory := 10
	commandHistory := ""
	//start := 0
	switch s {
	case "history":
		if Last != "" {
			History = append(History, Last)
		}
		if len(History) == 0 {
			return "No previous session."
		}
		if len(History) <= maxHistory {
			commandHistory = strings.Join(History, "\n")
		} else if len(History) > maxHistory {
			//History = append(History, Last)
			commandHistory = strings.Join(History[len(History)-10:], "\n")
		}
		return commandHistory

	case "clear":
		if History != nil && Last != "" {
			History = []string{}
			Last = ""
			return "Previous sessions cleared."
		}
		return "No previous sessions."
	}
	return "Press 'Enter' to return to previous page."
}
