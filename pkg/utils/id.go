package utils

import "github.com/lithammer/shortuuid/v4"

const SessionPrefix = "SE_"

func NewGuid(prefix string) string {
	return prefix + shortuuid.New()
}
