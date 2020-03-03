var app = new Vue({
    el: '#app',
    data: {
        items: null,
        keyword: '',
        message: ''
    },
    watch: {

    },
    created: function () {
        this.keyword = '静か'
        this.getAnswer()
    },
    methods: {
        getAnswer: function() {
            if(this.keyword === '') {
                this.items = null
                return
            }
            this.message = 'Loading...'
            var vm = this
            var params = { keyword: this.keyword, limit: 20 }
            axios.get('http://localhost:9000/web/search/suggestion', { params })
              .then(function(response){
                  console.log(response)
              })
              .catch(function (error) {
                  vm.message = 'Error!' + error
              })
              .finally(function () {
                  vm.message = ''
              })
        }
    }
})