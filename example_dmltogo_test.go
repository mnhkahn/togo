// Package dmltogo
package dmltogo_test

import (
	"log"

	"github.com/mnhkahn/dmltogo"
)

func ExampleNewLogger() {
	res, err := dmltogo.DmlToGo("CREATE TABLE `total_data` (`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id', " +
		"`region` varchar(32) NOT NULL COMMENT 'region name, like zh; th; kepler'," +
		"`data_size` bigint NOT NULL DEFAULT '0' COMMENT 'data size;'," +
		"`createtime` datetime NOT NULL COMMENT 'create time;'," +
		"`comment` varchar(100) NOT NULL DEFAULT '' COMMENT 'comment'," +
		"PRIMARY KEY (`id`))")
	log.Println(res, err)
}
