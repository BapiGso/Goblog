/*
 * Diaspora
 * @author LoeiFy
 * @url http://lorem.in
 */
let Home = location.href,
    Pages = 5,
    xhr,
    xhrUrl = '';
    newCommentIndex = 1;


let Diaspora = {

    //ajax请求封装
    L: function (url, f, err) {
        debugger
        if (url === xhrUrl) {
            return false
        }
        xhrUrl = url;

        if (xhr) {
            xhr.abort()
        }
        xhr = new XMLHttpRequest();
        xhr.open('GET', url, true);
        xhr.timeout = 10000;

        xhr.onload = function () {
            if (xhr.status >= 200 && xhr.status < 300) {
                f(xhr.responseText);
            } else {
                console.log(xhr.statusText, xhr.status);
            }
        };

        xhrUrl = '';
        xhr.send();
        // xhr = $.ajax({
        //     type: 'GET',
        //     url: url,
        //     timeout: 10000,
        //     success: function (data) {
        //         f(data)
        //         xhrUrl = '';
        //     },
        //     error: function (a, b, c) {
        //         if (b === 'abort') {
        //             err && err()
        //         } else {
        //             window.location.href = url
        //         }
        //         xhrUrl = '';
        //     }
        // })
    },

    SWLoader:function (){
      if(document.querySelector("#loader")){
          document.querySelector("#loader").remove()
      }else{
          let l=document.createElement("div")
          l.id="loader"
          document.body.insertAdjacentElement("afterbegin",l)
      }
    },

    SWPreview:function (){
        let l=document.createElement("div")
        l.id="preview"
        document.body.insertAdjacentElement("afterbegin",l)
    },

    AJAX:function (url){
        this.SWLoader()
        fetch(url)
            .then(function(response) {
                return response.text(); // 将响应转换为文本
            })
            .then(function(data) {
                let l=document.createElement("div")
                l.id="preview"
                l.innerHTML=data
                document.body.insertAdjacentElement("afterbegin",l)
                setTimeout(function(){document.querySelector("#container").style.display="none"}, 500);
            })
            .catch(function(error) {
                // 处理错误
                console.error('请求失败', error);
            });
        this.SWLoader()
    },

    PS: function () {
        if (!(window.history && history.pushState)) return;

        history.replaceState({ u: Home, t: document.title }, document.title, Home)

        window.addEventListener('popstate', function (e) {
            let state = e.state;

            if (!state) return;

            document.title = state.t;

            if (state.u === Home) {
                $('#preview').css('position', 'fixed')
                setTimeout(function () {
                    $('#preview').removeClass('show').addClass('trans')
                    $('#container').show()
                    window.scrollTo(0, parseInt($('#container').data('scroll')))
                    setTimeout(function () {
                        $('#preview').html('')
                        $(window).trigger('resize')
                    }, 300)
                }, 0)
            } else {
                Diaspora.SWLoader()

                Diaspora.L(state.u, function (data) {

                    document.title = state.t;

                    $('#preview').html($(data).filter('body'));

                    Diaspora.preview();

                    setTimeout(function () {
                        Diaspora.player(state.d);
                    }, 0);
                })
            }

        })
    },

    //传进来a标签和自定义flag
    HS: function (tag, flag) {
        debugger
        let id = tag.data('id') || 0,
            url = tag.attr('href'),
            title = tag.attr('title') || tag.text();

        if (!$('#preview').length || !(window.history && history.pushState)) location.href = url;

        Diaspora.SWLoader()

        let state = { d: id, t: title, u: url };
        debugger
        Diaspora.L(url, function (data) {
            if (!$(data).filter().length) {
                location.href = url;
                return
            }

            switch (flag) {

                case 'push':
                    history.pushState(state, title, url)
                    break;

                case 'replace':
                    history.replaceState(state, title, url)
                    break;

            }
            document.title = title;

            $('#preview').html($(data).filter())

            switch (flag) {
                case 'push':
                    Diaspora.preview()
                    break;
                case 'replace':
                    window.scrollTo(0, 0)
                    Diaspora.SWLoader()
                    break;
            }

            setTimeout(function () {
                if (!id) id = $('.icon-play').data('id');
                Diaspora.player(id)
                // get download link
                //给文章中的图片添加占位图用的
                $('.content img').each(function () {
                    if ($(this).attr('src').indexOf('/uploads/2014/downloading.png') > -1) {
                        $(this).hide()
                        $('.downloadlink').attr('href', $(this).parent().attr('href'))
                    }
                })

                if (flag === 'replace') {
                    $('#top').show()
                }
            }, 100)

        })
    },

    //只给preview添加动画用的，里面的数据是在这个方法之前调用
    preview: function () {
        setTimeout(function () {
            $('#preview').addClass('show')
            $('#container').data('scroll', window.scrollY)
            setTimeout(function () {
                $('#container').hide()
                setTimeout(function () {
                    $('#preview').css({
                        'position': 'static',
                        'overflow-y': 'auto'
                    }).removeClass('trans')
                    $('#top').show()

                    Diaspora.SWLoader()
                }, 500)
            }, 300)
        }, 0)
    },

    player: function (id) {
        let au = document.querySelector('#audio')
        let playButton = document.querySelector('.icon-play');
        let playbackBar = document.querySelector('.playbackbar');

        // 检查
        // if (!au.length()) {
        //     //todo 颜色变暗
        //     playButton.style.cursor = 'not-allowed';
        //     return
        // }
        //自动播放，暂时关闭
        // if (p.eq(0).data("autoplay") === false) {
        //     p[0].play();
        // }
        au.addEventListener('timeupdate', function() {
            let progress = (au.currentTime / au.duration) * 100;
            playbackBar.style.width = progress + '%';
        });

        au.addEventListener('ended', function() {
            playButton.classList.remove('icon-pause');
            playButton.classList.add('icon-play');
        });

        au.addEventListener('playing', function() {
            playButton.classList.remove('icon-play');
            playButton.classList.add('icon-pause');
        });
    },

    //给loader添加动画用的
    loading: function () {
        $('#loader').removeClass().addClass('loader890').show()
    },

    //给loader移除动画用的
    loaded: function () {
        $('#loader').removeClass().hide()
    },

    F: function (id, w, h) {
        let _height = $(id).parent().height(),
            _width = $(id).parent().width(),
            ratio = h / w;

        if (_height / _width > ratio) {
            id.style.height = _height + 'px';
            id.style.width = _height / ratio + 'px';
        } else {
            id.style.width = _width + 'px';
            id.style.height = _width * ratio + 'px';
        }

        id.style.left = (_width - parseInt(id.style.width)) / 2 + 'px';
        id.style.top = (_height - parseInt(id.style.height)) / 2 + 'px';
    },
}

