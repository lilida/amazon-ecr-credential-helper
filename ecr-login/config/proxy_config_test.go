package config

import (
	"testing"
	"io/ioutil"
	"fmt"
	//"os"
	"path/filepath"
	"github.com/stretchr/testify/assert"
	homedir "github.com/mitchellh/go-homedir"
)

func TestLoad(t *testing.T) {
	dir,_ :=homedir.Expand(GetCacheDir())
	proxyFile := filepath.Join(dir, "ecr_proxy.json")
	fmt.Println(proxyFile)
	err := ioutil.WriteFile(proxyFile, []byte(`{"proxies":{"test1":"ecr1"}}`), 0644)
	if err != nil {
		fmt.Println(err)
	}
	config, err := GetProxyConfig()
	if err != nil || config == nil {
		t.Error("Faile to read")

	} else {
		fmt.Print(config)
		assert.EqualValues(t, 1, len(config.Proxies))
		assert.EqualValues(t,"ecr1", GetRegistryURL("test1"))
	}

}
