<template>
  <div class="my-3 p-3 bg-body rounded shadow-sm">  
    <div class="row border-bottom p-1 my-1 ">
      <div class="col">
        <nav style="--bs-breadcrumb-divider: url(&#34;data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='8' height='8'%3E%3Cpath d='M2.5 0L1 1.5 3.5 4 1 6.5 2.5 8l4-4-4-4z' fill='currentColor'/%3E%3C/svg%3E&#34;);" aria-label="breadcrumb">
          <ol class="breadcrumb">
            <li class="breadcrumb-item"><router-link to="/home"> Home </router-link></li>
            <li class="breadcrumb-item active" aria-current="page">Absenteeism</li>
          </ol>
        </nav>
      </div>
      <div class="col text-end">
        <span class="btn btn-outline-primary position-relative">
          <svg style="margin-top:-4px" xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-calendar4" viewBox="0 0 16 16">
            <path d="M3.5 0a.5.5 0 0 1 .5.5V1h8V.5a.5.5 0 0 1 1 0V1h1a2 2 0 0 1 2 2v11a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V3a2 2 0 0 1 2-2h1V.5a.5.5 0 0 1 .5-.5zM2 2a1 1 0 0 0-1 1v1h14V3a1 1 0 0 0-1-1H2zm13 3H1v9a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V5z"/>
          </svg>
          {{ this.$route.params.period }}
          <span class="position-absolute top-0 start-100 translate-middle badge rounded-pill bg-danger">
            {{ this.$route.params.absenttype }}
          </span>
        </span>
      </div>
    </div>  
    <div class="text-muted pt-3">
      <form @submit.prevent="save">     
        <div class="row mb-3 border-bottom">
          <b-table :busy="isBusy" class="table table-sm custom-table" striped hover 
          :fields="fields"
          :items="list"
          :current-page="currentPage"
          :per-page="perPage"
          show-empty
          small>
            <template #table-busy>
              <div class="text-center text-primary my-2">
                <b-spinner class="align-middle"></b-spinner>
                <strong>Loading...</strong>
              </div>
            </template>

            <template #cell(actions)="row">
              <div class="text-center">
              <b-button style="font-size:12px;padding:0px;" size="sm" variant="outline-danger" @click="info(row.item, row.index, $event.target)">
                <b-icon icon="trash-fill" aria-hidden="true"></b-icon>
              </b-button>
              </div>
            </template>
          </b-table>

          <!-- Info modal -->
          <b-modal :id="infoModal.id" size="sm" @ok="handleDelete(infoModal.content)" buttonSize="sm" okVariant="danger" okTitle="YES" cancelTitle="NO" @hide="resetInfoModal">
            <template #modal-header="{}">
              <h5>{{infoModal.title}}</h5>
            </template>
            Delete the RFID <b>{{ infoModal.content.RFIDID }}</b> ?
          </b-modal>

          <b-row>
            <center>
            <b-col sm="7" md="6" class="my-1">
              <b-pagination
                v-model="currentPage"
                :total-rows="totalRows"
                :per-page="perPage"
                align="fill"
                size="sm"
                class="my-0"
              ></b-pagination>
            </b-col>
            <br>
            </center>
          </b-row>
        </div>
        <small class="d-block text-end mt-3">
           <div class="text-center text-primary my-2">
            <b-alert
              :show="dismissCountDown"
              fade
              
              :variant="validateVarian"
              @dismiss-count-down="countDownChanged"
            >
            <!-- variant="danger" -->
              <p style="margin-bottom:0">{{messageValidate}}</p>
            </b-alert>
          </div>
          <button type="submit" v-bind:disabled="hasSubmitted" class="btn btn-primary"><svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-save" viewBox="0 0 16 16">
            <path d="M2 1a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1H9.5a1 1 0 0 0-1 1v7.293l2.646-2.647a.5.5 0 0 1 .708.708l-3.5 3.5a.5.5 0 0 1-.708 0l-3.5-3.5a.5.5 0 1 1 .708-.708L7.5 9.293V2a2 2 0 0 1 2-2H14a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2h2.5a.5.5 0 0 1 0 1H2z"/>
          </svg> Submit</button>
        </small>
      </form>
    </div>
  </div>
</template>

