/*
 * Diaspora
 * @author LoeiFy
 * @url http://lorem.in
 */
let Home = location.href,
    Pages = LocalConst.MAX_PAGES,
    xhr,
    xhrUrl = '',
    scrollFunction,
    newCommentIndex = 1;

if(document.getElementById("mark")){var parallax=new Parallax(document.getElementById("mark"))}//Parallax Plugin js

console.time("1");
let Diaspora = {

    L: function (url, f, err) {
        if (url == xhrUrl) {
            return false
        }

        xhrUrl = url;

        if (xhr) {
            xhr.abort()
        }


        xhr = $.ajax({
            type: 'GET',
            url: url,
            timeout: 10000,
            success: function (data) {
                f(data)
                xhrUrl = '';
            },
            error: function (a, b, c) {
                if (b == 'abort') {
                    err && err()
                } else {
                    window.location.href = url
                }
                xhrUrl = '';
            }
        })
    },

    P: function () {
        return !!('ontouchstart' in window);
    },

    PS: function () {
        if (!(window.history && history.pushState)) return;

        history.replaceState({ u: Home, t: document.title }, document.title, Home)

        window.addEventListener('popstate', function (e) {
            let state = e.state;

            if (!state) return;

            document.title = state.t;

            if (state.u == Home) {
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
                Diaspora.loading()

                Diaspora.L(state.u, function (data) {

                    document.title = state.t;

                    $('#preview').html($(data).filter('#single'));

                    Diaspora.preview();

                    setTimeout(function () {
                        Diaspora.player(state.d);
                    }, 0);
                })
            }

        })
    },

    HS: function (tag, flag) {
        let id = tag.data('id') || 0,
            url = tag.attr('href'),
            title = tag.attr('title') || tag.text();

        if (!$('#preview').length || !(window.history && history.pushState)) location.href = url;

        Diaspora.loading()

        let state = { d: id, t: title, u: url };

        Diaspora.L(url, function (data) {

            if (!$(data).filter('#single').length) {
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

            $('#preview').html($(data).filter('#single'))

            switch (flag) {

                case 'push':
                    Diaspora.preview()
                    break;

                case 'replace':
                    window.scrollTo(0, 0)
                    Diaspora.loaded()
                    break;
            }

            setTimeout(function () {
                if (!id) id = $('.icon-play').data('id');
                Diaspora.player(id)

                // get download link
                $('.content img').each(function () {
                    if ($(this).attr('src').indexOf('/uploads/2014/downloading.png') > -1) {
                        $(this).hide()
                        $('.downloadlink').attr('href', $(this).parent().attr('href'))
                    }
                })

                if (flag == 'replace') {
                    $('#top').show()
                }
            }, 0)

        })
    },

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

                    Diaspora.loaded()
                }, 500)
            }, 300)
        }, 0)
    },

    player: function (id) {
        let p = $('#audio-' + id + '-1');

        if (!p.length) {
            $('.icon-play').css({
                'color': '#dedede',
                'cursor': 'not-allowed'
            })
            return
        }

        if (p.eq(0).data("autoplay") == false) {
            p[0].play();
        }

        p.on({
            'timeupdate': function () {
                $('.bar').css('width', p[0].currentTime / p[0].duration * 100 + '%')
            },
            'ended': function () {
                $('.icon-pause').removeClass('icon-pause').addClass('icon-play')
            },
            'playing': function () {
                $('.icon-play').removeClass('icon-play').addClass('icon-pause')
            }
        })
    },

    loading: function () {
        let w = window.innerWidth;
        let css = '<style class="loaderstyle" id="loaderstyle' + w + '">' +
            '@-moz-keyframes loader' + w + '{100%{background-position:' + w + 'px 0}}' +
            '@-webkit-keyframes loader' + w + '{100%{background-position:' + w + 'px 0}}' +
            '.loader' + w + '{-webkit-animation:loader' + w + ' 3s linear infinite;-moz-animation:loader' + w + ' 3s linear infinite;}' +
            '</style>';
        $('.loaderstyle').remove()
        $('head').append(css)

        $('#loader').removeClass().addClass('loader' + w).show()
    },

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
console.timeEnd("1");


