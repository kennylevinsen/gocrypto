package testdata

type KeyTestSet struct {
	Key      []byte
	LevelKey [28][4]uint64
	FinalKey [6][4]uint32
}

type ProcessTestSet struct {
	Key         []byte
	Data        []byte
	TreeBuffer  [28][4]uint64
	Buffer      [16]byte
	BitCount    uint64
	BufferIndex uint64
}

type CompleteTestSet struct {
	Key  []byte
	Iv   []byte
	Data []byte
	Mac  []byte
}

func KeyTestSet1() (s KeyTestSet) {
	s.Key = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	s.LevelKey[0][0] = 0xd7ec096b5540299c
	s.LevelKey[0][1] = 0x69f10783acb0a2fe
	s.LevelKey[0][2] = 0x391c2844d6a36562
	s.LevelKey[0][3] = 0xd0e49b2f1e5ecdc9
	s.LevelKey[1][0] = 0xa54a875ab82c480d
	s.LevelKey[1][1] = 0x919d7c879fd99751
	s.LevelKey[1][2] = 0x5be87185ac9e48a1
	s.LevelKey[1][3] = 0x83c1f48f2d2acdb7
	s.LevelKey[2][0] = 0xe7fd107337571fb9
	s.LevelKey[2][1] = 0x8788e9a8d2ecb611
	s.LevelKey[2][2] = 0x3491c2c0fbbcb3e1
	s.LevelKey[2][3] = 0xcd44ac2db9c309e1
	s.LevelKey[3][0] = 0xddc5862928f28ddb
	s.LevelKey[3][1] = 0xb09dd3236d34ce5a
	s.LevelKey[3][2] = 0xe59e17402f372c80
	s.LevelKey[3][3] = 0xe98a15b7b1a67cad
	s.LevelKey[4][0] = 0xb6408ac01ea644ef
	s.LevelKey[4][1] = 0xe073e35a178b42d4
	s.LevelKey[4][2] = 0x3ea425a7ed48742b
	s.LevelKey[4][3] = 0x48272cd2cce8fb84
	s.LevelKey[5][0] = 0x740221d1628f37b3
	s.LevelKey[5][1] = 0xcfa9dde2430a2063
	s.LevelKey[5][2] = 0xb92c2810482df091
	s.LevelKey[5][3] = 0x59c3ea57d8280e99
	s.LevelKey[6][0] = 0xec00eabfc5f43bd8
	s.LevelKey[6][1] = 0xf2d1b908b2e01568
	s.LevelKey[6][2] = 0x259a4bec4936bfe0
	s.LevelKey[6][3] = 0x88074257f89864e7
	s.LevelKey[7][0] = 0x49a3bdcb8e2a2f42
	s.LevelKey[7][1] = 0x5023b7a0d1c81464
	s.LevelKey[7][2] = 0x0e88d4f533c9faae
	s.LevelKey[7][3] = 0x9531bb9d0bfb8026
	s.LevelKey[8][0] = 0xf7502c22c96ee564
	s.LevelKey[8][1] = 0x78f168d0258360e2
	s.LevelKey[8][2] = 0x280cf08a43a7f822
	s.LevelKey[8][3] = 0x2def426348a09a2c
	s.LevelKey[9][0] = 0xf7fa4f98d11b1410
	s.LevelKey[9][1] = 0x480622fddfcc72aa
	s.LevelKey[9][2] = 0x83e048ec4c34b5d1
	s.LevelKey[9][3] = 0xff2042c2bbc817cd
	s.LevelKey[10][0] = 0x9a3290638b5f8393
	s.LevelKey[10][1] = 0x55a50abe913cb060
	s.LevelKey[10][2] = 0x0e641d5e6f15784d
	s.LevelKey[10][3] = 0x6f3b50f8e2e5bb06
	s.LevelKey[11][0] = 0xd4343b1e6ad5053a
	s.LevelKey[11][1] = 0x5134fb38bf7e7f8b
	s.LevelKey[11][2] = 0xd57946f0aaf032b8
	s.LevelKey[11][3] = 0xe316b53fb10e9303
	s.LevelKey[12][0] = 0x769c44091dd08012
	s.LevelKey[12][1] = 0x0519a1a230ebd2d4
	s.LevelKey[12][2] = 0x075bc2be2adcc667
	s.LevelKey[12][3] = 0xbdf9cdb86dab3c84
	s.LevelKey[13][0] = 0x3c29f0f2aa33aff9
	s.LevelKey[13][1] = 0x2f73bfc552e03907
	s.LevelKey[13][2] = 0x710e7df90cb021b7
	s.LevelKey[13][3] = 0x709b73dd6d6855b7
	s.LevelKey[14][0] = 0xadc181214e199ec9
	s.LevelKey[14][1] = 0x79cf30b1f1d0368d
	s.LevelKey[14][2] = 0x7934bed2f9324eea
	s.LevelKey[14][3] = 0x1d5a1598ccfc8c75
	s.LevelKey[15][0] = 0x954dd164b9b928a3
	s.LevelKey[15][1] = 0x82ff255f693b1c9b
	s.LevelKey[15][2] = 0xd9d7c44b9fa97b8f
	s.LevelKey[15][3] = 0x109e11a7cbec32e5
	s.LevelKey[16][0] = 0x573fdb6cac3f0fc9
	s.LevelKey[16][1] = 0x6d86d3c8d312db5a
	s.LevelKey[16][2] = 0xd4f7652a9898812c
	s.LevelKey[16][3] = 0x5f7f98a30cd00e52
	s.LevelKey[17][0] = 0xd780a7e2a59dc5a4
	s.LevelKey[17][1] = 0x8165366504ecf1ee
	s.LevelKey[17][2] = 0x1302e11c403d2848
	s.LevelKey[17][3] = 0x770594b221ce77a4
	s.LevelKey[18][0] = 0xf5330a2184b1773d
	s.LevelKey[18][1] = 0x10dd41b0651e1822
	s.LevelKey[18][2] = 0xc7d23401246ce0bd
	s.LevelKey[18][3] = 0x3a071798a55927f6
	s.LevelKey[19][0] = 0x06cb134e9783987f
	s.LevelKey[19][1] = 0x29ad77d9e1cd31b9
	s.LevelKey[19][2] = 0x0221a667dbad4458
	s.LevelKey[19][3] = 0x39e4b16912a9d419
	s.LevelKey[20][0] = 0xfd00f2a57ffbe326
	s.LevelKey[20][1] = 0xf94c405370df0c6c
	s.LevelKey[20][2] = 0x1d44c379fb35c746
	s.LevelKey[20][3] = 0x45fe5f6d4359c8b4
	s.LevelKey[21][0] = 0x5c8b9df73b038d4a
	s.LevelKey[21][1] = 0x401c6059381f6344
	s.LevelKey[21][2] = 0xc2695508be85dd67
	s.LevelKey[21][3] = 0x73d22b4dad63419e
	s.LevelKey[22][0] = 0x72d028baee3712d2
	s.LevelKey[22][1] = 0xfaf047ffbf49e5bb
	s.LevelKey[22][2] = 0xa732e483163a545e
	s.LevelKey[22][3] = 0x4d9cd9497df7c92f
	s.LevelKey[23][0] = 0x7e7a5d6d2ff206a1
	s.LevelKey[23][1] = 0x8a7a574c40d75f42
	s.LevelKey[23][2] = 0x97056084b673ecf3
	s.LevelKey[23][3] = 0xe27cff1a52af2e80
	s.LevelKey[24][0] = 0xa4dd5b5fb3621246
	s.LevelKey[24][1] = 0x07478d8bbdb0ecc9
	s.LevelKey[24][2] = 0xe1a47daa2c23d7eb
	s.LevelKey[24][3] = 0x8f547d143e1dca3f
	s.LevelKey[25][0] = 0xb5acb9f9d5bc99c4
	s.LevelKey[25][1] = 0x70a69126d76de8ed
	s.LevelKey[25][2] = 0x8fb7e2db14703d37
	s.LevelKey[25][3] = 0xc022f877756da821
	s.LevelKey[26][0] = 0xb111a39bf7001564
	s.LevelKey[26][1] = 0x57dad2fc1dbf9f82
	s.LevelKey[26][2] = 0x37271e8a0e780b3e
	s.LevelKey[26][3] = 0xf70ef25c7d3990ef
	s.LevelKey[27][0] = 0x6d265f7f5014710d
	s.LevelKey[27][1] = 0xd84699a0a930ffc1
	s.LevelKey[27][2] = 0x8a90216fb95fcefd
	s.LevelKey[27][3] = 0x15339cfbefa2db81

	s.FinalKey[0][0] = 0x00000000234f4008
	s.FinalKey[0][1] = 0x000000001702f02b
	s.FinalKey[0][2] = 0x00000000e997af5a
	s.FinalKey[0][3] = 0x00000000e55f6e2e
	s.FinalKey[1][0] = 0x0000000049266f2e
	s.FinalKey[1][1] = 0x000000007f025e7e
	s.FinalKey[1][2] = 0x00000000b0481f93
	s.FinalKey[1][3] = 0x000000009dc4518c
	s.FinalKey[2][0] = 0x0000000064d80470
	s.FinalKey[2][1] = 0x0000000051248fcc
	s.FinalKey[2][2] = 0x00000000c84c3ce0
	s.FinalKey[2][3] = 0x00000000544fc9c7
	s.FinalKey[3][0] = 0x000000004b9ca496
	s.FinalKey[3][1] = 0x00000000dd444169
	s.FinalKey[3][2] = 0x00000000e2bf334c
	s.FinalKey[3][3] = 0x000000006f25d811
	s.FinalKey[4][0] = 0x000000009be6f7a8
	s.FinalKey[4][1] = 0x000000008da74069
	s.FinalKey[4][2] = 0x00000000155c6a13
	s.FinalKey[4][3] = 0x000000005279154a
	s.FinalKey[5][0] = 0x000000005823e4a6
	s.FinalKey[5][1] = 0x000000002002e359
	s.FinalKey[5][2] = 0x00000000366468ea
	s.FinalKey[5][3] = 0x0000000053ef38bb
	return
}

