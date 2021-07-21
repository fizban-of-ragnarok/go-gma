/*
########################################################################################
#  _______  _______  _______                ___       ______       _____               #
# (  ____ \(       )(  ___  )              /   )     / ___  \     / ___ \              #
# | (    \/| () () || (   ) |             / /) |     \/   \  \   ( (___) )             #
# | |      | || || || (___) |            / (_) (_       ___) /    \     /              #
# | | ____ | |(_)| ||  ___  |           (____   _)     (___ (     / ___ \              #
# | | \_  )| |   | || (   ) | Game           ) (           ) \   ( (   ) )             #
# | (___) || )   ( || )   ( | Master's       | |   _ /\___/  / _ ( (___) )             #
# (_______)|/     \||/     \| Assistant      (_)  (_)\______/ (_) \_____/              #
#                                                                                      #
########################################################################################
*/

//
// Unit tests for the mapper authentication code
//

package auth

import (
	"testing"
)

func TestAuthenticatorRoundLimits(t *testing.T) {
	a := Authenticator{}
	for i := 0; i < 1000; i++ {
		a.Reset()
		a.GenerateChallenge()
		if a.Challenge[0]&0xf0 != 0 {
			t.Errorf("iteration %d, challenge high byte is %x", i, a.Challenge[0])
		}
		if a.Challenge[1]&0x40 != 0x40 {
			t.Errorf("iteration %d, challenge low byte is %x", i, a.Challenge[1])
		}
	}
}

