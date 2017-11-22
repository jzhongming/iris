package main

import (
	"github.com/kataras/iris/_benchmarks/iris-mvc-templates/controllers"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

const (
	// publicDir is the exactly the same path that .NET Core is using for its templates,
	// in order to reduce the size in the repository.
	// Change the "C\\mygopath" to your own GOPATH.
	// publicDir = "C:\\mygopath\\src\\github.com\\kataras\\iris\\_benchmarks\\netcore-mvc-templates\\wwwroot"
	publicDir = "/home/kataras/mygopath/src/github.com/kataras/iris/_benchmarks/netcore-mvc-templates/wwwroot"
)

func main() {
	app := iris.New()
	app.Configure(configure)

	// app.Controller("/", new(controllers.IndexController))
	// app.Controller("/about", new(controllers.AboutController))
	// app.Controller("/contact", new(controllers.ContactController))

	app.Controller("/", new(controllers.HomeController))

	app.Run(iris.Addr(":5000"), iris.WithoutVersionChecker)
}

func configure(app *iris.Application) {
	app.RegisterView(iris.HTML("./views", ".html").Layout("shared/layout.html"))
	app.StaticWeb("/public", publicDir)
	app.OnAnyErrorCode(onError)
}

type err struct {
	Title string
	Code  int
}

func onError(ctx context.Context) {
	ctx.ViewData("", err{"Error", ctx.GetStatusCode()})
	ctx.View("shared/error.html")
}
