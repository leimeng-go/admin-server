package utils

import (
	"context"
	"encoding/json"
	"errors"
)


func GetUserIdFromCtx(ctx context.Context) (int64, error) {
	value:= ctx.Value("user_id")
	userID, ok := value.(json.Number)
	if !ok {
		return 0, errors.New("user_id is not a uint64")
	}
	userIDInt, err := userID.Int64()
	if err != nil {
		return 0, err
	}
	return userIDInt, nil
}
