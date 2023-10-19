window.addEventListener('touchmove', function (e) {
    if (document.body.classList.contains('mu')) {
        e.preventDefault();
    }
});

document.body.addEventListener('click', function (e) {
    let target = e.target;
    let tag = target.className || '';
    let rel = target.getAttribute('rel') || '';
    let audio;
    switch (true) {
        // 点击菜单按钮
        case (tag.indexOf('switchmenu') !== -1):
            window.scrollTo(0, 0);
            document.documentElement.classList.toggle('mu');
            document.body.classList.toggle('mu');
            break;
        //加载更多
        case (tag.indexOf('more') !== -1):
            var moreButton = document.querySelector('.more');
            // 如果已经在加载了
            if (moreButton.dataset.status === 'loading') {
                return;
            }

            var num = parseInt(moreButton.dataset.page) || 1;

            if (num === 1) {
                moreButton.dataset.page = 1;
            }

            if (num >= Pages) {
                return;
            }

            // ajax加载第二页
            moreButton.innerHTML = '加载中..';
            moreButton.dataset.status = 'loading';
            Diaspora.SWLoader();

            Diaspora.L(moreButton.href, function (data) {
                var link = $(data).find('.more').attr('href');
                if (link !== undefined) {
                    moreButton.href = link;
                    moreButton.innerHTML = '加载更多';
                    moreButton.dataset.status = 'loaded';
                    moreButton.dataset.page = parseInt(moreButton.dataset.page) + 1;
                } else {
                    var pager = document.getElementById('pager');
                    if (pager) {
                        pager.remove();
                    }
                }

                var tempScrollTop = window.scrollY || document.documentElement.scrollTop;
                document.getElementById('primary').insertAdjacentHTML('beforeend', $(data).find('.post')[0].outerHTML);
                window.scrollTo(0, tempScrollTop);
                Diaspora.SWLoader();
                window.scrollTo(0, tempScrollTop + 400);
            }, function () {
                moreButton.innerHTML = '加载更多';
                moreButton.dataset.status = 'loaded';
            });
            break;

        // audio play
        case (tag.indexOf('icon-play') !== -1):
            audio = document.querySelector("audio");
            audio.play().then(function (){
                let iconPlay = document.querySelector('.icon-play');
                iconPlay.classList.remove('icon-play');
                iconPlay.classList.add('icon-pause');
            }).catch(function (err){
                console.log(err)
            });
            break;
        // audio pause
        case (tag.indexOf('icon-pause') !== -1):
            audio = document.querySelector("audio");
            audio.pause()
            var iconPause = document.querySelector('.icon-pause');
            iconPause.classList.remove('icon-pause');
            iconPause.classList.add('icon-play');
            break;
        //点击独立页面时
        case (tag.indexOf('pagelist') !== -1):
            //$('body').removeClass('mu')
            document.body.classList.remove('mu');

            setTimeout(function () {
                Diaspora.HS($(e.target), 'push')
            }, 300)
            break;
        // 其他情况的处理代码...
        // history state
        case (tag.indexOf('cover') !== -1):
            Diaspora.HS($(e.target).parent(), 'push')
            break;

        // history state
        case (tag.indexOf('posttitle') !== -1):
            e.preventDefault();
            Diaspora.AJAX(e.target.getAttribute('href'))
            break;
        default:
            return
    }
});




