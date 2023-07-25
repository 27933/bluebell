package redis

// redis key

// redis key注意使用命名空间的方式,方便查询和拆分

const (
	Prefix = "bluebell:" // 项目key前缀
	// KeyPostTimeZSet key: post:time offset: 帖子ID score: 发帖时间
	//可以根据不同帖子ID获取其不同的发帖时间 并且可以对其进行排序
	KeyPostTimeZSet = "post:time" // zset;贴子及发帖时间

	// KeyPostScoreZSet key: post:score offset: 帖子ID score: 投票的分数
	// 根据不同帖子ID获取其投票分数 并且可以对其进行排序
	KeyPostScoreZSet = "post:score" // zset;贴子及投票的分数

	// KeyPostVotedZSetPF key: KeyPostVotedZSetPF+postID offset: 帖子ID score: 投票类型 赞成票(1) 反对票(-1) 不投票(0)
	// 根据不同帖子ID获取用户对该帖子的投票情况 可以统计不同帖子的赞成票和反对票的个数
	KeyPostVotedZSetPF = "post:voted:" // zset;记录用户及投票类型;

	// KeyCommunitySetPF key: community+社区ID value: 帖子ID
	KeyCommunitySetPF = "community:" // set;保存每个分区下帖子的id
)

// 给redis key加上前缀
func getRedisKey(key string) string {
	return Prefix + key
}
