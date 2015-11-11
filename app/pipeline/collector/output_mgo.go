package collector

import (
	"github.com/ktktcom/pholcus/common/mgo"
	"github.com/ktktcom/pholcus/common/util"
	"github.com/ktktcom/pholcus/config"
	"github.com/ktktcom/pholcus/logs"
	mgov2 "gopkg.in/mgo.v2"
)

/************************ MongoDB 输出 ***************************/

func init() {
	Output["mgo"] = func(self *Collector, dataIndex int) {
		var err error
		//连接数据库
		mgoSession := mgo.MgoPool.GetOne().(*mgo.MgoSrc)
		defer mgo.MgoPool.Free(mgoSession)

		var db = mgoSession.DB(config.MGO.DB)
		var namespace = util.FileNameReplace(self.namespace())
		var collections = make(map[string]*mgov2.Collection)
		var dataMap = make(map[string][]interface{})

		for _, datacell := range self.DockerQueue.Dockers[dataIndex] {
			subNamespace := util.FileNameReplace(self.subNamespace(datacell))
			if _, ok := collections[subNamespace]; !ok {
				collections[subNamespace] = db.C(namespace + "__" + subNamespace)
			}
			for k, v := range datacell["Data"].(map[string]interface{}) {
				datacell[k] = v
			}
			delete(datacell, "Data")
			delete(datacell, "RuleName")
			dataMap[subNamespace] = append(dataMap[subNamespace], datacell)
		}

		for k, v := range dataMap {
			err = collections[k].Insert(v...)
			if err != nil {
				logs.Log.Error("%v", err)
			}
		}
	}
}
