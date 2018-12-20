package main

import "testing"

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func TestGetJiraIssueID(t *testing.T) {
	branchName := "DEV-537-remove-duplicate"
	jiraIssueID := getJiraIssueID(branchName)
	assertEqual(t, jiraIssueID, "DEV-537")
}

func TestFixCommitMessage(t *testing.T) {
	commitMessage := "Fix issues"
	jiraIssueID := "DEV-523"
	newCommitMessage := fixCommitMessage(commitMessage, jiraIssueID)
	assertEqual(t, newCommitMessage, jiraIssueID+" - "+commitMessage)
	commitMessage = jiraIssueID + commitMessage
	newCommitMessage = fixCommitMessage(commitMessage, jiraIssueID)
	assertEqual(t, newCommitMessage, "")
}
