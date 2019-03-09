{{template "layout/head.tpl" .}}
<div>
    <ul class="list-inline">
        <li><span class="button-dropdown" data-buttons="dropdown">
        <button class="button button-primary">
            <i class="fa fa-bars"></i> {{i18n .Lang "allcategory"}}
        </button>
        <ul class="button-dropdown-list is-below">
            {{range $index, $category := .categories}}
                <li><div>
                    <span class="badge-wrapper" style="background-color:#{{$category.Color}};"></span>
                    <span>{{$category.Name}}</span>
                    <span>{{$category.Count}}</span>
                    <span>{{$category.Introduction}}</span>
                </div></li>
        </ul>
        </span></li>
        <li>
            <a href="http://www.bootcss.com/" class="button button-primary button-small">{{i18n .Lang "latest"}}</a>
        </li>
        <li>
            <a href="http://www.bootcss.com/" class="button button-primary button-small">{{i18n .Lang "hottest"}}</a>
        </li>
    </ul>
</div>
<div>
    <table>
        <thead>
        <tr>
            <th>主题</th>
             <th>分类</th>
             <th>用户</th>
            <th>回复</th>
            <th>浏览</th>
            <th>上次活动</th>
        </tr>
        </thead>
        <tbody>
        {{range $index,$thread := .threads}}
        <td class="index-table-name"><a href="/detail/{{$thread.Id}}">{{$thread.Name}}</a></td>
        <td class="index-table-content">{{$thread.Category}}</td>
        <td class="index-table-content">{{$thread.Author}}</td>
        <td class="index-table-content">{{$thread.ReplyNum}}</td>
        <td class="index-table-content">{{$thread.ReadNum}}</td>
        <td class="index-table-content">{{$thread.LastActivity}}</td>
        </tbody>
    </table>
</div>
{{template "layout/footer.tpl" .}}