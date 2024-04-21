package eventsclient

import "context"

func GetUserRating(ctx context.Context, userId string) (int, error) {
	return 50, nil
}

func DecreaseUserRating(ctx context.Context, userId string, rating int) error {
	return nil
}
