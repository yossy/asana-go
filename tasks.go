package asana

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

func PickUpTaskID(body string) string {
	// 抽出する対象を見つける
	r := regexp.MustCompile(`\[Link Asana Task\]([^)]+)`)
	target := r.FindString(body)
	// 指定のPrefixを削除する
	target = strings.TrimPrefix(target, "[Link Asana Task](")
	// タスクのリンクをコピーからだとSuffixがついてるので削除
	target = strings.TrimSuffix(target, "/f")
	// 後方からtaskのIDまでを抜き取る。
	count := utf8.RuneCountInString(target)
	suflen := count - strings.LastIndex(target, "/")
	extra := count - (suflen + 1)
	return target[extra:]
}
