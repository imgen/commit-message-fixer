package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	filePath := os.Args[1]
	commitMessage := getFileAsString(filePath)

	branchName := getBranchName()
	if strings.HasPrefix(branchName, "master") {
		return
	}
	jiraIssueID := getJiraIssueID(branchName)

	if commitMessage = fixCommitMessage(commitMessage, jiraIssueID); commitMessage != "" {
		println("Commit message doesn't start with JIRA issue ID, correcting...")
		if err := ioutil.WriteFile(filePath, []byte(commitMessage), 0644); err != nil {
			println("Failed to write file " + filePath + ", error: " + err.Error())
		} else {
			println("Commit message fixed")
		}
	} else {
		println("Commit message is ok")
	}
}

func fixCommitMessage(commitMessage string, jiraIssueID string) string {
	if strings.HasPrefix(commitMessage, jiraIssueID) {
		return ""
	}
	return jiraIssueID + " - " + commitMessage
}

func getFileAsString(commitFilePath string) string {
	content, _ := ioutil.ReadFile(commitFilePath)
	commitMessage := string(content)
	return commitMessage
}

func getBranchName() string {
	cmdName := "git"
	cmdArgs := strings.Split("rev-parse --abbrev-ref HEAD", " ")
	var (
		cmdOut []byte
		err    error
	)
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).CombinedOutput(); err != nil {
		strOutput := string(cmdOut)
		println(strOutput)
		os.Exit(1)
	}

	strOutput := string(cmdOut)
	branchName := strings.Split(strOutput, "\n")[0]
	println("Branch name is " + branchName)
	return branchName
}

func getJiraIssueID(branchName string) string {
	parts := strings.Split(branchName, "-")
	jiraIssueID := parts[0] + "-" + parts[1]
	println("Jira issue id is " + jiraIssueID)
	return jiraIssueID
}
