<template>
  <div class="my-3 p-3 bg-body rounded shadow-sm">  
    <div class="border-bottom">
    <nav style="--bs-breadcrumb-divider: url(&#34;data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='8' height='8'%3E%3Cpath d='M2.5 0L1 1.5 3.5 4 1 6.5 2.5 8l4-4-4-4z' fill='currentColor'/%3E%3C/svg%3E&#34;);" aria-label="breadcrumb">
      <ol class="breadcrumb">
        <!-- <li class="breadcrumb-item"><router-link to="#"> Home </router-link></li> -->
        <li class="breadcrumb-item active" aria-current="page"><router-link to="#"> Home </router-link></li>
      </ol>
    </nav>
    </div>

    <div class="text-muted pt-3">
      <form v-on:submit.prevent="onSubmit">
        <div class="row mb-3">
          <label for="inputrfidType" class="col-sm-3 col-form-label">Select Type</label>
          <div class="col-sm-9">
            <div class="btn-group" role="group" aria-label="Basic radio toggle button group">
              <input type="radio" class="btn-check" name="rfidType" id="rfidType1" autocomplete="off" value="absenteeism" v-model="rfidType">
              <label class="btn btn-outline-primary radio-custom" for="rfidType1">Absenteeism</label>

              <input type="radio" class="btn-check" name="rfidType" id="rfidType2" autocomplete="off" value="productiontarget" v-model="rfidType">
              <label class="btn btn-outline-primary radio-custom" for="rfidType2">Production Target</label>
            </div>
          </div>
        </div>

        <div class="row mb-3" v-if="rfidType === 'absenteeism'">
          <label for="inputPassword3" class="col-sm-3 col-form-label"></label>
          <div class="col-sm-9">
            <div class="btn-group" role="group" aria-label="Basic radio toggle button group">
              <input  type="radio" class="btn-check" name="absenteeismType" id="absenteeismType1" autocomplete="off" value="Ijin" v-model="absenteeismType">
              <label class="btn btn-outline-primary radio-custom" for="absenteeismType1">Ijin</label>

              <input  type="radio" class="btn-check" name="absenteeismType" id="absenteeismType2" autocomplete="off" value="Alpa" v-model="absenteeismType">
              <label class="btn btn-outline-primary radio-custom" for="absenteeismType2">Alpa</label>
            </div>
          </div>
        </div>

        <div class="row mb-3">
          <label for="inputPassword3" class="col-sm-3 col-form-label">Date</label>
          <div class="col-sm-9">
          <date-picker class="radio-custom" v-model="rfidPeriod" :clearable="false" value-type="format" format="YYYY-MM-DD"></date-picker>
          </div>
        </div>

        <div class="row mb-3 border-bottom ">
          <div>
            <b-alert
              :show="dismissCountDown"
              fade
              variant="danger"
              @dismiss-count-down="countDownChanged"
            >
              {{messageSubmit}}
            </b-alert>
          </div>
        </div>

        <!-- <button type="submit"  class="btn btn-primary">Enter</button> -->
        <small class="d-block text-end mt-3">
          <button type="submit"  class="btn btn-primary">Enter</button>
        </small>

        <h3>{{message}}</h3>
        <a @click="getMessage">Press Me!</a>
      </form>
    </div>
  </div>
</template>

<script>
  import DatePicker from 'vue2-datepicker';
  import 'vue2-datepicker/index.css';
  import axios from 'axios';
  export default {
    components: { DatePicker },
    data() {
      return {
        message:'',
        dismissSecs: 5,
        dismissCountDown: 0,
        rfidType: '',
        absenteeismType: '',
        rfidPeriod:new Date().toISOString().slice(0,10), 
        messageSubmit:'',
      }
    },
    mounted:function() {
      
    },
    methods:{
      getMessage: function() {
        var self = this;
        // window.backend.basic('307820').then(result => {
        //   self.message = result;
        // });
        window.backend.RFID.GetEmployee('307820').then(result => {
          self.message = result;
        });
      },
      countDownChanged(dismissCountDown) {
        this.dismissCountDown = dismissCountDown
      },
      showAlert() {
        this.dismissCountDown = this.dismissSecs
      },
      onSubmit(){
        if(this.rfidType=='' || this.rfidPeriod==''){
          this.messageSubmit='Please choose type and date !'
          this.dismissCountDown = this.dismissSecs
        }else{
          if(this.rfidType=='absenteeism'){
            if(this.absenteeismType==''){
              this.messageSubmit='Please choose type and date !'
              this.dismissCountDown = this.dismissSecs
            }else{
              this.$router.push({ name: 'Absenteeism', params: { absenttype: this.absenteeismType, period: this.rfidPeriod } })
            }
          }else{
            this.$router.push({ name: 'ProductionTarget', params: { period: this.rfidPeriod } })
          }
        }
      }
      
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.radio-custom{
  width:250px;
}
</style>
