package migrate

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"text/template"
	"time"

	"github.com/binbinly/pkg/storage/orm"
	"github.com/binbinly/pkg/util/xfile"
	"github.com/spf13/cobra"

	"gin-chat/cmd/migrate/migration"
	_ "gin-chat/cmd/migrate/migration/version"
	"gin-chat/pkg/mysql"
)

var (
	host     string
	user     string
	password string
	name     string
	generate bool
	StartCmd = &cobra.Command{
		Use:     "migrate",
		Short:   "Initialize the database",
		Example: "gin-chat migrate",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&host, "host", "a", "127.0.0.1", "mysql host")
	StartCmd.PersistentFlags().StringVarP(&user, "user", "u", "root", "mysql user")
	StartCmd.PersistentFlags().StringVarP(&password, "password", "p", "root", "mysql password")
	StartCmd.PersistentFlags().StringVarP(&name, "name", "d", "chat", "mysql db name")
	StartCmd.PersistentFlags().BoolVarP(&generate, "generate", "g", false, "generate migration file")
}

func run() {
	if !generate {
		fmt.Println(`start init`)
		initDB()
	} else {
		fmt.Println(`generate migration file`)
		err := genFile()
		if err != nil {
			log.Fatal("err", err)
		}
	}
}

func initDB() {
	//3. 初始化数据库链接
	mysql.NewBasicDB(host, user, password, name)
	//4. 数据库迁移
	fmt.Println("数据库迁移开始")
	_ = migrateModel()
	fmt.Println(`数据库基础数据初始化成功`)
}

func migrateModel() error {
	db := mysql.DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")

	err := db.Debug().AutoMigrate(&orm.Migration{})
	if err != nil {
		return err
	}
	migration.Migrate.SetDb(db.Debug())
	migration.Migrate.Migrate()
	return err
}

func genFile() error {
	t1, err := template.ParseFiles("cmd/migrate/migrate.template")
	if err != nil {
		return err
	}
	m := map[string]string{}
	m["GenerateTime"] = strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	m["Package"] = "version"
	var b1 bytes.Buffer
	err = t1.Execute(&b1, m)
	if err != nil {
		return err
	}
	return xfile.Create(b1, "./cmd/migrate/migration/version/"+m["GenerateTime"]+"_migrate.go")
}
