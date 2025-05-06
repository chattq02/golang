package context

import (
	"context"
	"errors"

	"Go/internal/utils/cache"
)

type InfoUserUUID struct {
    UserId uint64 
	UserAccount string 
}


func getSubjectUUID(ctx context.Context) (string, error) { 
	sUUID, ok := ctx.Value("subjectUUUID").(string)
    if !ok {
        return "", errors.New("failed to get subject UUID")
    }
    return sUUID, nil
}
func GetUserIdFromUUID(ctx context.Context) (uint64, error) {
	sUUID, err := getSubjectUUID(ctx)

	if err != nil {
        return 0, err
    }

	// get infoUser Redis form uuid

	var infoUser InfoUserUUID

	if err := cache.GetCache(ctx, sUUID, &infoUser); err != nil {
		return 0, err
	}

	return infoUser.UserId, nil

}