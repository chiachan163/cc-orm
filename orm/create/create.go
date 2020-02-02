package create

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/chiachan163/cc-orm/orm/create/tpl"

	"github.com/henrylee2cn/goutil"

	"github.com/chiachan163/cc-orm/orm/info"
	"github.com/henrylee2cn/erpc/v6"
)

// MicroTpl template file name
const MicroTpl = "__tp-micro__tpl__.go"

// MicroGenLock the file is used to markup generated project
const MicroGenLock = "__tp-micro__gen__.lock"

func CreateModel(force, newdoc bool) {
	erpc.Infof("Generating project: %s", info.ProjPath())

	os.MkdirAll(info.AbsPath(), os.FileMode(0755))
	err := os.Chdir(info.AbsPath())
	if err != nil {
		erpc.Fatalf("[micro] Jump working directory failed: %v", err)
	}

	force = force || !goutil.FileExists(MicroGenLock)

	// creates base files
	if force {
		tpl.Create()
	}

	// read temptale file
	b, err := ioutil.ReadFile(MicroTpl)
	if err != nil {
		b = []byte(strings.Replace(__tpl__, "__PROJ_NAME__", info.ProjName(), -1))
	}
	// new project code
	proj := NewProject(b)
	//for _, v := range proj.models.mongo {
	//	//erpc.Infof("k::%s", k)
	//	erpc.Infof("v::%+v", v)
	//}
	//for _, v := range proj.models.mysql {
	//	//erpc.Infof("k::%s", k)
	//	erpc.Infof("v::%+v", v)
	//}
	//for k, v := range proj.codeFiles {
	//	erpc.Infof("k::%s", k)
	//	erpc.Infof("v::%+v", v)
	//}
	proj.Generator(force, force || newdoc)

	// write template file
	f, err := os.OpenFile(MicroTpl, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		erpc.Fatalf("[micro] Create files error: %v", err)
	}
	defer f.Close()
	f.Write(formatSource(b))

	tpl.RestoreAsset("./", MicroGenLock)

	erpc.Infof("Completed code generation!")
}
