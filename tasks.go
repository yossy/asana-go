package asana

import (
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/pkg/errors"
)

func PickUpTaskID(body string) (string, error) {
	// 抽出する対象を見つける
	r := regexp.MustCompile(`\[Link Asana Task\]([^)]+)`)
	target := r.FindString(body)
	// 指定のPrefixを削除する
	target = strings.TrimPrefix(target, "[Link Asana Task](")
	// タスクのリンクをコピーからだとSuffixがついてるので削除
	target = strings.TrimSuffix(target, "/f")
	if target == "" {
		return "", errors.New("[Link Asana Task]の()内に紐付けるAsanaのリンクを入力して下さい。")
	}

	// 後方からtaskのIDまでを抜き取る。
	count := utf8.RuneCountInString(target)
	suflen := count - strings.LastIndex(target, "/")
	extra := count - (suflen + 1)
	return target[extra:], nil
}
