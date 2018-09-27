# dmltogo
Generate go struct by create sql DML.

[Doc](https://godoc.org/github.com/mnhkahn/dmltogo)

### Example:

```
dmltogo.DmlToGo("CREATE TABLE `total_data` (`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id', " +
		"`region` varchar(32) NOT NULL COMMENT 'region name, like zh; th; kepler'," +
		"`data_size` bigint NOT NULL DEFAULT '0' COMMENT 'data size;'," +
		"`createtime` datetime NOT NULL COMMENT 'create time;'," +
		"`comment` varchar(100) NOT NULL DEFAULT '' COMMENT 'comment'," +
		"PRIMARY KEY (`id`))")
```
		
Output:

```

package model

import (
    "database/sql"
    "time"

    "github.com/guregu/null"
)

var (
    _ = time.Second
    _ = sql.LevelDefault
    _ = null.Bool{}
)


type TotalDatum struct {
    ID int `gorm:"column:id;primary_key" json:"id"`
    Region string `gorm:"column:region" json:"region"`
    DataSize int64 `gorm:"column:data_size" json:"data_size"`
    Createtime time.Time `gorm:"column:createtime" json:"createtime"`
    Comment string `gorm:"column:comment" json:"comment"`
    
}

// TableName sets the insert table name for this struct type
func (t *TotalDatum) TableName() string {
	return "TotalData"
}
```