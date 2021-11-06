package provider

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Dreamacro/clash/config"
	"github.com/Dreamacro/clash/hub"
	"github.com/Dreamacro/clash/log"

	C "github.com/Dreamacro/clash/constant"
)

var (
	sigCh = make(chan os.Signal, 1)
)

// Start starts the background process
func Start() error {
	if err := config.Init(C.Path.HomeDir()); err != nil {
		return fmt.Errorf("failed to initialize configuration: %s", err.Error())
	}

	if err := hub.Parse(); err != nil {
		return err
	}

	<-sigCh
	log.Infoln("Clash stopped")
	return nil
}

// Stop stops the daemon
func Stop() {
	signal.Notify(sigCh, syscall.SIGTERM)
}
