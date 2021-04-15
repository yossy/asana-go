package asana

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

func PickUpTaskID(body string) (string, error) {
	// 抽出する対象を見つける
	r := regexp.MustCompile(`\[Link Asana Task\]([^)]+)`)
	target := r.FindString(body)
	// 指定のPrefixを削除する
	target = strings.TrimPrefix(target, "[Link Asana Task](")
	// タスクのリンクをコピーからだとSuffixがついてるので削除
	target = strings.TrimSuffix(target, "/f")
	// Link Asana Taskに入力がない場合は正常終了する。不正な入力の場合のみエラー。
	if target == "" {
		return "", nil
	}

	// 後方からtaskのIDまでを抜き取る。
	count := utf8.RuneCountInString(target)
	suflen := count - strings.LastIndex(target, "/")
	// "/"が含まれるのでさらに1を引く
	extra := count - (suflen - 1)
	return target[extra:], nil
}
