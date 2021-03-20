//Make sure jQuery has been loaded
if (typeof jQuery === "undefined") {
    throw new Error("MultiTabs requires jQuery");
}((function ($) {
    "use strict";
    var NAMESPACE, tabIndex; //variable
    var MultiTabs, handler, getTabIndex, isExtUrl, sumDomWidth, trimText, supportStorage; //function
    var defaultLayoutTemplates, defaultInit; //default variable

    NAMESPACE = '.multitabs'; // namespace for on() function

    /**
     * splice namespace for on() function, and bind it
     * @param $selector         jQuery selector
     * @param event             event
     * @param childSelector     child selector (string), same as on() function
     * @param fn                function
     * @param skipNS            bool. If true skip splice namespace
     */
    handler = function ($selector, event, childSelector, fn, skipNS) {
        var ev = skipNS ? event : event.split(' ').join(NAMESPACE + ' ') + NAMESPACE;
        $selector.off(ev, childSelector, fn).on(ev, childSelector, fn);
    };

    /**
     * get index for tab
     * @param content   content type, for 'main' tab just can be 1
     * @param capacity  capacity of tab, except 'main' tab
     * @returns int     return index
     */
    getTabIndex = function (content, capacity) {
        if (content === 'main') return 0;
        capacity = capacity || 8; //capacity of maximum tab quantity, the tab will be cover if more than it
        tabIndex = tabIndex || 0;
        tabIndex++;
        tabIndex = tabIndex % capacity;
        return tabIndex;
    };

    /**
     * trim text, remove the extra space, and trim text with maxLength, add '...' after trim.
     * @param text          the text need to trim
     * @param maxLength     max length for text
     * @returns {string}    return trimed text
     */
    trimText = function (text, maxLength) {
        maxLength = maxLength || $.fn.multitabs.defaults.navTab.maxTitleLength;
        var words = (text + "").split(' ');
        var t = '';
        for (var i = 0; i < words.length; i++) {
            var w = $.trim(words[i]);
            t += w ? (w + ' ') : '';
        }

        if (t.length > maxLength) {
            t = t.substr(0, maxLength);
            t += '...'
        }
        return t;
    };

    supportStorage = function (is_cache) {
        return !(sessionStorage === undefined) && is_cache;
    }

    /**
     * Calculate the total width
     * @param JqueryDomObjList      the object list for calculate
     * @returns {number}        return total object width (int)
     */
    sumDomWidth = function (JqueryDomObjList) {
        var width = 0;
        $(JqueryDomObjList).each(function () {
            width += $(this).outerWidth(true)
        });
        return width
    };

    /**
     * Judgment is external URL
     * @param url           URL for judgment
     * @returns {boolean}   external URL return true, local return false
     */
    isExtUrl = function (url) {
        var absUrl = (function (url) {
            var a = document.createElement('a');
            a.href = url;
            return a.href;
        })(url);
        var webRoot = window.location.protocol + '//' + window.location.host + '/';
        var urlRoot = absUrl.substr(0, webRoot.length);
        return (!(urlRoot === webRoot));
    };

    /**
     * Layout Templates
     */
    defaultLayoutTemplates = {
        /**
         * Main Layout
         */
        default: '<div class="mt-wrapper {mainClass}" style="height: 100%;" >' +
            '<div class="mt-nav-bar {navClass}" style="background-color: {backgroundColor};">' +
            '<div class="mt-nav mt-nav-tools-left">' +
            '<ul  class="nav {nav-tabs}">' +
            '<li class="nav-item mt-move-left"><a class="nav-link"><i class="mdi mdi-skip-backward"></i></a></li>' +
            '</ul>' +
            '</div>' +
            '<nav class="mt-nav mt-nav-panel">' +
            '<ul  class="nav {nav-tabs}"></ul>' +
            '</nav>' +
            '<div class="mt-nav mt-nav-tools-right">' +
            '<ul  class="nav {nav-tabs}">' +
            '<li class="nav-item mt-move-right"><a class="nav-link"><i class="mdi mdi-skip-forward"></i></a></li>' +
            '<li class="nav-item mt-dropdown dropdown">' +
            '<a href="#" class="nav-link dropdown-toggle" data-toggle="dropdown">{dropdown}<span class="caret"></span></a>' +
            '<ul role="menu" class="dropdown-menu dropdown-menu-right">' +
            '<li class="mt-show-actived-tab"><a class="dropdown-item">{showActivedTab}</a></li>' +
            '<li class="dropdown-divider"></li>' +
            '<li class="mt-close-all-tabs"><a class="dropdown-item">{closeAllTabs}</a></li>' +
            '<li class="mt-close-other-tabs"><a class="dropdown-item">{closeOtherTabs}</a></li>' +
            '</ul>' +
            '</li>' +
            '</ul>' +
            '</div>' +
            '</div>' +
            '<div class="tab-content mt-tab-content " > </div>' +
            '</div>',
        classic: '<div class="mt-wrapper {mainClass}" style="height: 100%;" >' +
            '<div class="mt-nav-bar {navClass}" style="background-color: {backgroundColor};">' +
            '<nav class="mt-nav mt-nav-panel">' +
            '<ul  class="nav {nav-tabs}"> </ul>' +
            '</nav>' +
            '<div class="mt-nav mt-nav-tools-right">' +
            '<ul  class="nav {nav-tabs}">' +
            '<li class="mt-dropdown dropdown">' +
            '<a href="#"  class="dropdown-toggle dropdown-item" data-toggle="dropdown">{dropdown}<span class="caret"></span></a>' +
            '<ul role="menu" class="mt-hidden-list dropdown-menu dropdown-menu-right"></ul>' +
            '</li>' +
            '</ul>' +
            '</div>' +
            '</div>' +
            '<div class="tab-content mt-tab-content " > </div>' +
            '</div>',
        simple: '<div class="mt-wrapper {mainClass}" style="height: 100%;" >' +
            '<div class="mt-nav-bar {navClass}" style="background-color: {backgroundColor};">' +
            '<nav class="mt-nav mt-nav-panel">' +
            '<ul  class="nav {nav-tabs}"> </ul>' +
            '</nav>' +
            '</div>' +
            '<div class="tab-content mt-tab-content " > </div>' +
            '</div>',
        navTab: '<a data-id="{navTabId}" class="nav-link mt-nav-tab" data-type="{type}" data-index="{index}" data-url="{url}">{title}</a>',
        closeBtn: ' <i class="mt-close-tab mdi mdi-close" style="{style}"></i>',
        ajaxTabPane: '<div id="{tabPaneId}" class="tab-pane {class}">{content}</div>',
        iframeTabPane: '<iframe id="{tabPaneId}" class="tab-pane {class}"  width="100%" height="100%" frameborder="0" src="" seamless></iframe>'
    };

    /**
     * Default init page
     * @type {*[]}
     */
    defaultInit = [{ //default tabs in initial;
        type: 'main', //default is info;
        title: 'main', //default title;
        content: '<h1>Demo page</h1><h2>Welcome to use bootstrap multi-tabs :) </h2>' //default content
    }];

    /**
     * multitabs constructor
     * @param element       Primary container
     * @param options       options
     * @constructor
     */
    MultiTabs = function (element, options) {
        var self = this;
        self.$element = $(element);
        self._init(options)._listen()._final();
    };

    /**
     * MultiTabs's function
     */
    MultiTabs.prototype = {
        /**
         * constructor
         */
        constructor: MultiTabs,

        /**
         * create tab and return this.
         * @param obj           the obj to trigger multitabs
         * @param active        if true, active tab after create
         * @returns this        Chain structure.
         */
        create: function (obj, active) {
            var options = this.options;
            var param, $navTab;
            if (!(param = this._getParam(obj))) {
                return this; //return multitabs obj when is invaid obj
            }
            $navTab = this._exist(param)
            if ($navTab && !param.isNewTab) {
                this.active($navTab);
                return this;
            }
            param.active = !param.active ? active : param.active;
            //nav tab create
            $navTab = this._createNavTab(param);
            //tab-pane create
            this._createTabPane(param);
            //add tab to storage
            this._storage(param.did, param);
            if (param.active) {
                this.active($navTab);
            }
            return this;
        },

        /**
         * Create tab pane
         * @param param
         * @param index
         * @returns {*|{}}
         * @private
         */
        _createTabPane: function (param) {
            var self = this,
                $el = self.$element;
            $el.tabContent.append(self._getTabPaneHtml(param));
            return $el.tabContent.find('#' + param.did);
        },

        /**
         * get tab pane html
         * @param param
         * @param index
         * @returns {string}
         * @private
         */
        _getTabPaneHtml: function (param) {
            var self = this,
                options = self.options;
            if (!param.content && param.iframe) {
                return defaultLayoutTemplates.iframeTabPane
                    .replace('{class}', options.content.iframe.class)
                    .replace('{tabPaneId}', param.did);
            } else {
                return defaultLayoutTemplates.ajaxTabPane
                    .replace('{class}', options.content.ajax.class)
                    .replace('{tabPaneId}', param.did)
                    .replace('{content}', param.content);
            }
        },

        /**
         * create nav tab
         * @param param
         * @param index
         * @returns {*|{}}
         * @private
         */
        _createNavTab: function (param) {
            var self = this,
                $el = self.$element;
            var navTabHtml = self._getNavTabHtml(param);
            var $navTabLi = $el.navPanelList.find('a[data-type="' + param.type + '"][data-index="' + param.index + '"]').parent('li');
            if ($navTabLi.length) {
                $navTabLi.html(navTabHtml);
                self._getTabPane($navTabLi.find('a:first')).remove(); //remove old content pane directly
            } else {
                $el.navPanelList.append('<li class="nav-item">' + navTabHtml + '</li>');
            }
            return $el.navPanelList.find('a[data-type="' + param.type + '"][data-index="' + param.index + '"]:first');

        },

        /**
         * get nav tab html
         * @param param
         * @param index
         * @returns {string}
         * @private
         */
        _getNavTabHtml: function (param) {
            var self = this,
                options = self.options;
            var closeBtnHtml, display;

            display = options.nav.showCloseOnHover ? '' : 'display:inline;';
            closeBtnHtml = (param.type === 'main') ? '' : defaultLayoutTemplates.closeBtn.replace('{style}', display); //main content can not colse.
            return defaultLayoutTemplates.navTab
                .replace('{index}', param.index)
                .replace('{navTabId}', param.did)
                .replace('{url}', param.url)
                .replace('{title}', param.title)
                .replace('{type}', param.type) +
                closeBtnHtml;
        },

        /**
         * generate tab pane's id
         * @param param
         * @param index
         * @returns {string}
         * @private
         */
        _generateId: function (param) {
            return 'multitabs_' + param.type + '_' + param.index;
        },

        /**
         * active navTab
         * @param navTab
         * @returns self      Chain structure.
         */
        active: function (navTab, isNavBar) {
            var self = this,
                $el = self.$element;
                isNavBar = (isNavBar == false) ? false : true;
            var $navTab = self._getNavTab(navTab),
                $tabPane = self._getTabPane($navTab),
                $prevActivedTab = $el.navPanelList.find('li a.active');
            var prevNavTabParam = $prevActivedTab.length ? self._getParam($prevActivedTab) : {};
            var navTabParam = $navTab.length ? self._getParam($navTab) : {};
            //change storage active status
            var storage = self._storage();
            if (storage[prevNavTabParam.id]) {
                storage[prevNavTabParam.id].active = false;
            }
            if (storage[navTabParam.id]) {
                storage[navTabParam.id].active = true;
            }
            self._resetStorage(storage);
            //active navTab and tabPane
            $prevActivedTab.removeClass('active');
            $navTab.addClass('active');
            self._fixTabPosition($navTab);
            self._getTabPane($prevActivedTab).removeClass('active');
            $tabPane.addClass('active');
            self._fixTabContentLayout($tabPane);
            //fill tab pane
            self._fillTabPane($tabPane, navTabParam, isNavBar);
            
            return self;
        },
        /**
         * fill tab pane
         * @private
         */
        _fillTabPane: function (tabPane, param, isNavBar) {
            var self = this,
                options = self.options;
            var $tabPane = $(tabPane);
            //if navTab-pane empty, load content
            if (!$tabPane.html()) {
                if ($tabPane.is('iframe')) {
                  
                    //if (!$tabPane.attr('src')) {
                    if ((!$tabPane.attr('src') && options.refresh == 'no') || (options.refresh == 'nav' && isNavBar) || options.refresh == 'all'){
                        $tabPane.attr('src', param.url);
                    }
                } else {
                    $.ajax({
                        url: param.url,
                        dataType: "html",
                        success: function (callback) {
                            $tabPane.html(options.content.ajax.success(callback));
                        },
                        error: function (callback) {
                            $tabPane.html(options.content.ajax.error(callback));
                        }
                    });
                }

            }
        },
        /**
         * move left
         * @return self
         */
        moveLeft: function () {
            var self = this,
                $el = self.$element,
                navPanelListMarginLeft = Math.abs(parseInt($el.navPanelList.css("margin-left"))),
                navPanelWidth = $el.navPanel.outerWidth(true),
                sumTabsWidth = sumDomWidth($el.navPanelList.children('li')),
                leftWidth = 0,
                marginLeft = 0,
                $navTabLi;
            if (sumTabsWidth < navPanelWidth) {
                return self
            } else {
                $navTabLi = $el.navPanelList.children('li:first');
                while ((marginLeft + $navTabLi.width()) <= navPanelListMarginLeft) {
                    marginLeft += $navTabLi.outerWidth(true);
                    $navTabLi = $navTabLi.next();
                }
                marginLeft = 0;
                if (sumDomWidth($navTabLi.prevAll()) > navPanelWidth) {
                    while (((marginLeft + $navTabLi.width()) < navPanelWidth) && $navTabLi.length > 0) {
                        marginLeft += $navTabLi.outerWidth(true);
                        $navTabLi = $navTabLi.prev();
                    }
                    leftWidth = sumDomWidth($navTabLi.prevAll());
                }
            }
            $el.navPanelList.animate({
                marginLeft: 0 - leftWidth + "px"
            }, "fast");
            return self;
        },

        /**
         * move right
         * @return self
         */
        moveRight: function () {
            var self = this,
                $el = self.$element,
                navPanelListMarginLeft = Math.abs(parseInt($el.navPanelList.css("margin-left"))),
                navPanelWidth = $el.navPanel.outerWidth(true),
                sumTabsWidth = sumDomWidth($el.navPanelList.children('li')),
                leftWidth = 0,
                $navTabLi, marginLeft;
            if (sumTabsWidth < navPanelWidth) {
                return self;
            } else {
                $navTabLi = $el.navPanelList.children('li:first');
                marginLeft = 0;
                while ((marginLeft + $navTabLi.width()) <= navPanelListMarginLeft) {
                    marginLeft += $navTabLi.outerWidth(true);
                    $navTabLi = $navTabLi.next();
                }
                marginLeft = 0;
                while (((marginLeft + $navTabLi.width()) < navPanelWidth) && $navTabLi.length > 0) {
                    marginLeft += $navTabLi.outerWidth(true);
                    $navTabLi = $navTabLi.next();
                }
                leftWidth = sumDomWidth($navTabLi.prevAll());
                if (leftWidth > 0) {
                    $el.navPanelList.animate({
                        marginLeft: 0 - leftWidth + "px"
                    }, "fast");
                }
            }
            return self;
        },

        /**
         * close navTab
         * @param navTab
         * @return self     Chain structure.
         */
        close: function (navTab) {
            var self = this,
                $tabPane;
            var $navTab = self._getNavTab(navTab),
                $navTabLi = $navTab.parent('li');
            $tabPane = self._getTabPane($navTab);
            //close unsave tab confirm
            if ($navTabLi.length &&
                $tabPane.length &&
                $tabPane.hasClass('unsave') &&
                !self._unsaveConfirm()) {
                return self;
            }
            if ($navTabLi.find('a').hasClass("active")) {
                var $nextLi = $navTabLi.next("li:first"),
                    $prevLi = $navTabLi.prev("li:last");
                //if ($nextLi.size()) {
                if ($nextLi.length) {
                    self.active($nextLi);
                    self.activeMenu($nextLi.find('a'));
                //} else if ($prevLi.size()) {
                } else if ($prevLi.length) {
                    self.active($prevLi);
                    self.activeMenu($prevLi.find('a'));
                }
            }
            self._delStorage($navTab.attr('data-id')); //remove tab from session storage
            $navTabLi.remove();
            $tabPane.remove();
            return self;
        },

        /**
         * close others tab
         * @return self     Chain structure.
         */
        closeOthers: function (retainTab) {
            var self = this,
                $el = self.$element,
                findTab;
            
            if (!retainTab) {
                findTab = $el.navPanelList.find('li a:not([data-type="main"],.active)');
            } else {
                findTab = $el.navPanelList.find('a:not([data-type="main"])').filter(function(index){
                    if (retainTab != $(this).data('index')) return this;
                });
            }
          
            findTab.each(function () {
                var $navTab = $(this);
                self._delStorage($navTab.attr('data-id')); //remove tab from session storage
                self._getTabPane($navTab).remove(); //remove tab-content
                $navTab.parent('li').remove(); //remove navtab
            });
            if (retainTab) {
                self.active($el.navPanelList.find('a[data-index="' + retainTab + '"]'));
                self.activeMenu($el.navPanelList.find('a[data-index="' + retainTab + '"]'));
            }
            $el.navPanelList.css("margin-left", "0");
            return self;
        },

        /**
         * focus actived tab
         * @return self     Chain structure.
         */
        showActive: function () {
            var self = this,
                $el = self.$element;
            var navTab = $el.navPanelList.find('li a.active');
            self._fixTabPosition(navTab);
            return self;
        },

        /**
         * close all tabs, (except main tab)
         * @return self     Chain structure.
         */
        closeAll: function () {
            var self = this,
                $el = self.$element;
            $el.navPanelList.find('a:not([data-type="main"])').each(function () {
                var $navTab = $(this);
                self._delStorage($navTab.attr('data-id')); //remove tab from session storage
                self._getTabPane($navTab).remove(); //remove tab-content
                $navTab.parent('li').remove(); //remove navtab
            });
            self.active($el.navPanelList.find('a[data-type="main"]:first').parent('li'));
            self.activeMenu($el.navPanelList.find('a[data-type="main"]:first'));
            return self;
        },
        
        /**
         * 左侧导航变化
         */
        activeMenu: function(navTab) {
            // 点击选项卡时，左侧菜单栏跟随变化
            var $navObj       = $("a[href$='" + $(navTab).data('url') + "']"),   // 当前url对应的左侧导航对象
                $navHasSubnav = $navObj.parents('.nav-item'),
                $viSubHeight  = $navHasSubnav.siblings().find('.nav-subnav:visible').outerHeight();
            
            $('.nav-item').each(function(i){
                if ($(this).hasClass('active') && !$navObj.parents('.nav-item').last().hasClass('active')) {
                    $(this).removeClass('active').removeClass('open');
                    $(this).find('.nav-subnav:visible').slideUp(500);
                    if (window.innerWidth > 1024 && $('body').hasClass('lyear-layout-sidebar-close')) {
                        $(this).find('.nav-subnav').hide();
                    }
                }
            });
            
            $('.nav-drawer').find('li').removeClass('active');
            $navObj.parent('li').addClass('active');
            $navHasSubnav.first().addClass('active');
            
            // 当前菜单无子菜单
            if (!$navObj.parents('.nav-item').first().is('.nav-item-has-subnav')) {
                var hht = 48 * ( $navObj.parents('.nav-item').first().prevAll().length - 1 );
                $('.lyear-layout-sidebar-info').animate({scrollTop: hht}, 300);
            }
            
            if ($navObj.parents('ul.nav-subnav').last().is(':hidden')) {
                $navObj.parents('ul.nav-subnav').last().slideDown(500, function(){
                    $navHasSubnav.last().addClass('open');
		            var scrollHeight  = 0,
                        $scrollBox    = $('.lyear-layout-sidebar-info'),
		                pervTotal     = $navHasSubnav.last().prevAll().length,
		                boxHeight     = $scrollBox.outerHeight(),
	                    innerHeight   = $('.sidebar-main').outerHeight(),
                        thisScroll    = $scrollBox.scrollTop(),
                        thisSubHeight = $(this).outerHeight(),
                        footHeight    = 121;
		            
		            if (footHeight + innerHeight - boxHeight >= (pervTotal * 48)) {
		                scrollHeight = pervTotal * 48;
		            }
                    if ($navHasSubnav.length == 1) {
                        $scrollBox.animate({scrollTop: scrollHeight}, 300);
                    } else {
                        // 子菜单操作
                        if (typeof($viSubHeight) != 'undefined' && $viSubHeight != null) {
                            scrollHeight = thisScroll + thisSubHeight - $viSubHeight;
                            $scrollBox.animate({scrollTop: scrollHeight}, 300);
                        } else {
                            if ((thisScroll + boxHeight - $scrollBox[0].scrollHeight) == 0) {
                                scrollHeight = thisScroll - thisSubHeight;
                                $scrollBox.animate({scrollTop: scrollHeight}, 300);
                            }
                        }
                    }
                });
            }
        },

        /**
         * init function
         * @param options
         * @returns self
         * @private
         */
        _init: function (options) {
            var self = this,
                $el = self.$element;
            $el.html(defaultLayoutTemplates[options.nav.layout]
                .replace('{mainClass}', options.class)
                .replace('{navClass}', options.nav.class)
                .replace(/\{nav-tabs\}/g, options.nav.style)
                .replace(/\{backgroundColor\}/g, options.nav.backgroundColor)
                .replace('{dropdown}', options.language.nav.dropdown)
                .replace('{showActivedTab}', options.language.nav.showActivedTab)
                .replace('{closeAllTabs}', options.language.nav.closeAllTabs)
                .replace('{closeOtherTabs}', options.language.nav.closeOtherTabs)
            );
            $el.wrapper = $el.find('.mt-wrapper:first');
            $el.nav = $el.find('.mt-nav-bar:first');
            $el.navToolsLeft = $el.nav.find('.mt-nav-tools-left:first');
            $el.navPanel = $el.nav.find('.mt-nav-panel:first');
            $el.navPanelList = $el.nav.find('.mt-nav-panel:first ul');
            //$el.navTabMain    = $('#multitabs_main_0');
            $el.navToolsRight = $el.nav.find('.mt-nav-tools-right:first');
            $el.tabContent = $el.find('.tab-content:first');
            //hide tab-header if maxTabs less than 1
            if (options.nav.maxTabs <= 1) {
                options.nav.maxTabs = 1;
                $el.nav.hide();
            }
            //set the nav-panel width
            //var toolWidth = $el.nav.find('.mt-nav-tools-left:visible:first').width() + $el.nav.find('.mt-nav-tools-right:visible:first').width();
            $el.navPanel.css('width', 'calc(100% - 147px)');
            self.options = options;
            return self;
        },

        /**
         * final funcion for after init Multitabs
         * @returns self
         * @private
         */
        _final: function () {
            var self = this,
                $el = self.$element,
                options = self.options,
                storage, init = options.init,
                param;
            if (supportStorage(options.cache)) {
                storage = self._storage();
                self._resetStorage({});
                $.each(storage, function (k, v) {
                    self.create(v, false);
                })
            }
            if ($.isEmptyObject(storage)) {
                init = (!$.isEmptyObject(init) && init instanceof Array) ? init : defaultInit;
                for (var i = 0; i < init.length; i++) {
                    param = self._getParam(init[i]);
                    if (param) {
                        self.create(param);
                    }
                }
            }
            //if no any tab actived, active the main tab
            if (!$el.navPanelList.children('li a.active').length) {
                self.active($el.navPanelList.find('[data-type="main"]:first'));
            }
            return self;
        },

        /**
         * bind action
         * @return self
         * @private
         */
        _listen: function () {
            var self = this,
                $el = self.$element,
                options = self.options;
            //create tab
            handler($(document), 'click', options.selector, function () {
                self.create(this, true);
                if (!$(this).parent().parent('ul').hasClass('dropdown-menu')) {  // 20190402改，下拉菜单中的网址采用data-url，并且不阻止后面的动作
                    return false; //Prevent the default selector action
                }
            });
            //active tab
            handler($el.nav, 'click', '.mt-nav-tab', function () {
                self.active(this, false);
                self.activeMenu(this);
            });

            //drag tab
            if (options.nav.draggable) {
                handler($el.navPanelList, 'mousedown', '.mt-nav-tab', function (event) {
                    var $navTab = $(this),
                        $navTabLi = $navTab.closest('li');
                    var $prevNavTabLi = $navTabLi.prev();
                    var dragMode = true,
                        moved = false,
                        isMain = ($navTab.data('type') === "main");
                    var tmpId = 'mt_tmp_id_' + new Date().getTime(),
                        navTabBlankHtml = '<li id="' + tmpId + '" class="mt-dragging" style="width:' + $navTabLi.outerWidth() + 'px; height:' + $navTabLi.outerHeight() + 'px;"><a style="width: 100%;  height: 100%; "></a></li>';
                    var abs_x = event.pageX - $navTabLi.offset().left + $el.nav.offset().left;
                    $navTabLi.before(navTabBlankHtml);
                    $navTabLi.addClass('mt-dragging mt-dragging-tab').css({
                        'left': event.pageX - abs_x + 'px'
                    });

                    $(document).on('mousemove', function (event) {
                        if (dragMode && !isMain) {
                            $navTabLi.css({
                                'left': event.pageX - abs_x + 'px'
                            });
                            $el.navPanelList.children('li:not(".mt-dragging")').each(function () {
                                var leftWidth = $(this).offset().left + $(this).outerWidth() + 20; //20 px more for gap
                                if (leftWidth > $navTabLi.offset().left) {
                                    if ($(this).next().attr('id') !== tmpId) {
                                        moved = true;
                                        $prevNavTabLi = $(this);
                                        $('#' + tmpId).remove();
                                        $prevNavTabLi.after(navTabBlankHtml);
                                    }
                                    return false;
                                }
                            });
                        }
                    }).on("selectstart", function () { //disable text selection
                        if (dragMode) {
                            return false;
                        }
                    }).on('mouseup', function () {
                        if (dragMode) {
                            $navTabLi.removeClass('mt-dragging mt-dragging-tab').css({'left': 'auto'});
                            if (moved) {
                                $prevNavTabLi.after($navTabLi);
                            }
                            $('#' + tmpId).remove();
                        }
                        dragMode = false;
                    });
                });
            }
            
            // 右键菜单
            handler($el.nav, 'contextmenu', '.mt-nav-tab', function (event) {
                event.preventDefault();
                var menu     = $('<ul class="dropdown-menu" role="menu" id="contextify-menu"/>'),
                    $this    = $(this),
                    $nav     = $this.closest('li'),
                    $navTab  = self._getNavTab($nav),
                    $tabPane = self._getTabPane($navTab),
                    param    = $navTab.length ? self._getParam($navTab) : {};
                
                var menuData = [
                  {text: '刷新', onclick: function(){
                      var tempTabPane = $($tabPane);
                      
                      if (tempTabPane.is('iframe')) {
                          tempTabPane.attr('src', param.url);
                      } else {
                          $.ajax({
                              url: param.url,
                              dataType: "html",
                              success: function (callback) {
                                  tempTabPane.html(self.options.content.ajax.success(callback));
                              },
                              error: function (callback) {
                                  tempTabPane.html(self.options.content.ajax.error(callback));
                              }
                          });
                      }
                      menu.hide();
                      
                      return false;
                  }}
                ];
                
                var param = self._getParam($navTab);
                if (param.type !== 'main') {
                    menuData.push(
                        {text: '关闭', onclick: function(){
                            self.close($navTab);
                            menu.hide();
                            return false;
                        }}
                    );
                }
                
                menuData.push(
                    {text: '关闭其他', onclick: function(){
                       self.closeOthers($navTab.data('index'));
                       menu.hide();
                       return false;
                    }}
                );
                
                var l = menuData.length, i;
                
                for (i = 0; i < l; i++) {
                    var item = menuData[i],
                        el   = $('<li/>');
                    
                    el.append('<a class="dropdown-item" />');
                    
                    var a = el.find('a');
                    
                    a.on('click', item.onclick);
                    a.css('cursor', 'pointer');
                    a.html(item.text);
                    
                    menu.append(el);
                }
                
                var currentMenu = $("#contextify-menu");
                if (currentMenu.length > 0) {
                    if(currentMenu !== menu) {
                        currentMenu.replaceWith(menu);
                    }
                } else {
                    $('body').append(menu);
                }

                var clientTop = $(window).scrollTop() + event.clientY,
                    x = (menu.width() + event.clientX < $(window).width()) ? event.clientX : event.clientX - menu.width(),
                    y = (menu.height() + event.clientY < $(window).height()) ? clientTop : clientTop - menu.height();

                menu.css('top', y).css('left', x).css('position', 'fixed').show();
              
                
                $(this).parents().on('click', function () {
                    menu.hide();
                });
                $('#iframe-content').find('iframe').contents().find('body').on('click', function () {
                    menu.hide();
                });
            });
            
            // 双击事件
            handler($el.nav, 'dblclick', '.mt-nav-tab', function (event) {
                if (options.dbclickRefresh === true) {
                    var $this    = $(this),
                        $nav     = $this.closest('li'),
                        $navTab  = self._getNavTab($nav),
                        $tabPane = self._getTabPane($navTab),
                        param    = $navTab.length ? self._getParam($navTab) : {},
                        tempTabPane = $($tabPane);
                    
                    if (tempTabPane.is('iframe')) {
                        tempTabPane.attr('src', param.url);
                    } else {
                        $.ajax({
                            url: param.url,
                            dataType: "html",
                            success: function (callback) {
                                tempTabPane.html(self.options.content.ajax.success(callback));
                            },
                            error: function (callback) {
                                tempTabPane.html(self.options.content.ajax.error(callback));
                            }
                        });
                    }
                }
                
                return false;
            });

            //close tab
            handler($el.nav, 'click', '.mt-close-tab', function () {
                self.close($(this).closest('li'));
                return false; //Avoid possible BUG
            });
            //move left
            handler($el.nav, 'click', '.mt-move-left', function () {
                self.moveLeft();
                return false; //Avoid possible BUG
            });
            //move right
            handler($el.nav, 'click', '.mt-move-right', function () {
                self.moveRight();
                return false; //Avoid possible BUG
            });
            //show actived tab
            handler($el.nav, 'click', '.mt-show-actived-tab', function () {
                self.showActive();
                //return false; //Avoid possible BUG
            });
            //close all tabs
            handler($el.nav, 'click', '.mt-close-all-tabs', function () {
                self.closeAll();
                //return false; //Avoid possible BUG
            });
            //close other tabs
            handler($el.nav, 'click', '.mt-close-other-tabs', function () {
                self.closeOthers();
                //return false; //Avoid possible BUG
            });

            //fixed the nav-bar
            var navHeight = $el.nav.outerHeight();
            $el.tabContent.css('paddingTop', navHeight);
            if (options.nav.fixed) {
                handler($(window), 'scroll', function () {
                    var scrollTop = $(this).scrollTop();
                    scrollTop = scrollTop < ($el.wrapper.height() - navHeight) ? scrollTop + 'px' : 'auto';
                    $el.nav.css('top', scrollTop);
                    return false; //Avoid possible BUG
                });
            }
            //if layout === 'classic' show hide list in dropdown menu
            if (options.nav.layout === 'classic') {
                handler($el.nav, 'click', '.mt-dropdown:not(.open)', function () { //just trigger when dropdown not open.
                    var list = self._getHiddenList();
                    var $dropDown = $el.navToolsRight.find('.mt-hidden-list:first').empty();
                    if (list) { //when list is not empty
                        while (list.prevList.length) {
                            $dropDown.append(list.prevList.shift().clone());
                        }
                        while (list.nextList.length) {
                            $dropDown.append(list.nextList.shift().clone());
                        }
                    } else {
                        $dropDown.append('<li>empty</li>');
                    }
                    // return false; //Avoid possible BUG
                });
            }
            return self;
        },

        /**
         * get the multitabs object's param
         * @param obj          multitabs's object
         * @returns param      param
         * @private
         */
        _getParam: function (obj) {
            if ($.isEmptyObject(obj)) {
                return false;
            }
            var self = this,
                options = self.options,
                param = {},
                $obj = $(obj),
                data = $obj.data();

            //content
            param.content = data.content || obj.content || '';

            if (!param.content.length) {
                //url
                param.url = data.url || obj.url || $obj.attr('href') || $obj.attr('url') || '';
                param.url = $.trim(decodeURIComponent(param.url.replace('#', '')));
            } else {
                param.url = '';
            }
            if (!param.url.length && !param.content.length) {
                return false;
            }
            //isNewTab
            param.isNewTab = data.hasOwnProperty('isNewTab') || obj.hasOwnProperty('isNewTab') || options.isNewTab;
            //iframe
            param.iframe = data.iframe || obj.iframe || isExtUrl(param.url) || options.iframe;
            //type
            param.type = data.type || obj.type || options.type;
            //title
            param.title = data.title || obj.title || $obj.text() || param.url.replace('http://', '').replace('https://', '') || options.language.nav.title;
            param.title = trimText(param.title, options.nav.maxTitleLength);
            //active
            param.active = data.active || obj.active || false;
            //index
            param.index = data.index || obj.index || getTabIndex(param.type, options.nav.maxTabs);
            //id
            param.did = data.did || obj.did || this._generateId(param);
            return param;
        },

        /**
         * session storage for tab list
         * @param key
         * @param param
         * @returns storage
         * @private
         */
        _storage: function (key, param) {
            if (supportStorage(this.options.cache)) {
                var storage = JSON.parse(sessionStorage.multitabs || '{}');
                if (!key) {
                    return storage;
                }
                if (!param) {
                    return storage[key];
                }
                storage[key] = param;
                sessionStorage.multitabs = JSON.stringify(storage);
                return storage;
            }
            return {};
        },

        /**
         * delete storage by key
         * @param key
         * @private
         */
        _delStorage: function (key) {
            if (supportStorage(this.options.cache)) {
                var storage = JSON.parse(sessionStorage.multitabs || '{}');
                if (!key) {
                    return storage;
                }
                delete storage[key];
                sessionStorage.multitabs = JSON.stringify(storage);
                return storage;
            }
            return {};
        },

        /**
         * reset storage
         * @param storage
         * @private
         */
        _resetStorage: function (storage) {
            if (supportStorage(this.options.cache) && typeof storage === "object") {
                sessionStorage.multitabs = JSON.stringify(storage);
            }
        },

        /**
         * check if exist multitabs obj
         * @param param
         * @private
         */
        _exist: function (param) {
            if (!param || !param.url) {
                return false;
            }
            var self = this,
                $el = self.$element;
            var $navTab = $el.navPanelList.find('a[data-url="' + param.url + '"]:first');
            if ($navTab.length) {
                return $navTab;
            } else {
                return false;
            }
        },

        /**
         * get tab-pane from tab
         * @param tab
         * @returns {*}
         * @private
         */
        _getTabPane: function (navTab) {
            return $('#' + $(navTab).attr('data-id'));
        },

        /**
         * get real navTab in the panel list.
         * @param navTab
         * @returns navTab
         * @private
         */
        _getNavTab: function (navTab) {
            var self = this,
                $el = self.$element;
            var dataId = $(navTab).attr('data-id') || $(navTab).find('a').attr('data-id');
            return $el.navPanelList.find('a[data-id="' + dataId + '"]:first');
        },

        /**
         * fix nav navTab position
         * @param navTab
         * @private
         */
        _fixTabPosition: function (navTab) {
            var self = this,
                $el = self.$element,
                $navTabLi = $(navTab).parent('li'),
                tabWidth = $navTabLi.outerWidth(true),
                prevWidth = $navTabLi.prev().outerWidth(true),
                pprevWidth = $navTabLi.prev().prev().outerWidth(true),
                sumPrevWidth = sumDomWidth($navTabLi.prevAll()),
                sumNextWidth = sumDomWidth($navTabLi.nextAll()),
                navPanelWidth = $el.navPanel.outerWidth(true),
                sumTabsWidth = sumDomWidth($el.navPanelList.children('li')),
                leftWidth = 0;
            //all nav navTab's width no more than nav-panel's width
            if (sumTabsWidth < navPanelWidth) {
                leftWidth = 0
            } else {
                //when navTab and his right tabs sum width less or same as nav-panel, it means nav-panel can contain the navTab and his right tabs
                if ((prevWidth + tabWidth + sumNextWidth) <= navPanelWidth) {
                    leftWidth = sumPrevWidth; //sum width of left part
                    //add width from the left, calcular the maximum tabs can contained by nav-panel
                    while ((sumTabsWidth - leftWidth + prevWidth) < navPanelWidth) {
                        $navTabLi = $navTabLi.prev(); //change the left navTab
                        leftWidth -= $navTabLi.outerWidth(); //reduce the left part width
                    }
                } else { //nav-panel can not contain the navTab and his right tabs
                    //when the navTab and his left part tabs's sum width more than nav-panel, all the width of 2 previous tabs's width set as the nav-panel margin-left.
                    if ((sumPrevWidth + tabWidth) > navPanelWidth) {
                        leftWidth = sumPrevWidth - prevWidth - pprevWidth
                    }
                }
            }
            leftWidth = leftWidth > 0 ? leftWidth : 0; //avoid leftWidth < 0 BUG
            $el.navPanelList.animate({
                marginLeft: 0 - leftWidth + "px"
            }, "fast");
        },

        /**
         * hidden tab list
         * @returns hidden tab list, the prevList and nextList
         * @private
         */
        _getHiddenList: function () {
            var self = this,
                $el = self.$element,
                navPanelListMarginLeft = Math.abs(parseInt($el.navPanelList.css("margin-left"))),
                navPanelWidth = $el.navPanel.outerWidth(true),
                sumTabsWidth = sumDomWidth($el.navPanelList.children('li')),
                tabPrevList = [],
                tabNextList = [],
                $navTabLi, marginLeft;
            //all tab's width no more than nav-panel's width
            if (sumTabsWidth < navPanelWidth) {
                return false;
            } else {
                $navTabLi = $el.navPanelList.children('li:first');
                //overflow hidden left part
                marginLeft = 0;
                //from the first tab, add all left part hidden tabs
                while ((marginLeft + $navTabLi.width()) <= navPanelListMarginLeft) {
                    marginLeft += $navTabLi.outerWidth(true);
                    tabPrevList.push($navTabLi);
                    $navTabLi = $navTabLi.next();
                }
                //overflow hidden right part
                if (sumTabsWidth > marginLeft) { //check if the right part have hidden tabs
                    $navTabLi = $el.navPanelList.children('li:last');
                    marginLeft = sumTabsWidth; //set margin-left as the Rightmost, and reduce one and one.
                    while (marginLeft > (navPanelListMarginLeft + navPanelWidth)) {
                        marginLeft -= $navTabLi.outerWidth(true);
                        tabNextList.unshift($navTabLi); //add param from top
                        $navTabLi = $navTabLi.prev();
                    }
                }
                return {
                    prevList: tabPrevList,
                    nextList: tabNextList
                };
            }
        },



        /**
         * check if tab-pane is iframe, and add/remove class
         * @param tabPane
         * @private
         */
        _fixTabContentLayout: function (tabPane) {
            var $tabPane = $(tabPane);
            if ($tabPane.is('iframe')) {
                $('body').addClass('full-height-layout');
                /** fix chrome croll disappear bug **/
                $tabPane.css("height", "99%");
                window.setTimeout(function () {
                    $tabPane.css("height", "100%");
                }, 0);
            } else {
                $('body').removeClass('full-height-layout');
            }
        },
    };

    /**
     * Entry function
     * @param option
     */
    $.fn.multitabs = function (option, id) {
        var self = $(this),
            did = id ? id : 'multitabs',
            multitabs = $(document).data(did),
            options = typeof option === 'object' && option,
            opts;
        if (!multitabs) {
            opts = $.extend(true, {}, $.fn.multitabs.defaults, options, self.data());
            opts.nav.style = (opts.nav.style === 'nav-pills') ? 'nav-pills' : 'nav-tabs';
            multitabs = new MultiTabs(this, opts);
            $(document).data(did, multitabs);
        }
        return $(document).data(did);
    };

    /**
     * Default Options
     * @type {}
     */
    $.fn.multitabs.defaults = {
        selector: '.multitabs', // selector text to trigger multitabs.
        iframe: false, // Global iframe mode, default is false, is the auto mode (for the self page, use ajax, and the external, use iframe)
        cache: false, // 是否缓存当前打开的tab
        class: '', // class for whole multitabs
        type: 'info', // change the info content name, is not necessary to change.
        init: [],
        isNewTab: false, // 是否以新tab标签打开，为true时，每次点击都会打开新的tab
        refresh: 'no', // iframe中页面是否刷新，'no'：'从不刷新'，'nav'：'点击菜单刷新'，'all'：'菜单和tab点击都刷新'
        dbclickRefresh: false, // 双击刷新开启最好不要和refresh:'all'同时使用
        nav: {
            backgroundColor: '#f5f5f5', //default nav-bar background color
            class: '', //class of nav
            draggable: true, //nav tab draggable option
            fixed: false, //fixed the nav-bar
            layout: 'default', //it can be 'default', 'classic' (all hidden tab in dropdown list), and simple
            maxTabs: 15, //Max tabs number (without counting main tab), when is 1, hide the whole nav
            maxTitleLength: 25, //Max title length of tab
            showCloseOnHover: false, //while is true, show close button in hover, if false, show close button always
            style: 'nav-tabs' //can be nav-tabs or nav-pills
        },
        content: {
            ajax: {
                class: '', //Class for ajax tab-pane
                error: function (htmlCallBack) {
                    //modify html and return
                    return htmlCallBack;
                },
                success: function (htmlCallBack) {
                    //modify html and return
                    return htmlCallBack;
                }
            },
            iframe: {
                class: ''
            }
        },
        language: { //language setting
            nav: {
                title: 'Tab', //default tab's tittle
                dropdown: '<i class="mdi mdi-menu"></i>', //right tools dropdown name
                showActivedTab: '显示当前选项卡', //show active tab
                closeAllTabs: '关闭所有标签页', //close all tabs
                closeOtherTabs: '关闭其他标签页', //close other tabs
            }
        }
    };
})(jQuery));