import{S as t,i as s,s as e,e as l,t as o,c as r,b as n,g as a,d as h,h as c,k as f,l as i,m as g,a as u,q as p,f as d,n as m,o as b}from"./client.38965d01.js";function v(t,s,e){const l=t.slice();return l[1]=s[e],l}function E(t){let s,e,u,p,d=t[1].title+"";return{c(){s=l("li"),e=l("a"),u=o(d),this.h()},l(t){s=r(t,"LI",{});var l=n(s);e=r(l,"A",{rel:!0,href:!0});var o=n(e);u=a(o,d),o.forEach(h),l.forEach(h),this.h()},h(){c(e,"rel","prefetch"),c(e,"href",p="blogs/"+t[1].slug)},m(t,l){f(t,s,l),i(s,e),i(e,u)},p(t,s){1&s&&d!==(d=t[1].title+"")&&g(u,d),1&s&&p!==(p="blogs/"+t[1].slug)&&c(e,"href",p)},d(t){t&&h(s)}}}function j(t){let s,e,g,j,x,q=t[0],L=[];for(let s=0;s<q.length;s+=1)L[s]=E(v(t,q,s));return{c(){s=u(),e=l("h1"),g=o("Recent posts"),j=u(),x=l("ul");for(let t=0;t<L.length;t+=1)L[t].c();this.h()},l(t){p('[data-svelte="svelte-hfp9t8"]',document.head).forEach(h),s=d(t),e=r(t,"H1",{});var l=n(e);g=a(l,"Recent posts"),l.forEach(h),j=d(t),x=r(t,"UL",{class:!0});var o=n(x);for(let t=0;t<L.length;t+=1)L[t].l(o);o.forEach(h),this.h()},h(){document.title="Blog",c(x,"class","svelte-1frg2tf")},m(t,l){f(t,s,l),f(t,e,l),i(e,g),f(t,j,l),f(t,x,l);for(let t=0;t<L.length;t+=1)L[t].m(x,null)},p(t,[s]){if(1&s){let e;for(q=t[0],e=0;e<q.length;e+=1){const l=v(t,q,e);L[e]?L[e].p(l,s):(L[e]=E(l),L[e].c(),L[e].m(x,null))}for(;e<L.length;e+=1)L[e].d(1);L.length=q.length}},i:m,o:m,d(t){t&&h(s),t&&h(e),t&&h(j),t&&h(x),b(L,t)}}}function x({params:t,query:s}){return this.fetch("blogs.json").then(t=>t.json()).then(t=>({blogs:t}))}function q(t,s,e){let{blogs:l}=s;return t.$set=t=>{"blogs"in t&&e(0,l=t.blogs)},[l]}export default class extends t{constructor(t){super(),s(this,t,q,j,e,{blogs:0})}}export{x as preload};
