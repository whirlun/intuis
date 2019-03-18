{{template "layout/head.tpl" .}}
    <div class="row">
        <img class="col-md-2" src="{{.User.Avatar}}" />
        <p class="col-md-1">{{.User.Username}}</p>
        <p>{{i18n .Lang "emailaddr"}} {{.User.Email}}</p>
        <p>{{i18n .Lang "privkey"}} {{.User.Privkey}}</p>
        <p>{{i18n .Lang "credit"}} {{.User.Credit}}</p>
    </div>
        <div class="row">
            <a href="http://www.bootcss.com/" class="button button-glow button-border button-rounded button-primary">{{i18n .Lang "general"}}</a>
            <a href="http://www.bootcss.com/" class="button button-glow button-border button-rounded button-primary">{{i18n .Lang "activity"}}</a>
            <a href="http://www.bootcss.com/" class="button button-glow button-border button-rounded button-primary">{{i18n .Lang "notification"}}</a>
            <a href="http://www.bootcss.com/" class="button button-glow button-border button-rounded button-primary">{{i18n .Lang "pm"}}</a>
            <a href="http://www.bootcss.com/" class="button button-glow button-border button-rounded button-primary">{{i18n .Lang "budge"}}</a>
            <a href="http://www.bootcss.com/" class="button button-glow button-border button-rounded button-primary">{{i18n .Lang "settings"}}</a>
        </div>
        <div class="content">
            <div id="general">
                <div id="general-stat" class="row">
                    <div class="col-md-2">{{.User.Stats.Visited}} {{i18n .Lang "visited"}}</div>
                    <div class="col-md-2">{{.User.Stats.ReadThread}} {{i18n .Lang "readthread"}}</div>
                    <div class="col-md-2">{{.User.Stats.CreateThread}} {{i18n .Lang "createthread"}}</div>
                    <div class="col-md-2">{{.User.Stats.Commented}} {{i18n .Lang "commented"}}</div>
                    <div class="col-md-2">{{.User.Stats.Loved}} {{i18n .Lang "loved"}}</div>
                    <div class="col-md-2">{{.User.Stats.Beloved}} {{i18n .Lang "beloved"}}</div>
            </div>
                <div id="general-panel" class="row">
                    <div class="col-md-4 col-md-offset-2">
                        <p>{{i18n .Lang "hotcomment"}}</p>
                        <ul>
                        {{range $index, $thread := .User.HotComments}}
                            <li>
                                <span class="general-date">{{$thread.PostTime}}</span><br />
                                <a href="/detail/{{$thread.Id}}">{{$thread.Title}}</a>
                            </li>
                        {{end}}
                        </ul>
                    </div>
                    <div class="col-md-4 col-md-offset-2">
                        <p>{{i18n .Lang "hottopic"}}</p>
                        <ul>
                            {{range $index, $thread := .User.HotTopics}}
                                <li>
                                    <span class="general-date">{{$thread.PostTime}}</span><br />
                                    <a href="/detail/{{$thread.Id}}">{{$thread.Title}}</a>
                                </li>
                            {{end}}
                        </ul>
                    </div>
                    <div class="col-md-4 col-md-offset-2">
                        <p>{{i18n .Lang "belovedmost"}}</p>
                        <ul>
                            {{range $index, $user := .User.BelovedMost}}
                                <li>
                                    <div class="general-avatar">
                                        {{if $user.avatar}}
                                            <img src="{{$user.avatar}}" />
                                        {{else}}
                                            <img src="/static/img/default.webp" />
                                        {{end}}
                                    </div>
                                    <div class="general-user-detail">
                                        <span class="general-user-detail-name">
                                            <a href="/user/{{$user.Name}}">{{$user.Name}}</a>
                                        </span>
                                        <span class="glyphicons glyphicons-heart"></span>
                                        <span>{{$user.BelovedCount}}</span>
                                    </div>
                                </li>
                            {{end}}
                        </ul>
                    </div>
                    <div class="col-md-4 col-md-offset-2">
                        <p>{{i18n .Lang "lovedmost"}}</p>
                        <ul>
                            {{range $index, $user := .User.LovedMost}}
                                <li>
                                    <div class="general-avatar">
                                        {{if $user.avatar}}
                                            <img src="{{$user.avatar}}" />
                                        {{else}}
                                            <img src="/static/img/default.webp" />
                                        {{end}}
                                    </div>
                                    <div class="general-user-detail">
                                        <span class="general-user-detail-name">
                                            <a href="/user/{{$user.Name}}">{{$user.Name}}</a>
                                        </span>
                                        <span class="glyphicons glyphicons-heart"></span>
                                        <span>{{$user.LovedCount}}</span>
                                    </div>
                                </li>
                            {{end}}
                        </ul>
                    </div>
                </div>
        </div>
        </div>
{{template  "layout/footer.tpl" .}}