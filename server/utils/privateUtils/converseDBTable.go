package privateUtils

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

/*
*
@author lby
@description 将数据库表中字段转化为struct
*/
func ConverseTableField(name string, structBody *interface{}) error {
	// 首先检查该存放结构体的目标文件夹中是否有此文件，有就更新没有就创建
	isExist := global.GVA_DB.Migrator().HasTable(name)
	if !isExist {
		return errors.New("未查找到表信息")
	}

	return nil
}
