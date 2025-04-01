package main

import (
	"fmt"

	"github.com/rancher/machine/libmachine/drivers"
	"github.com/rancher/machine/libmachine/mcnflag"
	"github.com/rancher/machine/libmachine/state"
)

type FakeDriver struct {
	drivers.BaseDriver
}

func (f *FakeDriver) Create() error {
	return fmt.Errorf("NotImplemented")
}

func (f *FakeDriver) GetCreateFlags() []mcnflag.Flag {
	return []mcnflag.Flag{}
}

func (f *FakeDriver) GetSSHHostname() (string, error) {
	return "", fmt.Errorf("NotImplemented")
}

func (f *FakeDriver) GetState() (state.State, error) {
	return state.None, fmt.Errorf("NotImplemented")
}

func (f *FakeDriver) GetURL() (string, error) {
	return "", fmt.Errorf("NotImplemented")
}

func (f *FakeDriver) Kill() error {
	return fmt.Errorf("NotImplemented")
}

func (f *FakeDriver) PreCreateCheck() error {
	return fmt.Errorf("NotImplemented")
}

func (f *FakeDriver) Remove() error {
	return fmt.Errorf("NotImplemented")
}

func (f *FakeDriver) Restart() error {
	return fmt.Errorf("NotImplemented")
}

func (f *FakeDriver) SetConfigFromFlags(opts drivers.DriverOptions) error {
	return fmt.Errorf("NotImplemented")
}

func (f *FakeDriver) Start() error {
	return fmt.Errorf("NotImplemented")
}

func (f *FakeDriver) Stop() error {
	return fmt.Errorf("NotImplemented")
}

var _ drivers.Driver = &FakeDriver{}
