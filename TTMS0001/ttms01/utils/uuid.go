package utils

import (
	"crypto/rand"
	"fmt"
)

func CreatUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		fmt.Println(err)
	}
	u[8] = (u[8] | 0x40) & 0x7f
	u[6] = (u[6] & 0xf) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}
