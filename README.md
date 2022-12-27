# go-embyapi
Golang Emby API Client extracted from the official [Emby.SDK](https://github.com/MediaBrowser/Emby.SDK)

<table><tr />
    <tr>
        <th valign="top" align="left">Name</th>
        <td>embyapi</td>
    </tr>
    <tr>
        <th valign="top" align="left">Language</th>
        <td>Go</td>
    </tr>
    <tr>
        <th valign="top" align="left">SDK Folder</th>
        <td>https://github.com/MediaBrowser/Emby.SDK/tree/master/SampleCode/RestApi/Clients/Go</td>
    </tr>
</table>

## Changes

* Run `goimports` & `gofumpt`
* Change package name from `embyclient-rest-go` to `embyapi`
* Prefix the following constants with `_`:
    * `0RGB__EmbyMediaModelEnumsColorFormats` in `model_emby_media_model_enums_color_formats.go`
    * `0BGR__EmbyMediaModelEnumsColorFormats` in `model_emby_media_model_enums_color_formats.go`
    * `012V__EmbyMediaModelEnumsVideoMediaTypes` in `model_emby_media_model_enums_video_media_types.go`
    * `4XM__EmbyMediaModelEnumsVideoMediaTypes` in `model_emby_media_model_enums_video_media_types.go`
    * `8BPS__EmbyMediaModelEnumsVideoMediaTypes` in `model_emby_media_model_enums_video_media_types.go`
