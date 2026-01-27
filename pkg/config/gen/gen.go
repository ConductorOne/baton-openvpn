package main

import (
	cfg "github.com/ConductorOne/baton-openvpn/pkg/config"
	"github.com/conductorone/baton-sdk/pkg/config"
)

func main() {
	config.Generate("openvpn", cfg.Config)
}
