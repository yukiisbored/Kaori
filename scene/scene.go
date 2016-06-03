// Package Scene provides an interface for scenes
package scene

import "github.com/yukiisbored/Kaori/common"

// Scene is an interface the provides the needed functions for a usual scene
type Scene interface {
	// Enter executes when the scene is going to be used
	Enter()

	// Exit executes when the scene is going to be replaced or stopped
	Exit()

	common.Updater
	common.Drawer
	common.EventHandler
}
