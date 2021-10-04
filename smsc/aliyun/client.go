package aliyun

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/pkg/errors"
	"strings"
)

type Client interface {
	SendSms(msg *Message, telephones ...string) error
}

type client struct {
	kernel *dysmsapi20170525.Client
	config Config
}

func NewClient(config Config) (Client, error) {

	if config.AccessKeyId == "" || config.AccessKeySecret == "" {
		return nil, errors.New("create sms client must need access_key_id and access_key_secret")
	}
	if config.Endpoint == "" {
		config.Endpoint = "dysmsapi.aliyuncs.com"
	}

	cfg := &openapi.Config{
		AccessKeyId:     &config.AccessKeyId,
		AccessKeySecret: &config.AccessKeySecret,
		Endpoint:        &config.Endpoint,
	}
	c, err := dysmsapi20170525.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &client{kernel: c, config: config}, nil
}

// FailedInfo 发送失败的短信信息
type FailedInfo struct {
	Telephone string `json:"telephone"`
	Message   string `json:"message"`
}

func (c *client) SendSms(msg *Message, telephones ...string) error {
	phoneNumbers := strings.Join(telephones, ",")
	request := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  &phoneNumbers,
		SignName:      &msg.SignName,
		TemplateCode:  &msg.TemplateCode,
		TemplateParam: &msg.TemplateParamJson,
	}
	// 复制代码运行请自行打印 API 的返回值
	response, err := c.kernel.SendSms(request)
	_ = response
	if err != nil {
		return err
	}
	return nil
}
