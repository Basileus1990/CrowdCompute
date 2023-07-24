package utilitytest

import "math/rand"

// Generates a random string with length between 1 and n
func GenerateRandomString(n int) string {
	if n < 0 {
		panic("n must be not less than than 0")
	}
	length := rand.Intn(n) + 1
	ran_str := make([]byte, length)

	const space int = 32
	const tilde int = 126
	for i := 0; i < length; i++ {
		ran_str[i] = byte(space + rand.Intn(tilde-space))
	}

	return string(ran_str)
}
