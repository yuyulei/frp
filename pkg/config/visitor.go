// Copyright 2018 fatedier, fatedier@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"fmt"
	"reflect"

	"gopkg.in/ini.v1"
)

// DefaultVisitorConf creates a empty VisitorConf object by visitorType.
// If visitorType doesn't exist, return nil.
func DefaultVisitorConf(visitorType string) VisitorConf {
	v, ok := VisitorConfTypeMap[visitorType]
	if !ok {
		return nil
	}

	return reflect.New(v).Interface().(VisitorConf)
}

// Visitor loaded from ini
func NewVisitorConfFromIni(prefix string, name string, section *ini.Section) (VisitorConf, error) {
	// section.Key: if key not exists, section will set it with default value.
	visitorType := section.Key("type").String()

	if visitorType == "" {
		return nil, fmt.Errorf("visitor [%s] type shouldn't be empty", name)
	}

	conf := DefaultVisitorConf(visitorType)
	if conf == nil {
		return nil, fmt.Errorf("visitor [%s] type [%s] error", name, visitorType)
	}

	if err := conf.UnmarshalFromIni(prefix, name, section); err != nil {
		return nil, fmt.Errorf("visitor [%s] type [%s] error", name, visitorType)
	}

	if err := conf.Check(); err != nil {
		return nil, err
	}

	return conf, nil
}

// Base
func (cfg *BaseVisitorConf) GetBaseInfo() *BaseVisitorConf {
	return cfg
}

func (cfg *BaseVisitorConf) compare(cmp *BaseVisitorConf) bool {
	if cfg.ProxyName != cmp.ProxyName ||
		cfg.ProxyType != cmp.ProxyType ||
		cfg.UseEncryption != cmp.UseEncryption ||
		cfg.UseCompression != cmp.UseCompression ||
		cfg.Role != cmp.Role ||
		cfg.Sk != cmp.Sk ||
		cfg.ServerName != cmp.ServerName ||
		cfg.BindAddr != cmp.BindAddr ||
		cfg.BindPort != cmp.BindPort {
		return false
	}
	return true
}

func (cfg *BaseVisitorConf) check() (err error) {
	if cfg.Role != "visitor" {
		err = fmt.Errorf("invalid role")
		return
	}
	if cfg.BindAddr == "" {
		err = fmt.Errorf("bind_addr shouldn't be empty")
		return
	}
	if cfg.BindPort <= 0 {
		err = fmt.Errorf("bind_port is required")
		return
	}
	return
}

func (cfg *BaseVisitorConf) decorate(prefix string, name string, section *ini.Section) error {

	// proxy name
	cfg.ProxyName = prefix + name

	// server_name
	cfg.ServerName = prefix + cfg.ServerName

	// bind_addr
	if cfg.BindAddr == "" {
		cfg.BindAddr = "127.0.0.1"
	}

	return nil
}

// SUDP
var _ VisitorConf = &SUDPVisitorConf{}

func (cfg *SUDPVisitorConf) Compare(cmp VisitorConf) bool {
	cmpConf, ok := cmp.(*SUDPVisitorConf)
	if !ok {
		return false
	}

	if !cfg.BaseVisitorConf.compare(&cmpConf.BaseVisitorConf) {
		return false
	}

	// Add custom login equal, if exists

	return true
}

func (cfg *SUDPVisitorConf) UnmarshalFromIni(prefix string, name string, section *ini.Section) error {
	err := section.MapTo(cfg)
	if err != nil {
		return err
	}

	err = cfg.BaseVisitorConf.decorate(prefix, name, section)
	if err != nil {
		return err
	}

	// Add custom logic unmarshal, if exists

	return nil
}

func (cfg *SUDPVisitorConf) Check() (err error) {
	if err = cfg.BaseVisitorConf.check(); err != nil {
		return
	}

	// Add custom logic validate, if exists

	return
}

// STCP
var _ VisitorConf = &STCPVisitorConf{}

func (cfg *STCPVisitorConf) Compare(cmp VisitorConf) bool {
	cmpConf, ok := cmp.(*STCPVisitorConf)
	if !ok {
		return false
	}

	if !cfg.BaseVisitorConf.compare(&cmpConf.BaseVisitorConf) {
		return false
	}

	// Add custom login equal, if exists

	return true
}

func (cfg *STCPVisitorConf) UnmarshalFromIni(prefix string, name string, section *ini.Section) error {
	err := section.MapTo(cfg)
	if err != nil {
		return err
	}

	err = cfg.BaseVisitorConf.decorate(prefix, name, section)
	if err != nil {
		return err
	}

	// Add custom logic unmarshal, if exists

	return nil
}

func (cfg *STCPVisitorConf) Check() (err error) {
	if err = cfg.BaseVisitorConf.check(); err != nil {
		return
	}

	// Add custom logic validate, if exists

	return
}

// XTCP
var _ VisitorConf = &XTCPVisitorConf{}

func (cfg *XTCPVisitorConf) Compare(cmp VisitorConf) bool {
	cmpConf, ok := cmp.(*XTCPVisitorConf)
	if !ok {
		return false
	}

	if !cfg.BaseVisitorConf.compare(&cmpConf.BaseVisitorConf) {
		return false
	}

	// Add custom login equal, if exists

	return true
}

func (cfg *XTCPVisitorConf) UnmarshalFromIni(prefix string, name string, section *ini.Section) error {
	err := section.MapTo(cfg)
	if err != nil {
		return err
	}

	err = cfg.BaseVisitorConf.decorate(prefix, name, section)
	if err != nil {
		return err
	}

	// Add custom logic unmarshal, if exists

	return nil
}

func (cfg *XTCPVisitorConf) Check() (err error) {
	if err = cfg.BaseVisitorConf.check(); err != nil {
		return
	}

	// Add custom logic validate, if exists

	return
}
