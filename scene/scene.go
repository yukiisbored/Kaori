package scene

import "github.com/yukiisbored/Kaori/common"

type Scene interface {
	Enter()
	Exit()

	common.Updater
	common.Drawer
	common.EventHandler
}