func TestAuthenticator(t *testing.T) {
	a := Authenticator{}

	// Not testing the quality of random numbers we get,
	// just that we get them. We're not testing the library
	// function, just our own code.
	b := a.Challenge
	a.GenerateChallenge()
	if bytesEqual(b, a.Challenge) {
		t.Errorf("Challenge value didn't change")
	}

	// Here are a few challenges and responses we checked with
	// the original Python implementation. We must get the same
	// results with them.
	type testcase struct {
		Nonce     []byte
		Challenge string
		Response  string
		Valid     bool
	}

	testcases := []testcase{
		{Nonce: []byte{0xbc, 0x2f, 0x21, 0x09, 0xcb, 0x4b, 0xd8, 0x38, 0xa4, 0xb4,
			0xad, 0xc8, 0x3f, 0xe0, 0xc5, 0x30, 0x5f, 0x96, 0x3b, 0xb7,
			0xca, 0x62, 0x44, 0xec, 0x73, 0x1a, 0x45, 0x4a, 0x16, 0x8e,
			0x26, 0xa2},
			Challenge: "vC8hCctL2DiktK3IP+DFMF+WO7fKYkTscxpFShaOJqI=",
			Response:  "kGlo/jsAtM8t1ebu6YmHkCCHk9iDvTTssUo+k2hv0xY=",
			Valid:     true},
		{Nonce: []byte{0xbd, 0x27, 0x95, 0x10, 0xf6, 0x74, 0x77, 0x1a, 0x9b, 0x7f,
			0xed, 0xfe, 0xc0, 0x5e, 0xad, 0xe8, 0xb2, 0x03, 0x95, 0x2f,
			0xb1, 0x48, 0xee, 0x33, 0x26, 0xad, 0xfe, 0x6d, 0xa8, 0x32,
			0x74, 0x15},
			Challenge: "vSeVEPZ0dxqbf+3+wF6t6LIDlS+xSO4zJq3+bagydBU=",
			Response:  "kLQnscqqYKXe1SV1SXmmzqXx4bEo+eV1eR5jS++rY4c=",
			Valid:     true},
		{Nonce: []byte{0x0c, 0x72, 0xba, 0x37, 0xc2, 0x67, 0x15, 0x07, 0x22, 0x07,
			0x99, 0xa4, 0x40, 0xd9, 0x9c, 0xab, 0xa1, 0xfe, 0xa7, 0x15,
			0x5e, 0x8d, 0x10, 0x35, 0x50, 0xa3, 0x01, 0x4a, 0x07, 0x4f,
			0x60, 0xa2},
			Challenge: "DHK6N8JnFQciB5mkQNmcq6H+pxVejRA1UKMBSgdPYKI=",
			Response:  "q8nncBMkpv+O3zVVx2hs7NKqugvW43Hpvj5k0XCsua8=",
			Valid:     true},
		{Nonce: []byte{0xe6, 0x33, 0x8f, 0x48, 0x67, 0x62, 0x9d, 0x5d, 0x52, 0x4a,
			0xe9, 0xd8, 0xbd, 0x09, 0xd0, 0xdc, 0x54, 0x6e, 0xb2, 0x9b,
			0x6f, 0xb5, 0xa6, 0x00, 0x92, 0x47, 0x3c, 0x09, 0x18, 0xeb,
			0xe0, 0x4d},
			Challenge: "5jOPSGdinV1SSunYvQnQ3FRusptvtaYAkkc8CRjr4E0=",
			Response:  "aA7Pj/w/ONgHMH5JoRkGQ1tks1vUei37PwML9Bgi3Ak=",
			Valid:     true},
		{Nonce: []byte{0xb7, 0xc9, 0x6e, 0xb7, 0x82, 0x09, 0xcc, 0x27, 0x50, 0x8e,
			0x8b, 0xb9, 0x7c, 0xb1, 0x09, 0x2f, 0xc5, 0xf3, 0xe0, 0xb8,
			0x67, 0x02, 0xd2, 0x0a, 0x92, 0xcc, 0x0b, 0x70, 0x38, 0x2d,
			0x85, 0x5e},
			Challenge: "t8lut4IJzCdQjou5fLEJL8Xz4LhnAtIKkswLcDgthV4=",
			Response:  "IcFg2BuQCdUsUHeVzfU6mylVVS/K7jalFX0DPFH8U7c=",
			Valid:     true},
		{Nonce: []byte{0x38, 0xfe, 0xde, 0x2b, 0x16, 0xc5, 0xe8, 0x72, 0xb9, 0x06,
			0xca, 0x51, 0x05, 0xe5, 0xbe, 0x1a, 0x74, 0x25, 0xcb, 0x8e,
			0x6f, 0xe2, 0xd6, 0xc1, 0xf3, 0x88, 0xa0, 0xe9, 0x30, 0x1c,
			0xa4, 0xbe},
			Challenge: "OP7eKxbF6HK5BspRBeW+GnQly45v4tbB84ig6TAcpL4=",
			Response:  "90MpQRjZfPgZZNqz6emBMmYIPEDjqijAjmcXg6W0GwI=",
			Valid:     true},
		{Nonce: []byte{0xd9, 0x82, 0xc0, 0x19, 0x18, 0x8d, 0x57, 0xff, 0xaf, 0xa7,
			0xb3, 0xe1, 0xde, 0xc4, 0xd7, 0x6b, 0x0c, 0xf2, 0xcc, 0x71,
			0x48, 0xb3, 0x34, 0xce, 0x11, 0x13, 0x09, 0xab, 0x73, 0x02,
			0x45, 0xaa},
			Challenge: "2YLAGRiNV/+vp7Ph3sTXawzyzHFIszTOERMJq3MCRao=",
			Response:  "K9cEUf9TaAMvcHM3hxzFASf9JYo4piTUHnQovvkbtZo=",
			Valid:     true},
		{Nonce: []byte{0xa0, 0x24, 0x8f, 0x85, 0x88, 0x2f, 0x58, 0xe2, 0x22, 0x6a,
			0x1c, 0x54, 0x4c, 0x65, 0x5e, 0xb5, 0x9b, 0xea, 0x39, 0xe8,
			0xd4, 0x13, 0x4a, 0xee, 0xae, 0xfc, 0x15, 0xde, 0xa8, 0x00,
			0xb1, 0xbf},
			Challenge: "oCSPhYgvWOIiahxUTGVetZvqOejUE0rurvwV3qgAsb8=",
			Response:  "FAaUh3YI6I3ovNP5L8zlX6u1QbkkGkwBKkGP90Nsb5Q=",
			Valid:     true},
		{Nonce: []byte{0xdb, 0x34, 0x8f, 0xfc, 0x46, 0x7b, 0xd2, 0x57, 0x83, 0x40,
			0x79, 0x50, 0x7a, 0x41, 0xae, 0xb9, 0x30, 0xb4, 0x01, 0x92,
			0xcf, 0x7f, 0x04, 0xc7, 0xa2, 0x30, 0x03, 0xc9, 0x24, 0x36,
			0x30, 0xdf},
			Challenge: "2zSP/EZ70leDQHlQekGuuTC0AZLPfwTHojADySQ2MN8=",
			Response:  "biRPri/TOSfT4nJ/IR78synR82NL1pihYaWsMRSfioo=",
			Valid:     true},
		{Nonce: []byte{0xe9, 0xb9, 0x61, 0x81, 0xf8, 0xc9, 0xcb, 0x7e, 0xb8, 0x9c,
			0xdb, 0xe6, 0xb5, 0x49, 0x48, 0x24, 0x74, 0xfb, 0x34, 0x05,
			0x85, 0x65, 0x88, 0xa3, 0x43, 0x6c, 0x34, 0x1f, 0xfd, 0x01,
			0xe6, 0x88},
			Challenge: "6blhgfjJy364nNvmtUlIJHT7NAWFZYijQ2w0H/0B5og=",
			Response:  "lNvqSPX7fLh6rsSLAfP1aZ5UZjxR5U9lLztyvrlUK9g=",
			Valid:     true},
		{Nonce: []byte{0xb5, 0x37, 0x54, 0x4d, 0xdc, 0xaf, 0x4b, 0xc5, 0xab, 0x66,
			0x0d, 0xde, 0xad, 0xae, 0x09, 0x83, 0x11, 0x35, 0x1f, 0x75,
			0x34, 0x6f, 0x97, 0x8b, 0x17, 0xa1, 0x40, 0xb0, 0x40, 0x06,
			0x74, 0x41},
			Challenge: "tTdUTdyvS8WrZg3era4JgxE1H3U0b5eLF6FAsEAGdEE=",
			Response:  "FBdzziLBdqDxgs0jBaTuEbckrCe1UyQucoA+PCV2Pbs=",
			Valid:     true},
		{Nonce: []byte{0x46, 0x02, 0x65, 0x0f, 0xaf, 0x6a, 0x45, 0x67, 0x7b, 0x1e,
			0x1b, 0x3e, 0xc1, 0xad, 0x3a, 0x66, 0x7c, 0x3c, 0x20, 0x3f,
			0x13, 0xaf, 0xb9, 0x93, 0xea, 0x17, 0x8d, 0x1c, 0xc1, 0x91,
			0x28, 0x20},
			Challenge: "RgJlD69qRWd7Hhs+wa06Znw8ID8Tr7mT6heNHMGRKCA=",
			Response:  "IDs7YP2vbhYUjuokMneDFumcd3GhOXikSQyjrMxTIcg=",
			Valid:     true},
		{Nonce: []byte{0xe4, 0xcf, 0x3f, 0x6a, 0xe8, 0xd1, 0x1e, 0x75, 0xf6, 0x77,
			0x71, 0xca, 0x49, 0xef, 0x5d, 0x8f, 0xa1, 0x85, 0x70, 0x58,
			0x30, 0xa7, 0xed, 0xc6, 0xd7, 0x32, 0xd1, 0xf3, 0xa6, 0xea,
			0x3b, 0xf8},
			Challenge: "5M8/aujRHnX2d3HKSe9dj6GFcFgwp+3G1zLR86bqO/g=",
			Response:  "nAwZNHISt+pV2zDa/BQPTxgyIkyl06BLNrjPL8EVqSM=",
			Valid:     true},
		{Nonce: []byte{0x01, 0x0e, 0x95, 0x6b, 0xc1, 0xe6, 0x23, 0x23, 0x52, 0x3e,
			0x62, 0xd5, 0x76, 0x7c, 0xeb, 0x18, 0xa7, 0x2d, 0xdb, 0x7f,
			0x6c, 0xab, 0xfa, 0x72, 0x71, 0xd1, 0x31, 0xa6, 0x38, 0xc9,
			0x06, 0x5c},
			Challenge: "AQ6Va8HmIyNSPmLVdnzrGKct239sq/pycdExpjjJBlw=",
			Response:  "vjjJTHzBVSva0Zd2XzNFOpDCTTv/LSQWc3rVnaejdvo=",
			Valid:     true},
		{Nonce: []byte{0xb8, 0xc0, 0x0a, 0x6a, 0x55, 0xd3, 0xcf, 0x47, 0xa8, 0xa8,
			0x55, 0xaf, 0x5c, 0x72, 0x4b, 0xad, 0x55, 0xf6, 0x75, 0xb1,
			0xfd, 0x32, 0xf8, 0x6c, 0x4a, 0x23, 0x6c, 0xc2, 0x6e, 0x06,
			0x13, 0x78},
			Challenge: "uMAKalXTz0eoqFWvXHJLrVX2dbH9MvhsSiNswm4GE3g=",
			Response:  "YlMRscFkHyPiSaC8RbtS9kepY+ituzGZIAwnmayR0oo=",
			Valid:     true},
		{Nonce: []byte{0x6a, 0xbe, 0xe8, 0xb3, 0x1b, 0x93, 0x5d, 0x86, 0xfb, 0x88,
			0xf9, 0x32, 0xdf, 0xf8, 0xe6, 0xc1, 0xe1, 0x75, 0x8a, 0x05,
			0x6b, 0xcd, 0xa8, 0xc8, 0xcb, 0xf3, 0x77, 0x93, 0xee, 0x4e,
			0x57, 0x9d},
			Challenge: "ar7osxuTXYb7iPky3/jmweF1igVrzajIy/N3k+5OV50=",
			Response:  "t5QHcwGQc10/atdWVxoYBHfNg94dCpflDuo4VgSSuMw=",
			Valid:     true},
		{Nonce: []byte{0x68, 0x91, 0x75, 0xbd, 0xda, 0x87, 0x6b, 0x4d, 0x81, 0xbd,
			0x2d, 0xd0, 0x39, 0x8c, 0xe3, 0x25, 0x49, 0x32, 0x2c, 0xea,
			0x37, 0xde, 0x15, 0x48, 0x7e, 0xc6, 0x9a, 0xec, 0x65, 0x5e,
			0x40, 0xa6},
			Challenge: "aJF1vdqHa02BvS3QOYzjJUkyLOo33hVIfsaa7GVeQKY=",
			Response:  "NZBLdygw/v6e9aj5pclbb6gZKIsvaMQ8y0dLMcnZX18=",
			Valid:     true},
		{Nonce: []byte{0x1d, 0xbf, 0x14, 0x55, 0xc4, 0x0b, 0xff, 0x03, 0x49, 0xb6,
			0xf2, 0x49, 0xdf, 0x86, 0xfe, 0x58, 0xf0, 0x9c, 0x73, 0xfd,
			0xfc, 0xa0, 0x33, 0xd8, 0xfd, 0x59, 0xb2, 0x7e, 0xad, 0x12,
			0x1a, 0xd3},
			Challenge: "Hb8UVcQL/wNJtvJJ34b+WPCcc/38oDPY/Vmyfq0SGtM=",
			Response:  "wICpk7dVj09ve+GMesM+XsI78LNoGJ0LdcdN6jE8NJ0=",
			Valid:     true},
		{Nonce: []byte{0x22, 0xdf, 0x2c, 0xde, 0xce, 0xcd, 0xee, 0x56, 0x6a, 0xdc,
			0xdf, 0xe6, 0x3f, 0x0c, 0x60, 0xf4, 0x20, 0x1b, 0x20, 0xc1,
			0x48, 0x52, 0x6e, 0xb5, 0x0b, 0xfa, 0x56, 0x2b, 0x50, 0xed,
			0x83, 0x6f},
			Challenge: "It8s3s7N7lZq3N/mPwxg9CAbIMFIUm61C/pWK1Dtg28=",
			Response:  "D6M4hQxL7jCiopms6tePDC5PFpwPxw7xyh9ZgFiWOjU=",
			Valid:     true},
		{Nonce: []byte{0x5e, 0x3f, 0xcb, 0xad, 0xe0, 0x69, 0xd1, 0x24, 0x20, 0x81,
			0x35, 0x94, 0x00, 0xeb, 0xb7, 0xc6, 0xb3, 0x79, 0xe2, 0xd4,
			0x05, 0x1e, 0x41, 0x04, 0x6f, 0xc7, 0x66, 0xb7, 0xba, 0xf3,
			0x12, 0x25},
			Challenge: "Xj/LreBp0SQggTWUAOu3xrN54tQFHkEEb8dmt7rzEiU=",
			Response:  "H5UP7uK+GbEbDHf7rxdUTKg86hPq7zWj/Ymj6iODc6I=",
			Valid:     true},
		{Nonce: []byte{0xea, 0xc2, 0x79, 0xa5, 0x0c, 0x6c, 0x40, 0x0a, 0xd5, 0x28,
			0x67, 0x98, 0x38, 0xae, 0xcc, 0xf0, 0x2e, 0xf8, 0xa8, 0x36,
			0xb2, 0x91, 0x51, 0x9a, 0xf8, 0xa5, 0xdb, 0x48, 0xfb, 0x63,
			0x4d, 0x36},
			Challenge: "6sJ5pQxsQArVKGeYOK7M8C74qDaykVGa+KXbSPtjTTY=",
			Response:  "Xl6P/M1EcEjtAnbYEww/RasdeDqBJ5Qt4JtH2WM5glw=",
			Valid:     true},
		{Nonce: []byte{0x83, 0x2f, 0xc0, 0x18, 0xcb, 0xde, 0xc4, 0xbb, 0xcf, 0x2f,
			0x14, 0xfa, 0x85, 0xd1, 0xcb, 0xed, 0x30, 0x8e, 0x69, 0xf7,
			0x64, 0xba, 0xad, 0x11, 0xbb, 0xd1, 0x5a, 0x98, 0xad, 0x85,
			0xb3, 0x44},
			Challenge: "gy/AGMvexLvPLxT6hdHL7TCOafdkuq0Ru9FamK2Fs0Q=",
			Response:  "GnzzJSuSG66apoQD1nkKI8nYQ0i5Ip1rkGgQTojzM98=",
			Valid:     true},
		{Nonce: []byte{0x33, 0x87, 0xcf, 0x20, 0x13, 0xee, 0x29, 0x4b, 0x83, 0xb1,
			0x21, 0x77, 0xd9, 0x5e, 0x64, 0xe8, 0x8c, 0x8d, 0xf8, 0x11,
			0xe1, 0x40, 0x62, 0x56, 0x7e, 0xe8, 0xc0, 0x6a, 0xec, 0xfb,
			0xd5, 0x64},
			Challenge: "M4fPIBPuKUuDsSF32V5k6IyN+BHhQGJWfujAauz71WQ=",
			Response:  "lv8L1XFm+Z8sNEbcg1hyO9CxXobAjvI+Rs0lc3JhnpQ=",
			Valid:     true},
		{Nonce: []byte{0x82, 0x10, 0x22, 0xdb, 0x56, 0xf2, 0xb9, 0x4c, 0x2c, 0x5b,
			0x34, 0x7f, 0xda, 0x6f, 0xb2, 0xf9, 0x68, 0x39, 0x55, 0xb1,
			0x57, 0xdb, 0xa1, 0x75, 0x93, 0x51, 0x7a, 0x09, 0x05, 0xb6,
			0x42, 0xf7},
			Challenge: "ghAi21byuUwsWzR/2m+y+Wg5VbFX26F1k1F6CQW2Qvc=",
			Response:  "ii2qgX820Muret7RinJNf0hrer//XGlP0hZ3dETILw0=",
			Valid:     true},
		{Nonce: []byte{0xa7, 0x1d, 0x0a, 0xfe, 0x7e, 0x90, 0x9a, 0x16, 0x80, 0x79,
			0x52, 0xc9, 0x2c, 0xb1, 0x57, 0x45, 0xc4, 0x14, 0x60, 0x13,
			0x0f, 0xba, 0x15, 0x97, 0x2b, 0x68, 0x20, 0xd5, 0x3b, 0x78,
			0xd2, 0xb0},
			Challenge: "px0K/n6QmhaAeVLJLLFXRcQUYBMPuhWXK2gg1Tt40rA=",
			Response:  "1KMl3BDV/+opDA/+J5mg9AjDjCZILRc+yRQl+taxiIE=",
			Valid:     true},
		{Nonce: []byte{0xc1, 0x92, 0xab, 0xb4, 0xd6, 0xb4, 0x7d, 0x12, 0x7b, 0xc4,
			0xf2, 0x13, 0x67, 0xdd, 0xf8, 0xef, 0xe0, 0x38, 0xe1, 0x03,
			0xba, 0x01, 0x41, 0xa5, 0x73, 0x1b, 0xea, 0x4e, 0xb8, 0xc8,
			0x7b, 0x14},
			Challenge: "wZKrtNa0fRJ7xPITZ9347+A44QO6AUGlcxvqTrjIexQ=",
			Response:  "UF9XDxqK8/Fr+CF3gFHuyDRFUy7skuRceR5gRiHbdkU=",
			Valid:     true},
		{Nonce: []byte{0xa2, 0xa7, 0x72, 0xb1, 0x49, 0x36, 0x96, 0x44, 0x44, 0x20,
			0x8d, 0x10, 0xba, 0xbe, 0x80, 0x0d, 0xa1, 0xd0, 0xd9, 0xb1,
			0xbc, 0xfe, 0x2e, 0x37, 0xd5, 0x7f, 0xd6, 0xb8, 0x89, 0xd9,
			0xd3, 0x58},
			Challenge: "oqdysUk2lkREII0Qur6ADaHQ2bG8/i431X/WuInZ01g=",
			Response:  "WtkmdJsXC5AnJq0BtlsndqE3H+7NNfYiZIGDvnar79I=",
			Valid:     true},
		{Nonce: []byte{0x3b, 0x59, 0x43, 0x7d, 0x94, 0x58, 0x75, 0x0c, 0xfa, 0x63,
			0x7a, 0x34, 0x45, 0x73, 0xe3, 0x22, 0x87, 0x82, 0xff, 0xbc,
			0x4d, 0x88, 0x9d, 0xb2, 0x79, 0x21, 0x83, 0xb0, 0xb3, 0x54,
			0xde, 0xbb},
			Challenge: "O1lDfZRYdQz6Y3o0RXPjIoeC/7xNiJ2yeSGDsLNU3rs=",
			Response:  "UqcSDjtNmgG4kVViGp0/qWCKNLTvVlQrpRN4WbQ9R8k=",
			Valid:     true},
		{Nonce: []byte{0xd6, 0x15, 0x8a, 0xa0, 0x29, 0x09, 0x3c, 0xbd, 0x63, 0xf3,
			0x92, 0xe2, 0xac, 0xf6, 0x3e, 0x62, 0xa2, 0x27, 0x5d, 0x12,
			0x7f, 0xdf, 0xb1, 0x4e, 0x21, 0x30, 0xb7, 0xd7, 0x97, 0xc4,
			0x89, 0xb0},
			Challenge: "1hWKoCkJPL1j85LirPY+YqInXRJ/37FOITC315fEibA=",
			Response:  "Lv8xlK30r9tt9eebsyvyGD2aL4l/SuD82gMNRm34QnM=",
			Valid:     true},
		{Nonce: []byte{0x2e, 0x17, 0xe9, 0x59, 0xfc, 0xd3, 0x93, 0x00, 0x2a, 0xe1,
			0xd6, 0xd8, 0x3a, 0xc4, 0x03, 0x28, 0x75, 0x70, 0xa6, 0x73,
			0xc5, 0xc7, 0xd9, 0xb2, 0x7e, 0x46, 0xea, 0x8d, 0xe9, 0xde,
			0x6a, 0x8b},
			Challenge: "LhfpWfzTkwAq4dbYOsQDKHVwpnPFx9myfkbqjeneaos=",
			Response:  "K/j0yJqX0U+VUiurE6lrrnrc67i4FVOpAfyfFtp/q4I=",
			Valid:     true},
		{Nonce: []byte{0x76, 0x03, 0xfb, 0x25, 0x02, 0xfc, 0x1f, 0xd3, 0xdc, 0x97,
			0xe1, 0xf1, 0x51, 0x04, 0x57, 0xe7, 0x5e, 0xd0, 0xc4, 0x2f,
			0x58, 0xeb, 0x50, 0xaf, 0x35, 0x94, 0xa5, 0x59, 0x42, 0xf8,
			0x6e, 0x36},
			Challenge: "dgP7JQL8H9Pcl+HxUQRX517QxC9Y61CvNZSlWUL4bjY=",
			Response:  "lOxFWFcyFIgyryopBEnDHHhAI0ZlYUIW0V0Mo5bZRW4=",
			Valid:     true},
		{Nonce: []byte{0xae, 0x7d, 0x20, 0x5e, 0x25, 0x7b, 0x18, 0xf0, 0x77, 0x29,
			0x8b, 0xb6, 0x21, 0xbc, 0x7f, 0xf9, 0x21, 0x4b, 0x22, 0x0d,
			0xe4, 0x17, 0xf7, 0x81, 0xb9, 0x5f, 0x5d, 0xc9, 0x5f, 0xc6,
			0x2e, 0xa9},
			Challenge: "rn0gXiV7GPB3KYu2Ibx/+SFLIg3kF/eBuV9dyV/GLqk=",
			Response:  "B+I5rUig6jSb3i+kBUNWkz2qHm28UMtOG9pWT12+zVg=",
			Valid:     true},
		{Nonce: []byte{0xb9, 0xa4, 0x13, 0x04, 0xcf, 0x54, 0xe3, 0xec, 0xab, 0x51,
			0xcc, 0x45, 0xfb, 0x81, 0x8d, 0xe6, 0x56, 0x60, 0xb6, 0xd7,
			0xda, 0x7b, 0xdd, 0xca, 0x48, 0x54, 0x23, 0x24, 0x36, 0xea,
			0x7a, 0x0a},
			Challenge: "uaQTBM9U4+yrUcxF+4GN5lZgttfae93KSFQjJDbqego=",
			Response:  "eNmLUgWYdqohlqFuCT9epMpcERYqfLBw/Az4SDRidHs=",
			Valid:     true},
		{Nonce: []byte{0xbf, 0xd5, 0x53, 0x74, 0xc8, 0xf0, 0xf5, 0xb2, 0x5c, 0x30,
			0x78, 0x27, 0x57, 0xf1, 0x8b, 0xef, 0x1c, 0x07, 0x04, 0x78,
			0xb5, 0x41, 0x3d, 0x43, 0x16, 0x69, 0xfc, 0xf6, 0x77, 0xed,
			0xc1, 0xcd},
			Challenge: "v9VTdMjw9bJcMHgnV/GL7xwHBHi1QT1DFmn89nftwc0=",
			Response:  "PabyQnr3KtQIgGVAokz0SyIDublQ0rho8lWcBUsfQ9s=",
			Valid:     true},
		{Nonce: []byte{0x5f, 0xd2, 0x5e, 0x5b, 0x32, 0xcf, 0x9d, 0xf1, 0x30, 0xed,
			0x9c, 0xfc, 0x37, 0x66, 0xdd, 0x79, 0x73, 0x10, 0x35, 0x91,
			0x8f, 0xb4, 0x3a, 0x36, 0x33, 0xa2, 0x93, 0x0e, 0x97, 0x3c,
			0xcb, 0x05},
			Challenge: "X9JeWzLPnfEw7Zz8N2bdeXMQNZGPtDo2M6KTDpc8ywU=",
			Response:  "c9d0pS1yROJb4wevLycIXCXW1BCb/TlmAc7KFiq9v6o=",
			Valid:     true},
		{Nonce: []byte{0x75, 0xbe, 0xab, 0xd2, 0x26, 0xa4, 0xac, 0x64, 0xb8, 0x7d,
			0x60, 0x8a, 0x2a, 0x72, 0x61, 0x84, 0x06, 0xd0, 0xb6, 0x80,
			0xf1, 0x58, 0x7e, 0xb4, 0xa8, 0x92, 0x3d, 0x38, 0xda, 0x90,
			0xf1, 0xd1},
			Challenge: "db6r0iakrGS4fWCKKnJhhAbQtoDxWH60qJI9ONqQ8dE=",
			Response:  "4rgUB+OyNjwKmTgK2LeIs4VUAJlbR1ACCsHc5lIP9uU=",
			Valid:     true},
		{Nonce: []byte{0xad, 0x71, 0xc8, 0x5f, 0xed, 0x4e, 0x02, 0x32, 0xaa, 0xdd,
			0xc6, 0xea, 0xea, 0xa5, 0xae, 0xf2, 0xa6, 0x1d, 0x22, 0xf7,
			0x17, 0xa9, 0x76, 0xbe, 0x09, 0x36, 0x1b, 0x57, 0x14, 0x89,
			0x3a, 0x01},
			Challenge: "rXHIX+1OAjKq3cbq6qWu8qYdIvcXqXa+CTYbVxSJOgE=",
			Response:  "rJqypsyUpjqMGQcq11+EVa9Z8ncmiKizr/Jx3tLLBJE=",
			Valid:     true},
		{Nonce: []byte{0x2e, 0x5a, 0xd7, 0x69, 0x20, 0xd9, 0x7b, 0xd3, 0x3c, 0x52,
			0xdc, 0x67, 0x08, 0xcc, 0xcb, 0xcf, 0x95, 0x4a, 0x04, 0x20,
			0xb7, 0x3f, 0x63, 0xb8, 0x47, 0x48, 0xaa, 0x6c, 0x41, 0xdd,
			0x2e, 0xb3},
			Challenge: "LlrXaSDZe9M8UtxnCMzLz5VKBCC3P2O4R0iqbEHdLrM=",
			Response:  "p33GLC6m/XQy/fytOr+9GWp043m7U1nRm/wmXPw0V98=",
			Valid:     true},
		{Nonce: []byte{0xb5, 0xb9, 0x4a, 0x2a, 0xba, 0xaa, 0x34, 0xf5, 0xd4, 0xa8,
			0xca, 0xec, 0xb0, 0x5a, 0x62, 0x8c, 0x59, 0x13, 0xbd, 0xcd,
			0xd2, 0xfd, 0x9c, 0xed, 0xc8, 0x66, 0xa9, 0x3e, 0x7b, 0x9b,
			0x45, 0x18},
			Challenge: "tblKKrqqNPXUqMrssFpijFkTvc3S/ZztyGapPnubRRg=",
			Response:  "GDxfHqFYKb/hloffJtUQZmfjpxQvJLJlkAMHgG0+rok=",
			Valid:     true},
		{Nonce: []byte{0x38, 0x18, 0x08, 0x19, 0x30, 0x90, 0xe0, 0x9a, 0x8d, 0xdb,
			0x04, 0x43, 0x01, 0x34, 0x55, 0x0c, 0x14, 0x0c, 0xe2, 0x88,
			0x99, 0xda, 0x68, 0x92, 0xba, 0x02, 0x53, 0x91, 0xbd, 0xbd,
			0xb7, 0xba},
			Challenge: "OBgIGTCQ4JqN2wRDATRVDBQM4oiZ2miSugJTkb29t7o=",
			Response:  "4jqlQysIG4MV77I9DnAvr2nfQ6RS5eJYRH4G+dIzLlQ=",
			Valid:     true},
		{Nonce: []byte{0x7b, 0xa5, 0xc6, 0x19, 0x0c, 0x9b, 0x73, 0x81, 0x94, 0xc7,
			0x3e, 0xc6, 0x12, 0x66, 0x7d, 0xcc, 0x66, 0x04, 0x63, 0xbb,
			0x77, 0x5e, 0x70, 0x9f, 0xac, 0xae, 0x3f, 0x6b, 0xb3, 0xa3,
			0x30, 0x55},
			Challenge: "e6XGGQybc4GUxz7GEmZ9zGYEY7t3XnCfrK4/a7OjMFU=",
			Response:  "xNqK9/ItpZNLVXtw0Fc/47q21oko7LXF7mZcoWhFDAM=",
			Valid:     true},
		{Nonce: []byte{0xba, 0x51, 0xd1, 0x51, 0xfa, 0xca, 0x47, 0x48, 0x2c, 0xfa,
			0x02, 0x52, 0xf0, 0xca, 0xcf, 0x1f, 0x8f, 0xe0, 0x06, 0x3c,
			0x7e, 0x96, 0xc3, 0x60, 0x97, 0x7d, 0xb9, 0x58, 0x3b, 0x7a,
			0xa6, 0x95},
			Challenge: "ulHRUfrKR0gs+gJS8MrPH4/gBjx+lsNgl325WDt6ppU=",
			Response:  "ejWeH/uuCgQnFyd5hb95+SVq5qoXVJl67D+xns0YJcE=",
			Valid:     true},
		{Nonce: []byte{0xd2, 0xb0, 0x00, 0x6f, 0x2d, 0xa1, 0x96, 0x11, 0x06, 0x3e,
			0x89, 0x44, 0x97, 0xfd, 0x0c, 0x8d, 0x5a, 0xd2, 0x32, 0x05,
			0xc0, 0x4c, 0x75, 0xf2, 0x49, 0xf9, 0x1f, 0x67, 0x8c, 0x4d,
			0x58, 0xfd},
			Challenge: "0rAAby2hlhEGPolEl/0MjVrSMgXATHXySfkfZ4xNWP0=",
			Response:  "MvDMFmFH3rq2jRxWkdeWHYfOQ92/xKfgOq37/8THhhU=",
			Valid:     true},
		{Nonce: []byte{0x8e, 0xd2, 0xcd, 0xf9, 0x5c, 0x99, 0xc7, 0x1e, 0x75, 0x4f,
			0x3d, 0x55, 0xb8, 0x10, 0xf5, 0x28, 0x7f, 0xcb, 0x2f, 0x07,
			0x87, 0x0b, 0x53, 0x37, 0xb8, 0x2c, 0x69, 0xca, 0xbe, 0x60,
			0xcc, 0xd8},
			Challenge: "jtLN+VyZxx51Tz1VuBD1KH/LLweHC1M3uCxpyr5gzNg=",
			Response:  "RQ/g9I+ZDVcdeY/qhwEFSDwfx0QYI8h7TsgCqg4hQqw=",
			Valid:     true},
		{Nonce: []byte{0x4c, 0xe0, 0xe3, 0xdb, 0x6b, 0x5f, 0x0e, 0xfe, 0xd0, 0x7b,
			0x97, 0x01, 0x7c, 0xd2, 0xe0, 0x22, 0x05, 0x40, 0x98, 0x01,
			0x48, 0xb0, 0x84, 0x0a, 0x9a, 0x17, 0x84, 0x2e, 0x03, 0x97,
			0x2a, 0xff},
			Challenge: "TODj22tfDv7Qe5cBfNLgIgVAmAFIsIQKmheELgOXKv8=",
			Response:  "Lh/AP5pB1dataeIfxJqg1pV+pDKFTh0wLsi1Gd/9TRU=",
			Valid:     true},
		{Nonce: []byte{0x09, 0x19, 0xe2, 0x02, 0x8d, 0x04, 0x3a, 0x60, 0x50, 0x02,
			0xb9, 0x29, 0x6e, 0x4d, 0xdb, 0x34, 0xb2, 0x9d, 0x38, 0xe2,
			0x21, 0xf0, 0xb5, 0x3a, 0x2f, 0xfe, 0x4e, 0x5b, 0x93, 0x83,
			0x52, 0x72},
			Challenge: "CRniAo0EOmBQArkpbk3bNLKdOOIh8LU6L/5OW5ODUnI=",
			Response:  "7NSwdNpZjUMLn5KEFyF+p6O1f0Z8h6I26+Lmh/TK7Ew=",
			Valid:     true},
		{Nonce: []byte{0x86, 0x82, 0x73, 0xfa, 0x17, 0xfd, 0xba, 0x1d, 0xe8, 0x54,
			0x7c, 0x21, 0x51, 0x4e, 0xa2, 0xe5, 0x84, 0xd6, 0x86, 0x64,
			0xec, 0x77, 0x67, 0x04, 0x75, 0xbe, 0x58, 0x79, 0x5e, 0x2e,
			0xad, 0xe7},
			Challenge: "hoJz+hf9uh3oVHwhUU6i5YTWhmTsd2cEdb5YeV4urec=",
			Response:  "6xyuzY17NeQVz75/7+xe/GczckJ75o0LuYsUCtUq0O0=",
			Valid:     true},
		{Nonce: []byte{0x87, 0x05, 0xf9, 0xe0, 0x9a, 0x29, 0x6d, 0x3f, 0xd2, 0xd7,
			0xb1, 0xd6, 0x1e, 0x56, 0x83, 0x44, 0xfb, 0x5c, 0xfb, 0xcb,
			0xa1, 0xe5, 0xce, 0x62, 0x2f, 0x74, 0xbe, 0x97, 0xad, 0x44,
			0x1f, 0x06},
			Challenge: "hwX54JopbT/S17HWHlaDRPtc+8uh5c5iL3S+l61EHwY=",
			Response:  "S41H0GR7FRmTo1FI/HoxA6zqpeh24IcHVrcnM9kQHx0=",
			Valid:     true},
		{Nonce: []byte{0xa0, 0x64, 0x7e, 0x2a, 0x34, 0xca, 0xc5, 0xf8, 0xc8, 0x79,
			0xde, 0x40, 0x96, 0x35, 0xaf, 0x8b, 0x5e, 0x77, 0xd6, 0x39,
			0xff, 0xba, 0xf0, 0x7b, 0x9e, 0xa6, 0x78, 0x86, 0x7f, 0xb6,
			0x30, 0x5b},
			Challenge: "oGR+KjTKxfjIed5AljWvi1531jn/uvB7nqZ4hn+2MFs=",
			Response:  "sTRxAtf7/YtikA7VVK2NjqaAbRke8LrYPz++14sUIks=",
			Valid:     true},
		{Nonce: []byte{0x1d, 0xee, 0xbf, 0x74, 0xf2, 0x0b, 0x46, 0x68, 0xda, 0xad,
			0x6d, 0x28, 0x3a, 0xd6, 0x37, 0xe4, 0x5f, 0x3d, 0x53, 0x56,
			0xe8, 0x00, 0xf3, 0x6b, 0x1b, 0xb6, 0x23, 0x8e, 0xec, 0xdd,
			0x53, 0xd8},
			Challenge: "He6/dPILRmjarW0oOtY35F89U1boAPNrG7YjjuzdU9g=",
			Response:  "Pdc8yR7DVOjqQlPZtgH21Ein6OJIW6job6me6uVfZzY=",
			Valid:     true},
		{Nonce: []byte{0x1d, 0xee, 0xbf, 0x74, 0xf2, 0x0b, 0x46, 0x68, 0xda, 0xad,
			0x6d, 0x28, 0x3a, 0xd6, 0x37, 0xe4, 0x5f, 0x3d, 0x53, 0x56,
			0xe8, 0x00, 0xf3, 0x6b, 0x1b, 0xb6, 0x23, 0x8e, 0xec, 0xdd,
			0x53, 0xd8},
			Challenge: "He6/dPILRmjarW0oOtY35F89U1boAPNrG7YjjuzdU9g=",
			Response:  "Pdc8yR7DVOjqQlPZtgH21Ein6OJIW6job6me6uVfZzY=",
			Valid:     true},
		{Nonce: []byte{0x1d, 0xee, 0xbf, 0x74, 0xf2, 0x0b, 0x46, 0x68, 0xda, 0xad,
			0x6d, 0x28, 0x3a, 0xd6, 0x37, 0xe4, 0x5f, 0x3d, 0x53, 0x56,
			0xe8, 0x00, 0xf3, 0x6b, 0x1b, 0xb6, 0x23, 0x8e, 0xec, 0xdd,
			0x53, 0xd8},
			Challenge: "He6/dPILRmjarW0oOtY35F89U1boAPNrG7YjjuzdU9g=",
			Response:  "Pdc8yR7DVOjqQlPZtgH2aEin6OJIW6job6me6uVfZzY=",
			Valid:     false},
	}
	for i, test := range testcases {
		a := Authenticator{Secret: []byte("abc123**sekret**XXXyyyZZZ")}
		a.Challenge = test.Nonce
		if a.CurrentChallenge() != test.Challenge {
			t.Errorf("Challenge iteration %d was %s, expected %s", i, a.CurrentChallenge(), test.Challenge)
		}
		r, err := a.ValidateResponse(test.Response)
		if err != nil {
			t.Errorf("iteration %d, error %v in validation", i, err)
		}
		if r != test.Valid {
			t.Errorf("iteration %d, got %v result, expected %v", i, r, test.Valid)
		}
	}
}

