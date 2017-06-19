import Vue from 'vue/dist/vue.js'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'
import '../src/style.css'
import Nav from '../src/components/Nav.vue'
import UserTable from '../src/components/UserTable.vue'
import Footer from '../src/components/Footer.vue'


Vue.use(Buefy)

new Vue({
    el: '#nav',
    components: { 
        'nav-bar' : Nav 
    }
})

new Vue({
    el: '#footer',
    components: { 
        'foot' : Footer 
    }
})