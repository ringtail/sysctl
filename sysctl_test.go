package sysctl

import (
	"testing"
)

func TestAll(t *testing.T) {
	kernel_settings := All()
	if kernel_settings != nil {
		t.Log("fetch kernel_settings successfully \n")
		t.Log(kernel_settings)
		return
	}
	t.Error("Failed to fetch any kernel_setttings")
}

func TestFind(t *testing.T) {
	kernel_settings := Find("dev.cdrom.debug")
	if kernel_settings != nil {
		t.Log("fetch kernel_settings successfully \n")
		t.Log(kernel_settings)
		return
	}
	t.Error("Failed to fetch any kernel_setttings")
}

func TestFindOneMore(t *testing.T) {
	kernel_settings := Find("dev.cdrom.debug", "vm.laptop_mode")
	if kernel_settings != nil {
		t.Log("fetch kernel_settings successfully \n")
		t.Log(kernel_settings)
		return
	}
	t.Error("Failed to fetch any kernel_setttings")
}

func TestApply(t *testing.T) {
	err := Apply("vm.laptop_mode", "1")
	if err != nil {
		t.Errorf("Failed to apply kernel settings %v", err.Error())
		return
	}
	t.Log("Apply kernel settings successfully")
}
