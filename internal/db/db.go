// Copyright 2022 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package db

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/tamakyi/TamaBox/internal/conf"
)

var AllTables = []interface{}{
	&User{}, &Question{}, &CensorLog{}, &UploadImage{}, &UploadImageQuestion{},
}

func Init(typ, dsn string) (*gorm.DB, error) {
	var dialector gorm.Dialector

	switch typ {
	case "mysql", "":
		dialector = mysql.Open(dsn)
	case "postgres":
		dialector = postgres.Open(dsn)
	default:
		return nil, errors.Errorf("unknown database type: %q", typ)
	}

	db, err := gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, errors.Wrap(err, "connect to database")
	}

	if err := db.AutoMigrate(AllTables...); err != nil {
		return nil, errors.Wrap(err, "auto migrate")
	}

	Users = NewUsersStore(db)
	Questions = NewQuestionsStore(db)
	CensorLogs = NewCensorLogsStore(db)
	UploadImages = NewUploadImagesStore(db)

	if err := db.Use(otelgorm.NewPlugin(
		otelgorm.WithDBName(conf.Database.Name),
	)); err != nil {
		return nil, errors.Wrap(err, "register otelgorm plugin")
	}

    if typ == "postgres" {
        // 为每个表创建序列或修复序列
        if err := fixPostgresSequences(db); err != nil {
            return nil, errors.Wrap(err, "fix postgres sequences")
        }
    }

	return db, nil
}

func fixPostgresSequences(db *gorm.DB) error {
    // 只处理确实有id字段的表
    tablesWithID := []string{"users", "questions", "censor_logs", "upload_images"}
    // 注意：移除了upload_image_questions，因为它没有id字段
    
    for _, table := range tablesWithID {
        seqName := table + "_id_seq"
        
        // 简化版本：直接尝试修复序列，如果表没有id字段会失败
        // 但对于有id字段的表，这样可以减少一次查询
        resetQuery := fmt.Sprintf(`
            -- 确保序列存在
            CREATE SEQUENCE IF NOT EXISTS nekobox.%s;
            
            -- 设置默认值
            ALTER TABLE nekobox.%s 
            ALTER COLUMN id SET DEFAULT nextval('nekobox.%s');
            
            -- 设置序列值，添加类型转换
            SELECT setval('nekobox.%s', 
                         COALESCE((SELECT MAX(id) FROM nekobox.%s), 1)::bigint, 
                         true);
        `, seqName, table, seqName, seqName, table)
        
        if err := db.Exec(resetQuery).Error; err != nil {
            // 如果表没有id字段，这里会失败
            // 但我们已经从列表中移除了这样的表
            return fmt.Errorf("修复表 %s 序列时出错: %w", table, err)
        }
    }
    return nil
}
