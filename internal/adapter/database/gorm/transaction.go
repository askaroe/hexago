package gorm

import (
	"context"
	"time"
)

func (c *appDB) GetFirst(ctx context.Context, dest interface{}, conds ...interface{}) (err error) {
	tx := c.db.WithContext(ctx)
	err = tx.First(&dest, conds).Error

	return
}

func (c *appDB) GetAll(ctx context.Context, dest interface{}, conds ...interface{}) (err error) {
	tx := c.db.WithContext(ctx)
	err = tx.Find(&dest, conds...).Error

	return
}

func (c *appDB) Create(ctx context.Context, value interface{}) (err error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, c.defaultTimeout*time.Second)
	defer cancel()

	tx := c.db.WithContext(ctxWithTimeout)
	err = tx.Create(&value).Error

	return
}

func (c *appDB) Update(ctx context.Context, value interface{}) (err error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, c.defaultTimeout*time.Second)
	defer cancel()

	tx := c.db.WithContext(ctxWithTimeout)
	err = tx.Updates(&value).Error

	return
}

func (c *appDB) Delete(ctx context.Context, value interface{}) (err error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, c.defaultTimeout*time.Second)
	defer cancel()

	tx := c.db.WithContext(ctxWithTimeout)
	err = tx.Delete(&value).Error

	return
}
