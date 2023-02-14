package redis

import (
	"strconv"
)

func InquireFavorite(key string) (int, error) {
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return -1, err
	}
	n, _ := strconv.Atoi(val)
	return n, nil
}

func UpdateFavorite(key string, n int) error {
	val := strconv.Itoa(n)
	err := client.Set(ctx, key, val, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