func KeyTestSet2() (s KeyTestSet) {
	s.Key = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	s.LevelKey[0][0] = 0xe52bdb8a47bf8f66
	s.LevelKey[0][1] = 0x2ade822b29de6c1e
	s.LevelKey[0][2] = 0x22795956c62a8db4
	s.LevelKey[0][3] = 0x986057e7a709c90e
	s.LevelKey[1][0] = 0xe645e6b98ea12439
	s.LevelKey[1][1] = 0x53e30951648a0aba
	s.LevelKey[1][2] = 0x9da3facd6efad1d1
	s.LevelKey[1][3] = 0x8dc2309420ae7ed9
	s.LevelKey[2][0] = 0x783e02fc530539d2
	s.LevelKey[2][1] = 0x623b4213868f1585
	s.LevelKey[2][2] = 0x563d9b7a11b42d81
	s.LevelKey[2][3] = 0x3be6e34b4f05fcc4
	s.LevelKey[3][0] = 0x2fa5812f7e65ef4f
	s.LevelKey[3][1] = 0x6194ce0aad4ad465
	s.LevelKey[3][2] = 0x5c89105ec73d411e
	s.LevelKey[3][3] = 0x87ede67007a1ce70
	s.LevelKey[4][0] = 0x3d8b4995d4b06aa3
	s.LevelKey[4][1] = 0xdc7551c1f9c62710
	s.LevelKey[4][2] = 0x4059d154c3fe9278
	s.LevelKey[4][3] = 0x5255843c581633a5
	s.LevelKey[5][0] = 0x49bc31c4d4306670
	s.LevelKey[5][1] = 0x46b9fa98bf9c7de0
	s.LevelKey[5][2] = 0xbb58b3d99d290871
	s.LevelKey[5][3] = 0x5f1e7ff64ae6aeb4
	s.LevelKey[6][0] = 0xbee1cea007ba3dec
	s.LevelKey[6][1] = 0x0681bb10ab0a2f2e
	s.LevelKey[6][2] = 0xaaacb820da03f435
	s.LevelKey[6][3] = 0x20a8d04655d8e4cd
	s.LevelKey[7][0] = 0xfe7020a314b65db3
	s.LevelKey[7][1] = 0x454c1878b367ac8c
	s.LevelKey[7][2] = 0x57d0e67ed3534735
	s.LevelKey[7][3] = 0xc4dae55b0d6eb3ec
	s.LevelKey[8][0] = 0x70b6a210da36cc11
	s.LevelKey[8][1] = 0x7f3704882c1b20bc
	s.LevelKey[8][2] = 0x3e4f2929ed99c525
	s.LevelKey[8][3] = 0x6b2768cc6d867c26
	s.LevelKey[9][0] = 0x6c5734cce60cf04f
	s.LevelKey[9][1] = 0x2a9260df1d6bad7e
	s.LevelKey[9][2] = 0x65de4a61cc8a19d4
	s.LevelKey[9][3] = 0x83c597ec89e32403
	s.LevelKey[10][0] = 0xc4b3eb1d46b21028
	s.LevelKey[10][1] = 0x6dbafeed92c9e318
	s.LevelKey[10][2] = 0x2da826492c11b5d5
	s.LevelKey[10][3] = 0x82b5cda1a376fb53
	s.LevelKey[11][0] = 0xfacfb83732c796b3
	s.LevelKey[11][1] = 0x36ebd6281218d430
	s.LevelKey[11][2] = 0x8f1cf1365f47553d
	s.LevelKey[11][3] = 0x7e8349c2d77f1c14
	s.LevelKey[12][0] = 0x5aa5540f07dbfe85
	s.LevelKey[12][1] = 0x19a3c24258722adf
	s.LevelKey[12][2] = 0xa149828b71699aef
	s.LevelKey[12][3] = 0x9545905b6e5a3ca7
	s.LevelKey[13][0] = 0x8cd76dd463559d10
	s.LevelKey[13][1] = 0x0c303600d9f70529
	s.LevelKey[13][2] = 0x85a31a3bbe6e3260
	s.LevelKey[13][3] = 0xfd1b31c7c5f175fb
	s.LevelKey[14][0] = 0xe5336c710c96da6d
	s.LevelKey[14][1] = 0x400c2006192b5a1c
	s.LevelKey[14][2] = 0xbf412a78b52ff6cf
	s.LevelKey[14][3] = 0x441234c939766f81
	s.LevelKey[15][0] = 0xb087980fef95c54f
	s.LevelKey[15][1] = 0xc9a95d381a69706f
	s.LevelKey[15][2] = 0x364939c6fef8e785
	s.LevelKey[15][3] = 0x59ff5fb20552f40a
	s.LevelKey[16][0] = 0xdcc83f91cc7f5c80
	s.LevelKey[16][1] = 0xb545655d06e00f81
	s.LevelKey[16][2] = 0x9e098940843622f9
	s.LevelKey[16][3] = 0x39563b63498a4954
	s.LevelKey[17][0] = 0xf6d0e0596a6b0e09
	s.LevelKey[17][1] = 0x48b2d9f99477fc56
	s.LevelKey[17][2] = 0x302dd5622f0e4b2b
	s.LevelKey[17][3] = 0x768ec8b84779c122
	s.LevelKey[18][0] = 0xba8807ec1c7270ea
	s.LevelKey[18][1] = 0xe16c13b9a4c660cb
	s.LevelKey[18][2] = 0xa5c5b92b0bc75622
	s.LevelKey[18][3] = 0x6a7f8d973b413b69
	s.LevelKey[19][0] = 0x87de2ed54b6df04d
	s.LevelKey[19][1] = 0xa4e1b1b27688589f
	s.LevelKey[19][2] = 0x8751dc7bc97dc6b9
	s.LevelKey[19][3] = 0x54a7b46e64c0b818
	s.LevelKey[20][0] = 0xf080ab1ce10bb027
	s.LevelKey[20][1] = 0xfe4d88c616655a95
	s.LevelKey[20][2] = 0x8669158c3c2fd44f
	s.LevelKey[20][3] = 0xb0d7106ad0dbe6e1
	s.LevelKey[21][0] = 0xb1e176002a8831f0
	s.LevelKey[21][1] = 0x379933df169b8e7d
	s.LevelKey[21][2] = 0xc1b7c03410230057
	s.LevelKey[21][3] = 0xb0e3599f28682d94
	s.LevelKey[22][0] = 0x79943621bb94fa49
	s.LevelKey[22][1] = 0xee2ad476ceeb23e7
	s.LevelKey[22][2] = 0x31fbe13e4259fce9
	s.LevelKey[22][3] = 0xc90553053f9bbf82
	s.LevelKey[23][0] = 0xd27e3bc10db5bb50
	s.LevelKey[23][1] = 0x5633962f56dd32ed
	s.LevelKey[23][2] = 0x4426a1808339adcf
	s.LevelKey[23][3] = 0xf62d7edc3cfe056b
	s.LevelKey[24][0] = 0xf5f41a950bbae431
	s.LevelKey[24][1] = 0x39c4af72d3ab71f2
	s.LevelKey[24][2] = 0x2ba4f37bfe54785f
	s.LevelKey[24][3] = 0xf2cdb911e5f8558e
	s.LevelKey[25][0] = 0xaf7c643c67085719
	s.LevelKey[25][1] = 0x88a50f26970931ca
	s.LevelKey[25][2] = 0x21ea0cce906998f3
	s.LevelKey[25][3] = 0xd4d6468b7a1a774f
	s.LevelKey[26][0] = 0x7b624a9b328ef427
	s.LevelKey[26][1] = 0x45e73970c876f2b4
	s.LevelKey[26][2] = 0xf70b3c0ba3adc22b
	s.LevelKey[26][3] = 0xa9b917f1fa60f8a5
	s.LevelKey[27][0] = 0x505e476bd5e48eb3
	s.LevelKey[27][1] = 0x10444b112047f184
	s.LevelKey[27][2] = 0x8c561de59885bbe3
	s.LevelKey[27][3] = 0x3089654ca9766aa2

	s.FinalKey[0][0] = 0x000000001c4af702
	s.FinalKey[0][1] = 0x00000000f56b4526
	s.FinalKey[0][2] = 0x0000000036a5d6ec
	s.FinalKey[0][3] = 0x00000000b15754f0
	s.FinalKey[1][0] = 0x0000000089c68aa7
	s.FinalKey[1][1] = 0x000000007b696c47
	s.FinalKey[1][2] = 0x00000000c59c0c39
	s.FinalKey[1][3] = 0x0000000088e8d815
	s.FinalKey[2][0] = 0x000000001673d696
	s.FinalKey[2][1] = 0x00000000da68d188
	s.FinalKey[2][2] = 0x00000000700cd451
	s.FinalKey[2][3] = 0x00000000f416a1c3
	s.FinalKey[3][0] = 0x000000008c38699d
	s.FinalKey[3][1] = 0x0000000093397810
	s.FinalKey[3][2] = 0x00000000d56d8025
	s.FinalKey[3][3] = 0x000000001b653730
	s.FinalKey[4][0] = 0x000000006705b7ed
	s.FinalKey[4][1] = 0x000000007ccd5d37
	s.FinalKey[4][2] = 0x00000000f85495d8
	s.FinalKey[4][3] = 0x00000000c6a7275e
	s.FinalKey[5][0] = 0x0000000070dc4a8d
	s.FinalKey[5][1] = 0x000000007b8f2932
	s.FinalKey[5][2] = 0x0000000004f5efd4
	s.FinalKey[5][3] = 0x000000005f29a6ac
	return
}

