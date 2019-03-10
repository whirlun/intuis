<script type="text/javascript">
    //<![CDATA[
    $("#navbar").xq_navbar({
        "bgcolor": "#000",
        "type": "underline",
        "liwidth": "avg",
        "hcolor": "#82CEF2"
    })
    {{if .notseed}}
    window.onload = function() {
        if (!window.localStorage) {
            ;
        } else {
            var storage = window.localStorage;
            if (storage["lasturl"] != undefined ||
                    (window.location.host + window.location.pathname) != {{.mainsiteurl}} +"/newseed") {
                localState.clear()
            }
        }
    }
    {{end}}
    //]]>
</script>
</body>
</html>
