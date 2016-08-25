package utils

import (
	"distribute_file_system/dao"
	"distribute_file_system/log"
	"fmt"
	"io/ioutil"
	//	"math/rand"
	"path"
	"time"
)

func SaveBlockToNode(fileId, dir string) bool {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Errorf("read folder fail: " + dir)
		return false
	}
	log.Infof("1")
	for i, fileBlock := range files {
		log.Infof(fmt.Sprintf("%d", i))

		fileBlockPath := path.Join(dir, fileBlock.Name())
		log.Infof("block:" + fileBlockPath)
		//		fileBlockData, err := ioutil.ReadFile(fileBlockPath)
		if err != nil {
			log.Errorf("read file block fail: " + fileBlockPath)
			continue
		}
		// choose node
		for i := 0; i < 3; i++ {
			nodeList, _ := dao.GetNodeCandicates(fileId, fileBlock.Name())
			if len(nodeList) == 0 {
				break
			}
			log.Infof(fmt.Sprintf("i:%d len:%d", i, len(nodeList)))
			//			sed := time.Now().Unix()
			//			index := rand.Intn(len(nodeList))

			index := int(time.Now().Second() % (len(nodeList)))

			node := nodeList[index]
			log.Infof(fmt.Sprintf("index:%d", index))
			dao.AddBlockToNode(fileId, fileBlock.Name(), node.Ip, fileBlock.Size())
			//copy data to node

		}
	}
	return true
}
