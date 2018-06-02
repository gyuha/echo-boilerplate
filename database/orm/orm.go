package orm

import (
	"echo-boilerplate/conf"
	"fmt"

	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

type (
	DBFunc func(tx *gorm.DB) error // func type which accept *gorm.DB and return error
)

// Init database
func Init() *gorm.DB {
	databaseConn()
	return db
}

// mysqlConn: setup mysql database connection using the configuration from config.toml
func databaseConn() {
	var (
		connectionString string
		err              error
	)

	if conf.Conf.Database.Adapter == "mysql" {
		connectionString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
			conf.Conf.Database.UserName, conf.Conf.Database.Pwd,
			conf.Conf.Database.Host, conf.Conf.Database.Port, conf.Conf.Database.Name)
	} else if conf.Conf.Database.Adapter == "postgres" {
		connectionString = fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
			conf.Conf.Database.UserName, conf.Conf.Database.Pwd,
			conf.Conf.Database.Host, conf.Conf.Database.Port, conf.Conf.Database.Name,
			"disable")
	} else if conf.Conf.Database.Adapter == "sqlite3" {
		connectionString = conf.Conf.Database.Name
	} else {
		panic("known database adapter..(" + conf.Conf.Database.Adapter + ")")
	}

	if db, err = gorm.Open(conf.Conf.Database.Adapter, connectionString); err != nil {
		fmt.Println("데이터 베이스 연결 오류")
		panic(err)
	}
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	db.LogMode(conf.Conf.Database.LogMode)
	db.DB().SetMaxIdleConns(conf.Conf.Database.IdleConns)
	db.DB().SetMaxOpenConns(conf.Conf.Database.OpenConns)
}

// Db : return GORM's postgres database connection instance.
func DB() *gorm.DB {
	return db
}

// Create Helper function to insert gorm model to database by using 'WithinTransaction'
func Create(v interface{}) error {
	return WithinTransaction(func(tx *gorm.DB) (err error) {
		// check new object
		if !db.NewRecord(v) {
			return err
		}
		if err = tx.Create(v).Error; err != nil {
			tx.Rollback() // rollback
			return err
		}
		return err
	})
}

// Save Helper function to save gorm model to database by using 'WithinTransaction'
func Save(v interface{}) error {
	return WithinTransaction(func(tx *gorm.DB) (err error) {
		// check new object
		if db.NewRecord(v) {
			return err
		}
		if err = tx.Save(v).Error; err != nil {
			tx.Rollback() // rollback
			return err
		}
		return err
	})
}

// Delete Helper function to save gorm model to database by using 'WithinTransaction'
func Delete(v interface{}) error {
	return WithinTransaction(func(tx *gorm.DB) (err error) {
		// check new object
		if err = tx.Delete(v).Error; err != nil {
			tx.Rollback() // rollback
			return err
		}
		return err
	})
}

// FindOneByID : Helper function to find a record by using 'WithinTransaction'
func FindOneByID(v interface{}, id uint) (err error) {
	return WithinTransaction(func(tx *gorm.DB) error {
		if err = tx.Last(v, id).Error; err != nil {
			tx.Rollback() // rollback db transaction
			return err
		}
		return err
	})
}

// FindAll : Helper function to find records by using 'WithinTransaction'
func FindAll(v interface{}) (err error) {
	return WithinTransaction(func(tx *gorm.DB) error {
		if err = tx.Find(v).Error; err != nil {
			tx.Rollback() // rollback db transaction
			return err
		}
		return err
	})
}

// FindOneByQuery : Helper function to find a record by using 'WithinTransaction'
// Example : orm.FindOneByQuery(&user, map[string]interface{}{"email": email})
func FindOneByQuery(v interface{}, params map[string]interface{}) (err error) {
	return WithinTransaction(func(tx *gorm.DB) error {
		if err = tx.Where(params).Last(v).Error; err != nil {
			tx.Rollback() // rollback db transaction
			return err
		}
		return err
	})
}

// FindByQuery : Helper function to find records by using 'WithinTransaction'
func FindByQuery(v interface{}, query interface{}, args ...interface{}) (err error) {
	return WithinTransaction(func(tx *gorm.DB) error {
		if err = tx.Where(query, args...).Find(v).Error; err != nil {
			tx.Rollback() // rollback db transaction
			return err
		}
		return err
	})
}

// FindByQueryPaging : Helper function to find records by using 'WithinTransaction'
func FindByQueryPaging(v interface{}, page int, pageSize int, query interface{}, args ...interface{}) (err error) {
	return WithinTransaction(func(tx *gorm.DB) error {
		if err = tx.Where(query, args...).Limit(pageSize).Offset((page - 1) * pageSize).Find(v).Error; err != nil {
			tx.Rollback() // rollback db transaction
			return err
		}
		return err
	})
}

// WithinTransaction : accept DBFunc as parameter
// call DBFunc function within transaction begin, and commit and return error from DBFunc
func WithinTransaction(fn DBFunc) (err error) {
	tx := db.Begin() // start db transaction
	defer tx.Commit()
	err = fn(tx)
	// close db transaction
	return err
}

// CountByQuery
func CountByQuery(table string, params map[string]interface{}) (count uint, err error) {
	return count, WithinTransaction(func(tx *gorm.DB) error {
		if err = tx.Table(table).Where(params).Count(&count).Error; err != nil {
			tx.Rollback() // rollback db transaction
			return err
		}
		return err
	})
}
