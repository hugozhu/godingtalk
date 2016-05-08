/**
 * Created by liqiao on 8/10/15.
 */

logger.i('Here we go...');

logger.i(location.href);

/**
 * _config comes from server-side template. see views/index.jade
 */
DingTalkPC.config({
    agentId: _config.agentId,
    corpId: _config.corpId,
    timeStamp: _config.timeStamp,
    nonceStr: _config.nonceStr,
    signature: _config.signature,
    jsApiList: [
        'runtime.permission.requestAuthCode',
        'device.notification.alert',
        'device.notification.confirm',
        'biz.contact.choose',
        'device.notification.prompt',
        'biz.ding.post'
        ] // 必填，需要使用的jsapi列表
});
DingTalkPC.userid=0;
DingTalkPC.ready(function(res){
    logger.i('dd.ready rocks!');

    DingTalkPC.runtime.permission.requestAuthCode({
        corpId: _config.corpId, //企业ID
        onSuccess: function(info) {
            /*{
             code: 'hYLK98jkf0m' //string authCode
             }*/
            logger.i('authcode: ' + info.code);
	    $.ajax({
                url: '/sendMsg.php',
                type:"POST",
                data: {"event":"get_userinfo","code":info.code},
                dataType:'json',
                timeout: 900,
                success: function (data, status, xhr) {
                    var info = JSON.parse(data);
                    if (info.errcode === 0) {
                        logger.i('user id: ' + info.userid);
                        DingTalkPC.userid = info.userid;
                    }
                    else {
                        logger.e('auth error: ' + data);
                    }
                },
                error: function (xhr, errorType, error) {
                    logger.e(errorType + ', ' + error);
                }
            });
        },
        onFail : function(err) {
	logger.e(JSON.stringify(err));
	}

    });
    $('.chooseonebtn').on('click', function() {

        DingTalkPC.biz.contact.choose({
            multiple: false, //是否多选： true多选 false单选； 默认true
            users: [], //默认选中的用户列表，工号；成功回调中应包含该信息
            corpId: _config.corpId, //企业id
            max: 1, //人数限制，当multiple为true才生效，可选范围1-1500
            onSuccess: function(data) {
                if(data&&data.length>0){
                    var selectUserId = data[0].emplId;
                    if(selectUserId>0){
                        DingTalkPC.device.notification.prompt({
                            message: "发送消息",
                            title: data[0].name,
                            buttonLabels: ['发送', '取消'],
                            onSuccess : function(result) {
                                var textContent = result.value;
                                if(textContent==''){
                                    return false;
                                }
                                DingTalkPC.biz.ding.post({
                                    users : [selectUserId],//用户列表，工号
                                    corpId: _config.corpId, //加密的企业id
                                    type: 1, //钉类型 1：image  2：link
                                    alertType: 2,
                                    alertDate: {"format":"yyyy-MM-dd HH:mm","value":"2016-05-09 08:00"},
                                    attachment: {
                                        images: [] //只取第一个image
                                    }, //附件信息
                                    text: textContent, //消息体
                                    onSuccess : function(info) {
                                        logger.i('DingTalkPC.biz.ding.post: info' + JSON.stringify(info));
                                    },
                                    onFail : function(err) {
                                        logger.e('DingTalkPC.biz.ding.post: info' + JSON.stringify(err));
                                    }
                                })
                                /*
                                 {
                                 buttonIndex: 0, //被点击按钮的索引值，Number类型，从0开始
                                 value: '' //输入的值
                                 }
                                 */
                            },
                            onFail : function(err) {}
                        });
                    }
                }
        },
        onFail : function(err) {}
    });
    });
    /*DingTalkPC.biz.util.uploadImage({
        multiple: false, //是否多选，默认false
        max: 5, //最多可选个数
        onSuccess : function(result) {
            logger.i(result);
        },
        onFail : function() {}
    });*/
    /*DingTalkPC.device.notification.alert({
        message: "亲爱的",
        title: "提示",//可传空
        buttonName: "收到",
        onSuccess : function() {
        },
        onFail : function(err) {}
    });*/
});


