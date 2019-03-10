{{template "layout/head.tpl" .}}
    <div class="row">
        <img class="col-md-2" src="{{.User.Avatar}}" />
        <p class="col-md-1">{{.User.Username}}</p>
    </div>
        <div class="row">
            <a href="http://www.bootcss.com/" class="button button-glow button-border button-rounded button-primary">{{i18n .Lang general}}</a>
            <a href="http://www.bootcss.com/" class="button button-glow button-border button-rounded button-primary">{{i18n .Lang activity}}</a>
            <a href="http://www.bootcss.com/" class="button button-glow button-border button-rounded button-primary">{{i18n .Lang notification}}</a>
            <a href="http://www.bootcss.com/" class="button button-glow button-border button-rounded button-primary">{{i18n .Lang pm}}</a>
            <a href="http://www.bootcss.com/" class="button button-glow button-border button-rounded button-primary">{{i18n .Lang budge}}</a>
            <a href="http://www.bootcss.com/" class="button button-glow button-border button-rounded button-primary">{{i18n .Lang settings}}</a>
        </div>
        <div class="content">
        </div>
{{template  "layout/footer.tpl" .}}