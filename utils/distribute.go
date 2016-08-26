package utils

import (
	"distribute_file_system/dao"
	"distribute_file_system/log"
	"fmt"
	"io/ioutil"
	//	"math/rand"
	"os"
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
		fileBlockData, err := ioutil.ReadFile(fileBlockPath)
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
			tmpFileName := fileId + "_" + fileBlock.Name()
			storageFilePath := path.Join(nodeToStoragePath[node.Ip], tmpFileName)
			err = ioutil.WriteFile(storageFilePath, 777)
			if err != nil {
				log.Errorf(storageFilePath + " block data storage operation fail")
			}
		}
	}
	//set dynamic file upload change

	return true
}

const (
	RootPath = "static/nodes" //all node root storage folder
)

var nodeToStoragePath map[string]string //node ip : node storage folder

// init node storage file block data directory
func InitNodeStorageDir() error {
	nodeList, err := dao.GetNodeList(nil)
	if err != nil {
		log.Debugf("can't init node storage diretory")
		return err
	}
	for _, node := range nodeList {
		ip := node.Ip
		dirPath := path.Join(RootPath, node.Ip)
		//chech dirpath is exist
		if !isDirExists(dirPath) {
			os.Mkdir(dirPath, 777) // create node storage folder
		}
		nodeToStoragePath[ip] = dirPath
	}

}

func isDirExists(path string) bool {
	fi, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}

	panic("not reached")
}
