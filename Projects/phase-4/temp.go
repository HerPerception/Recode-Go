package myDir

import "strings"

func PrintHistory(s string) string {
	maxHistory := 10
	commandHistory := ""
	newHistory := []string{}
	switch s {
	case "history":
		if len(History) == maxHistory {
			if Last != "" {
				start := len(History) - maxHistory
				History = append(History, Last)
				commandHistory = strings.Join(newHistory[start:], "\n")
			}
		}
		if commandHistory == "" && Last == "" {
			return "No previous sessions."
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
