package timer

import "time"

type InitTimer struct {
}

func (i *InitTimer) ConvertCreationTimeToStandardFormat(createdAt time.Time) string {
	// 将字符串解析为时间对象，使用数据库返回的时间格式
	//t, err := time.Parse("2006-01-02 15:04:05", createdAt)
	//if err != nil {
	//	// 处理解析错误
	//	return ""
	//}

	// 将时间对象格式化为标准格式
	standardFormat := createdAt.Format("2006-01-02 15:04:05")
	return standardFormat
}
