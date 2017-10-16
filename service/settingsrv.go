// Solo.go - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package service

import (
	"sync"

	"github.com/b3log/solo.go/model"
	log "github.com/sirupsen/logrus"
)

var Setting = &settingService{
	mutex: &sync.Mutex{},
}

type settingService struct {
	mutex *sync.Mutex
}

func (srv *settingService) GetSetting(category, name string, blogID uint) *model.Setting {
	ret := &model.Setting{}
	if err := db.Where("category = ? AND name = ? AND blog_id = ?", category, name, blogID).Find(ret).Error; nil != err {
		log.Errorf("get setting failed: " + err.Error())

		return nil
	}

	return ret
}

func (srv *settingService) GetCategorySettings(blogID uint, category string) []*model.Setting {
	ret := []*model.Setting{}

	if err := db.Where("category = ? AND blog_id = ?", category, blogID).Find(&ret).Error; nil != err {
		log.Errorf("get category settings failed: " + err.Error())

		return nil
	}

	return ret
}

func (srv *settingService) GetAllSettings(blogID uint) []*model.Setting {
	ret := []*model.Setting{}

	if err := db.Where("category != ? AND blog_id = ?", model.SettingCategoryStatistic, blogID).Find(&ret).Error; nil != err {
		log.Errorf("get all settings failed: " + err.Error())

		return nil
	}

	return ret
}

func (srv *settingService) GetSettings(blogID uint, category string, names []string) map[string]*model.Setting {
	ret := map[string]*model.Setting{}
	settings := []*model.Setting{}
	if err := db.Where("category = ? AND name IN (?) AND blog_id = ?", category, names, blogID).Find(&settings).Error; nil != err {
		log.Errorf("get settings failed: " + err.Error())

		return nil
	}

	for _, setting := range settings {
		ret[setting.Name] = setting
	}

	return ret
}

func (srv *settingService) UpdateSettings(category string, settings []*model.Setting) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	tx := db.Begin()
	for _, setting := range settings {
		if err := tx.Model(&model.Setting{}).Where("category = ? AND name = ?", category, setting.Name).Updates(setting).Error; nil != err {
			tx.Rollback()

			return err
		}
	}
	tx.Commit()

	return nil
}
