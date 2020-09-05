package buffalo

import (
	// add buffalo adapter
	_ "github.com/marerm27/go-admin/adapter/buffalo"
	"github.com/marerm27/go-admin/modules/config"
	"github.com/marerm27/go-admin/modules/language"
	"github.com/marerm27/go-admin/plugins/admin/modules/table"
	"github.com/marerm27/themes/adminlte"

	// add mysql driver
	_ "github.com/marerm27/go-admin/modules/db/drivers/mysql"
	// add postgresql driver
	_ "github.com/marerm27/go-admin/modules/db/drivers/postgres"
	// add sqlite driver
	_ "github.com/marerm27/go-admin/modules/db/drivers/sqlite"
	// add mssql driver
	_ "github.com/marerm27/go-admin/modules/db/drivers/mssql"
	// add adminlte ui theme
	_ "github.com/marerm27/themes/adminlte"

	"github.com/marerm27/go-admin/template"
	"github.com/marerm27/go-admin/template/chartjs"

	"net/http"
	"os"

	"github.com/marerm27/go-admin/engine"
	"github.com/marerm27/go-admin/plugins/admin"
	"github.com/marerm27/go-admin/plugins/example"
	"github.com/marerm27/go-admin/tests/tables"
	"github.com/gobuffalo/buffalo"
)

func newHandler() http.Handler {
	bu := buffalo.New(buffalo.Options{
		Env:  "test",
		Addr: "127.0.0.1:9033",
	})

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(tables.Generators)
	adminPlugin.AddGenerator("user", tables.GetUserTable)

	examplePlugin := example.NewExample()

	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin, examplePlugin).Use(bu); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	bu.ServeFiles("/uploads", http.Dir("./uploads"))

	return bu
}

func NewHandler(dbs config.DatabaseList, gens table.GeneratorList) http.Handler {
	bu := buffalo.New(buffalo.Options{
		Env:  "test",
		Addr: "127.0.0.1:9033",
	})

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(gens)

	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfig(config.Config{
		Databases: dbs,
		UrlPrefix: "admin",
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language:    language.EN,
		IndexUrl:    "/",
		Debug:       true,
		ColorScheme: adminlte.ColorschemeSkinBlack,
	}).
		AddPlugins(adminPlugin).Use(bu); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	bu.ServeFiles("/uploads", http.Dir("./uploads"))

	return bu
}
