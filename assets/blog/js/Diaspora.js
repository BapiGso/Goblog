/*
 * Diaspora
 * @author LoeiFy
 * @url http://lorem.in
 */
let Home = location.href,
    preview = Object.assign(document.createElement("div"), { id: "preview" });


let Diaspora = {
    SWLoader:function (){
      if(document.getElementById("loader")){
          document.getElementById("loader").remove()
      }else{
          let l=Object.assign(document.createElement("div"), { id: "loader" });
          document.body.insertAdjacentElement("afterbegin",l)
      }
    },

    // 传ajax的url和响应成功后需要执行的函数
    AJAX:async function (url,func,errfunc){
        fetch(url)
            .then(function(response) {
                return response.text(); // 将响应转换为文本
            })
            .then(function(data) {
                    func(data)
            })
            .catch(function(error) {
                errfunc()
                // 处理错误
                console.error('请求失败', error);
                debugger;
                window.location.href = url;
            });
    },

    //传进来元素和flag状态
    SingleLoader:function (tag,flag) {
        let id = tag.getAttribute('data-id') || 0,
            url = tag.getAttribute('href'),
            title = tag.getAttribute('title') || tag.innerText || tag.textContent; // 根据需要获取文本内容
        let state = { d: id, t: title, u: url };

        Diaspora.AJAX(url,function (data) {
            let parser = new DOMParser();
            let htmlDoc = parser.parseFromString(data, 'text/html');
            preview.innerHTML = htmlDoc.body.innerHTML;
            document.body.insertAdjacentElement("afterbegin", preview)
            Diaspora.player(id)
            Diaspora.PS()
            setTimeout(function () {
                // 将window.scrollY的值设置为#container元素的data-scroll属性
                document.getElementById('container').dataset.scroll = window.scrollY.toString();
                preview.style.transform = "unset"
                setTimeout(function () {
                    document.querySelector("#container").style.display = "none"
                }, 100)
            },500)
            document.title = title;
            switch (flag) {
                case 'push':
                    debugger
                    history.pushState(state, title, url)
                    break;
                //评论翻页的时候会用这个
                case 'replace':
                    history.replaceState(state, title, url)
                    break;
            }
        })
    },

    PS: function () {
        if (!(window.history && history.pushState)) return;

        history.replaceState({ u: Home, t: document.title }, document.title, Home)

        window.addEventListener('popstate', function (e) {
            let state = e.state;
            if (!state) return;
            document.title = state.t;

            if (state.u === Home) {
                preview.style.transform="translateX(100%)"
                document.getElementById('container').style.display = 'block';
                // window.scrollTo(0, parseInt(document.getElementById('container').dataset.scroll));
            } else {//后退之后又前进
                document.body.insertAdjacentElement("afterbegin", preview)
                //todo 后退后前进滚动条会跳动
                // window.scrollTo(0, parseInt(document.getElementById('container').dataset.scroll));
                setTimeout(function () {
                    preview.style.transform = "unset"
                    setTimeout(function () {
                        document.querySelector("#container").style.display = "none"
                    }, 100)
                },500)
            }

        })
    },

    player: function (id) {
        let au = document.getElementById('audio')
        let playButton = document.querySelector('.icon-play');
        let playbackBar = document.querySelector('.playbackbar');

        if(!au){
            return
        }
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

    SubmitComment:function (e) {
        if (e.preventDefault) e.preventDefault();
        // 获取表单元素
        const form = document.querySelector('form');

        // 设置表单action
        const localpath = window.location.pathname
        form.action = localpath + '/comment';
        // 构建额外的参数
        const extraData = {
            timestamp: Date.now(),
        };
        // 构建FormData对象
        const formData = new FormData(form);
        if (form.parentElement.firstChild !== form) {

        } else {
            extraData.parent = "0"
        }
        extraData.cid = localpath.split("/")[localpath.length - 1]

        // 追加额外参数
        Object.keys(extraData).forEach(key => {
            formData.append(key, extraData[key]);
        });
        // 提交表单
        fetch(form.action, {
            method: 'POST',
            body: formData
        })
            .then(response => {
                if (response.status === 200) {
                    // 成功提交
                    alert('评论发表成功!');
                } else {
                    // 请求失败
                    alert('评论发表失败,请重试!');
                }
            });

        // 阻止表单默认提交行为
        return false;
    },

    ReplyComment:function () {
    let commentForm = document.getElementById("comment-form")
    const buttongroup = commentForm.querySelector("#buttongroup")
    const cancelButton = document.createElement('button');
    cancelButton.className = "cancelButton";
    cancelButton.innerText = "取消回复";
    cancelButton.addEventListener('click', function () {
        const commentWrap = document.querySelector(".comment-wrap")
        commentWrap.insertAdjacentElement("afterbegin", commentForm)
        cancelButton.remove()
    });
    if (!buttongroup.querySelector(".cancelButton")) {
        buttongroup.insertAdjacentElement("afterbegin", cancelButton);
    }

    commentForm.action = new URL(window.location.href).pathname + "/comment"
    // 获取当前点击的按钮
    const replyButton = event.target;
    // 获取按钮的父节点的父节点
    const grandParentNode = replyButton.parentNode.parentNode;
    // 将表单插入到按钮的父节点的父节点后面
    grandParentNode.insertAdjacentElement("afterend", commentForm);
},

    scroller: function (){
        let scrollbar = document.querySelector('.scrollbar');
        let subtitle = document.querySelector('.subtitle');
        let article = document.getElementsByTagName('article');

        if (scrollbar && !('ontouchstart' in window)) {
            let st = window.scrollY||document.documentElement.scrollTop ;
            if (document.getElementById('preview')){
                st= document.getElementById('preview').scrollTop;
            }
            let ct = article.offsetHeight;
            if (st > ct) {
                st = ct;
            }
            scrollbar.style.width = ((50 + st) / ct * 100) + '%';

            if (st > 80 && window.innerWidth > 800) {
                subtitle.style.visibility = 'visible';
            } else {
                subtitle.style.visibility = 'hidden';
            }
        }
    },

    init:function () {
        if (('ontouchstart' in window)) {
            $('body').addClass('touch')
        }
        if(document.getElementById("container")){
            Diaspora.PS()
        }else {
            window.addEventListener('popstate', function (e) {
                if (e.state) location.href = e.state.u;
            })
            document.querySelector(".logo-min").href="/";
        }

    }
}

Diaspora.init()