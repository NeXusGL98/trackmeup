package jira

type JiraClient interface {
	createIssue(issue *Issue) error
}
