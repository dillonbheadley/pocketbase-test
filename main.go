package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/template"
)


func main() {
    countId := "fh1a59tqqs3y0s"
    app := pocketbase.New()

    // serves static files from the provided public dir (if exists)
    app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
        registry := template.NewRegistry()

        e.Router.GET("/", func(c echo.Context) error {
            record, err := app.Dao().FindRecordById("count", countId)

            html, err := registry.LoadFiles(
                "pages/_layout.html",
                "pages/index.html",
            ).Render(record.GetInt("count"))

            if err != nil {
                return apis.NewNotFoundError("", err)
            }
            return c.HTML(http.StatusOK, html)
        })

        e.Router.POST("/count-up", func(c echo.Context) error {
            record, err := app.Dao().FindRecordById("count", countId)
            count := record.GetInt("count")
            record.Set("count", count + 1)
            if err := app.Dao().SaveRecord(record); err != nil {
                return err
            }
            html := record.GetString("count")
            if err != nil {
                return apis.NewNotFoundError("", err)
            }
            return c.HTML(http.StatusOK, html)
        })

        e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
        return nil
    })

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}