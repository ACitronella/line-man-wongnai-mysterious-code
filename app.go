package main

import (
	"encoding/base64"
	"fmt"
)

func reverseBytes(xs []byte) []byte {
	// i love pure function, even its cost some performance, so it be.
	xsCopy := make([]byte, len(xs))
	copy(xsCopy, xs)

	tmpSdLen := len(xsCopy)
	tmpSdLenHalf := len(xsCopy) / 2
	for idx := 0; idx < tmpSdLenHalf; idx++ {
		mirroedIdx := tmpSdLen - idx - 1
		tmpByte := xsCopy[idx]
		xsCopy[idx] = xsCopy[mirroedIdx]
		xsCopy[mirroedIdx] = tmpByte
	}
	return xsCopy
}

func main() {
	var whatIsIt string

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	sd, _ := base64.StdEncoding.DecodeString(secret)

	// tmpStr := string(sd[:])
	// fmt.Println(tmpStr) // got "iangnoW:NAM:ENIL:ta:su:nioJ"
	// so i try to reverse []byte
	reversedSd := reverseBytes(sd)
	whatIsIt = string(reversedSd[:])
	fmt.Println(whatIsIt) // got "Join:us:at:LINE:MAN:Wongnai"
	// I'd say that I'm not familia with golang but if you also looking for AI engineer internship, I'm your a guy.
	// Check out my CV for more infomation.
}
