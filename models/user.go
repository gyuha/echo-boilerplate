package models

import (
	"echo-boilerplate/database/orm"
	"echo-boilerplate/utils/shortid"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User 사용자
type User struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
	Role      int        `gorm:"default:200"`           // 사용자 접근 권한. 0:관리자, 100:일반 사용자, 200:미인증
	UserCode  string     `gorm:"size:100;unique_index"` // 사용자용 코드, 이메일 인증코드로도 사용 됨.
	Name      string     `gorm:"size:255"`              // Default size for string is 255, reset it with this tag
	Image     string     `gorm:"size:50"`               // 사용자 이미지
	Birthday  *time.Time
	Email     string `gorm:"size:100"`
	Password  string `gorm:"size:100" json:"-"`
}

// BeforeCreate : before create user
func (u *User) BeforeCreate() (err error) {
	u.CreatedAt = time.Now()
	guid, _ := shortid.Generate()
	u.UserCode = guid
	return err
}

// PasswordHash : 패스워드 암호화 하기
func (u User) PasswordHash(palinPwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(palinPwd), bcrypt.DefaultCost)
	return string(hash), err
}

// ComparePassword : 패스워드 풀기
func (u User) ComparePassword(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	pain := []byte(plainPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, pain)
	if err != nil {
		return false
	}
	return true
}

// GetByCode 사용자 존재 체크
func (u *User) GetByCode(userCode string) error {
	db := orm.DB()
	return db.Where("binary(user_code) = ?", userCode).First(u).Error
}

// GetByCodes 다수의 사용자 불러오기
func (u User) GetByCodes(userCodes []string) ([]User, error) {
	db := orm.DB()
	users := []User{}
	err := db.Where("binary(user_code) in (?)", userCodes).Find(&users).Error
	return users, err
}
