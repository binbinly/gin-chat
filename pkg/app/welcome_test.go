package app

import "testing"

func TestWelcomeText(t *testing.T) {
	text := WelcomeText()
	t.Logf("text:%s", text)
}
