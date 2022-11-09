package initialize

import (
	"aixinge/config"
	"aixinge/global"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func Gorm() *gorm.DB {
	return GormMysql()
}

func IsMysql() bool {
	return global.CONFIG.System.DbType == "mysql"
}

func GormMysql() *gorm.DB {
	m := global.CONFIG.Database
	if m.Dbname == "" {
		return nil
	}
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置

	}
	db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig())
	return GormInit(db, err, m)
}

func deleteCallback(db *gorm.DB) {
	if db.Error != nil {
		return
	}

	db.Statement.SQL.Grow(100)
	db.Statement.Clauses = make(map[string]clause.Clause)
	db.Statement.AddClauseIfNotExists(clause.Update{Table: clause.Table{Name: clause.CurrentTable}})
	db.Statement.AddClause(clause.Set{
		clause.Assignment{
			Column: clause.Column{Name: "deleted_on"}, // Change field name to anything you want
			Value:  time.Now().Unix(),
		},
	})
	db.Statement.BuildClauses = []string{"UPDATE", "SET"}
	db.Statement.Build(db.Statement.BuildClauses...)

	res, err := db.Statement.ConnPool.ExecContext(db.Statement.Context, db.Statement.SQL.String(), db.Statement.Vars...)
	if db.AddError(err) == nil {
		db.RowsAffected, _ = res.RowsAffected()
	}

	log.Printf("SQL: %v\n", db.Statement.SQL.String()) // Like UPDATE `xxx` SET `deleted_on`=xxx
}

func GormInit(db *gorm.DB, err error, m config.Database) *gorm.DB {
	if err != nil {
		global.LOG.Error("database start error", zap.Any("err", err))
		os.Exit(0)
		return nil
	}
	sqlDB, _ := db.DB()

	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}

func gormConfig() *gorm.Config {
	config := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "axg_", // 设置表前缀
			SingularTable: true,   // 使用单数表名
		},
	}
	switch global.CONFIG.Database.LogMode {
	case "silent", "Silent":
		config.Logger = logger.Default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = logger.Default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = logger.Default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = logger.Default.LogMode(logger.Info)
	default:
		config.Logger = logger.Default.LogMode(logger.Info)
	}
	return config
}
