package utils

import "strings"

func FormatProjectName(rawName string) string {
	newName := strings.ReplaceAll(rawName, " ", "_")
	newName = strings.ReplaceAll(newName, "-", "_")
	return strings.TrimSpace(newName)
}
