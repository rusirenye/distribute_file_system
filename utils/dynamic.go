package dynamic

import (
	"math/rand"
	"time"

	"distribute_file_system/dao"
	"distribute_file_system/models"
)

var nodetList []models.Node

func init() {
	nodetList = dao.GetStandardNodeList()
}

// getRandom return random float between 0.7 and 1.3
func getRandom() float32 {
	rand.Seed(time.Now().Unix())
	randomPara := rand.Intn(7)
	randomPara = randomPara + 7
	return randomPara / 10
}
func UpdateNodeStatus() {
	go func(){
		for {
			for _, node := range nodetList {
				// update target node by nodeList IP
				ip := node.Ip
				//node.DiskUsed *= getRandom()
				node.CpuUsed *= getRandom()
				node.MemoryUsed *= getRandom()
				node.BrandwidthUsed *= getRandom()
				dao.UpdateNodeStatus(node)
			}
			// update node network used percent every 100ms
			time.Sleep(500 * time.Millisecond)
		}
	}

}
