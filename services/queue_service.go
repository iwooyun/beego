package services

import (
	"beego/library/queue"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/goinggo/mapstructure"
)

var queueServiceImpl IQueueService

type IQueueService interface {
	Consume(messageStr string, consumer *queue.BeeConsumer, callback func(message *queue.Message)) bool
}

type QueueService struct {
}

func NewQueueService() IQueueService {
	if queueServiceImpl == nil {
		queueServiceImpl = &QueueService{}
	}

	return queueServiceImpl
}

// Consume 消息消费.
func (c *QueueService) Consume(messageStr string, IConsumer *queue.BeeConsumer, callback func(message *queue.Message)) bool {
	if messageStr == "" {
		IConsumer.ErrorCode = queue.ParamError
		return false
	}
	var messageMap map[string]interface{}
	err := json.Unmarshal([]byte(messageStr), &messageMap)
	if err != nil {
		logs.Error("消息内容json解析失败， 错误信息 => {%s}", err)
		IConsumer.ErrorCode = queue.ParamError
		return false
	}

	if _, ok := messageMap["msg_content"]; !ok {
		messageMap = map[string]interface{}{
			"msg_type":    queue.DefaultMessageType,
			"msg_content": messageMap,
		}
	}

	var messageBody queue.MessageBody
	err = mapstructure.Decode(messageMap["msg_content"], &messageBody)
	if err != nil {
		logs.Error("消息对象实例化失败， 错误信息 => {%s}", err)
		IConsumer.ErrorCode = queue.ParamError
		return false
	}

	message := &queue.Message{
		Type:    messageMap["msg_type"].(string),
		Content: messageBody,
	}

	callback(message)

	return IConsumer.Consume(message)
}