func ProcessTest1() (s ProcessTestSet) {
	s.Key = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	s.Data = []byte{0}
	s.BitCount = 0
	s.BufferIndex = 1
	return
}

func ProcessTest2() (s ProcessTestSet) {
	s.Key = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	s.Data = make([]byte, 1024)
	for i, _ := range s.Data {
		s.Data[i] = 0xAA
	}
	s.BitCount = 8192
	s.BufferIndex = 0
	s.TreeBuffer = [28][4]uint64{}
	s.TreeBuffer[0][0] = 0x2d369214b3334c68
	s.TreeBuffer[0][1] = 0xb1b2edd0a66e2132
	s.TreeBuffer[0][2] = 0x1db74fcff7af6dd2
	s.TreeBuffer[0][3] = 0x0bb2baae2a65c325
	s.TreeBuffer[1][0] = 0x85811fb0f4b983ae
	s.TreeBuffer[1][1] = 0xc42dc6b62d2413b7
	s.TreeBuffer[1][2] = 0x6bc693652428fe8e
	s.TreeBuffer[1][3] = 0x3cc4650b628eab91
	s.TreeBuffer[2][0] = 0x9859edbd61552ac3
	s.TreeBuffer[2][1] = 0xc432bdceca37ab27
	s.TreeBuffer[2][2] = 0x7fc1249ba0751299
	s.TreeBuffer[2][3] = 0x3de096e8d67cee81
	s.TreeBuffer[3][0] = 0xd827fa2004d602b7
	s.TreeBuffer[3][1] = 0xdd7d0c61287e9d98
	s.TreeBuffer[3][2] = 0xd1fd5afa42a0cffc
	s.TreeBuffer[3][3] = 0x52d6b6b3938b6813
	s.TreeBuffer[4][0] = 0xebe566c655f64bf7
	s.TreeBuffer[4][1] = 0x0d009c4b497a607c
	s.TreeBuffer[4][2] = 0xd51a290352ef2c83
	s.TreeBuffer[4][3] = 0x8d3c635c268d0a86
	s.TreeBuffer[5][0] = 0x3105d58fff0f453d
	s.TreeBuffer[5][1] = 0x86200e185a454daf
	s.TreeBuffer[5][2] = 0x2b4eea3a17c3a8ff
	s.TreeBuffer[5][3] = 0x73121a55b0bbee33
	s.TreeBuffer[6][0] = 0x475c6d430e6e59a3
	s.TreeBuffer[6][1] = 0x8c55fef71f9f918f
	s.TreeBuffer[6][2] = 0x49f5836c72054419
	s.TreeBuffer[6][3] = 0x19288e585ecb1bab
	return
}