function old () {

    if (('ontouchstart' in window)) {
        $('body').addClass('touch')
    }
    debugger
    if ($('#preview').length) {

        let cover = {};
        cover.t = $('#cover');
        cover.w = cover.t.attr('width');
        cover.h = cover.t.attr('height');

        ; (cover.o = function () {
            $('#mark').height(window.innerHeight)
        })();

        if (cover.t.prop('complete')) {
            // why setTimeout ?
            setTimeout(function () { cover.t.load() }, 0)
        }

        cover.t.on('load', function () {

            ; (cover.f = function () {

                let _w = $('#mark').width(), _h = $('#mark').height(), x, y, i, e;

                e = (_w >= 1000 || _h >= 1000) ? 1000 : 500;

                if (_w >= _h) {
                    i = _w / e * 50;
                    y = i;
                    x = i * _w / _h;
                } else {
                    i = _h / e * 50;
                    x = i;
                    y = i * _h / _w;
                }

                $('.layer').css({
                    'width': _w + x,
                    'height': _h + y,
                    'marginLeft': - 0.5 * x,
                    'marginTop': - 0.5 * y
                })

                if (!cover.w) {
                    cover.w = cover.t.width();
                    cover.h = cover.t.height();
                }

                Diaspora.F($('#cover')[0], cover.w, cover.h)

            })();


            let vibrant = new Vibrant(cover.t[0]);
            let swatches = vibrant.swatches()

            if (swatches['DarkVibrant']) {
                $('#vibrant polygon').css('fill', swatches['DarkVibrant'].getHex())
            }
            if (swatches['Vibrant']) {
                $('.icon-menu').css('color', swatches['Vibrant'].getHex())
            }

        })

        if (!cover.t.attr('src')) {
            alert('Please set the post thumbnail')
        }

        $('#preview').css('min-height', window.innerHeight)

        Diaspora.PS()


        let T;
        $(window).on('resize', function () {
            clearTimeout(T)

            T = setTimeout(function () {
                if (!('ontouchstart' in window) && location.href == Home) {
                    cover.o()
                    cover.f()
                }

                if ($('#loader').attr('class')) {
                    Diaspora.SWLoader()
                }
            }, 500)
        })

    } else {

        $('body').css('min-height', window.innerHeight)

        window.addEventListener('popstate', function (e) {
            if (e.state) location.href = e.state.u;
        })

        Diaspora.player($('.icon-play').data('id'))

        $('.icon-icon, .image-icon').attr('href', '/')
    }



    $('body').on('click', function (e) {

        let tag = $(e.target).attr('class') || '',
            rel = $(e.target).attr('rel') || '';

        if (!tag && !rel) return;

        switch (true) {

            // history state
            case (tag.indexOf('cover') !== -1):
                Diaspora.HS($(e.target).parent(), 'push')
                return false;
                break;

            // history state
            case (tag.indexOf('posttitle') !== -1):
                Diaspora.HS($(e.target), 'push')
                return false;
                break;

            // prev, next post
            case (rel === 'prev' || rel === 'next'):
                if (rel === 'prev') {
                    let t = $('#prev_next a')[0].text
                } else {
                    let t = $('#prev_next a')[1].text
                }
                $(e.target).attr('title', t)

                Diaspora.HS($(e.target), 'replace')
                return false;
                break;

            default:
                return;
        }
    });
}