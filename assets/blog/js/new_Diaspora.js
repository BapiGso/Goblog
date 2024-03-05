window.addEventListener('touchmove', function (e) {
    if (document.body.classList.contains('mu')) {
        e.preventDefault();
    }
});


// window.addEventListener('click', function (e) {
//     let target = e.target;
//     let tag = target.className || '';
//     let rel = target.getAttribute('rel') || '';
//     let audio;
//     switch (true) {
        // 点击菜单按钮
        // case (tag.indexOf('switchmenu') !== -1):
        //     window.scrollTo(0, 0);
        //     document.documentElement.classList.toggle('mu');
        //     break;
        //加载更多
        // case (tag.indexOf('more') !== -1):
        //     e.preventDefault();
        //     let moreButton = document.querySelector('.more');
        //     // 如果已经在加载了
        //     if (moreButton.dataset.status === 'loading') {
        //         return;
        //     }
        //
        //     // ajax加载第二页
        //     moreButton.innerHTML = '加载中..';
        //     moreButton.dataset.status = 'loading';
        //     Diaspora.SWLoader();
        //
        //     Diaspora.AJAX(moreButton.href, function (data) {
        //         let tmpScrollTop = window.scrollY || document.documentElement.scrollTop;
        //         document.getElementById('pager').remove()
        //         document.getElementById('primary').insertAdjacentHTML('beforeend', data);
        //     }, function () {
        //         moreButton.innerHTML = '加载更多';
        //         moreButton.dataset.status = 'loaded';
        //     });
        //     Diaspora.SWLoader();
        //     break;

        // audio play
        // case (tag.indexOf('icon-play') !== -1):
        //     audio = document.querySelector("audio");
        //     audio.play().then(function (){
        //         let iconPlay = document.querySelector('.icon-play');
        //     }).catch(function (err){
        //         console.log(err)
        //     });
        //     break;
        // // audio pause
        // case (tag.indexOf('icon-pause') !== -1):
        //     audio = document.querySelector("audio");
        //     audio.pause()
        //     const iconPause = document.querySelector('.icon-pause');
        //     iconPause.classList.remove('icon-pause');
        //     iconPause.classList.add('icon-play');
        //     break;
        //点击独立页面时
        // case (tag.indexOf('pagelist') !== -1):
        //     e.preventDefault();
        //     document.documentElement.classList.remove('mu');
        //     Diaspora.SingleLoader(e.target,'push')
        //     break;
        // history state
        // case (tag.indexOf('cover') !== -1):
        //     e.preventDefault();
        //     debugger
        //     Diaspora.SingleLoader(e.target.parentElement,'push')
        //     break;

        // history state
        // case (tag.indexOf('posttitle') !== -1):
        //     e.preventDefault();
        //     Diaspora.SingleLoader(e.target,'push')
        //     debugger
        //     break;
//         default:
//             return
//     }
// });

// 在document上添加滚动事件监听器
// window.addEventListener('scroll',Diaspora.scroller);
// preview.addEventListener('scroll',Diaspora.scroller);



//todo 页面加载完后检查localstorage
