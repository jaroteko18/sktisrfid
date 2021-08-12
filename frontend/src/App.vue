<template>
  <div id="app" class="bg-light">
    <!-- <router-link to="/home"> Home </router-link> | 
    <router-link to="/absenteeism"> Absenteeism </router-link> | 
    <router-link to="/productiontarget"> Production Target </router-link> | 
    <router-link to="/test"> TEST </router-link>

    <hr> -->

    <div >
      <main class="container">
        <div class="d-flex align-items-center p-3 my-3 text-white bg-blue rounded shadow-sm">
          <div class="d-flex align-items-center p-1 text-white bg-white rounded shadow-sm">
          <img class="img-thumbnail" src="./assets/images/sampoerna-logo-small.png" alt="" width="48" height="38">
          </div>
          
          <div class="lh-1 pl-10">
            <h1 class="h6 mb-0 text-white lh-1">SKT-IS</h1>
            <small>RFID</small>
          </div>
        </div>

        <router-view></router-view>

        <br>
      </main>
    </div>

    
    
  </div>
</template>

<script>
// import RFID from "./components/RFID.vue";
import "./assets/css/main.css";


export default {
  name: "app",
  // components: {
  //   RFID
  // }
};
</script>


<style scoped>
.bd-example {
    padding: 1.5rem;
    margin-right: 0;
    margin-left: 0;
    border-width: 1px;
    border-top-left-radius: .25rem;
    border-top-right-radius: .25rem;
    position: relative;
    margin: 1rem -.75rem 0;
    border: solid #dee2e6;
}
.bg-blue {
    background-color: #1B4C7B;
}
.bg-white {
    background-color: white;
}
.pl-10{
  padding-left: 10px;
}
body{
  background-color: #f8f9fa !important;
}
</style>

