/** @license
 * eventsource.js
 * Available under MIT License (MIT)
 * https://github.com/Yaffle/EventSource/
 */
!function(a){"use strict";function b(a){this.withCredentials=!1,this.responseType="",this.readyState=0,this.status=0,this.statusText="",this.responseText="",this.onprogress=p,this.onreadystatechange=p,this._contentType="",this._xhr=a,this._sendTimeout=0,this._abort=p}function c(a){this._xhr=new b(a)}function d(){this._listeners=Object.create(null)}function e(a){j(function(){throw a},0)}function f(a){this.type=a,this.target=void 0}function g(a,b){f.call(this,a),this.data=b.data,this.lastEventId=b.lastEventId}function h(a,b){d.call(this),this.onopen=void 0,this.onmessage=void 0,this.onerror=void 0,this.url=void 0,this.readyState=void 0,this.withCredentials=void 0,this._close=void 0,i(this,a,b)}function i(a,b,d){b=String(b);var h=void 0!=d&&Boolean(d.withCredentials),i=D(1e3),n=D(45e3),o="",p=i,A=!1,B=void 0!=d&&void 0!=d.headers?JSON.parse(JSON.stringify(d.headers)):void 0,F=void 0!=d&&void 0!=d.Transport?d.Transport:void 0!=m?m:l,G=new c(new F),H=0,I=q,J="",K="",L="",M="",N=v,O=0,P=0,Q=function(b,c,d){if(I===r)if(200===b&&void 0!=d&&z.test(d)){I=s,A=!0,p=i,a.readyState=s;var g=new f("open");a.dispatchEvent(g),E(a,a.onopen,g)}else{var h="";200!==b?(c&&(c=c.replace(/\s+/g," ")),h="EventSource's response has a status "+b+" "+c+" that is not 200. Aborting the connection."):h="EventSource's response has a Content-Type specifying an unsupported type: "+(void 0==d?"-":d.replace(/\s+/g," "))+". Aborting the connection.",e(new Error(h)),T();var g=new f("error");a.dispatchEvent(g),E(a,a.onerror,g)}},R=function(b){if(I===s){for(var c=-1,d=0;d<b.length;d+=1){var e=b.charCodeAt(d);(e==="\n".charCodeAt(0)||e==="\r".charCodeAt(0))&&(c=d)}var f=(-1!==c?M:"")+b.slice(0,c+1);M=(-1===c?M:"")+b.slice(c+1),""!==f&&(A=!0);for(var h=0;h<f.length;h+=1){var e=f.charCodeAt(h);if(N===u&&e==="\n".charCodeAt(0))N=v;else if(N===u&&(N=v),e==="\r".charCodeAt(0)||e==="\n".charCodeAt(0)){if(N!==v){N===w&&(P=h+1);var l=f.slice(O,P-1),m=f.slice(P+(h>P&&f.charCodeAt(P)===" ".charCodeAt(0)?1:0),h);"data"===l?(J+="\n",J+=m):"id"===l?K=m:"event"===l?L=m:"retry"===l?(i=C(m,i),p=i):"heartbeatTimeout"===l&&(n=C(m,n),0!==H&&(k(H),H=j(function(){U()},n)))}if(N===v){if(""!==J){o=K,""===L&&(L="message");var q=new g(L,{data:J.slice(1),lastEventId:K});if(a.dispatchEvent(q),"message"===L&&E(a,a.onmessage,q),I===t)return}J="",L=""}N=e==="\r".charCodeAt(0)?u:v}else N===v&&(O=h,N=w),N===w?e===":".charCodeAt(0)&&(P=h+1,N=x):N===x&&(N=y)}}},S=function(){if(I===s||I===r){I=q,0!==H&&(k(H),H=0),H=j(function(){U()},p),p=D(Math.min(16*i,2*p)),a.readyState=r;var b=new f("error");a.dispatchEvent(b),E(a,a.onerror,b)}},T=function(){I=t,G.cancel(),0!==H&&(k(H),H=0),a.readyState=t},U=function(){if(H=0,I!==q)return void(A?(A=!1,H=j(function(){U()},n)):(e(new Error("No activity within "+n+" milliseconds. Reconnecting.")),G.cancel()));A=!1,H=j(function(){U()},n),I=r,J="",L="",K=o,M="",O=0,P=0,N=v;var a=b;"data:"!==b.slice(0,5)&&"blob:"!==b.slice(0,5)&&(a=b+(-1===b.indexOf("?",0)?"?":"&")+"lastEventId="+encodeURIComponent(o));var c={};if(c.Accept="text/event-stream",void 0!=B)for(var d in B)Object.prototype.hasOwnProperty.call(B,d)&&(c[d]=B[d]);try{G.open(Q,R,S,a,h,c)}catch(f){throw T(),f}};a.url=b,a.readyState=r,a.withCredentials=h,a._close=T,U()}var j=a.setTimeout,k=a.clearTimeout,l=a.XMLHttpRequest,m=a.XDomainRequest,n=a.EventSource,o=a.document;null==Object.create&&(Object.create=function(a){function b(){}return b.prototype=a,new b});var p=function(){};b.prototype.open=function(a,b){this._abort(!0);var c=this,d=this._xhr,e=1,f=0;this._abort=function(a){0!==c._sendTimeout&&(k(c._sendTimeout),c._sendTimeout=0),(1===e||2===e||3===e)&&(e=4,d.onload=p,d.onerror=p,d.onabort=p,d.onprogress=p,d.onreadystatechange=p,d.abort(),0!==f&&(k(f),f=0),a||(c.readyState=4,c.onreadystatechange())),e=0};var g=function(){if(1===e){var a=0,b="",f=void 0;if("contentType"in d)a=200,b="OK",f=d.contentType;else try{a=d.status,b=d.statusText,f=d.getResponseHeader("Content-Type")}catch(g){a=0,b="",f=void 0}0!==a&&(e=2,c.readyState=2,c.status=a,c.statusText=b,c._contentType=f,c.onreadystatechange())}},h=function(){if(g(),2===e||3===e){e=3;var a="";try{a=d.responseText}catch(b){}c.readyState=3,c.responseText=a,c.onprogress()}},i=function(){h(),(1===e||2===e||3===e)&&(e=4,0!==f&&(k(f),f=0),c.readyState=4,c.onreadystatechange())},m=function(){void 0!=d&&(4===d.readyState?i():3===d.readyState?h():2===d.readyState&&g())},n=function(){f=j(function(){n()},500),3===d.readyState&&h()};d.onload=i,d.onerror=i,d.onabort=i,"sendAsBinary"in l.prototype||"mozAnon"in l.prototype||(d.onprogress=h),d.onreadystatechange=m,"contentType"in d&&(b+=(-1===b.indexOf("?",0)?"?":"&")+"padding=true"),d.open(a,b,!0),"readyState"in d&&(f=j(function(){n()},0))},b.prototype.abort=function(){this._abort(!1)},b.prototype.getResponseHeader=function(a){return this._contentType},b.prototype.setRequestHeader=function(a,b){var c=this._xhr;"setRequestHeader"in c&&c.setRequestHeader(a,b)},b.prototype.send=function(){if(!("ontimeout"in l.prototype)&&void 0!=o&&void 0!=o.readyState&&"complete"!==o.readyState){var a=this;return void(a._sendTimeout=j(function(){a._sendTimeout=0,a.send()},4))}var b=this._xhr;b.withCredentials=this.withCredentials,b.responseType=this.responseType;try{b.send(void 0)}catch(c){throw c}},c.prototype.open=function(a,b,c,d,e,f){var g=this._xhr;g.open("GET",d);var h=0;g.onprogress=function(){var a=g.responseText,c=a.slice(h);h+=c.length,b(c)},g.onreadystatechange=function(){if(2===g.readyState){var b=g.status,d=g.statusText,e=g.getResponseHeader("Content-Type");a(b,d,e)}else 4===g.readyState&&c()},g.withCredentials=e,g.responseType="text";for(var i in f)Object.prototype.hasOwnProperty.call(f,i)&&g.setRequestHeader(i,f[i]);g.send()},c.prototype.cancel=function(){var a=this._xhr;a.abort()},d.prototype.dispatchEvent=function(a){a.target=this;var b=this._listeners[a.type];if(void 0!=b)for(var c=b.length,d=0;c>d;d+=1){var f=b[d];try{"function"==typeof f.handleEvent?f.handleEvent(a):f.call(this,a)}catch(g){e(g)}}},d.prototype.addEventListener=function(a,b){a=String(a);var c=this._listeners,d=c[a];void 0==d&&(d=[],c[a]=d);for(var e=!1,f=0;f<d.length;f+=1)d[f]===b&&(e=!0);e||d.push(b)},d.prototype.removeEventListener=function(a,b){a=String(a);var c=this._listeners,d=c[a];if(void 0!=d){for(var e=[],f=0;f<d.length;f+=1)d[f]!==b&&e.push(d[f]);0===e.length?delete c[a]:c[a]=e}},g.prototype=Object.create(f.prototype);var q=-1,r=0,s=1,t=2,u=-1,v=0,w=1,x=2,y=3,z=/^text\/event\-stream;?(\s*charset\=utf\-8)?$/i,A=1e3,B=18e6,C=function(a,b){var c=parseInt(a,10);return c!==c&&(c=b),D(c)},D=function(a){return Math.min(Math.max(a,A),B)},E=function(a,b,c){try{"function"==typeof b&&b.call(a,c)}catch(d){e(d)}};h.prototype=Object.create(d.prototype),h.prototype.CONNECTING=r,h.prototype.OPEN=s,h.prototype.CLOSED=t,h.prototype.close=function(){this._close()},h.CONNECTING=r,h.OPEN=s,h.CLOSED=t,h.prototype.withCredentials=void 0,a.EventSourcePolyfill=h,a.NativeEventSource=n,void 0==l||void 0!=n&&"withCredentials"in n.prototype||(a.EventSource=h)}("undefined"!=typeof window?window:this);

