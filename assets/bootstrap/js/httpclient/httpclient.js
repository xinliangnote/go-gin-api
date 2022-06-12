document.write('<script type="text/javascript" src="/assets/bootstrap/js/authorization/ksort.js"></script>');
document.write('<script type="text/javascript" src="/assets/bootstrap/js/authorization/crypto-js.min.js"></script>');
document.write('<script type="text/javascript" src="/assets/bootstrap/js/authorization/hmac-sha256.js"></script>');
document.write('<script type="text/javascript" src="/assets/bootstrap/js/authorization/enc-base64.min.js"></script>');
document.write('<script type="text/javascript" src="/assets/bootstrap/js/jquery.cookie.min.js"></script>');
document.write('<div style="display:none"><script type="text/javascript">document.write(unescape("%3Cspan id=\'cnzz_stat_icon_1279911342\'%3E%3C/span%3E%3Cscript src=\'https://v1.cnzz.com/z_stat.php%3Fid%3D1279911342%26\' type=\'text/javascript\'%3E%3C/script%3E"));</script></div>');

function GenerateAuthorization(path, method, params) {
    let key = "admin";
    let secret = "12878dd962115106db6d";

    let date = new Date();
    let datetime = date.getFullYear() + "-" // "年"
        + ((date.getMonth() + 1) >= 10 ? (date.getMonth() + 1) : "0" + (date.getMonth() + 1)) + "-" // "月"
        + (date.getDate() < 10 ? "0" + date.getDate() : date.getDate()) + " " // "日"
        + (date.getHours() < 10 ? "0" + date.getHours() : date.getHours()) + ":" // "小时"
        + (date.getMinutes() < 10 ? "0" + date.getMinutes() : date.getMinutes()) + ":" // "分钟"
        + (date.getSeconds() < 10 ? "0" + date.getSeconds() : date.getSeconds()); // "秒"

    let sortParamsEncode = decodeURIComponent(jQuery.param(ksort(params)));
    let encryptStr = path + "|" + method.toUpperCase() + "|" + sortParamsEncode + "|" + datetime;
    let digest = CryptoJS.enc.Base64.stringify(CryptoJS.HmacSHA256(encryptStr, secret));
    return {authorization: key + " " + digest, date: datetime};
}

function IsJson(str) {
    if (typeof str == 'string') {
        try {
            let obj = JSON.parse(str);
            if (typeof obj == 'object' && obj) {
                return true;
            } else {
                return false;
            }
        } catch (e) {
            console.log('error：' + str + '!!!' + e);
            return false;
        }
    }
    console.log('It is not a string!')
}

function AjaxError(response) {
    let errCode = response.status;
    let errMsg = response.responseText;

    if (errCode === 401) {
        // 跳转到登录页
        if (window.frames.length !== parent.frames.length) {
            parent.window.open("/login",'_self');
        }else{
            window.open("/login",'_self');
        }
        return;
    }

    if (IsJson(response.responseText)) {
        const errInfo = JSON.parse(response.responseText);
        errCode = errInfo.code;
        errMsg = errInfo.message;
    }

    $.alert({
        title: '错误提示',
        icon: 'mdi mdi-alert',
        type: 'red',
        content: '错误码：' + errCode + '<br/>' + '错误信息：' + errMsg,
    });
}

function AjaxForm(method, url, params, beforeSendFunction, successFunction, errorFunction) {
    let authorizationData = GenerateAuthorization(url, method, params);

    $.ajax({
        url: url,
        type: method,
        data: params,
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded; charset=utf-8',
            'Authorization': authorizationData.authorization,
            'Authorization-Date': authorizationData.date,
            'Token': $.cookie("_login_token_"),
        },
        beforeSend: beforeSendFunction,
        success: successFunction,
        error: errorFunction,
    });
}

function AjaxFormNoAsync(method, url, params, beforeSendFunction, successFunction, errorFunction) {
    let authorizationData = GenerateAuthorization(url, method, params);

    $.ajax({
        url: url,
        type: method,
        data: params,
        async: false,
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded; charset=utf-8',
            'Authorization': authorizationData.authorization,
            'Authorization-Date': authorizationData.date,
            'Token': $.cookie("_login_token_"),
        },
        beforeSend: beforeSendFunction,
        success: successFunction,
        error: errorFunction,
    });
}

function AjaxPostJson(url, params, beforeSendFunction, successFunction, errorFunction) {
    let authorizationData = GenerateAuthorization(url, "POST", params);

    $.ajax({
        url: url,
        type: "POST",
        data: JSON.stringify(params),
        headers: {
            'Content-Type': 'application/json; charset=utf-8',
            'Authorization': authorizationData.authorization,
            'Authorization-Date': authorizationData.date,
            'Token': $.cookie("_login_token_"),
        },
        beforeSend: beforeSendFunction,
        success: successFunction,
        error: errorFunction,
    });
}
