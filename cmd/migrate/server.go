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
	"gin-chat/pkg/dbs"
)

var (
	dsn      string
	driver   string
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
	StartCmd.PersistentFlags().StringVarP(&dsn, "dsn", "d", "root:root@127.0.0.1:3306/chat", "dbs dsn data source name")
	StartCmd.PersistentFlags().StringVarP(&driver, "driver", "t", "mysql", "db driver")
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
	dbs.NewBasicDB(driver, dsn)
	//4. 数据库迁移
	fmt.Println("数据库迁移开始")
	_ = migrateModel()
	fmt.Println(`数据库基础数据初始化成功`)
}

func migrateModel() error {
	db := dbs.DB
	err := db.Debug().AutoMigrate(new(orm.Migration))
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
