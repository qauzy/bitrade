package config

import (
	"bitrade/core/log"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

// 监听nacos配置
func getNacosConfig() (result int) {
	// 从控制台命名空间管理的"命名空间详情"中拷贝 End Point、命名空间 ID
	var endpoint = "127.0.0.1:8848/nacos"
	var namespaceId = ""
	// 推荐使用 RAM 用户的 accessKey、secretKey
	//var accessKey = "${accessKey}"
	//var secretKey = "${secretKey}"
	clientConfig := constant.ClientConfig{
		Endpoint:    endpoint,
		NamespaceId: namespaceId,
		//AccessKey:      accessKey,
		//SecretKey:      secretKey,
		TimeoutMs:      10 * 1000, //http请求超时时间，单位毫秒
		ListenInterval: 30 * 1000, //监听间隔时间，单位毫秒（仅在ConfigClient中有效）
		BeatInterval:   5 * 1000,  //心跳间隔时间，单位毫秒（仅在ServiceClient中有效）
	}

	// 至少一个
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "127.0.0.1",
			ContextPath: "/nacos",
			Port:        8848,
		},
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	var dataId = "sungyTest"
	var group = "DEFAULT_GROUP"

	// 获取配置
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group})

	fmt.Println("Get config：" + content)
	if len(content) == 0 {
		log.Error("getNacosConfigData - LoadErr , len is 0")
		result = 0
	}
	//监听nacos配置
	go ListenConfig()
	return
}

// 监听nacos配置
func ListenConfig() {
	log.Info("开启nacos配置监听")
	// 从控制台命名空间管理的"命名空间详情"中拷贝 End Point、命名空间 ID
	var endpoint = "127.0.0.1:8848/nacos"
	clientConfig := constant.ClientConfig{
		Endpoint:    endpoint,
		NamespaceId: "",
		//AccessKey:      accessKey,
		//SecretKey:      secretKey,
		TimeoutMs:      10 * 1000, //http请求超时时间，单位毫秒
		ListenInterval: 30 * 1000, //监听间隔时间，单位毫秒（仅在ConfigClient中有效）
		BeatInterval:   5 * 1000,  //心跳间隔时间，单位毫秒（仅在ServiceClient中有效）
	}

	// 至少一个
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "127.0.0.1",
			ContextPath: "/nacos",
			Port:        8848,
		},
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	var dataId = "sungyTest"
	var group = "DEFAULT_GROUP"
	_ = configClient.ListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("ListenConfig group:" + group + ", dataId:" + dataId + ", data:" + data)
			if len(data) == 0 {
				log.Error("getNacosConfigData - LoadErr , len is 0")
			}
		},
	})
}
