package securitykey

var Key []byte // read only

func init() {
	var err error

	Key, err = getKey()
	if err != nil {
		panic(err)
	}
}

// 从私密存储上获取安全key
func getKey() ([]byte, error) {
	key := []byte{51, 1, 149, 66, 134, 82, 96, 95, 230, 181, 208, 63, 97, 97, 134, 78,
		11, 37, 39, 3, 60, 20, 125, 133, 152, 106, 67, 24, 74, 123, 228, 130,
		48, 23, 209, 88, 6, 35, 31, 70, 91, 32, 154, 221, 209, 82, 225, 166,
		159, 238, 116, 176, 163, 154, 99, 115, 1, 17, 38, 163, 17, 6, 123, 245,
		218, 143, 247, 255, 41, 102, 166, 157, 94, 122, 242, 139, 1, 168, 15, 60,
		163, 143, 52, 190, 146, 138, 52, 107, 90, 227, 101, 198, 76, 240, 33, 7,
		210, 96, 158, 204, 159, 95, 234, 49, 222, 150, 43, 166, 58, 80, 198, 250,
		81, 207, 26, 204, 85, 50, 113, 232, 67, 169, 160, 247, 173, 244, 136, 38,
	}
	return key, nil
}
