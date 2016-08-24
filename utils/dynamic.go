package utils

import (
	"math/rand"
	"time"

	"distribute_file_system/dao"

	//"distribute_file_system/models"
)

// getRandom return random float between 0.7 and 1.3
func getRandom() float32 {
	rand.Seed(time.Now().Unix())

	var randomPara float32
	randomPara += float32(rand.Intn(7))
	randomPara = randomPara + 7.0
	randomPara = (randomPara / 10.0001)
	return randomPara
}
func UpdateNodeStatus() {
	go func() {
		nodeList, _ := dao.GetStandardNodeList()
		for {
			for _, node := range nodeList {
				node.CpuUsed *= getRandom()
				node.MemoryUsed *= getRandom()
				node.BrandwidthUsed *= getRandom()
				dao.UpdateNodeStatus(node)
				time.Sleep(500 * time.Millisecond)
			}
			// update node network used percent every 100ms
			time.Sleep(500 * time.Millisecond)
		}
	}()

}
