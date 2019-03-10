{{template "layout/head.tpl" .}}
<div class="row">
    <div class="col-md-6 col-md-offset-3">
        <form method="post">
        {{.xsrfdata}}
        {{i18n .Lang "username"}} <input class="form-control" type="text" name="username"/>

        {{i18n .Lang "email"}} <input class="form-control" type="email" name="email" />

        {{i18n .Lang "password"}} <input class="form-control" type="password" name="password" />

        {{i18n .Lang "passwordconfirmation"}} <input class="form-control" type="password" name="password_confirmation" />

        {{i18n .Lang "passwordquestion"}} <input class="form-control" type="text" name="password_question" />

        {{i18n .Lang "passwordanswer"}} <input class="form-control" type="text" name="password_answer" />

        {{i18n .Lang "invitationcode"}} <input class="form-control" type="text" name="invitation_code" />
        <br />
        <input class="form-control btn btn-primary" type="submit" value="{{i18n .Lang "register"}}" name="submit" />
        </form>
    </div>
</div>

{{template "layout/footer.tpl" .}}