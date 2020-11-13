package main

import (
	"fmt"
	"os"
	"task/dao"
	"task/models"
	"task/routers"
	"task/setting"
)

func main() {
	/*
	运行提示
	*/
	if len(os.Args) < 2 {
		// 那么问题来了,怎么将 Go 生成可执行文件呢 ?
		fmt.Println("Usage on windows: task.exe conf/config.ini")
		fmt.Println("Usage on mac or unix: ./task conf/config.ini")
		return
	}

	/*
	加载配置文件
	 */
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}

	/*
	连接 sqlite3 数据库
	 */
	initErr := dao.InitSqlite(setting.Conf.SQLite3Config)
	if initErr != nil {
		fmt.Printf("init mysql failed, err:%v\n", initErr)
		return
	}

	/*
	程序退出关闭数据库连接
	*/
	defer dao.Close()

	/*
	模型绑定
	 */
	autoErr := dao.DB.AutoMigrate(&models.Todo{})
	if autoErr != nil {
		fmt.Printf("failed auto migrate the model...")
		return
	}

	/*
	注册路由
	 */
	r := routers.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
