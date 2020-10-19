package notification

import (
	"github.com/go-toast/toast"
)

func Notification(title string, message string) error {
	notification := toast.Notification{
		AppID:   "Microsoft.Windows.Shell.RunDialog",
		Title:   title,
		Message: message,
		Icon:    "", // 文件必须存在
		Actions: []toast.Action{},
	}
	err := notification.Push()
	if err != nil {
		return err
	}
	return nil
}
