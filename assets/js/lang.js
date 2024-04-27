var dataReload = document.querySelectorAll("[reload]");

var language = {

    eng: {
        menu_home:"Home",
        menu_profile:"profile",
        menu_edit:"Edit Profile",
        menu_logout:"Logout"
    },

    ar:{
        menus_home:"الرئسية",
        menus_profile:"الصفحة الشخصية",
        menus_edit:"تعديل الصفحة الشخصية",
        menus_logout:"خروج"
    }


};

if (window.location.hash) {
    if(window.location.hash ==="ar") {
        m_home.texContent = language.ar.menus_home;
        m_edit.textContent = language.ar.menu_edit;
        m_logout.textContent = language.ar.menu_logout;
        m_profile.textContent = language.ar.menu_profile;
    }
}