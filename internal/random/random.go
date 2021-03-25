package random

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var list = []string{
	"(){}[]<>/+=-_,.~:#!@;$%^&*?|",
	"0123456789",
	"abcdefghijklmnopqrstuvwxyz",
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
}

// GeneratePlainPassword generates a random plain password string.
// It consists of four groups which are certain symbols, numbers,
// lower and upper case letters. As per the length and the remainder,
// character selecting is evenly distributed within these groups.
// Although it is unlikely, a final password might contain duplicated
// characters. If a specific character was selected previously, it gets
// dropped in order to increase randomness. Depending on the length
// this is repeated set amount of times to prevent infinite loop and
// applies to all four groups. Greater the length, greater chance in
// duplicated characters.
func GeneratePlainPassword(length int) (string, error) {
	if length < 1 {
		return "", nil
	}

	var (
		retry int
		group int
		chars string
	)

	if length <= 4 {
		retry = 1
	} else if length > 20 && length <= 50 {
		retry = 20
	} else if length > 50 {
		retry = 30
	} else {
		retry = 10
	}

	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= length; i++ {
		switch group {
		case 0:
			pick(0, 31, retry, &chars)
			group++
		case 1:
			pick(1, 10, retry, &chars)
			group++
		case 2:
			pick(2, 26, retry, &chars)
			group++
		case 3:
			pick(3, 26, retry, &chars)
			group = 0
		default:
			return "", fmt.Errorf("out of group range")
		}
	}

	psw := []byte(chars)
	rand.Shuffle(length, func(i, j int) { psw[i], psw[j] = psw[j], psw[i] })

	return string(psw), nil
}

func pick(key, len, retry int, chars *string) {
	var try int

	for {
		try++

		idx := rand.Intn(len)
		char := list[key][idx : idx+1]
		if !strings.Contains(*chars, char) || try == retry {
			*chars += char

			break
		}
	}
}
