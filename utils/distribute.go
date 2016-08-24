package distribute

import (
	"io/ioutil"
	"path"
	"math/rand"
	"time"
	"distribute_file_system/log"
)
const(
	REPLICA = 3
)

func SaveBlockToNode(fileId, dir string) bool {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Errorf("read folder fail: " + dir )
		return false
	}
	for _, fileBlock := range files {
		fileBlockPath := path.Join(dir, fileBlock.Name())
		fileBlockData, err := ioutil.ReadFile(fileBlockPath)
		if err != nil {
			log.Errorf("read file block fail: " + fileBlockPath)
			continue
		}
		// choose node
		for i:=range 1...REPLICA {
			//nodeList,err:=dao.GetNodeCandicates(fileId,fileBlock.Name())
			sed:=time.Now().Unix() 
			rand.Seed(sed)
			index:=rand.Intn(len(nodeList))
			err:= dao.AddBlockToNode(fileId,fileBlock.Name())
			//copy data to node 
			node:=nodeList[index]
			ip:=node.Ip
						
		}
	}
}
