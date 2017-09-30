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

package console

import (
	"net/http"

	"github.com/b3log/solo.go/service"
	"github.com/b3log/solo.go/util"
	"github.com/gin-gonic/gin"
)

type ConsoleComment struct {
	ID            uint    `json:"id"`
	Author        *Author `json:"author"`
	ArticleAuthor *Author `json:"articleAuthor"`
	CreatedAt     string  `json:"createdAt"`
	Title         string  `gorm:"size:128" json:"title"`
	Content       string  `gorm:"type:text" json:"content"`
	Permalink     string  `json:"permalink"`
}

func GetCommentsCtl(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	sessionData := util.GetSession(c)
	commentModels, pagination := service.Comment.ConsoleGetComments(c.GetInt("p"), sessionData.BID)

	comments := []*ConsoleComment{}
	for _, commentModel := range commentModels {
		author := &Author{
			Name:      commentModel.AuthorName,
			AvatarURL: commentModel.AuthorAvatarURL,
		}
		articleAuthor := &Author{
			Name:      "article author name",
			AvatarURL: "article author avatar URL",
		}

		comment := &ConsoleComment{
			ID:            commentModel.ID,
			Author:        author,
			ArticleAuthor: articleAuthor,
			CreatedAt:     commentModel.CreatedAt.Format("2006-01-02"),
			Title:         "article title",
			Content:       commentModel.Content,
			Permalink:     sessionData.BPath + "todo comment link",
		}

		comments = append(comments, comment)
	}

	data := map[string]interface{}{}
	data["comments"] = comments
	data["pagination"] = pagination
	result.Data = data
}