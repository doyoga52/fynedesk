package desktop

// Screens provides information about available physical screens for Fyne desktop
type Screens interface {
	Screens() []*Screen // Screens returns a Screen type slice of each available physical screen
	Active()  *Screen    // Active returns the screen index of the currently active screen
	Primary() *Screen    // Primary returns the screen index of the primary screen
	Scale()   float32    // Return the scale calculated for use across screens
}

// Screen provides relative information about a single physical screen
type Screen struct {
	Name                      string // Name is the randr provided name of the screen
	Index                     int    // Index is the position of the screen in the screens slice.
	X, Y, Width, Height       int    // Geometry of the screen
	ScaledX, ScaledY          int    // Scaled position of the screen
	ScaledWidth, ScaledHeight int    // Scaled size of the screen
}
