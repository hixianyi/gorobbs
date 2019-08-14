package util

import (
	"time"

	"github.com/mojocn/base64Captcha/store"
)

var (
	// GCLimitNumber The number of captchas created that triggers garbage collection used by default store.
	// 默认图像验证GC清理的上限个数
	GCLimitNumber = 10240
	// Expiration time of captchas used by default store.
	// 内存保存验证码的时限
	Expiration = 5 * time.Minute
	// globalStore is a shared storage for captchas, generated by New function.
	// 默认内存储存
	globalStore = store.NewMemoryStore(GCLimitNumber, Expiration)
)

func GlobalSet(key, value string) {
	globalStore.Set(key, value)
}

func GlobalGet(key string) {
	globalStore.Get(key, true)
}
