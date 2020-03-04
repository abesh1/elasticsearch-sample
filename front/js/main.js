var app = new Vue({
    el: '#app',
    data: {
        selected: 1,
        options: [
            { text: '前方一致', value: 1 },
            { text: '部分一致', value: 2 },
            { text: 'Web', value: 3 }
        ],
        titles: null,
        authors: null,
        keyword: '',
        message: '',
    },
    watch: {
        keyword: function (newKeyword, oldKeyword) {
            this.debouncedGetAnswer()
        },
        selected: function (newSelected, oldSelected) {
            this.debouncedGetAnswer()
        }
    },
    created: function () {
        this.debouncedGetAnswer = _.debounce(this.getAnswer, 1000)
    },
    methods: {
        getAnswer: function() {
            const V5_PREFIX_REQUEST_URL = 'http://localhost:9000/v5_prefix/search/suggestion'
            const V5_PARTIAL_REQUEST_URL = 'http://localhost:9000/v5_partial/search/suggestion'
            const WEB_REQUEST_URL = 'http://localhost:9000/web/search/suggestion'
            if(this.keyword === '') {
                this.titles = null
                this.authors = null
                return
            }
            this.message = 'Loading...'
            var vm = this
            var params = { keyword: this.keyword, limit: 1000 }
            var url = ''
            switch (this.selected) {
                case 1:
                    url = V5_PREFIX_REQUEST_URL
                    break
                case 2:
                    url = V5_PARTIAL_REQUEST_URL
                    break
                case 3:
                    url = WEB_REQUEST_URL
                    break
            }
            axios.get(url, { params })
              .then(function(response){
                  vm.authors = response.data.author.items
                  vm.titles = response.data.product.items
                  console.log(vm.titles)
                  console.log(vm.authors)
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