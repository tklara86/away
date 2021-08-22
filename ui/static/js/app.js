const dropdown = document.querySelector('.js-nav');

dropdown.addEventListener('click', function(e) {
    const dropDownTarget = e.target.closest('.nav-item');


    if (e.target.id === 'dropDownMenuButton') {
        // Hide the dropdown
        dropDownTarget.nextElementSibling.classList.toggle('hidden');

        if (!dropDownTarget.nextElementSibling.classList.contains('hidden')) {
            e.target.style.color = '#181614';
        }

        // Toggle aria expanded attributes
        ariaToggle(e.target)

        // Click outside
        hideOnClickOutside(e.target);

        e.preventDefault();
    }

})

const hideOnClickOutside = (element) => {
    const outsideClickListener = (event) => {
       if (!element.contains(event.target)) {
           if (!!element && !!(element.offsetWidth || element.offsetHeight || element.getClientRects().length)) {
                element.nextElementSibling.classList.toggle('hidden');
                removeClickListener();
                element.style.color = '';

           }
       }
    }
    const removeClickListener = () => {
        document.removeEventListener('click', outsideClickListener);
    }
    document.addEventListener('click', outsideClickListener);
}



const ariaToggle = (toggle) => {
    toggle.getAttribute('aria-expanded') === 'false' ?
        toggle.setAttribute('aria-expanded', 'true') :
        toggle.setAttribute('aria-expanded', 'false');
}


const toggleNav = () => {
    const toggleNav = document.querySelector('.nav-brand-toggle');

    toggleNav.addEventListener('click', function (e) {
       const navContentWrapper = document.querySelector('.navContentWrapper');
       navContentWrapper.classList.toggle('hidden')
    })
}

toggleNav();





































































































































































































































//# sourceMappingURL=app.js.map
