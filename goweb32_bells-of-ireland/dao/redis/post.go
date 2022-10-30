package redis

import "goweb32_bells-of-ireland/models"

func GetPostIDsByPage(params *models.ParamPostData) ([]string, error) {

	//排序参数
	key := GetRedisKey(KeyPostTimeZSet)
	if params.Order == models.OrderByScore {
		key = GetRedisKey(KeyPostScoreZSet)
	}
	start := (params.Page - 1) * params.Size
	stop := start + params.Size - 1

	//获取分页参数
	return client.ZRevRange(key, start, stop).Result()
}
