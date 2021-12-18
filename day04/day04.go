package day04

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

func Part1(secretKey string) (int, error) {
	return findNumber(secretKey, "00000")
}

func Part2(secretKey string) (int, error) {
	return findNumber(secretKey, "000000")
}

func findNumber(secretKey, prefix string) (int, error) {
	h := md5.New()
	for i := 0; ; i++ {
		if _, err := h.Write([]byte(secretKey)); err != nil {
			return 0, err
		}

		if _, err := h.Write([]byte(strconv.FormatInt(int64(i), 10))); err != nil {
			return 0, err
		}

		if strings.HasPrefix(hex.EncodeToString(h.Sum(nil)), prefix) {
			return i, nil
		}

		h.Reset()
	}
}
