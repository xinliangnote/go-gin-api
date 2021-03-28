function IsJson(str) {
    if (typeof str == 'string') {
        try {
            var obj = JSON.parse(str);
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
    var errCode = response.status;
    var errMsg = response.responseText;

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
