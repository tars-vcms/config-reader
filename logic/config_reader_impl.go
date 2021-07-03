package logic

import (
	"github.com/tars-vcms/config-reader/repo/configWriter"
	"github.com/tars-vcms/config-reader/tars-protocol/vcms"
)

type ConfigReaderImpl struct {
}


func (c *ConfigReaderImpl) GetConfig(req *vcms.GetConfigReq, rsp *vcms.GetConfigRsp) (int32, error) {
	rsp.Status = vcms.ConfigStatus_Online
	rsp.Key = req.Key
	rsp.Value = configWriter.Impl.GetConfig(req.Key)
	return 0, nil
}

func (c *ConfigReaderImpl) PutConfig(req *vcms.PutConfigReq, rsp *vcms.PutConfigRsp) (int32, error) {
	panic("Not Implement Panic")
}