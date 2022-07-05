package gct

// IsBlank 是否空(空白)字符串.
func (ts *TString) IsBlank(str string) bool {
	// Check length
	if len(str) > 0 {
		// Iterate string
		for i := range str {
			// Check about char different from whitespace
			// 227为全角空格
			if str[i] > 32 && str[i] != 227 {
				return false
			}
		}
	}
	return true
}

// IsNotBlank 是否非空(非空白)字符串.
func (ts *TString) IsNotBlank(str string) bool {
	return !ts.IsBlank(str)
}

// IsEmpty 是否空字符串.
func (ts *TString) IsEmpty(str string) bool {
	return len(str) == 0
}

// IsNotEmpty 是否非空字符串.
func (ts *TString) IsNotEmpty(str string) bool {
	return !ts.IsEmpty(str)
}

// IsBlankOrEmpty 是否空白或空字符串.
func (ts *TString) IsBlankOrEmpty(str string) bool {
	return ts.IsBlank(str) || ts.IsEmpty(str)
}

// IsNotBlankOrEmpty 是否非空白或非空字符串.
func (ts *TString) IsNotBlankOrEmpty(str string) bool {
	return !ts.IsBlankOrEmpty(str)
}

// IsBlankOrEmptyOrNull 是否空白或空字符串或空指针.
func (ts *TString) IsBlankOrEmptyOrNull(str string) bool {
	return ts.IsBlankOrEmpty(str) || str == ""
}

// IsNotBlankOrEmptyOrNull 是否非空白或非空字符串或非空指针.
func (ts *TString) IsNotBlankOrEmptyOrNull(str string) bool {
	return !ts.IsBlankOrEmptyOrNull(str)
}
