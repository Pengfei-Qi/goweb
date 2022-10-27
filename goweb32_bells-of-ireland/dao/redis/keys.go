package redis

//redis 注意使用命名空间, 方便查询和使用

const (
	Prefix                 = "bellsIreland:"
	KeyPostTimeZSet        = "post:time"   //帖子及发帖时间
	KeyPostScoreZSet       = "post:score"  //帖子及分数
	KeyPostVotedZSetPrefix = "post:voted:" //记录用户及投票类型, 参数是post id
)

func GetRedisKey(key string) string {
	return Prefix + key
}