console.time("2");
$(function () {

    if (Diaspora.P()) {
        $('body').addClass('touch')
    }

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

            setTimeout(function () {
                $('html, body').removeClass('loading')
            }, 100)



            let vibrant = new Vibrant(cover.t[0]);
            let swatches = vibrant.swatches()

            if (swatches['DarkVibrant']) {
                $('#vibrant polygon').css('fill', swatches['DarkVibrant'].getHex())
                $('#vibrant div').css('background-color', swatches['DarkVibrant'].getHex())
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

        $('.pview a').addClass('pviewa')

        let T;
        $(window).on('resize', function () {
            clearTimeout(T)

            T = setTimeout(function () {
                if (!Diaspora.P() && location.href == Home) {
                    cover.o()
                    cover.f()
                }

                if ($('#loader').attr('class')) {
                    Diaspora.loading()
                }
            }, 500)
        })

    } else {

        $('#single').css('min-height', window.innerHeight)
        $('html, body').removeClass('loading')

        window.addEventListener('popstate', function (e) {
            if (e.state) location.href = e.state.u;
        })

        Diaspora.player($('.icon-play').data('id'))

        $('.icon-icon, .image-icon').attr('href', '/')


        $('#top').show()

    }

    $(window).on('scroll', function () {
        if ($('.scrollbar').length && !Diaspora.P() && !$('.icon-images').hasClass('active')) {
            let st = $(window).scrollTop(),
                ct = $('.content').height();

            if (st > ct) {
                st = ct
            }

            $('.scrollbar').width((50 + st) / ct * 100 + '%')

            if (st > 80 && window.innerWidth > 800) {
                $('.subtitle').fadeIn()
            } else {
                $('.subtitle').fadeOut()
            }
        }
    })

    // $(window).on('scroll', scrollFunction = function (e) {
    //     if ($('#single').length) {
    //         if ((window.innerHeight + window.scrollY) >= document.body.offsetHeight) {
    //             // Diaspora.loadDisqus ();
    //         }
    //     }
    // })


    $(window).on('touchmove', function (e) {
        if ($('body').hasClass('mu')) {
            e.preventDefault()
        }
    })


    $('body').on('click', function (e) {

        let tag = $(e.target).attr('class') || '',
            rel = $(e.target).attr('rel') || '';

        if (!tag && !rel) return;

        switch (true) {

            // nav menu
            case (tag.indexOf('switchmenu') != -1):
                window.scrollTo(0, 0)
                $('html, body').toggleClass('mu')
                break;

            // next page
            case (tag.indexOf('more') != -1):
                tag = $('.more');

                if (tag.data('status') == 'loading') {
                    return false
                }

                let num = parseInt(tag.data('page')) || 1;

                if (num == 1) {
                    tag.data('page', 1)
                }

                if (num >= Pages) {
                    return
                }

                tag.html('加载中..').data('status', 'loading')
                Diaspora.loading()

                Diaspora.L(tag.attr('href'), function (data) {
                    let link = $(data).find('.more').attr('href');
                    if (link != undefined) {
                        tag.attr('href', link).html('加载更多').data('status', 'loaded')
                        tag.data('page', parseInt(tag.data('page')) + 1)
                    } else {
                        $('#pager').remove()
                    }

                    let tempScrollTop = $(window).scrollTop();
                    $('#primary').append($(data).find('.post'))
                    $(window).scrollTop(tempScrollTop);
                    Diaspora.loaded()
                    $('html,body').animate({ scrollTop: tempScrollTop + 400 }, 500);
                }, function () {
                    tag.html('加载更多').data('status', 'loaded')
                })

                return false;
                break;

            // qrcode这里要改成vue的
            case (tag.indexOf('icon-wechat') != -1):
                var qrcode = new QRCode(document.getElementById("qr"), {
                    width: 80,
                    height: 80,
                    useSVG: true,
                    text: window.location.href
                });
                $('#qr').toggle()
                break;

            // audio play
            case (tag.indexOf('icon-play') != -1):
                $('#audio-' + $('.icon-play').data('id') + '-1')[0].play()
                $('.icon-play').removeClass('icon-play').addClass('icon-pause')
                break;

            // audio pause
            case (tag.indexOf('icon-pause') != -1):
                $('#audio-' + $('.icon-pause').data('id') + '-1')[0].pause()
                $('.icon-pause').removeClass('icon-pause').addClass('icon-play')
                break;

            // post like
            case (tag.indexOf('icon-like') != -1):
                let t = $(e.target).parent(),
                    classes = t.attr('class'),
                    id = t.attr('id').split('like-');

                if (t.prev().hasClass('icon-view')) return;
                //if ((t.prev()).classList.contains('icon-view')) return;

                classes = classes.split(' ');
                if (classes[1] == 'active') return;

                fetch(window.location.origin + '/action/like', {
                    "headers": {"content-type": "application/x-www-form-urlencoded; charset=UTF-8"},
                    "body": "cid="+id[1],
                    "method": "POST",
                })
                    .then(function(response) {
                        return response.text();
                    })
                    .then(function(text) {
                        console.log('Request successful', text);
                        let text1 = $('#like-' + id[1]).html();
                        localStorage.setItem('isLike-' + id[1], 1);
                        t.addClass('active');
                        $('.count').html(JSON.parse(text).data.count);
                    })
                    .catch(function(error) {
                        console.log('Request failed', error)
                    });;
                // $.ajax({
                //     type: 'POST',
                //     url: window.location.origin + '/action/like',
                //     data: {
                //         cid: id[1]
                //     },
                //     dataType: 'json',
                //     success: function (ret) {
                //         let text = $('#like-' + id[1]).html();

                //         localStorage.setItem('isLike-' + id[1], 1);
                //         t.addClass('active');
                //         $('.count').html(ret.data.count);
                //     }
                // })
                break;

            // history state
            case (tag.indexOf('cover') != -1):
                Diaspora.HS($(e.target).parent(), 'push')
                return false;
                break;

            // history state
            case (tag.indexOf('posttitle') != -1):
                Diaspora.HS($(e.target), 'push')
                return false;
                break;

            // prev, next post
            case (rel == 'prev' || rel == 'next'):
                if (rel == 'prev') {
                    let t = $('#prev_next a')[0].text
                } else {
                    let t = $('#prev_next a')[1].text
                }
                $(e.target).attr('title', t)

                Diaspora.HS($(e.target), 'replace')
                return false;
                break;

            // quick view
            case (tag.indexOf('pviewa') != -1):
                //$('body').removeClass('mu')
                document.body.classList.remove('mu');

                setTimeout(function () {
                    Diaspora.HS($(e.target), 'push')
                }, 300)

                return false;
                break;

            default:
                return;
                break;
        }

    });
    // console.log("%c Github %c","background:#24272A; color:#ffffff","","https://github.com/LoeiFy/Diaspora")
    console.log("\n %c Diaspora For Typecho %c Jin < https://jcl.moe/ > \n", "color:rgb(255, 242, 242);background:rgb(244, 164, 164);padding:5px 0;border-radius:3px 0 0 3px;", "color:rgb(244, 164, 164);background:rgb(255, 242, 242);padding:5px 0;border-radius:0 3px 3px 0;")

});
