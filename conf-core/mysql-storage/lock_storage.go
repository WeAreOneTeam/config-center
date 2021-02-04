package mysql_storage

import (
	"config-center/conf-core/exception"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type ProfileLockDao struct {
	Name string `gorm:"column:name"`

	Owner string `gorm:"column:owner"`

	CreatedAt time.Time `gorm:"column:created_at"`

	ExpiredAt time.Time `gorm:"column:expired_at"`
}

func (ProfileLockDao) TableName() string {
	return "profile_lock"
}

type LockStorage struct {
	db *gorm.DB
}

func (s *LockStorage) Insert(l ProfileLockDao) error {
	if err := s.db.Model(ProfileLockDao{}).Create(&l); err != nil {
		return exception.LOCK_EXIST
	}
	return nil
}

func (s *LockStorage) Delete(name, owner string) error {
	if err := s.db.Model(ProfileLockDao{}).Delete(ProfileLockDao{}).Where(&ProfileLockDao{
		Name:  name,
		Owner: owner,
	}); err != nil {
		return exception.INTERVAL_ERROR
	}

	return nil
}

func expiredLock(db *gorm.DB) *gorm.DB {
	return db.Where("ExpiredAt < ?", time.Now())
}

func (s * LockStorage) DeleteExpiredLock() {
	if err := s.db.Model(ProfileLockDao{}).Delete(ProfileLockDao{}).Scopes(expiredLock);err != nil {
		log.Error("delete expired lock failed")
	}
}
