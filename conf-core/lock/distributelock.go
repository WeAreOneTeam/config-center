package lock

import (
	mysqlstorage "config-center/conf-core/mysql-storage"
	"time"
	"github.com/robfig/cron"
)

type DistributeLock interface {
	TryLock(duration time.Duration) bool
	Unlock()
}

type ProfileDbLock struct {
	Name string

	Owner string

	lockStorage mysqlstorage.LockStorage
}

func (l *ProfileDbLock) TryLock(duration time.Duration) bool {
	lock := mysqlstorage.ProfileLockDao{
		Name:      l.Name,
		Owner:     l.Owner,
		CreatedAt: time.Now().UTC(),
		ExpiredAt: time.Now().Add(duration),
	}

	if err := l.lockStorage.Insert(lock); err != nil {
		return false
	}

	return true
}

func (l *ProfileDbLock) Unlock() {
	_ = l.lockStorage.Delete(l.Name, l.Owner)
}

func LockCleaner() {
	c := cron.New()
	_ = c.AddFunc("0/1 * * * * *", func() {
		//删除过期的锁，一秒执行一次
		//DeleteExpiredLock()
	})
	c.Start()
}

