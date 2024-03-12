package adapter

import (
	"axxonsoft/data/entity"
	"axxonsoft/data/model"
	"strings"
)

func Task(task *entity.Task) model.Task {
	var headers string
	if task.ResponseHeaders != nil {
		headers = *task.ResponseHeaders
	}

	return model.Task{
		ID:          task.ID,
		Status:      task.Status,
		StatusCode:  task.ResponseStatusCode,
		Headers:     SplitHeaders(headers),
		Length:      task.ResponseLength,
		ErrorReason: task.ResponseErrorReason,
	}
}

func JoinHeaders(headers map[string]string) string {
	join := strings.Builder{}
	for key, value := range headers {
		join.WriteString(key)
		join.WriteString("=")
		join.WriteString(value)
		join.WriteString(";")
	}
	return join.String()
}

func SplitHeaders(headers string) map[string]string {
	headerPairs := strings.Split(headers, ";")

	headersMap := make(map[string]string)
	for _, pair := range headerPairs {
		before, after, _ := strings.Cut(pair, "=")
		if before == "" {
			continue
		}
		headersMap[before] = after
	}

	return headersMap
}
