package evaluation

import "github.com/bytedance/sonic"

// ExtractJSONFromResponse 从文本中提取 JSON 字符串
func ExtractJSONFromResponse(text string) string {
	// 开始位置和大括号计数器
	start := -1
	braceCount := 0

	// 遍历文本，找到第一个完整的 JSON 对象
	for i := 0; i < len(text); i++ {
		// 根据大括号的数量来判断 JSON 对象的开始和结束
		switch text[i] {
		case '{':
			if start == -1 {
				start = i
			}
			braceCount++
		case '}':
			braceCount--
			if start != -1 && braceCount == 0 {
				jsonStr := text[start : i+1]
				var temp any
				if err := sonic.Unmarshal([]byte(jsonStr), &temp); err == nil {
					return jsonStr
				}
			}
		}
	}

	// 如果没有找到有效的 JSON 对象，返回空字符串
	return ""
}
