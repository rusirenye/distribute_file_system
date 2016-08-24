package utils

import (
	"distribute_file_system/dao"
	"distribute_file_system/log"
	"io/ioutil"
	"math/rand"
	"path"
	"time"
)

func SaveBlockToNode(fileId, dir string) bool {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Errorf("read folder fail: " + dir)
		return false
	}
	for _, fileBlock := range files {
		fileBlockPath := path.Join(dir, fileBlock.Name())
		//		fileBlockData, err := ioutil.ReadFile(fileBlockPath)
		if err != nil {
			log.Errorf("read file block fail: " + fileBlockPath)
			continue
		}
		// choose node
		for i := 0; i < 3; i++ {
			nodeList, _ := dao.GetNodeCandicates(fileId, fileBlock.Name())
			sed := time.Now().Unix()
			rand.Seed(sed)
			index := rand.Intn(len(nodeList))
			node := nodeList[index]
			dao.AddBlockToNode(fileId, fileBlock.Name(), node.Ip, fileBlock.Size())
			//copy data to node

		}
	}
	return true
}
