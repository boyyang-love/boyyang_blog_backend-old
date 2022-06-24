/**
 * @Author: boyyang
 * @Date: 2022-06-09 09:46:53
 * @LastEditTime: 2022-06-24 09:52:51
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\server\setupDatabase\registerTables.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package setupDatabase

import (
	"blog/models"

	"github.com/jinzhu/gorm"
)

func RegisterTables(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.Article{},
		&models.Upload{},
		&models.PictureWall{},
		&models.ImagesTag{},
		&models.ThumbsUp{},
	)
}
