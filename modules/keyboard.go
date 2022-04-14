package modules

import (
	"github.com/micmonay/keybd_event"
	"runtime"
	"time"
)

func Keyboard() {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	// For linux, it is very important to wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	// Select keys to be pressed
	kb.SetKeys(keybd_event.VK_SPACE)

	// Set shift to be pressed
	kb.HasSuper(true)

	// Press the selected keys
	err = kb.Launching()
	if err != nil {
		panic(err)
	}
  // Select keys to be pressed
	kb.SetKeys(keybd_event.VK_T, keybd_event.VK_E, keybd_event.VK_R, keybd_event.VK_M, keybd_event.VK_ENTER)

	// Set shift to be pressed
	kb.HasSuper(false)

	// Press the selected keys
	err = kb.Launching()
	if err != nil {
		panic(err)
	}
}
