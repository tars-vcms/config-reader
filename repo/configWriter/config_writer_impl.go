package configWriter

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	"github.com/tars-vcms/vcms-common/tars-protocol/vcmsconfig"
)
var comm *tars.Communicator
var app *vcmsconfig.ConfigWriter
type impl struct {}
var Impl *impl
func init() {
	comm = tars.NewCommunicator()
	obj := fmt.Sprintf("") //TODO 添加服务地址
	app = new(vcmsconfig.ConfigWriter)
	comm.StringToProxy(obj, app)
}

func(this *impl) GetConfig(key string) string{
	//TODO: 扩充
	req := new(vcmsconfig.GetConfigReq)
	rsp := new(vcmsconfig.GetConfigRsp)
	req.Ident = key
	_ , err := app.GetConfig(req, rsp)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return rsp.Config.Comment
}





