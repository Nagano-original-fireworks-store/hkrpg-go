//go:build !push
// +build !push

package client

import (
	"github.com/gucooing/hkrpg-go/pkg/constant"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

func NewPushClient(addr string) {
	logger.Info("push client start blocked")
}

func PushServer(message constant.PushMessageAll) {}
