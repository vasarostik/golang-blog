package pgsql

import (
	"strings"

	"github.com/go-pg/pg"

	"github.com/go-pg/pg/orm"
	go_blog "github.com/vasarostik/go_blog/pkg/utl/model"
)

// NewUser returns a new user database instance
func NewPost() *Post {
	return &Post{}
}

// User represents the client for user table
type Post struct{}

// Create creates a new user on database
func (u *Post) Create(db orm.DB, pst go_blog.Post) (*go_blog.Post, error) {
	var post = new(go_blog.Post)
	err := db.Model(post).Where("lower(title) = ? and deleted_at is null", strings.ToLower(pst.Title)).Select()

	if err != nil && err != pg.ErrNoRows {
		return nil, go_blog.ErrAlreadyExists

	}

	if err := db.Insert(&pst); err != nil {
		return nil, err
	}
	return &pst, nil
}

// View returns single post by ID
func (u *Post) View(db orm.DB, id int) (*go_blog.Post, error) {
	var post = new(go_blog.Post)
	sql := `SELECT "post".* FROM "posts" AS "post" WHERE ("id" = ? and deleted_at is null)`
	_, err := db.QueryOne(post, sql, id)
	if err != nil {
		return nil, err
	}

	return post, nil
}

// Update updates user's contact info
func (u *Post) Update(db orm.DB, post *go_blog.Post) error {
	_, err := db.Model(post).Where("id = ?",post.Base.ID).UpdateNotZero()
	return err
}

// MyList returns list of user`s posts
func (u *Post) MyList(db orm.DB, qp int, p *go_blog.Pagination) ([]go_blog.Post, error) {
	var posts []go_blog.Post
	q := db.Model(&posts).Column("post.*").Limit(p.Limit).Offset(p.Offset).Where("deleted_at is null and user_id = ?",qp).Order("post.id desc")

	if err := q.Select(); err != nil {
		return nil, err
	}
	return posts, nil
}

// List returns list of all posts
func (u *Post) List(db orm.DB, p *go_blog.Pagination) ([]go_blog.Post, error) {
	var posts []go_blog.Post
	q := db.Model(&posts).Column("post.*").Limit(p.Limit).Offset(p.Offset).Where("deleted_at is null").Order("post.id desc")

	if err := q.Select(); err != nil {
		return nil, err
	}
	return posts, nil
}

// Delete sets deleted_at for a user
func (u *Post) Delete(db orm.DB, post *go_blog.Post) error {
	return db.Delete(post)
}
