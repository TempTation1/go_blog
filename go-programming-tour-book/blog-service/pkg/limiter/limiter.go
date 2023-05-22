package limiter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

//限流器的接口，实现三个方法，key用来从context中获取对应的接口名，GetBucket通过接口名获取对应的bucket，AddBuckets添加令牌桶
type LimiterIface interface {
	Key(c *gin.Context) string
	GetBucket(key string) (*ratelimit.Bucket, bool)
	AddBuckets(rules ...LimiterBucketRule) LimiterIface
}

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

type LimiterBucketRule struct {
	Key          string
	FillInterval time.Duration //间隔多久放N个令牌
	Capacity     int64         //令牌桶容量
	Quantum      int64         //每次放令牌的数量
}
