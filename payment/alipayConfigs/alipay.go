package alipayConfigs

type alipayConf struct {
	AppId        string
	PartnerId    string
	AliPublicKey string
	PrivateKey   string
}

var Confs map[string]alipayConf

func init() {
	Confs = make(map[string]alipayConf)

}
func GetAppConf(appId string) alipayConf {
	return Confs[appId]
}