/*!
 * smallPop 0.1.2 | https://github.com/silvio-r/spop
 * Copyright (c) 2015 Sílvio Rosa @silvior_
 * MIT license
 */

$(function() {


	;(function() {
		'use strict';

		var animationTime = 390;
		var options, defaults, container, icon, layout, popStyle, positions, close;

		var SmallPop = function(template, style) {

			this.defaults = {
				template  : null,
				style     : 'info',
				autoclose : false,
				position  : 'top-right',
				icon      : true,
				group     : false,
				onOpen    : false,
				onClose   : false
			};

			defaults = extend(this.defaults, spop.defaults);

			if ( typeof template === 'string' || typeof style === 'string' ) {
				options = { template: template, style: style || defaults.style};
			}
			else if (typeof template === 'object') {
				options = template;
			}
			else {
				console.error('Invalid arguments.');
				return false;
			}

			this.opt = extend( defaults, options);

			if ($('spop--' + this.opt.group)) {

				this.remove($('spop--' + this.opt.group));
			}

			this.open();
		};

		SmallPop.prototype.create = function(template) {

			container = $(this.getPosition('spop--', this.opt.position));

			icon = (!this.opt.icon) ? '' : '<i class="spop-icon '+
						this.getStyle('spop-icon--', this.opt.style) +'"></i>';

			layout ='<div class="spop-close" data-spop="close" aria-label="Close">&times;</div>' +
							icon +
						'<div class="spop-body">' +
							template +
						'</div>';

			if (!container) {

				this.popContainer = document.createElement('div');

				this.popContainer.setAttribute('class', 'spop-container ' +
					this.getPosition('spop--', this.opt.position));

				this.popContainer.setAttribute('id', this.getPosition('spop--', this.opt.position));

				document.body.appendChild(this.popContainer);

				container = $(this.getPosition('spop--', this.opt.position));
			}

			this.pop = document.createElement('div');

			this.pop.setAttribute('class', 'spop spop--out spop--in ' + this.getStyle('spop--', this.opt.style) );

			if (this.opt.group && typeof this.opt.group === 'string') {
				this.pop.setAttribute('id', 'spop--' + this.opt.group);
			}


			this.pop.setAttribute('role', 'alert');

			this.pop.innerHTML = layout;

			container.appendChild(this.pop);
		};

		SmallPop.prototype.getStyle = function(sufix, arg) {

			popStyle = {
				'success': 'success',
				'error'  : 'error',
				'warning': 'warning'
			};
			return sufix + (popStyle[arg] || 'info');
		};

		SmallPop.prototype.getPosition = function(sufix, position) {

			positions = {
				'top-left'     : 'top-left',
				'top-center'   : 'top-center',
				'top-right'    : 'top-right',
				'bottom-left'  : 'bottom-left',
				'bottom-center': 'bottom-center',
				'bottom-right' : 'bottom-right'
			};
			return sufix + (positions[position] || 'top-right');
		};

		SmallPop.prototype.open = function() {

			this.create(this.opt.template);

			if (this.opt.onOpen) { this.opt.onOpen();}

			this.close();
		};

		SmallPop.prototype.close = function () {

			if (this.opt.autoclose && typeof this.opt.autoclose === 'number') {

				this.autocloseTimer = setTimeout( this.remove.bind(null, this.pop), this.opt.autoclose);
			}

			this.pop.addEventListener('click', this.addListeners.bind(this) , false);
		};

		SmallPop.prototype.addListeners = function(event) {

			close = event.target.getAttribute('data-spop');

			if (close === 'close') {

				if (this.autocloseTimer) { clearTimeout(this.autocloseTimer);}

				if (this.opt.onClose) { this.opt.onClose();}

				this.remove(this.pop);
			}
		};

		SmallPop.prototype.remove = function(elm) {

			removeClass(elm, 'spop--in');

			setTimeout( function () {

				if(document.body.contains(elm)) {
					elm.parentNode.removeChild(elm);
				}

			}, animationTime);
		};


		// Helpers

		function $(el, con) {
			return typeof el === 'string'? (con || document).getElementById(el) : el || null;
		}

		function removeClass(el, className) {
			if (el.classList) {
				el.classList.remove(className);
			}
			else {
				el.className = el.className.replace(new RegExp('(^|\\b)' +
								className.split(' ').join('|') +
								'(\\b|$)', 'gi'), ' ');
			}
		}

		function extend(obj, src) {

			for (var key in src) {
				if (src.hasOwnProperty(key)) obj[key] = src[key];
			}

			return obj;
		}

		window.spop = function (template, style) {
			if ( !template || !window.addEventListener ) { return false;}

			return new SmallPop(template, style);
		};

		spop.defaults = {};
	}());

	$.getJSON("/index.php/apps/realtime_notifications/settings", function (data) {
		if ("backend_host" in data) {
			var backendHost = data['backend_host'];
			$.getJSON("/index.php/apps/realtime_notifications/token", function(token) {
				console.log(token);
				var source = new EventSourcePolyfill(backendHost + "/events", {headers: {"Authorization": "BEARER " + token["token"]}});
				source.addEventListener($('head').data('user'), function (ev) {
					FileList.reload();
					var data = JSON.parse(ev.data);
					spop({
						template: data["uidOwner"] + " shared " + data["fileTarget"],
						position  : 'bottom-right',
						style: 'success',
						autoclose: 5000
					});
				}, false);
			});
		}
	});
});
