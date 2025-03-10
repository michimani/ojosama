package ojosama

// equalsFeatures は Features が等しいかどうかを判定する。
//
// featuresが空の時は * が設定されているため、 * が出現したら以降は無視する。
//
// 例えば {名詞,代名詞,一般,*,*} と {名詞,代名詞,一般} を比較したとき、単純にス
// ライスを完全一致で比較すると false になるが、この関数に関しては * 以降を無視
// するため true になる。
func equalsFeatures(a, b []string) bool {
	var a2 []string
	for _, v := range a {
		if v == "*" {
			break
		}
		a2 = append(a2, v)
	}

	a = a2
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

// containsFeatures は a の中に b が含まれるかを判定する。
//
// features用。
func containsFeatures(a [][]string, b []string) bool {
	for _, a2 := range a {
		if equalsFeatures(b, a2) {
			return true
		}
	}
	return false
}

// containsString は a の中に b が含まれるかを判定する。
func containsString(a []string, b string) bool {
	for _, a2 := range a {
		if a2 == b {
			return true
		}
	}
	return false
}
