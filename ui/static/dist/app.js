let attention=Prompt();const dropdown=document.querySelector(".js-nav");dropdown.addEventListener("click",(function(e){const t=e.target.closest(".nav-item");"dropDownMenuButton"===e.target.id&&(t.nextElementSibling.classList.toggle("hidden"),t.nextElementSibling.classList.contains("hidden")||(e.target.style.color="#181614"),ariaToggle(e.target),hideOnClickOutside(e.target),e.preventDefault())}));const hideOnClickOutside=e=>{const t=t=>{e.contains(t.target)||e&&(e.offsetWidth||e.offsetHeight||e.getClientRects().length)&&(e.nextElementSibling.classList.toggle("hidden"),s(),e.style.color="")},s=()=>{document.removeEventListener("click",t)};document.addEventListener("click",t)},ariaToggle=e=>{"false"===e.getAttribute("aria-expanded")?e.setAttribute("aria-expanded","true"):e.setAttribute("aria-expanded","false")},toggleNav=()=>{document.querySelector(".nav-brand-toggle").addEventListener("click",(function(e){document.querySelector(".navContentWrapper").classList.toggle("hidden")}))};document.querySelector(".nav-brand-toggle").addEventListener("click",(function(e){document.querySelector(".navContentWrapper").classList.toggle("hidden")}));class FormValidator{constructor(e,t){this.form=e,this.fields=t}validateOnSubmit(){let e=this;this.form.addEventListener("submit",(t=>{t.preventDefault(),Array.from(e.fields,(t=>{e.validateFields(t)}))}))}validateOnEntry(){let e=this;Array.from(e.fields,(t=>{const s=t.getAttribute("id");document.querySelector(`#${s}`).addEventListener("input",(s=>{e.validateFields(t)}))}))}validateFields(e){if(""===e.value.trim()?this.setStatus(e,`${e.previousElementSibling.innerHTML} cannot be blank`,"error"):this.setStatus(e,null,"success"),"email"===e.type){/^(([^<>()[\]\.,;:\s@\"]+(\.[^<>()[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s@\"]+\.)+[^<>()[\]\.,;:\s@\"]{2,})$/i.test(e.value)?this.setStatus(e,null,"success"):this.setStatus(e,"Please enter a valid email address","error")}if("passwordConfirmation"===e.id){const t=this.form.querySelector("#password");""===e.value.trim()?this.setStatus(e,"Password confirmation required","error"):e.value!==t.value?this.setStatus(e,"Password does not match","error"):this.setStatus(e,null,"success")}if("endDate"===e.id){const t=this.form.querySelector("#startDate"),s=new Date(t.value),i=new Date(e.value);s.getTime()>=i.getTime()?this.setStatus(e,"Select date after your arrival","error"):this.setStatus(e,null,"success")}}setStatus(e,t,s){const i=e.parentElement.querySelector(".error-message"),r=e.parentElement.querySelector(".form-alert");"success"===s&&(r.classList.remove("is-danger"),r.classList.add("is-success"),e.classList.remove("is-danger"),e.classList.add("is-success"),i&&(i.innerHTML="")),"error"===s&&(r.classList.remove("is-success"),i.innerHTML=t,r.classList.add("is-danger"),e.classList.add("is-danger"))}init(){this.validateOnSubmit(),this.validateOnEntry()}}const form=document.querySelector(".js-validateForm"),fields=document.querySelectorAll(".js-validateForm input"),validator=new FormValidator(form,fields);function notify(e,t){notie.alert({type:t,text:e})}function notifyModal(e,t,s,i){Swal.fire({title:e,html:t,icon:s,confirmButtonText:i})}function Prompt(){return{toast:function(e){const{msg:t="",icon:s="success",position:i="top-end"}=e;Swal.mixin({toast:!0,title:t,position:i,showConfirmButton:!1,timer:3e3,timerProgressBar:!0,didOpen:e=>{e.addEventListener("mouseenter",Swal.stopTimer),e.addEventListener("mouseleave",Swal.resumeTimer)}}).fire({icon:s})},success:function(e){const{msg:t="",title:s="",footer:i=""}=e;Swal.fire({icon:"success",title:s,text:t,footer:i})},error:function(e){const{msg:t="",title:s="",footer:i=""}=e;Swal.fire({icon:"error",title:s,text:t,footer:i})}}}null!=form&&validator.init();
//# sourceMappingURL=app.js.map