// @[00]@| GMA 4.3.8
// @[01]@|
// @[10]@| Copyright © 1992–2021 by Steven L. Willoughby
// @[11]@| (AKA Software Alchemy), Aloha, Oregon, USA. All Rights Reserved.
// @[12]@| Distributed under the terms and conditions of the BSD-3-Clause
// @[13]@| License as described in the accompanying LICENSE file distributed
// @[14]@| with GMA.
// @[15]@|
// @[20]@| Redistribution and use in source and binary forms, with or without
// @[21]@| modification, are permitted provided that the following conditions
// @[22]@| are met:
// @[23]@| 1. Redistributions of source code must retain the above copyright
// @[24]@|    notice, this list of conditions and the following disclaimer.
// @[25]@| 2. Redistributions in binary form must reproduce the above copy-
// @[26]@|    right notice, this list of conditions and the following dis-
// @[27]@|    claimer in the documentation and/or other materials provided
// @[28]@|    with the distribution.
// @[29]@| 3. Neither the name of the copyright holder nor the names of its
// @[30]@|    contributors may be used to endorse or promote products derived
// @[31]@|    from this software without specific prior written permission.
// @[32]@|
// @[33]@| THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND
// @[34]@| CONTRIBUTORS “AS IS” AND ANY EXPRESS OR IMPLIED WARRANTIES,
// @[35]@| INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF
// @[36]@| MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// @[37]@| DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS
// @[38]@| BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY,
// @[39]@| OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
// @[40]@| PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// @[41]@| PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// @[42]@| THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR
// @[43]@| TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF
// @[44]@| THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
// @[45]@| SUCH DAMAGE.
// @[46]@|
// @[50]@| This software is not intended for any use or application in which
// @[51]@| the safety of lives or property would be at risk due to failure or
// @[52]@| defect of the software.
//
