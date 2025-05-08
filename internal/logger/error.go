package logger

func HandleError(err error, shouldBubble bool) error {
	LogError(err)
	if shouldBubble {
		return err
	}
	return nil
}
