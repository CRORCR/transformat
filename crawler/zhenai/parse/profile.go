package parse

import (
	"regexp"
	"strconv"
	"transformat/crawler/engine"
	"transformat/crawler/model"
)

//先编译,否则每次都得获得body去编译,很耗费时间
var ageRq = regexp.MustCompile(`<td><span class="label">年龄：</span>[\d]+岁</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>[^<]+</td>`)
var hightRe = regexp.MustCompile(`<td><span class="label">身高：</span>[\d]+CM</td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>[^<]+元</td>`)

func ParseProfile(contents []byte) engine.ParseResult {
	//1.创建对象,赋值
	profile := model.Profile{}
	//2.使用正则匹配各项
	age, err := strconv.Atoi(string(extractString(contents, ageRq)))
	if err == nil {
		profile.Age = age
	}
	//2.2 婚姻状态
	profile.Marriage = extractString(contents, marriageRe)

	//2.3身高
	heightStr, err := strconv.Atoi(extractString(contents, hightRe))
	if err == nil {
		profile.Height = heightStr
	}
	//2.4月收入
	profile.Income = extractString(contents, incomeRe)

	result := engine.ParseResult{
		Item: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	//必须大于2 0:正则匹配整串字符,然后才是分段字符
	if len(match) >= 2 {
		//0:整个ageRq匹配上的字符串  1:第一个正则匹配字符,然后字符串转为int
		return string(match[1])
	} else {
		return ""
	}
}
