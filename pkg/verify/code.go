package verify

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	verifyCodeKeyPrefix = "verify:code:"
	verifyCodeExpire    = 5 * time.Minute
)

type Code struct {
	rdb *redis.Client
}

func NewCode(rdb *redis.Client) *Code {
	return &Code{
		rdb: rdb,
	}
}

// Generate 生成6位数字验证码
func (c *Code) Generate() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

// Store 将验证码存储到 Redis
func (c *Code) Store(ctx context.Context, email, code string) error {
	key := verifyCodeKeyPrefix + email
	return c.rdb.Set(ctx, key, code, verifyCodeExpire).Err()
}

// Verify 验证验证码
func (c *Code) Verify(ctx context.Context, email, code string) (bool, error) {
	key := verifyCodeKeyPrefix + email
	storedCode, err := c.rdb.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil
		}
		return false, err
	}

	// 验证成功后删除验证码
	if storedCode == code {
		c.rdb.Del(ctx, key)
		return true, nil
	}

	return false, nil
}

// IsEmailVerified 检查邮箱是否已验证
func (c *Code) IsEmailVerified(ctx context.Context, email string) (bool, error) {
	key := verifyCodeKeyPrefix + "verified:" + email
	exists, err := c.rdb.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}

// MarkEmailAsVerified 标记邮箱为已验证
func (c *Code) MarkEmailAsVerified(ctx context.Context, email string) error {
	key := verifyCodeKeyPrefix + "verified:" + email
	return c.rdb.Set(ctx, key, "1", 0).Err()
}
