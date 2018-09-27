package dmltogo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDmlToDo(t *testing.T) {
	res, err := DmlToGo("CREATE TABLE `total_data` (`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id', " +
		"`region` varchar(32) NOT NULL COMMENT 'region name, like zh; th; kepler'," +
		"`data_size` bigint NOT NULL DEFAULT '0' COMMENT 'data size;'," +
		"`createtime` datetime NOT NULL COMMENT 'create time;'," +
		"`comment` varchar(100) NOT NULL DEFAULT '' COMMENT 'comment'," +
		"PRIMARY KEY (`id`))")
	assert.Nil(t, err)
	assert.Equal(t, "package model\n\nimport (\n    \"database/sql\"\n    \"time\"\n\n    \"github.com/guregu/null\"\n)\n\nvar (\n    _ = time.Second\n    _ = sql.LevelDefault\n    _ = null.Bool{}\n)\n\n\ntype TotalDatum struct {\n    \n}\n\n// TableName sets the insert table name for this struct type\nfunc (t *TotalDatum) TableName() string {\n\treturn \"TotalData\"\n}\n", res)
}
