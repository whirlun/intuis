<!DOCTYPE html>
<html>
<head>
    <title>{{.title}}</title>
    <style type="text/css">
        #sortable .file_p img{
            width: 240px;
        }
    </style>
    <meta name="csrf-token" content="{{.csrftoken}}" />
    <meta name="csrf-param" content="_xsrf" />
    {{if .editor}}
    <link rel="stylesheet"href="/static/bower_components/editor.md/css/editormd.min.css" />
    {{end}}
    <link rel="stylesheet" href="/static/css/glyphicons.css" type="text/css" />
    <link rel="stylesheet" href="/static/css/main.css" type="text/css" />
    <link rel="stylesheet" href="/static/css/style.css" type="text/css"/>
    <link rel="stylesheet" href="/static/css/xq.css" type="text/css" />
    <link rel="stylesheet" href="/static/css/xq_navbar.css" type="text/css" />
    <link rel="stylesheet" href="/static/css/buttons.css" type="text/css" />
    <link rel="stylesheet" href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">

    <script src="/static/js/jquery.min.js" type="text/javascript"></script>
    <script src="/static/js/cropbox.js" type="application/javascript"></script>
    <script src="/static/js/xq_navbar.js" type="text/javascript" ></script>
    {{if .hmackey}}
    <script type="text/javascript" src="/static/bower_components/crypto-js/crypto-js.js"></script>
    {{end}}
    <script src="https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>


</head>
<body>
    <div class="header">
        <div id="logo" class="col-md-3">
            <img src="http://cdn-x-w-0.momentcdn.net/54ccee4f2bf9483a6c00001a/0.d.0.0.0/hdchina.org/styles/images/logo_hdchina.club.png.CM_WP.webp">
        </div>
        <div id="userinfo" class="col-md-5">
            {{if .isLogin}}
            <ul>
                <li>
                    <p>魔力值</p>
                    <p>30000000.0</p>
                </li>
                <li>
                    <p>上传量</p>
                    <p>30.0T</p>
                </li>
                <li>
                    <p>下载量</p>
                    <p>30.0T</p>
                </li>
                <li>
                    <p>分享率</p>
                   <p>1.00</p>
                </li>
                <li>
                    <i class="glyphicons glyphicons-arrow-up"></i>
                    <p>20</p>
                </li>
                <li>
                    <i class="glyphicons glyphicons-arrow-down"></i>
                    <p>5</p>
                </li>
            </ul>
            {{end}}
        </div>
        <div id="functionalbtn" class="col-md-4">
            <ul>
            {{if .isLogin}}
                <li>
                {{if .user.avatar_url}}
                    <img src="{{.user.avatar_url}}" />
                {{else}}
                    <img src="/static/img/default.webp" />
            {{end}}
            </li>
            <li>
                <div class="dropdown">
                    <a id="dLabel" data-target="#" href="http://example.com" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">
                        Dropdown trigger
                        <span class="caret"></span>
                    </a>
                    <ul class="dropdown-menu" aria-labelledby="dLabel">
                        <li><a href="{{urlfor "UserController.Exit"}}"></a>登出</li> </ul>
                </div>
            </li>
        </ul>
        {{end}}
    </div>
    <nav class="xq_bag" id="navbar">
        <ul class="xq_navbar">
            {{if .isLogin}}
            <li class="xq_navli"><a href="{{urlfor "IndexController.Get"}}">首页</a> </li>
            <li class="xq_navli"><a href="{{urlfor "SeedController.Get"}}">种子</a></li>
            <!--<li class="xq_navli"><%= link_to "论坛", forum_path %></li>-->
            <li class="xq_navli"><a href="{{urlfor "ForumController.Get"}}">论坛</a></li>
            <li class="xq_navli"><a href="{{urlfor "SeedController.New"}}">发布</a></li>
            <li class="xq_navli"><a href="{{urlfor "RuleController.Get"}}">规则</a></li>
            <li class="xq_navli"><a href="{{urlfor "VoteController.Get"}}">候选</a></li>
            {{else}}
            <li class="xq_navli"><a href="{{urlfor "SessionController.Get"}}">登入</a></li>
            <li class="xq_navli"><a href="{{urlfor "RegisterController.Get"}}">注册</a></li>
            {{end}}
        </ul>
    </nav>

    <div class="container-fluid">
        <div id="notices">
        {{if .flash.error }}
            <div class="alert alert-danger">{{str2html .flash.error}}</div>
        {{else if .flash.notice}}
            <div class="alert alert-success">{{.flash.success}}</div>
        {{else if .flash.warning}}
            <div class="alert alert-warning">{{.flash.warning}}</div>
        {{end}}
        </div>
