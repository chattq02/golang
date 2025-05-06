package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func GetUserKey(key string) string {
	return fmt.Sprintf("u:%s:otp", key)
}

func GenerateCliTokenUUID(userId int) string {
	newUUID := uuid.New()
	// convert UUID to string, remove

	uuidString := strings.ReplaceAll((newUUID).String(), "", "")

    return strconv.Itoa(userId) + "clitoken" + uuidString
}