package kafka

const ConsumerGroupTask string = "group-task"
const ConsumerGroupBackup string = "group-backup"
const TopicTaskAllRequest string = "task-all-request"
const TopicTaskAllResponse string = "task-all-response"

type TaskAllRequestMessage struct {
	RequestId string
}

type TaskAllResponseMessage struct {
	RequestId string
	Data      string
}
