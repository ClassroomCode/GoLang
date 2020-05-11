package notify

type notifier interface {
	Notify(string) error
}

func Update(n notifier) error {
	return n.Notify("update sent")
}
