package builtin

import (
	"log"
	"os/exec"
	"strings"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"

	"fyne.io/desktop"
	wmtheme "fyne.io/desktop/theme"
)

const (
	networkNameEthernet = "Ethernet"
)

type network struct {
	name *widget.Label
	icon *widget.Icon
}

func (n *network) wirelessName() (string, error) {
	out, err := exec.Command("bash", []string{"-c", "iw dev `iw dev | grep Interface | cut -d \" \" -f2` info | grep ssid | cut -d \" \" -f 2"}...).Output()
	if err != nil {
		log.Println("Error running iw", err)
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}

func (n *network) isEthernetConnected() (bool, error) {
	out, err := exec.Command("bash", []string{"-c", "ip link | grep \",UP,\" | grep -v LOOPBACK | grep -v \": wl\""}...).Output()
	if err != nil {
		log.Println("Error running iw", err)
		return false, err
	}
	return strings.TrimSpace(string(out)) != "", nil
}

func (n *network) networkName() string {
	name, _ := n.wirelessName()
	if name != "" {
		return name
	}

	ether, _ := n.isEthernetConnected()
	if ether {
		return networkNameEthernet
	}
	return ""
}

func (n *network) tick() {
	tick := time.NewTicker(time.Second * 10)
	go func() {
		for {
			val := n.networkName()
			if val != n.name.Text {
				n.name.SetText(val)

				if val == "" {
					n.icon.SetResource(wmtheme.WifiOffIcon)
				} else if val == networkNameEthernet {
					n.icon.SetResource(wmtheme.EthernetIcon)
				} else {
					n.icon.SetResource(wmtheme.WifiIcon)
				}
			}
			<-tick.C
		}
	}()
}

func (n *network) StatusAreaWidget() fyne.CanvasObject {
	if _, err := n.wirelessName(); err != nil {
		if _, err = n.isEthernetConnected(); err != nil {
			return nil
		}
	}

	n.name = widget.NewLabel("")
	n.icon = widget.NewIcon(wmtheme.WifiIcon)
	n.tick()

	return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, n.icon, nil), n.icon, n.name)
}

func (n *network) Metadata() desktop.ModuleMetadata {
	return desktop.ModuleMetadata{
		Name: "Network",
	}
}

// NewNetwork creates a new module that will show network information in the status area
func NewNetwork() desktop.Module {
	return &network{}
}
