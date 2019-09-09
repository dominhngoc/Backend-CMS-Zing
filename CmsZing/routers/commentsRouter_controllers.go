package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:ArtistSongController"] = append(beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:ArtistSongController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:ArtistSongController"] = append(beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:ArtistSongController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:ArtistSongController"] = append(beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:ArtistSongController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:ArtistSongController"] = append(beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:ArtistSongController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:ArtistSongController"] = append(beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:ArtistSongController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:ArtistsController"] = append(beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:ArtistsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:ArtistsController"] = append(beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:ArtistsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:ArtistsController"] = append(beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:ArtistsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:ArtistsController"] = append(beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:ArtistsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:ArtistsController"] = append(beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:ArtistsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:MigrationsController"] = append(beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:MigrationsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:MigrationsController"] = append(beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:MigrationsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:MigrationsController"] = append(beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:MigrationsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:MigrationsController"] = append(beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:MigrationsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:MigrationsController"] = append(beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:MigrationsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:SongsController"] = append(beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:SongsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:SongsController"] = append(beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:SongsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:SongsController"] = append(beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:SongsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:SongsController"] = append(beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:SongsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:SongsController"] = append(beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:SongsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/create`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:SongsController"] = append(beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:SongsController"],
        beego.ControllerComments{
            Method: "UploadFileImage",
            Router: `/file/upload/image`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:SongsController"] = append(beego.GlobalControllerRouter["backend-cms-zing/CmsZing/controllers:SongsController"],
        beego.ControllerComments{
            Method: "UploadFileMusic",
            Router: `/file/upload/music`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
