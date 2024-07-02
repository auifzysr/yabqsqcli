package domain

import "fmt"

type PubSubTopic struct {
	ProjectID string
	TopicID   string
}

func (t *PubSubTopic) Name() (string, error) {
	if t.ProjectID == "" {
		return "", fmt.Errorf("project id cannot be empty")
	}
	if t.TopicID == "" {
		return "", fmt.Errorf("topic id cannot be empty")
	}
	return fmt.Sprintf("projects/%s/topics/%s", t.ProjectID, t.TopicID), nil
}
