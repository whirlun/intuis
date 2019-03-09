{{template "layout/head.tpl" .}}
<h1> Edit Seed </h1>
<div class="col-md-offset-2 col-md-8">
<form id="seedform" method="POST" enctype="multipart/form-data" onsubmit="dosubmit()">
    {{.xsrfdata}}
    {{i18n .Lang "torrent"}} <input class="form-control" accept="application/x-bittorrent" type="file" name="torrent" id="torrent" data-toggle="tooltip" title="{{i18n .Lang "torrenterror"}}"/>

    {{i18n .Lang "title"}}   <input class="form-control" type="text" name="title" id="title" data-toggle="tooltip" title="{{i18n .Lang "titleerror"}}"/>

    {{i18n .Lang "subtitle"}} <input class="form-control" type="text" name="subtitle" id="subtitle" data-toggle="tooltip" title="{{i18n .Lang "subtitleerror"}}"/>

    {{i18n .Lang "3D"}} <input class="form-control checkbox" type="checkbox" name="is3D" id="is3D"/>

    {{i18n .Lang "poster"}} <input class="form-control" type="file" accept="image/jpeg, image/png" name="poster" id="poster" data-toggle="tooltip" title="{{i18n .Lang "postererror"}}"/>

    {{i18n .Lang "nfo"}} <input class="form-control" type="file" accept="text/x-nfo" name="nfo" id="nfo" data-toggle="tooltip" title="{{i18n .Lang "nfoerror"}}"/>

    {{i18n .Lang "referrule"}} <select class="form-control" id="referrule" name="referrule" data-toggle="tooltip" title="{{i18n .Lang "referruleerror"}}">
                                   {{range $index, $r := $.refer_rule}}
                                       <option value="{{$index}}" >{{$r}}</option>
                                   {{end}}
                                </select>

    {{i18n .Lang "medium"}} <select class="form-control" id="medium" name="medium" data-toggle="tooltip" title="{{i18n .Lang "mediumerror"}}">
                                {{range $index,$m := $.medium}}
                                    <option value="{{$index}}">{{$m}}</option>
                                {{end}}
                            </select>

    {{i18n .Lang "videoencode"}} <select class="form-control" id="videoencode" name="videoencode" data-toggle="tooltip" title="{{i18n .Lang "videoencodeerror"}}">
                                     {{range $index, $v := $.video_encode}}
                                     <option value="{{$index}}">{{$v}}</option>
                                     {{end}}
                                 </select>
    {{i18n .Lang "audioencode"}} <select class="form-control" id="audioencode" name="audioencode" data-toggle="tooltip" title="{{i18n .Lang "audioencodeerror"}}">
                                     {{range $index, $a := $.audio_encode}}
                                         <option value="{{$index}}">{{$a}}</option>
                                     {{end}}
                                 </select>

    {{i18n .Lang "category"}} <select class="form-control" id="category" name="category" data-toggle="tooltip" title="{{i18n .Lang "categoryerror"}}">
                                  {{range $index, $c := $.category}}
                                      <option value="{{$index}}">{{$c}}</option>
                                  {{end}}
                              </select>

    {{i18n .Lang "productiongroup"}} <select class="form-control" id="productiongroup" name="productiongroup" data-toggle="tooltip" title="{{i18n .Lang "productiongrouperror"}}">
                                      {{range $index, $m := $.production_group}}
                                          <option value="{{$index}}">{{$m}}</option>
                                      {{end}}
                                  </select>

    {{i18n .Lang "format"}} <select class="form-control" id="format" name="format" data-toggle="tooltip" title="{{i18n .Lang "formaterror"}}">
                                {{range $index, $f := $.format}}
                                    <option value="{{$index}}">{{$f}}</option>
                                {{end}}
                            </select>

    {{i18n .Lang "imdblink"}} <input class="form-control" type="text" name="imdblink" id="imdblink" data-toggle="tooltip" title="{{i18n .Lang "imdblinkerror"}}"/>
    {{i18n .Lang "doubanlink"}} <input class="form-control" type="text" name="doubanlink" id="doubanklink" data-toggle="tooltip" title="{{i18n .Lang "doubanlinkerror"}}"/>
    <br />
    <div id="content">
        <textarea name="content-markdown-doc" id="content-markdown-doc" style="display:none;" data-toggle="tooltip" title="{{i18n .Lang "contenterror"}}" data-placement="left"></textarea>
        <textarea name="content-html-code" id="content-html-code" style="display:none;"></textarea>
    </div>
        <input class="form-control" type="submit" value="{{i18n .Lang "submit"}}" onclick="return dosubmit()" />

</form>
</div>
<script src="/static/bower_components/editor.md/editormd.min.js"></script>
<script src="/static/bower_components/approvejs/dist/approve.min.js"></script>
<script src="https://cdn.bootcss.com/bootstrap/3.2.0/js/tooltip.min.js"></script>
<script type="text/javascript">
    var ruleset = {
        torrent: {required: true},
        poster: {required: true},
        nfo: {required: true},
        title: {required: true},
        subtitle: {required: true},
        is3D: {required: true},
        referrule: {required: true},
        medium: {required: true},
        videoencode: {required: true},
        audioencode: {required: true},
        category: {required: true},
        productiongroup: {required: true},
        format: {required: true},
        imdblink: {required: true},
        doubanlink: {required: true},
        "content-markdown-doc": {required: true}
    };

    var dosubmit = function() {
        var results = [];
        var inputs = [];
        for(var i in ruleset) {
            inputs.push(i);
        }
        $.each(inputs, function(index,element) {
            result = approve.value(element, ruleset[element]);
            if(result.failed.length > 0) {
                results.push(element)
            }
        });
        if(results.length != 0){
            for (var i = 0; i<results.length;i++) {
                $("#" + results[i]).tooltip("show");
            }
            return false;
        }
        return true;
    };
    $(function() {
        var editor = editormd("content", {
            height : "640",
            syncScrolling : "single",
            path : "/static/bower_components/editor.md/lib/",
            saveHTMLToTextarea : true,
            imageUpload : true,
            imageFormats : ["jpg", "jpeg", "gif", "png", "webp", "bmp"],
            imageUploadURL : "/newseed/image"
        });
    });
</script>
{{template "layout/footer.tpl" .}}