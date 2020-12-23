function toggleFAB(){var fabContainer=document.querySelector(".fab-message"),fabButton=fabContainer.querySelector(".fab-message-button a"),fabToggle=document.getElementById("fab-message-toggle");fabContainer.classList.toggle("is-open"),fabButton.classList.toggle("toggle-icon")}$(document).ready(function(){var fabContainer=document.querySelector(".fab-message"),messages=document.querySelector(".fab-message-content h3");messages&&(fabContainer.style.display="initial")}),$(document).ready(function(){var currentTheme=localStorage.getItem("theme");snow(currentTheme)});var snowFlakes=null;function snow(theme){if(snowFlakes)try{snowFlakes.destroy()}catch(error){}snowFlakes=new Snowflakes({color:theme==="dark"?"#fff":"#5ECDEF",count:9,minOpacity:theme==="dark"?.1:.33,maxOpacity:.95,minSize:10,maxSize:20,rotation:!0,speed:1,wind:!1,zIndex:100})}function switchTheme(e){var d1=document.getElementById("app-theme");e.target.checked?(d1.href="/theme/css/beacon-light.min.css",document.documentElement.setAttribute("data-theme","light"),localStorage.setItem("theme","light"),snow("light")):(d1.href="/theme/css/beacon-dark.min.css",document.documentElement.setAttribute("data-theme","dark"),localStorage.setItem("theme","dark"),snow("dark"))}$("#toggleSwitch").on("change",switchTheme),$(document).ready(function(){formatTimestamps(),$('[data-toggle="tooltip"]').tooltip();var bhValidators=new Bloodhound({datumTokenizer:Bloodhound.tokenizers.whitespace,queryTokenizer:Bloodhound.tokenizers.whitespace,identify:function(obj){return obj.pubkey},remote:{url:"/search/validators/%QUERY",wildcard:"%QUERY"}}),bhBlocks=new Bloodhound({datumTokenizer:Bloodhound.tokenizers.whitespace,queryTokenizer:Bloodhound.tokenizers.whitespace,identify:function(obj){return obj.slot},remote:{url:"/search/blocks/%QUERY",wildcard:"%QUERY"}}),bhGraffiti=new Bloodhound({datumTokenizer:Bloodhound.tokenizers.whitespace,queryTokenizer:Bloodhound.tokenizers.whitespace,identify:function(obj){return obj.graffiti},remote:{url:"/search/graffiti/%QUERY",wildcard:"%QUERY"}}),bhEpochs=new Bloodhound({datumTokenizer:Bloodhound.tokenizers.whitespace,queryTokenizer:Bloodhound.tokenizers.whitespace,identify:function(obj){return obj.epoch},remote:{url:"/search/epochs/%QUERY",wildcard:"%QUERY"}}),bhEth1Accounts=new Bloodhound({datumTokenizer:Bloodhound.tokenizers.whitespace,queryTokenizer:Bloodhound.tokenizers.whitespace,identify:function(obj){return obj.account},remote:{url:"/search/eth1_addresses/%QUERY",wildcard:"%QUERY"}});$(".typeahead").typeahead({minLength:1,highlight:!0,hint:!1,autoselect:!1},{limit:5,name:"validators",source:bhValidators,display:"pubkey",templates:{header:'<h3 class="h5">Validators</h3>',suggestion:function(data){return`<div class="text-monospace text-truncate">${data.index}: ${data.pubkey}</div>`}}},{limit:5,name:"blocks",source:bhBlocks,display:"blockroot",templates:{header:'<h3 class="h5">Blocks</h3>',suggestion:function(data){return`<div class="text-monospace text-truncate">${data.slot}: ${data.blockroot}</div>`}}},{limit:5,name:"epochs",source:bhEpochs,display:"epoch",templates:{header:'<h3 class="h5">Epochs</h3>',suggestion:function(data){return`<div>${data.epoch}</div>`}}},{limit:5,name:"addresses",source:bhEth1Accounts,display:"address",templates:{header:'<h3 class="h5">ETH1 Addresses</h3>',suggestion:function(data){return`<div class="text-monospace text-truncate">0x${data.address}</div>`}}},{limit:5,name:"graffiti",source:bhGraffiti,display:"graffiti",templates:{header:'<h3 class="h5">Blocks by Graffiti</h3>',suggestion:function(data){return`<div class="text-monospace" style="display:flex"><div class="text-truncate" style="flex:1 1 auto;">${data.graffiti}</div><div style="max-width:fit-content;white-space:nowrap;">${data.count}</div></div>`}}}),$(".typeahead").on("focus",function(event){event.target.value!==""&&$(this).trigger($.Event("keydown",{keyCode:40}))}),$(".typeahead").on("input",function(input){$(".tt-suggestion").first().addClass("tt-cursor")}),$(".tt-menu").on("mouseenter",function(){$(".tt-suggestion").first().removeClass("tt-cursor")}),$(".tt-menu").on("mouseleave",function(){$(".tt-suggestion").first().addClass("tt-cursor")}),$(".typeahead").on("typeahead:select",function(ev,sug){if(sug.slot!==void 0)window.location="/block/"+sug.slot;else if(sug.index!==void 0)sug.index==="deposited"?window.location="/validator/"+sug.pubkey:window.location="/validator/"+sug.index;else if(sug.epoch!==void 0)window.location="/epoch/"+sug.epoch;else if(sug.address!==void 0)window.location="/validators/eth1deposits?q="+sug.address;else if(sug.graffiti!==void 0){var el=document.createElement("textarea");el.innerHTML=sug.graffiti,window.location="/blocks?q="+encodeURIComponent(el.value)}else console.log("invalid typeahead-selection",sug)})}),$("[aria-ethereum-date]").each(function(item){var dt=$(this).attr("aria-ethereum-date"),format=$(this).attr("aria-ethereum-date-format");format||(format="ff"),format==="FROMNOW"?($(this).text(luxon.DateTime.fromMillis(dt*1e3).toRelative({style:"short"})),$(this).attr("title",luxon.DateTime.fromMillis(dt*1e3).toFormat("ff"))):$(this).text(luxon.DateTime.fromMillis(dt*1e3).toFormat(format))}),$(document).ready(function(){var clipboard=new ClipboardJS("[data-clipboard-text]");clipboard.on("success",function(e){var title=$(e.trigger).attr("data-original-title");$(e.trigger).tooltip("hide").attr("data-original-title","Copied!").tooltip("show"),setTimeout(function(){$(e.trigger).tooltip("hide").attr("data-original-title",title)},1e3)}),clipboard.on("error",function(e){var title=$(e.trigger).attr("data-original-title");$(e.trigger).tooltip("hide").attr("data-original-title","Failed to Copy!").tooltip("show"),setTimeout(function(){$(e.trigger).tooltip("hide").attr("data-original-title",title)},1e3)})}),$(".nav-tabs a").on("shown.bs.tab",function(e){history.replaceState?history.pushState(null,null,e.target.hash):window.location.hash=e.target.hash});var url=document.location.toString();url.match("#")&&$('.nav-tabs a[href="#'+url.split("#")[1]+'"]').tab("show");function formatTimestamps(selStr){var sel=$(document);selStr!==void 0&&(sel=$(selStr)),sel.find(".timestamp").each(function(){var ts=$(this).data("timestamp"),tsLuxon=luxon.DateTime.fromMillis(ts*1e3);$(this).attr("data-original-title",tsLuxon.toFormat("ff")),$(this).text(tsLuxon.toRelative({style:"short"}))}),sel.find('[data-toggle="tooltip"]').tooltip()}
