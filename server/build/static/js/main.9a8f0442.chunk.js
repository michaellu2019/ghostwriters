(this.webpackJsonpclient=this.webpackJsonpclient||[]).push([[0],{37:function(e,t,n){},38:function(e,t,n){},39:function(e,t,n){},40:function(e,t,n){"use strict";n.r(t);var s=n(1),c=n(23),r=n.n(c),a=n(4),i=n.n(a),o=n(13),u=n(16),l=n(10),j=n(8),d=n(9),h=n(2),b=n(15),p=n(0);var O=function(e){return Object(p.jsx)(d.b,{to:"/story/"+e.id,children:Object(p.jsxs)("div",{className:"story-card",children:[Object(p.jsxs)("header",{children:[Object(p.jsx)("h3",{children:Object(p.jsx)("i",{children:e.title})}),Object(p.jsxs)("h5",{children:["by ",e.author]})]}),Object(p.jsx)("img",{src:e.imageURL}),Object(p.jsxs)("footer",{children:["Posts: ",null!=e.content?e.content.length:0]})]})})};var x=function(){var e=Object(s.useState)([]),t=Object(j.a)(e,2),n=t[0],c=t[1];function r(){return(r=Object(l.a)(i.a.mark((function e(){var t,s,r;return i.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return t={method:"GET",mode:"cors",headers:{"Content-Type":"application/json"}},e.next=3,fetch("/api/get-stories",t);case 3:return s=e.sent,e.next=6,s.json();case 6:"OK"==(r=e.sent).status&&c([].concat(Object(b.a)(n),Object(b.a)(r.data.stories)));case 8:case"end":return e.stop()}}),e)})))).apply(this,arguments)}return Object(s.useEffect)((function(){!function(){r.apply(this,arguments)}()}),[]),Object(p.jsxs)("article",{children:[Object(p.jsxs)("header",{children:[Object(p.jsx)("h1",{children:"Public Stories"}),Object(p.jsx)("h3",{children:"Click on a story to view and edit it!"})]}),Object(p.jsx)("div",{className:"content",children:Object(p.jsx)("div",{className:"stories-container",children:n.map((function(e){return Object(p.jsx)(O,{id:e.id,author:e.author,title:e.title,imageURL:e.image_url,content:e.content,createdAt:e.createdAt},e.id)}))})})]})};var f=function(e){var t=Object(s.useState)(!1),n=Object(j.a)(t,2),c=n[0],r=n[1],a=Object(s.useState)(!1),i=Object(j.a)(a,2),o=i[0],u=i[1],l=e.createdAt.split(/[-TZ:]/),d=new Date(Date.UTC(l[0],l[1],l[2],l[3],l[4],l[5])),h="".concat(d.getMonth(),"/").concat(+d.getDate(),"/").concat(d.getFullYear()),b=e.author;return b.indexOf(":")>-1&&(b=b.split(":")[0]),Object(p.jsxs)("div",{className:"post",children:[Object(p.jsxs)("div",{className:"post-content",onMouseEnter:function(){return r(!0)},onMouseLeave:function(){return r(!1)},onClick:function(){return u(!o)},children:[Object(p.jsxs)("span",{className:"heart-container",children:[Object(p.jsx)("svg",{className:e.liked?"heart-button active":"heart-button",onClick:function(){return e.likePost(e.id)},children:Object(p.jsx)("path",{d:"M17.027 2.21c-2.248 0-4.166 1.786-5.027 3.704C11.139 3.995 9.222 2.21 6.973 2.21 3.931 2.21 1.416 4.725 1.416 7.766c0 6.218 6.283 7.872 10.584 14.024 4.035-6.152 10.584-8.005 10.584-14.024C22.584 4.725 20.072 2.21 17.027 2.21z"})}),e.likes]}),Object(p.jsx)("span",{className:c||o?"post-text active":"post-text",children:e.text})]}),Object(p.jsx)("div",{className:c||o?"post-info active":"post-info",children:Object(p.jsxs)("div",{children:["Written by ",b," on ",h+" - "+d.toLocaleTimeString()]})})]})},m=n(25),k=n.n(m);var v=function(e){return Object(p.jsx)(k.a,{appId:"850650752207328",autoLoad:!1,fields:"name,picture",onClick:function(){console.log("Facebook button clicked.",e)},callback:function(t){"unknown"!=t.status&&e.login({username:t.name,picture:t.picture.data.url,userId:t.userID})}})};var y=function(e){var t=Object(h.e)().id,n=Object(s.useState)([]),c=Object(j.a)(n,2),r=c[0],a=c[1],o=Object(s.useState)(""),u=Object(j.a)(o,2),d=u[0],O=u[1],x=Object(s.useRef)();function m(){return k.apply(this,arguments)}function k(){return(k=Object(l.a)(i.a.mark((function e(){var n,s,c;return i.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return n={method:"GET",mode:"cors",headers:{"Content-Type":"application/json"}},e.next=3,fetch("/api/get-story?id="+t,n);case 3:return s=e.sent,e.next=6,s.json();case 6:"OK"==(c=e.sent).status&&a(Object(b.a)(c.data.stories));case 8:case"end":return e.stop()}}),e)})))).apply(this,arguments)}function y(e){return g.apply(this,arguments)}function g(){return(g=Object(l.a)(i.a.mark((function t(n){var s,c,r,a;return i.a.wrap((function(t){for(;;)switch(t.prev=t.next){case 0:if(!e.loggedIn){t.next=10;break}return s={id:n,author:e.user.username+":"+e.user.userId},c={method:"POST",mode:"cors",headers:{"Content-Type":"application/json"},body:JSON.stringify(s)},t.next=5,fetch(e.user.likedPosts.hasOwnProperty(n)?"api/unlike-post":"/api/like-post",c);case 5:return r=t.sent,t.next=8,r.json();case 8:"OK"==(a=t.sent).status?(m(),e.login(e.user),e.userLikePost(n)):O(a.data);case 10:case"end":return t.stop()}}),t)})))).apply(this,arguments)}function w(){return(w=Object(l.a)(i.a.mark((function n(s){var c,r,a,o;return i.a.wrap((function(n){for(;;)switch(n.prev=n.next){case 0:return s.preventDefault(),c={story_id:parseInt(t),author:e.user.username+":"+e.user.userId,text:x.current.value},r={method:"POST",mode:"cors",headers:{"Content-Type":"application/json"},body:JSON.stringify(c)},n.next=5,fetch("/api/create-post",r);case 5:return a=n.sent,n.next=8,a.json();case 8:"OK"==(o=n.sent).status?(m(),x.current.value="",O("")):O(o.data);case 10:case"end":return n.stop()}}),n)})))).apply(this,arguments)}return Object(s.useEffect)((function(){m()}),[]),Object(p.jsxs)("article",{className:"story",children:[r.map((function(t){return Object(p.jsxs)("div",{children:[Object(p.jsxs)("header",{children:[Object(p.jsx)("h1",{children:Object(p.jsx)("i",{children:t.title})}),Object(p.jsxs)("h3",{children:["by ",t.author]})]}),Object(p.jsx)("div",{className:"content",children:null!=t.content?t.content.map((function(t){return Object(p.jsx)(f,{id:t.id,author:t.author,text:t.text,likes:t.likes,dislikes:t.dislikes,createdAt:t.created_at,liked:e.user.hasOwnProperty("likedPosts")&&e.user.likedPosts.hasOwnProperty(t.id),likePost:y},t.id)})):""})]},t.id)})),Object(p.jsx)("div",{className:"post-section",children:e.loggedIn?Object(p.jsxs)("form",{children:[Object(p.jsx)("span",{className:"error-message",children:d}),Object(p.jsx)("textarea",{ref:x,placeholder:"Hi ".concat(e.user.username,"! Contribute to the story here!")}),Object(p.jsx)("button",{onClick:function(e){return w.apply(this,arguments)},children:"SUBMIT"})]}):Object(p.jsxs)("div",{className:"login-section",children:[Object(p.jsx)("h2",{children:"Want to contribute to the story?"}),Object(p.jsx)(v,{login:e.login})]})})]})};n(37),n(38),n(39);var g=function(){var e=Object(s.useState)(!1),t=Object(j.a)(e,2),n=t[0],c=t[1],r=Object(s.useState)({}),a=Object(j.a)(r,2),b=a[0],O=a[1];function f(e){return m.apply(this,arguments)}function m(){return(m=Object(l.a)(i.a.mark((function e(t){var n,s,r;return i.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:if(!t.hasOwnProperty("userId")||!t.hasOwnProperty("username")){e.next=12;break}return console.log(t),c(!0),t.likedPosts={},n={method:"GET",mode:"cors",headers:{"Content-Type":"application/json"}},e.next=7,fetch("/api/get-author-post-likes?author="+t.username+":"+t.userId,n);case 7:return s=e.sent,e.next=10,s.json();case 10:"OK"==(r=e.sent).status&&null!=r.data.post_likes&&r.data.post_likes.forEach((function(e){t.likedPosts[e.post_id]=e.author}));case 12:O(t);case 13:case"end":return e.stop()}}),e)})))).apply(this,arguments)}function k(e){b.hasOwnProperty("likedPosts")&&O((function(t){return Object(u.a)(Object(u.a)({},t),{},Object(o.a)({},t.likedPosts,Object(u.a)(Object(u.a)({},t.likedPosts),{},Object(o.a)({},e,t.username+":"+t.userId))))}))}return Object(p.jsxs)("div",{className:"wrapper",children:[Object(p.jsx)("header",{className:"primary",children:Object(p.jsxs)("nav",{children:[Object(p.jsx)("span",{className:"heading",children:Object(p.jsx)(d.b,{className:"nav-link",to:"/",children:"Ghostwriters"})}),Object(p.jsx)("ul",{className:"nav-buttons links",children:Object(p.jsx)("li",{children:Object(p.jsx)(d.b,{className:"nav-link",to:"/",children:"Stories"})})})]})}),Object(p.jsxs)("main",{className:"primary",children:[Object(p.jsx)(h.a,{exact:!0,path:"/",render:function(){return Object(p.jsx)(x,{loggedIn:n,user:b,login:f,userLikePost:k})}}),Object(p.jsx)(h.a,{exact:!0,path:"/story/:id",render:function(){return Object(p.jsx)(y,{loggedIn:n,user:b,login:f,userLikePost:k})}})]})]})};r.a.render(Object(p.jsx)(d.a,{children:Object(p.jsx)(g,{})}),document.getElementById("root"))}},[[40,1,2]]]);
//# sourceMappingURL=main.9a8f0442.chunk.js.map