func CompleteTest1() (s CompleteTestSet) {
	s.Key = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	s.Iv = []byte{0, 1, 2, 3, 4, 5, 6, 7}
	s.Data = []byte{0}
	s.Mac = []byte{0x5F, 0xAA, 0xAB, 0x85, 0xAC, 0xBE, 0x04, 0x48, 0x1D, 0xD6, 0x34, 0xD0, 0xFA, 0xD9, 0xFA, 0xFA}
	return
}

func CompleteTest2() (s CompleteTestSet) {
	s.Key = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	s.Iv = []byte{0, 1, 2, 3, 4, 5, 6, 7}
	s.Data = []byte{1}
	s.Mac = []byte{0x47, 0xEA, 0x18, 0xA1, 0x99, 0xAE, 0x07, 0x31, 0x7C, 0xA5, 0xAC, 0xC9, 0x37, 0x2F, 0x55, 0x85}
	return
}

func CompleteTest3() (s CompleteTestSet) {
	s.Key = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	s.Iv = []byte{0, 1, 2, 3, 4, 5, 6, 7}
	s.Data = datn[0:9]
	s.Mac = []byte{0xF7, 0x02, 0x3D, 0x65, 0xCF, 0x66, 0x69, 0x23, 0x47, 0xA0, 0x8B, 0x5F, 0x93, 0x55, 0x84, 0x27}
	return
}

