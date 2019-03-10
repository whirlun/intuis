{{template "layout/head.tpl" .}}
<div class="row">
    <div class="col-md-6 col-md-offset-3">
        <form method="post" onsubmit="return false">
        <input type="hidden" name="_hmackey" id="_hmackey" value="{{.hmackey}}"/>
        {{i18n .Lang "username"}} <input class="form-control" type="text" name="username" id="username"/>

        {{i18n .Lang "password"}} <input class="form-control" type="password" name="password" id="password"/>
        <br />
            <button class="form-control btn btn-primary" onclick="login()">{{i18n .Lang "login"}}</button>
        </form>
    </div>
</div>
<script type="text/javascript">
    var login = function() {
        var hmac = CryptoJS.algo.HMAC.create(CryptoJS.algo.SHA256, $("#_hmackey").val())
        var encrypted = CryptoJS.enc.Base64.stringify(CryptoJS.SHA256($("#password").val()))
        var hmaced = hmac.update(encrypted)
        var hash = CryptoJS.enc.Base64.stringify(hmac.finalize())
        $.ajax({
            url: "/login",
            type: "post",
            dataType: "json",
            data: {
                username: $("#username").val(),
                password: hash,
                _hmackey: $("#_hmackey").val(),
                _xsrf:{{.xsrftoken}}
            },
            success: function(data) {
                console.log(data)
                if(data.Reason != "success") {
                    $("#notices").html("<div class='alert alert-danger'>"+data.Reason+"</div>")
                }
                else {
                    window.location.href = "/"
                }
            }
        })
    }
</script>
{{template  "layout/footer.tpl" .}}