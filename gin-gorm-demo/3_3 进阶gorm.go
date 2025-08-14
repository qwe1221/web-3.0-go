package main

/**
题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
    要求 ：使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。

题目2：关联查询
基于上述博客系统的模型定义。
    要求 ：编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。编写Go代码，使用Gorm查询评论数量最多的文章信息。

题目3：钩子函数
继续使用博客系统的模型。
    要求 ：为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/
import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

// User 用户模型
type User struct {
	gorm.Model
	Name      string `gorm:"size:50;not null"`
	Email     string `gorm:"size:100;unique;not null"`
	Password  string `gorm:"size:100;not null"`
	Posts     []Post `gorm:"foreignKey:UserID"` // 一对多关系
	PostCount int    `gorm:"default:0"`         // 文章数量统计
}

// Post 文章模型
type Post struct {
	gorm.Model
	Title         string    `gorm:"size:200;not null"`
	Content       string    `gorm:"type:text;not null"`
	UserID        uint      `gorm:"not null"` // 外键
	User          User      `gorm:"foreignKey:UserID"`
	Comments      []Comment `gorm:"foreignKey:PostID"`     // 一对多关系
	CommentCount  int       `gorm:"default:0"`             // 评论数量统计
	CommentStatus string    `gorm:"size:20;default:'无评论'"` // 评论状态
}

// Comment 评论模型
type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null"`
	PostID  uint   `gorm:"not null"` // 外键
	Post    Post   `gorm:"foreignKey:PostID"`
	UserID  uint   `gorm:"not null"`
	User    User   `gorm:"foreignKey:UserID"`
}

func initDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 打印SQL日志
	})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 自动迁移表结构
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	return db
}

// 查询用户所有文章及其评论
func GetUserPostsWithComments(db *gorm.DB, userID uint) ([]Post, error) {
	var posts []Post

	err := db.Preload("Comments").Where("user_id = ?", userID).Find(&posts).Error
	if err != nil {
		return nil, fmt.Errorf("查询用户文章失败: %w", err)
	}

	return posts, nil
}

// 查询评论最多的文章
func GetPostWithMostComments(db *gorm.DB) (*Post, error) {
	var post Post

	// 方法1: 使用子查询
	err := db.Preload("Comments").
		Order("(SELECT COUNT(*) FROM comments WHERE comments.post_id = posts.id) DESC").
		First(&post).Error

	if err != nil {
		return nil, fmt.Errorf("查询评论最多的文章失败: %w", err)
	}

	return &post, nil
}

// Post 模型的钩子函数,在创建文章前更新用户的文章计数
func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	// 增加用户的文章计数
	return tx.Model(&User{}).
		Where("id = ?", p.UserID).
		Update("post_count", gorm.Expr("post_count + 1")).Error
}

// Comment 模型的钩子函数,删除时检查评论数量
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var commentCount int64

	// 查询该文章剩余评论数量
	err = tx.Model(&Comment{}).
		Where("post_id = ?", c.PostID).
		Count(&commentCount).Error
	if err != nil {
		return err
	}

	// 更新文章的评论状态
	status := "有评论"
	if commentCount == 0 {
		status = "无评论"
	}

	return tx.Model(&Post{}).
		Where("id = ?", c.PostID).
		Updates(map[string]interface{}{
			"comment_count":  commentCount,
			"comment_status": status,
		}).Error
}

func main() {
	// 初始化数据库
	db := initDB()

	//  创建用户和文章
	user := User{
		Name:     "张三",
		Email:    "zhangsan@example.com",
		Password: "encryptedpassword",
	}
	db.Create(&user)

	post := Post{
		Title:   "GORM钩子函数使用指南",
		Content: "本文介绍如何在GORM中使用钩子函数...",
		UserID:  user.ID,
	}
	db.Create(&post) // 会自动触发BeforeCreate钩子更新用户文章计数

	// 添加评论（测试AfterDelete钩子）
	comment1 := Comment{
		Content: "很好的文章！",
		PostID:  post.ID,
		UserID:  user.ID,
	}
	db.Create(&comment1)

	comment2 := Comment{
		Content: "学到了很多",
		PostID:  post.ID,
		UserID:  user.ID,
	}
	db.Create(&comment2)

	// 删除一个评论，触发AfterDelete钩子
	db.Delete(&comment1)

	// 查询用户文章及其评论
	posts, err := GetUserPostsWithComments(db, user.ID)
	if err == nil {
		fmt.Printf("用户 %s 的文章:\n", user.Name)
		for _, p := range posts {
			fmt.Printf("- %s (评论数: %d, 状态: %s)\n",
				p.Title, p.CommentCount, p.CommentStatus)
			for _, c := range p.Comments {
				fmt.Printf("  - %s\n", c.Content)
			}
		}
	}

	// 查询评论最多的文章
	topPost, err := GetPostWithMostComments(db)
	if err == nil && topPost != nil {
		fmt.Printf("\n评论最多的文章: %s (评论数: %d)\n",
			topPost.Title, topPost.CommentCount)
	}
}
