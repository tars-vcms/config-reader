// Package vcms comment
// This file was generated by tars2go 1.1.4
// Generated from config_reader.tars
package vcms

import (
	"fmt"

	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8

type ConfigStatus int32

const (
	ConfigStatus_Online  = 1
	ConfigStatus_Offline = 2
)

// ConfigCacheEntity struct implement
type ConfigCacheEntity struct {
	Key    string       `json:"Key"`
	Status ConfigStatus `json:"Status"`
	Value  string       `json:"Value"`
}

func (st *ConfigCacheEntity) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *ConfigCacheEntity) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_string(&st.Key, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32((*int32)(&st.Status), 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Value, 2, true)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *ConfigCacheEntity) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require ConfigCacheEntity, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *ConfigCacheEntity) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.Key, 0)
	if err != nil {
		return err
	}

	err = _os.Write_int32(int32(st.Status), 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Value, 2)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *ConfigCacheEntity) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// PutConfigRsp struct implement
type PutConfigRsp struct {
	Config ConfigCacheEntity `json:"Config"`
	TTL    int32             `json:"TTL"`
}

func (st *PutConfigRsp) ResetDefault() {
	st.Config.ResetDefault()
	st.TTL = 86400
}

//ReadFrom reads  from _is and put into struct.
func (st *PutConfigRsp) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = st.Config.ReadBlock(_is, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.TTL, 1, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *PutConfigRsp) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require PutConfigRsp, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *PutConfigRsp) WriteTo(_os *codec.Buffer) error {
	var err error

	err = st.Config.WriteBlock(_os, 0)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.TTL, 1)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *PutConfigRsp) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// PutConfigReq struct implement
type PutConfigReq struct {
	Key string `json:"Key"`
	TTL int32  `json:"TTL"`
}

func (st *PutConfigReq) ResetDefault() {
	st.TTL = 86400
}

//ReadFrom reads  from _is and put into struct.
func (st *PutConfigReq) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_string(&st.Key, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.TTL, 1, true)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *PutConfigReq) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require PutConfigReq, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *PutConfigReq) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.Key, 0)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.TTL, 1)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *PutConfigReq) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// GetConfigRsp struct implement
type GetConfigRsp struct {
	Key    string       `json:"Key"`
	Status ConfigStatus `json:"Status"`
	Value  string       `json:"Value"`
}

func (st *GetConfigRsp) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *GetConfigRsp) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_string(&st.Key, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32((*int32)(&st.Status), 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Value, 2, true)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *GetConfigRsp) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require GetConfigRsp, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *GetConfigRsp) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.Key, 0)
	if err != nil {
		return err
	}

	err = _os.Write_int32(int32(st.Status), 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Value, 2)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *GetConfigRsp) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}

// GetConfigReq struct implement
type GetConfigReq struct {
	Key string `json:"Key"`
}

func (st *GetConfigReq) ResetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *GetConfigReq) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.ResetDefault()

	err = _is.Read_string(&st.Key, 0, true)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *GetConfigReq) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.ResetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require GetConfigReq, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(_is)
	if err != nil {
		return err
	}

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *GetConfigReq) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.Key, 0)
	if err != nil {
		return err
	}

	_ = err

	return nil
}

//WriteBlock encode struct
func (st *GetConfigReq) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(_os)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}