func CompleteTest4() (s CompleteTestSet) {
	s.Key = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	s.Iv = []byte{0, 1, 2, 3, 4, 5, 6, 7}
	s.Data = datn[0:17]
	s.Mac = []byte{0x9B, 0x05, 0x8C, 0x92, 0xB7, 0x55, 0x0D, 0xF9, 0x0E, 0xD8, 0x67, 0x63, 0xC8, 0x95, 0x4B, 0x89}
	return
}

func CompleteTest5() (s CompleteTestSet) {
	s.Key = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	s.Iv = []byte{0, 1, 2, 3, 4, 5, 6, 7}
	s.Data = datn[0:27]
	s.Mac = []byte{0x39, 0x37, 0x0F, 0x67, 0x99, 0x7B, 0x03, 0x21, 0x58, 0x2B, 0x1F, 0x2C, 0x76, 0x68, 0xE6, 0x39}
	return
}

func CompleteTest6() (s CompleteTestSet) {
	s.Key = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	s.Iv = []byte{0, 1, 2, 3, 4, 5, 6, 7}
	s.Data = datn[0:33]
	s.Mac = []byte{0x38, 0x70, 0x96, 0x82, 0x7E, 0x0B, 0xD2, 0x2C, 0x37, 0xB8, 0xE1, 0xD2, 0xF7, 0xCC, 0xF2, 0x38}
	return
}

func CompleteTest7() (s CompleteTestSet) {
	s.Key = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	s.Iv = []byte{0, 1, 2, 3, 4, 5, 6, 7}
	s.Data = make([]byte, 1024)
	for i, _ := range s.Data {
		s.Data[i] = 0xAA
	}
	s.Mac = []byte{0x69, 0x51, 0x3D, 0x8F, 0x70, 0xB5, 0x62, 0x1E, 0xA9, 0x7C, 0x28, 0x16, 0xCE, 0x76, 0x65, 0x1E}
	return
}

func CompleteTest8() (s CompleteTestSet) {
	s.Key = []byte{0x76, 0xfd, 0x42, 0xb0, 0xb9, 0x86, 0x2e, 0xf1, 0x87, 0x80, 0x7b, 0x10, 0x11, 0xd0, 0xec, 0x95}
	s.Iv = []byte{0x88, 0x77, 0x66, 0x55, 0x44, 0x33, 0x22, 0x11}
	s.Data = blob1
	s.Mac = []byte{0xae, 0x91, 0xcc, 0x3c, 0x8f, 0x5f, 0xdf, 0xf8, 0xf6, 0xb1, 0xe5, 0x49, 0x47, 0x15, 0x99, 0x80}
	return
}
