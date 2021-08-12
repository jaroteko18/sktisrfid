// router.js
import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from "./components/Home.vue";

import Absenteeism from "./components/Absenteeism.vue";
import ProductionTarget from "./components/ProductionTarget.vue";
import Test from "./components/Test.vue";

Vue.use(VueRouter)

const routes = [
  { 
    component: Home, 
    name: 'Home', 
    path: '/home'
  },
  { 
    component: Absenteeism, 
    name: 'Absenteeism', 
    path: '/absenteeism' 
  },
  { 
    component: ProductionTarget, 
    name: 'ProductionTarget', 
    path: '/productiontarget' 
  },
  { 
    component: Test, 
    name: 'Test', 
    path: '/test' 
  }
]

const router = new VueRouter({
  mode: 'abstract', // mode must be set to 'abstract'
  routes,
})

export default router