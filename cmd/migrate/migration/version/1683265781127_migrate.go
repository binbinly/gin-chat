package version

import (
	"runtime"

	"github.com/binbinly/pkg/storage/orm"
	"gorm.io/gorm"

	"gin-chat/cmd/migrate/migration"
	"gin-chat/internal/model"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1683265781127Up)
}

func _1683265781127Up(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {

		err := tx.Debug().Migrator().AutoMigrate(
			new(model.ApplyModel),
			new(model.FriendModel),
			new(model.GroupModel),
			new(model.GroupUserModel),
			new(model.CollectModel),
			new(model.UserTagModel),
			new(model.ReportModel),
			//new(model.MessageModel),
			new(model.MomentModel),
			new(model.MomentCommentModel),
			new(model.MomentLikeModel),
			new(model.MomentTimelineModel),
			new(model.EmoticonModel),
			new(model.UserModel),
		)
		if err != nil {
			return err
		}

		return tx.Create(&orm.Migration{
			Version: version,
		}).Error
	})
}
