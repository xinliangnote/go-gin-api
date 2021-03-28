/**
 * pagination.js 1.5.1
 * A jQuery plugin to provide simple yet fully customisable pagination.
 * @version 1.5.1
 * @author mss
 * @url https://github.com/Maxiaoxiang/jQuery-plugins
 *
 * @调用方法
 * $(selector).pagination(option, callback);
 * -此处callback是初始化调用，option里的callback是点击页码后调用
 * 
 * -- example --
 * $(selector).pagination({
 *     ... // 配置参数
 *     callback: function(api) {
 *         console.log('点击页码调用该回调'); //切换页码时执行一次回调
 *     }
 * }, function(){
 *     console.log('初始化'); //插件初始化时调用该方法，比如请求第一次接口来初始化分页配置
 * });
 */
;
(function (factory) {
    if (typeof define === "function" && (define.amd || define.cmd) && !jQuery) {
        // AMD或CMD
        define(["jquery"], factory);
    } else if (typeof module === 'object' && module.exports) {
        // Node/CommonJS
        module.exports = function (root, jQuery) {
            if (jQuery === undefined) {
                if (typeof window !== 'undefined') {
                    jQuery = require('jquery');
                } else {
                    jQuery = require('jquery')(root);
                }
            }
            factory(jQuery);
            return jQuery;
        };
    } else {
        //Browser globals
        factory(jQuery);
    }
}(function ($) {

    //配置参数
    var defaults = {
        totalData: 0, //数据总条数
        showData: 0, //每页显示的条数
        pageCount: 9, //总页数,默认为9
        current: 1, //当前第几页
        prevCls: 'prev', //上一页class
        nextCls: 'next', //下一页class
        prevContent: '<', //上一页内容
        nextContent: '>', //下一页内容
        activeCls: 'active', //当前页选中状态
        coping: false, //首页和尾页
        isHide: false, //当前页数为0页或者1页时不显示分页
        homePage: '', //首页节点内容
        endPage: '', //尾页节点内容
        keepShowPN: false, //是否一直显示上一页下一页
        mode: 'unfixed', //分页模式，unfixed：不固定页码数量，fixed：固定页码数量
        count: 4, //mode为unfixed时显示当前选中页前后页数，mode为fixed显示页码总数
        jump: false, //跳转到指定页数
        jumpIptCls: 'jump-ipt', //文本框内容
        jumpBtnCls: 'jump-btn', //跳转按钮
        jumpBtn: '跳转', //跳转按钮文本
        callback: function () {} //回调
    };

    var Pagination = function (element, options) {
        //全局变量
        var opts = options, //配置
            current, //当前页
            $document = $(document),
            $obj = $(element); //容器

        /**
         * 设置总页数
         * @param {int} page 页码
         * @return opts.pageCount 总页数配置
         */
        this.setPageCount = function (page) {
            return opts.pageCount = page;
        };

        /**
         * 获取总页数
         * 如果配置了总条数和每页显示条数，将会自动计算总页数并略过总页数配置，反之
         * @return {int} 总页数
         */
        this.getPageCount = function () {
            return opts.totalData && opts.showData ? Math.ceil(parseInt(opts.totalData) / opts.showData) : opts.pageCount;
        };

        /**
         * 获取当前页
         * @return {int} 当前页码
         */
        this.getCurrent = function () {
            return current;
        };

        /**
         * 填充数据
         * @param {int} 页码
         */
        this.filling = function (index) {
            var html = '';
            current = parseInt(index) || parseInt(opts.current); //当前页码
            var pageCount = this.getPageCount(); //获取的总页数
            switch (opts.mode) { //配置模式
                case 'fixed': //固定按钮模式
                    html += '<li class="page-item"><a href="javascript:;" class="page-link ' + opts.prevCls + '">' + opts.prevContent + '</a></li>';
                    if (opts.coping) {
                        var home = opts.coping && opts.homePage ? opts.homePage : '1';
                        html += '<li class="page-item"><a class="page-link" href="javascript:;" data-page="1">' + home + '</a></li>';
                    }
                    var start = current > opts.count - 1 ? current + opts.count - 1 > pageCount ? current - (opts.count - (pageCount - current)) : current - 2 : 1;
                    var end = current + opts.count - 1 > pageCount ? pageCount : start + opts.count;
                    for (; start <= end; start++) {
                        if (start != current) {
                            html += '<li class="page-item"><a class="page-link" href="javascript:;" data-page="' + start + '">' + start + '</a></li>';
                        } else {
                            html += '<li class="page-item active"><span class="page-link ' + opts.activeCls + '">' + start + '</span></li>';
                        }
                    }
                    if (opts.coping) {
                        var _end = opts.coping && opts.endPage ? opts.endPage : pageCount;
                        html += '<li class="page-item"><a class="page-link" href="javascript:;" data-page="' + pageCount + '">' + _end + '</a></li>';
                    }
                    html += '<li class="page-item"><a href="javascript:;" class="page-link ' + opts.nextCls + '">' + opts.nextContent + '</a></li>';
                    break;

                    // if (opts.keepShowPN || current > 1) { //上一页
                    //     html += '<a href="javascript:;" class="' + opts.prevCls + '">' + opts.prevContent + '</a>';
                    // } else {
                    //     if (opts.keepShowPN == false) {
                    //         $obj.find('.' + opts.prevCls) && $obj.find('.' + opts.prevCls).remove();
                    //     }
                    // }
                    // if (current >= opts.count + 2 && current != 1 && pageCount != opts.count) {
                    //     var home = opts.coping && opts.homePage ? opts.homePage : '1';
                    //     html += opts.coping ? '<a href="javascript:;" data-page="1">' + home + '</a><span>...</span>' : '';
                    // }
                    // var start = (current - opts.count) <= 1 ? 1 : (current - opts.count);
                    // var end = (current + opts.count) >= pageCount ? pageCount : (current + opts.count);
                    // for (; start <= end; start++) {
                    //     if (start <= pageCount && start >= 1) {
                    //         if (start != current) {
                    //             html += '<a href="javascript:;" data-page="' + start + '">' + start + '</a>';
                    //         } else {
                    //             html += '<span class="' + opts.activeCls + '">' + start + '</span>';
                    //         }
                    //     }
                    // }
                    // if (current + opts.count < pageCount && current >= 1 && pageCount > opts.count) {
                    //     var end = opts.coping && opts.endPage ? opts.endPage : pageCount;
                    //     html += opts.coping ? '<span>...</span><a href="javascript:;" data-page="' + pageCount + '">' + end + '</a>' : '';
                    // }
                    // if (opts.keepShowPN || current < pageCount) { //下一页
                    //     html += '<a href="javascript:;" class="' + opts.nextCls + '">' + opts.nextContent + '</a>';
                    // } else {
                    //     if (opts.keepShowPN == false) {
                    //         $obj.find('.' + opts.nextCls) && $obj.find('.' + opts.nextCls).remove();
                    //     }
                    // }
                    // break;

                case 'unfixed': //不固定按钮模式
                    if (opts.keepShowPN || current > 1) { //上一页
                        html += '<a href="javascript:;" class="' + opts.prevCls + '">' + opts.prevContent + '</a>';
                    } else {
                        if (opts.keepShowPN == false) {
                            $obj.find('.' + opts.prevCls) && $obj.find('.' + opts.prevCls).remove();
                        }
                    }
                    if (current >= opts.count + 2 && current != 1 && pageCount != opts.count) {
                        var home = opts.coping && opts.homePage ? opts.homePage : '1';
                        html += opts.coping ? '<a href="javascript:;" data-page="1">' + home + '</a><span>...</span>' : '';
                    }
                    var start = (current - opts.count) <= 1 ? 1 : (current - opts.count);
                    var end = (current + opts.count) >= pageCount ? pageCount : (current + opts.count);
                    for (; start <= end; start++) {
                        if (start <= pageCount && start >= 1) {
                            if (start != current) {
                                html += '<a href="javascript:;" data-page="' + start + '">' + start + '</a>';
                            } else {
                                html += '<span class="' + opts.activeCls + '">' + start + '</span>';
                            }
                        }
                    }
                    if (current + opts.count < pageCount && current >= 1 && pageCount > opts.count) {
                        var end = opts.coping && opts.endPage ? opts.endPage : pageCount;
                        html += opts.coping ? '<span>...</span><a href="javascript:;" data-page="' + pageCount + '">' + end + '</a>' : '';
                    }
                    if (opts.keepShowPN || current < pageCount) { //下一页
                        html += '<a href="javascript:;" class="' + opts.nextCls + '">' + opts.nextContent + '</a>';
                    } else {
                        if (opts.keepShowPN == false) {
                            $obj.find('.' + opts.nextCls) && $obj.find('.' + opts.nextCls).remove();
                        }
                    }
                    break;
                case 'easy': //简单模式
                    break;
                default:
            }
            html += opts.jump ? '<input type="text" class="' + opts.jumpIptCls + '"><a href="javascript:;" class="' + opts.jumpBtnCls + '">' + opts.jumpBtn + '</a>' : '';
            $obj.empty().html(html);
        };

        //绑定事件
        this.eventBind = function () {
            var that = this;
            var pageCount = that.getPageCount(); //总页数
            var index = 1;
            $obj.off().on('click', 'a', function () {
                if ($(this).hasClass(opts.nextCls)) {
                    if (parseInt($obj.find('.' + opts.activeCls).text()) >= pageCount) {
                        $(this).addClass('disabled');
                        return false;
                    } else {
                        index = parseInt($obj.find('.' + opts.activeCls).text()) + 1;
                    }
                } else if ($(this).hasClass(opts.prevCls)) {
                    if (parseInt($obj.find('.' + opts.activeCls).text()) <= 1) {
                        $(this).addClass('disabled');
                        return false;
                    } else {
                        index = parseInt($obj.find('.' + opts.activeCls).text()) - 1;
                    }
                } else if ($(this).hasClass(opts.jumpBtnCls)) {
                    if ($obj.find('.' + opts.jumpIptCls).val() !== '') {
                        index = parseInt($obj.find('.' + opts.jumpIptCls).val());
                    } else {
                        return;
                    }
                } else {
                    index = parseInt($(this).data('page'));
                }
                that.filling(index);
                typeof opts.callback === 'function' && opts.callback(that);
            });
            //输入跳转的页码
            $obj.on('input propertychange', '.' + opts.jumpIptCls, function () {
                var $this = $(this);
                var val = $this.val();
                var reg = /[^\d]/g;
                if (reg.test(val)) $this.val(val.replace(reg, ''));
                (parseInt(val) > pageCount) && $this.val(pageCount);
                if (parseInt(val) === 0) $this.val(1); //最小值为1
            });
            //回车跳转指定页码
            $document.keydown(function (e) {
                if (e.keyCode == 13 && $obj.find('.' + opts.jumpIptCls).val()) {
                    var index = parseInt($obj.find('.' + opts.jumpIptCls).val());
                    that.filling(index);
                    typeof opts.callback === 'function' && opts.callback(that);
                }
            });
        };

        //初始化
        this.init = function () {
            this.filling(opts.current);
            this.eventBind();
            if (opts.isHide && this.getPageCount() == '1' || this.getPageCount() == '0') {
                $obj.hide();
            } else {
                $obj.show();
            }
        };
        this.init();
    };

    $.fn.pagination = function (parameter, callback) {
        if (typeof parameter == 'function') { //重载
            callback = parameter;
            parameter = {};
        } else {
            parameter = parameter || {};
            callback = callback || function () {};
        }
        var options = $.extend({}, defaults, parameter);
        return this.each(function () {
            var pagination = new Pagination(this, options);
            callback(pagination);
        });
    };

}));