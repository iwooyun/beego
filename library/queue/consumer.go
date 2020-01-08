package queue

import (
	"beego/library/utils/base"
	"github.com/astaxie/beego/logs"
)

const (
	SuccessCode      = 0     //消费成功
	ParamError       = 22002 //参数错误
	DataEmpty        = 22003 //消息体为空
	TypeMissMatch    = 22004 //消息类型不匹配
	ConsumeForbidden = 22005 //消息禁止消费
	SourceMissMatch  = 22007 //消息来源不匹配
	SystemError      = 22100 //系统错误
	WriteFailed      = 22101 //写入失败
	UnexpectedError  = 22102 //未知错误
	MsgBodyTransErr  = 22103 //消息内容转换失败
)

// errorMessageMap 错误码与错误信息对应关系集合.
var errorMessageMap = map[uint16]string{
	SuccessCode:      "success",
	ParamError:       "param error",
	DataEmpty:        "message data empty",
	TypeMissMatch:    "message type miss match",
	SystemError:      "system error",
	WriteFailed:      "message data write failed",
	UnexpectedError:  "unexpected error",
	SourceMissMatch:  "consumer source miss match",
	ConsumeForbidden: "consumer forbidden",
	MsgBodyTransErr:  "message data trans error",
}

const DefaultMessageType = "DEFAULT_TYPE"

type newConsumerFunc func() IConsumer

// IConsumer 消费者接口.
type IConsumer interface {
	SetExtendErrorCodeMap(bc *BeeConsumer)
	ConsumeAble(bc *BeeConsumer) bool
	SetConsumeRouterMap(bc *BeeConsumer)
}

var adapters = make(map[string]newConsumerFunc)

// Register 注册消费者
func Register(name string, consumer newConsumerFunc) {
	if consumer == nil {
		panic("BeeConsumer: Register provide is nil")
	}
	if _, dup := adapters[name]; dup {
		panic("BeeConsumer: Register called twice for provider " + name)
	}
	adapters[name] = consumer
}

// MessageBody 消息内容.
type MessageBody = map[string]interface{}

// BeeConsumer 消费者基础结构体.
type BeeConsumer struct {
	Message          *Message
	ConsumeRouterMap map[string]func(MessageBody)
	ErrorCode        uint16
	ErrorMessage     string
	ErrorCodeArr     []uint16
	Adapter          newConsumerFunc
}

// ConsumeResult 消费结果集.
type ConsumeResult struct {
	ErrorCode    uint16 `json:"code"`
	ErrorMessage string `json:"msg"`
}

// Message 消息结构体.
type Message struct {
	Type    string      `json:"msg_type"`
	Content MessageBody `json:"msg_content"`
}

// NewBeeConsumer 创建消费者对象.
func NewBeeConsumer(adapterName string) *BeeConsumer {
	adapter, ok := adapters[adapterName]
	if !ok {
		logs.Error("BeeConsumer: unknown adapter name %q (forgotten Register?)", adapterName)
	}
	return &BeeConsumer{
		ConsumeRouterMap: make(map[string]func(messageBody MessageBody)),
		ErrorCode:        SuccessCode,
		ErrorCodeArr:     []uint16{WriteFailed, SystemError},
		Message:          &Message{Type: DefaultMessageType},
		Adapter:          adapter,
	}
}

//	AddErrorCodes 添加返回消费失败错误码到错误码集合.
func (c *BeeConsumer) AddErrorCodes(errorCodes []uint16) {
	c.ErrorCodeArr = append(c.ErrorCodeArr, errorCodes...)
}

// setErrorMessage 根据错误码，设置错误信息.
func (c *BeeConsumer) setErrorMessage() {
	if errorMsg, ok := errorMessageMap[c.ErrorCode]; ok {
		c.ErrorMessage = errorMsg
	} else {
		c.ErrorMessage = errorMessageMap[UnexpectedError]
	}
}

//	isSuccess 根据错误码，返回是否消费失败.
func (c *BeeConsumer) isSuccess() bool {
	if base.Contains(c.ErrorCodeArr, c.ErrorCode) != -1 {
		return false
	}
	return true
}

//	DataResult 消费成功，返回结果对象.
func (c *BeeConsumer) DataResult() *ConsumeResult {
	if c.ErrorMessage == "" {
		c.setErrorMessage()
	}

	if c.isSuccess() {
		c.ErrorCode = SuccessCode
	}

	return &ConsumeResult{
		ErrorCode:    c.ErrorCode,
		ErrorMessage: c.ErrorMessage,
	}
}

//	FailureResult 消费失败，设置消费返回对象.
func (c *BeeConsumer) FailureResult(errorCode uint16, errorMessage string, isReturn bool) interface{} {
	c.ErrorCode = errorCode
	if errorMessage == "" {
		c.setErrorMessage()
	} else {
		c.ErrorMessage = errorMessage
	}

	if isReturn {
		return c.DataResult()
	}

	return nil
}

//	checkMessage 校验消息及路由配置.
func (c *BeeConsumer) checkMessage() bool {
	if c.Message.Type == "" || c.Message.Content == nil {
		c.ErrorCode = ParamError
		return false
	}

	if _, ok := c.ConsumeRouterMap[c.Message.Type]; !ok {
		c.ErrorCode = TypeMissMatch
		return false
	}

	return true
}

//	invoke 根据消息类型路由，执行消费函数.
func (c *BeeConsumer) invoke() {
	defer func() {
		err := recover()
		if err != nil {
			logs.Error("消息执行消费函数失败， 错误信息 => [%s]", err)
		}
	}()
	invokeMethod := c.ConsumeRouterMap[c.Message.Type]
	invokeMethod(c.Message.Content)
}

//	AddConsumeRouterMap 添加消费路由映射.
func (c *BeeConsumer) AddConsumeRouterMap(callback func(messageBody MessageBody), messageType string) {
	if messageType == "" {
		messageType = DefaultMessageType
	}
	c.ConsumeRouterMap[messageType] = callback
}

//	Consume 根据消息类型路由，执行消费函数.
func (c *BeeConsumer) Consume(message *Message) bool {
	consumerAdapter := c.Adapter()
	c.Message = message
	//设置错误码
	consumerAdapter.SetExtendErrorCodeMap(c)
	//设置路由
	consumerAdapter.SetConsumeRouterMap(c)
	//校验消息及路由配置
	if !c.checkMessage() {
		return false
	}
	//是否可消费
	if !consumerAdapter.ConsumeAble(c) {
		return false
	}
	//根据消费路由，执行消费函数
	c.invoke()

	return true
}
