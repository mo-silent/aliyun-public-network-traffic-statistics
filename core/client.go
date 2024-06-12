package core

import (
	cms20190101 "github.com/alibabacloud-go/cms-20190101/v8/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/credentials-go/credentials"
)

type Client interface {
	CreateClient() (*cms20190101.Client, error)
}

type ClientConfig struct {
	Type            string `json:"type"`
	AccessKeyId     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	RoleArn         string `json:"role_arn"`
	RoleSessionName string `json:"role_session_name"`
	RoleName        string `json:"role_name"`
	SecurityToken   string `json:"security_token"`
}

// CreateClient
// Description: initialize account client
//
// @return Client
//
// @throws Exception
func (cc ClientConfig) CreateClient() (client *cms20190101.Client, err error) {
	// Get Credential
	credential, err := createCredential(cc)
	if err != nil {
		return nil, err
	}

	// set openapi config
	openapiConfig := new(openapi.Config).SetCredential(credential)
	// Endpoint: https://api.aliyun.com/product/Cms
	openapiConfig.Endpoint = tea.String("metrics.cn-hangzhou.aliyuncs.com")
	// cms client
	client = &cms20190101.Client{}
	client, err = cms20190101.NewClient(openapiConfig)
	return
}

func createCredential(cc ClientConfig) (credential credentials.Credential, err error) {
	var config *credentials.Config
	switch cc.Type {
	case "access_key":
		config = new(credentials.Config).
			SetType("access_key").
			SetAccessKeyId(cc.AccessKeyId).
			SetAccessKeySecret(cc.AccessKeySecret)
	case "sts":
		config = new(credentials.Config).
			SetType("sts").
			SetAccessKeyId(cc.AccessKeyId).
			SetAccessKeySecret(cc.AccessKeySecret).
			SetSecurityToken(cc.SecurityToken)
	case "ecs_ram_role":
		config = new(credentials.Config).
			SetType("ecs_ram_role").
			// 选填，该ECS角色的角色名称，不填会自动获取，但是建议加上以减少请求次数
			SetRoleName(cc.RoleName)
	case "ram_role_arn":
		config = new(credentials.Config).
			SetType("ram_role_arn").
			SetAccessKeyId(cc.AccessKeyId).
			SetAccessKeySecret(cc.AccessKeySecret).
			SetRoleArn(cc.RoleArn).
			SetRoleSessionName(cc.RoleSessionName).
			SetRoleSessionExpiration(3600)
	}
	credential, err = credentials.NewCredential(config)
	return
}