<script>
import Wails from '@wailsapp/runtime';
import axios from 'axios';
import moment from 'moment';
export default {
  data() {
    return {
      hasSubmitted:false,
      dismissSecs: 5,
      dismissCountDown: 0,
      messageValidate:'',
      validateVarian:'danger',
      isBusy: false,
      totalRows: 1,
      currentPage: 1,
      perPage: 5,
      pageOptions: [5, 10, 15, { value: 100, text: "Show more" }],
      fields: [
        { key: 'RFIDID', label: 'RFID ID', sortable: false, class: 'text-center' },
        { key: 'EmployeeID', label: 'Emp ID', sortable: false, class: 'text-center' },
        { key: 'EmployeeNumber', label: 'Emp No', sortable: false, class: 'text-center' },
        { key: 'EmployeeName', label: 'Emp Name', sortable: false, class: 'text-center' },
        { key: 'LocationCode', label: 'Location', sortable: false, class: 'text-center' },
        { key: 'GroupCode', label: 'Group', sortable: false, class: 'text-center' },
        { key: 'UnitCode', label: 'Unit', sortable: false, class: 'text-center' },
        { key: 'CreatedDate', label: 'Created Date', sortable: false, class: 'text-center',
          formatter: (value, key, item) => {
              return moment(value).format("YYYY-MM-DD hh:mm:ss")
          } 
        },
        { key: 'actions', label: 'Actions', sortable: false, class: 'text-center' }
        
      ],
      form:{
        RFIDID:'',
        EmployeeID:'',
        EmployeeNumber:'',
        EmployeeName:'',
        LocationCode:'',
        GroupCode:'',
        UnitCode:'',
        CreatedDate: moment().format('YYYY-MM-DD hh:mm:ss')
      },
      list: [],
      listDelete: [],
      listInsert: [],
      infoModal: {
        id: 'info-modal',
        title: '',
        content: ''
      }
    };
  },
  mounted:function() {
    this.load();
    Wails.Events.On("rfid", rfid => {
      if (rfid) {
        this.form = {
          RFIDID: rfid.id,
          EmployeeID: '',
          EmployeeNumber: '',
          EmployeeName: '',
          LocationCode: '',
          GroupCode: '',
          UnitCode: '',
          CreatedDate: moment().format('YYYY-MM-DD hh:mm:ss')
        };
        if( this.list.some(item => item.RFIDID === rfid.id) ){
          this.validateVarian='danger'
          this.messageValidate='Data already exist !'
          this.dismissCountDown = this.dismissSecs
        }else{
          this.validate()
        }
      }
    });

  },
  methods:{
    handleDelete(item) {
      item.AbsentDate = this.$route.params.period,
      item.AbsentType = this.$route.params.absenttype 
      this.listDelete.push(item)
      this.deleteFromList(item.RFIDID)
    },
    deleteFromList(key){
      for (let [i, item] of this.listInsert.entries()) {
          if (item.RFIDID == key) {
              this.listInsert.splice(i, 1);
          }
      }

      for (let [i, item] of this.list.entries()) {
          if (item.RFIDID == key) {
              this.list.splice(i, 1);
          }
      }
    },
    info(item, index, button) {
      // this.infoModal.title = `Row index: ${index}`
      this.infoModal.title = `Confirmation Delete`
      this.infoModal.content = item
      this.$root.$emit('bv::show::modal', this.infoModal.id, button)
    },
    resetInfoModal() {
      this.infoModal.title = ''
      this.infoModal.content = ''
    },
    countDownChanged(dismissCountDown) {
      this.dismissCountDown = dismissCountDown
    },
    showAlert() {
      this.dismissCountDown = this.dismissSecs
    },
    toggleBusy() {
      this.isBusy = !this.isBusy
    },
    
    load(){
      this.toggleBusy()
      var params = {
        AbsentDate: this.$route.params.period,
        AbsentType: this.$route.params.absenttype 
      }
      window.backend.RFID.GetListAbsenteeism(params).then(result => {
        this.list = result
        this.totalRows = this.list.length
        this.toggleBusy()
      });

      // axios.get('http://localhost:9090/id/RFIDAPI/GetListAbsenteeism',{
      //   params: {
      //     AbsentDate: this.$route.params.period,
      //     AbsentType: this.$route.params.absenttype 
      //   }
      // }).then(resp => {
      //   this.list = resp.data.data
      //   this.totalRows = this.list.length
      //   this.toggleBusy()
      // })
    },

    validate(){
      axios.get('http://localhost:9090/id/RFIDAPI/ValidateRFID',{
        params: {
          Date: this.$route.params.period,
          AbsentType: 'Absenteeism', 
          RFIDID: this.form.RFIDID
        }
      }).then(resp => {
        if(resp.data.status=="success"){
          resp.data.data.CreatedDate = moment().format('YYYY-MM-DD hh:mm:ss')
          resp.data.data.AbsentDate = this.$route.params.period
          resp.data.data.AbsentType = this.$route.params.absenttype 
          this.listValidate=resp.data.data
          this.listInsert.push(this.listValidate)
          this.list.unshift(this.listValidate)
          console.log(this.list)
          this.validateVarian='info'
          this.messageValidate='RFID '+resp.data.data.RFIDID+' - '+resp.data.data.EmployeeName
          this.dismissCountDown = this.dismissSecs
        }else{
          this.validateVarian='danger'
          this.messageValidate=resp.data.message
          this.dismissCountDown = this.dismissSecs
        }
      })
    },
    save(){
      this.hasSubmitted=true
      
      let payload = {
        insert: this.listInsert,
        delete:this.listDelete
      };
      axios({
        url: 'http://localhost:9090/RFIDAPI/InsertDeleteRFIDAbsenteeism',
        method: 'post',
        data: payload
      }).then(resp => {
        this.hasSubmitted=false

        this.validateVarian='success'
        this.messageValidate='Data was successfully saved !'
        this.dismissCountDown = this.dismissSecs
        
        this.load();
        this.listInsert=[]
        this.listDelete=[]
        
      }).catch(function (error) {
        console.log(error);
      });
    }

  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
 
.custom-table{
  font-size: 0.875em;
}

.badge {
     border-radius: 0;
     font-size: 12px;
     line-height: 1;
     padding: .375rem .5625rem;
     font-weight: normal
 }

 .badge-outline-primary {
     color: #405189;
     border: 1px solid #405189
 }

 .badge.badge-pill {
     border-radius: 10rem
 }

 .badge-outline-info {
     color: #3da5f4;
     border: 1px solid #3da5f4
 }

 .badge-outline-danger {
     color: #f1536e;
     border: 1px solid #f1536e
 }

 .badge-outline-success {
     color: #00c689;
     border: 1px solid #00c689
 }

 .badge-outline-warning {
     color: #fda006;
     border: 1px solid #fda006
 }
</style>
