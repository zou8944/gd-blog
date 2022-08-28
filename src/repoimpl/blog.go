package repoimpl

import (
	"database/sql"
	"gd-blog/src/domain/entity"
)

type BlogRepoImpl struct {
	db *sql.DB
}

func NewBlogRepoImpl(db *sql.DB) BlogRepoImpl {
	return BlogRepoImpl{db: db}
}

func (b BlogRepoImpl) SelectOne(id int) (*entity.Blog, error) {
	stmt, err := b.db.Prepare("SELECT * FROM blog WHERE id = #{id}")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	blog := &entity.Blog{}
	err = stmt.QueryRow(id).Scan(blog)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return blog, err
}

func (b BlogRepoImpl) Select(separateId int, limit int) ([]*entity.Blog, error) {
	stmt, err := b.db.Prepare("SELECT * FROM blog WHERE id < ? ORDER BY id DESC LIMIT ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	blogs := []*entity.Blog{}
	rows, err := stmt.Query(separateId, limit)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		blog := &entity.Blog{}
		err := rows.Scan(blog)
		if err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}
	return blogs, nil
}

func (b BlogRepoImpl) Search(keyword string, limit int) ([]*entity.Blog, error) {
	//TODO implement me
	panic("implement me")
}

func (b BlogRepoImpl) Insert(blog *entity.Blog) (*entity.Blog, error) {
	//TODO implement me
	panic("implement me")
}

func (b BlogRepoImpl) Update(blog *entity.Blog) (*entity.Blog, error) {
	//TODO implement me
	panic("implement me")
}

func (b BlogRepoImpl) Delete(id int) (*entity.Blog, error) {
	//TODO implement me
	panic("implement me")
}
