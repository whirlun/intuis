{{template "layout/head.tpl" .}}
<table class="index-stripe-table">
    <thead>
    <tr>
        <td class="index-table-name">名称</td>
        <td class="index-table-wide-heading"></td>
        <td class="index-table-heading">大小</td>
        <td class="index-table-heading">做种</td>
        <td class="index-table-heading">下载</td>
        <td class="index-table-wide-heading">发布者</td>
    </tr>
    </thead>


    <tbody>
    {{ range $index, $seed := .seeds }}
    <tr>
        <td class="index-table-name"><a href="/detail/{{$seed.Id}}">{{$seed.Title}}</a></td>
        <td class="index-table-wide-content"><a href="#"><i class="icon-download-alt"></i></a><a href="#"><i class="icon-star"></i></a></td>
        <td class="index-table-content">/</td>
        <td class="index-table-content">/</td>
        <td class="index-table-content">/</td>
        <td class="index-table-wide-content"><a href="#">{{$seed.Author}}</a></td>
    </tr>
    {{end}}
    <%= will_paginate @seeds%>
    </tbody>
</table>
{{template "layout/footer.tpl" .}}