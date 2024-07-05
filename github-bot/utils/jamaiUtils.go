package utils

import (
	"github.com/wenjielee1/github-bot/models"
)

func GetColumnMessage(columnId string) []models.Message {
	const prompt = `
# Instructions

Based on the content provided, categorize the issue and provide the appropriate labels from the following options:

- Labels: "BUG", "FEATURE", "DOCUMENTATION", "ENHANCEMENT", "GOOD FIRST ISSUE", "HELP WANTED", "INVALID", "QUESTION"
- Priority: "LOW", "MEDIUM", "HIGH", "CRITICAL"

# Examples

## Example 1
### Issue Body
Dummy issue #1

### Response
{
  "labels": ["BUG", "HELP WANTED"],
  "priority": "HIGH",
  "response": "Jambo! I am Jambu, your github assistant. We appreciate your report. It seems there's a critical bug that needs immediate attention. Our team will prioritize this and work on a fix. Thank you for your help!"
}

## Example 2
### Issue Body
IssueBody2

### Response
{
  "labels": ["FEATURE", "GOOD FIRST ISSUE"],
  "priority": "MEDIUM",
  "response": "Thank you for the feature suggestion! This is a great idea for a first-time contributor to "Jam" on. We will add it to our development roadmap."
}

# Your Task

Analyze the issue described by User Input and respond in the same format as the examples above.

Ensure your response is JSON-friendly for parsing and includes both key-value pairs for "labels", "priority", and "response". Do NOT add any additional words or content other than the specified to make your response parse-able. Do NOT use markdown syntax for your response.

# User Input
${IssueBody}
`
	if columnId == "IssueResponse" {
		return []models.Message{
			{
				Role:    "system",
				Content: "You are Jambu, a github issue bot. Keep your responses brief and short and adhere to the response templates given to you. You will not mention anything else other than the requested response. In your responses, provide a quick pun of Jam or your own name.",
			},
			{
				Role:    "user",
				Content: prompt,
			},
		}
	} else if columnId == "PullReqResponse" {
		const changelogPrompt = `# Instructions

Based on the content provided, suggest how the "CHANGELOG.md" could be updated. The "PullReqBody" will be a git diff.

# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

- Each section is broken into:

  - "ADDED": New features.
  - "CHANGED / FIXED": Changes in existing functionality, or any bug fixes.
  - "DEPRECATED": Soon-to-be removed features.
  - "REMOVED": Removed features.
  - "SECURITY": Anything related to vulnerabilities.

- The version number mentioned here refers to the cloud version.

# Examples

## Example 1

### Pull Request Body
File: dummy-file-1.txt
Changes: @@ -0,0 +1 @@
+This is a dummy file for branch dummy-branch-1
\ No newline at end of file

File: dummy-secrets.txt
Changes: @@ -0,0 +1 @@
+SECRET_KEY=12345-abcde-67890-fghij
\ No newline at end of file

### Your response:

# Suggested Changelog:

## [Unreleased]

### ADDED

"Embeddings" endpoint

- Get a vector representation of a given input that can be easily consumed by machine learning models and algorithms. Note that the vectors are NOT normalized.
- Similar to OpenAI's embeddings endpoint
- Resolves #86

# Your Task

Analyze the git diff described in User Input and suggest appropriate updates to the "CHANGELOG.md" in the same format as the examples above.

Ensure your response is clear and concise, providing meaningful and insightful updates that accurately reflect the changes described in the git diff.

# User Input
${PullReqBody}`

		return []models.Message{
			{
				Role:    "system",
				Content: "You are a github bot managing pull requests. Keep your responses brief and short.",
			},
			{
				Role:    "user",
				Content: changelogPrompt,
			},
		}
	} else if columnId == "PullReqSecretsResponse" {
		const checkSecretsPrompt = `# Instructions

Based on the diff provided, check if there are any sensitive keys, secrets, passwords, or information accidentally added. Note that the leak can be in any type of variable name or type. If there is, provide the commit SHA where it was leaked, and the suspected section and file name.

# Response Template

Your response must be in the template of:

{
  "leak": true or false,
  "commit": "the SHA of the leaked commit, if any.",
  "response": "The file name, your response on the suspected leak."
}

Your entire response must adhere to a JSON-friendly format for parsing, that includes both key-value pairs.

# Examples

## Example 1
### Pull Request Secrets Body
PullReqSecretsBody1

### Response

{
  "leak": true,
  "commit": "abc123def456",
  "response": "Hello! I am Jambu, your github assistant. It seems that you made a very bad mistake! I suspect a secret key leaked in dummy-secrets.txt. If this is not a false positive, please squash your commits and make a PR!"
}

## Example 2
### Pull Request Secrets Body
PullReqSecretsBody2

### Response

{
  "leak": false,
  "commit": "",
  "response": "No suspected leaks found."
}

# Your Task

Analyze the git diff described in the User Input and respond in the same format as the examples above.

Ensure your response is clear and concise, providing meaningful and accurate information about any suspected leaks. Adhere to the JSON-friendly format for parsing.

# User Input
${PullReqSecretsBody}`

		return []models.Message{
			{
				Role:    "system",
				Content: "You are Jambu, a github bot. Your job is to find if the provided content contains any secrets, keys, passwords or sensitive information. Keep your responses brief and short and adhere to the response templates given to you. You will not mention anything else other than the requested response. In your responses, it should include puns with Jam and your name. ",
			},
			{
				Role:    "user",
				Content: checkSecretsPrompt,
			},
		}

	}
	return nil
}