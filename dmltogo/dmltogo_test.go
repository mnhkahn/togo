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
	assert.Equal(t, "package model\n\nimport \"time\"\n\ntype TotalDatum struct {\n\tComment    string    `gorm:\"column:comment\" json:\"comment\"`\n\tCreatetime time.Time `gorm:\"column:createtime\" json:\"createtime\"`\n\tDataSize   int64     `gorm:\"column:data_size\" json:\"data_size\"`\n\tID         int       `gorm:\"column:id\" json:\"id\"`\n\tRegion     string    `gorm:\"column:region\" json:\"region\"`\n}\n\n// TableName sets the insert table name for this struct type\nfunc (t *TotalDatum) TableName() string {\n\treturn \"TotalData\"\n}\n", string(res))
}

func TestDmlToGoErrorSQL(t *testing.T) {
	res, err := DmlToGo("CREATE TABLE Persons (PersonID  int;);")
	assert.Nil(t, err)
	assert.NotEmpty(t, string(res))
}
