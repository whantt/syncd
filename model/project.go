// Copyright 2019 syncd Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package model

import(
    "time"
)

type Project struct {
    ID                  int     `gorm:"primary_key"`
    Name                string  `gorm:"type:varchar(100);not null;default:''"`
    Description         string  `gorm:"type:varchar(500);not null;default:''"`
    NeedAudit           int     `gorm:"type:int(11);not null;default:0"`
    RepoUrl             string  `gorm:"type:varchar(500);not null;default:''"`
    RepoBranch          string  `gorm:"type:varchar(100);not null;default:''"`
    PreReleaseCluster   int     `gorm:"type:int(11);not null;default:0"`
    OnlineCluster       string  `gorm:"type:varchar(1000);not null;default:''"`
    DeployUser          string  `gorm:"type:varchar(100);not null;default:''"`
    DeployPath          string  `gorm:"type:varchar(500);not null;default:''"`
    PreDeployCmd        string  `gorm:"type:text;not null"`
    AfterDeployCmd      string  `gorm:"type:text;not null"`
    DeployTimeout       int     `gorm:"type:int(11);not null;default:0"`
    Ctime               int     `gorm:"type:int(11);not null;default:0"`
}

func (m *Project) TableName() string {
    return "syd_project"
}

func (m *Project) Create() bool {
    m.Ctime = int(time.Now().Unix())
    return Create(m)
}

func (m *Project) Update() bool {
    return UpdateByPk(m)
}

func (m *Project) List(query QueryParam) ([]Project, bool) {
    var data []Project
    ok := GetMulti(&data, query)
    return data, ok
}

func (m *Project) Count(query QueryParam) (int, bool) {
    var count int
    ok := Count(m, &count, query)
    return count, ok
}

func (m *Project) Delete() bool {
    return DeleteByPk(m)
}

func (m *Project) Get(id int) bool {
    return GetByPk(m, id)
}