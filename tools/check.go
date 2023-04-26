package tools

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// 中文检查
func CheckChinese(input *string) error {
	// 去除首位空格
	userInput := strings.TrimSpace(*input)
	// 限制输入中只能出现指定字符集合[\u4e00-\u9fa5\d\s]+
	if ok, err := regexp.MatchString("^[a-zA-Z0-9\u4e00-\u9fa5\\s]+$", userInput); !ok {
		fmt.Println(err, ok)
		return errors.New("仅允许中文、英文字母、数字和空白字符")
	}
	return nil
}

// 手机号码检查
func CheckPhoneNumber(input *string) error {
	// 去除首位空格
	userInput := strings.TrimSpace(*input)
	// 手机号码检查
	if ok, _ := regexp.MatchString(`^1(3\d|4[57]|5[^4\D]|6[67]|7[^249\D]|8\d|9[89])\d{8}$`, userInput); !ok {
		return errors.New(" 是非法的手机号")
	}
	return nil
}
