var app = new Vue({
    el: '#app',
    data: {
        prefix: {
            titles: null,
            authors: null,
        },
        partial: {
            titles: null,
            authors: null,
        },
        prefixAndPartial: {
            titles: null,
            authors: null,
        },
        web: {
            titles: null,
            authors: null,
        },
        keyword: '',
        message: '',
        searchMethod: '',
        limit: 20,
    },
    watch: {
        keyword: function (newKeyword, oldKeyword) {
            this.debouncedGetSearchSuggestion()
        },
    },
    created: function () {
        this.debouncedGetSearchSuggestion = _.debounce(this.getSearchSuggestion, 1000)
    },
    methods: {
        getSearchSuggestion: function() {
            const V5_PREFIX_REQUEST_URL = 'https://www.dev.nagisagames.jp/v5_prefix/search/suggestion'
            const V5_PARTIAL_REQUEST_URL = 'https://www.dev.nagisagames.jp/v5_partial/search/suggestion'
            const V5_PREFIX_PARTIAL_REQUEST_URL = 'https://www.dev.nagisagames.jp/v5_prefix_partial/search/suggestion'
            const WEB_REQUEST_URL = 'https://www.dev.nagisagames.jp/web/search/suggestion'

            if(this.keyword === '') {
                this.prefix.titles = null
                this.prefix.authors = null
                this.partial.titles = null
                this.partial.authors = null
                this.prefixAndPartial.titles = null
                this.prefixAndPartial.authors = null
                this.web.titles = null
                this.web.authors = null
                this.searchMethod = ''
                return
            }
            this.message = 'Loading...'
            var vm = this
            var params = { keyword: this.keyword, limit: this.limit }
            this.getPrefixSearchSuggestion(V5_PREFIX_REQUEST_URL, vm, params)
            this.getPartialSearchSuggestion(V5_PARTIAL_REQUEST_URL, vm, params)
            this.getPrefixPartialSearchSuggestion(V5_PREFIX_PARTIAL_REQUEST_URL, vm, params)
            this.getWebSearchSuggestion(WEB_REQUEST_URL, vm, params)
            this.searchMethod = 'サジェスト検索結果：'
        },
        getSearch: function() {
            const V5_PREFIX_REQUEST_URL = 'https://www.dev.nagisagames.jp/v5_prefix/search'
            const V5_PARTIAL_REQUEST_URL = 'https://www.dev.nagisagames.jp/v5_partial/search'
            const V5_PREFIX_PARTIAL_REQUEST_URL = 'https://www.dev.nagisagames.jp/v5_prefix_partial/search'
            const WEB_REQUEST_URL = 'https://www.dev.nagisagames.jp/web/search'

            if(this.keyword === '') {
                this.prefix.titles = null
                this.prefix.authors = null
                this.partial.titles = null
                this.partial.authors = null
                this.prefixAndPartial.titles = null
                this.prefixAndPartial.authors = null
                this.web.titles = null
                this.web.authors = null
                this.searchMethod = ''
                return
            }
            this.message = 'Loading...'
            var vm = this
            var params = { keyword: this.keyword, limit: this.limit }
            this.getPrefixSearch(V5_PREFIX_REQUEST_URL, vm, params)
            this.getPartialSearch(V5_PARTIAL_REQUEST_URL, vm, params)
            this.getPrefixPartialSearch(V5_PREFIX_PARTIAL_REQUEST_URL, vm, params)
            this.getWebSearch(WEB_REQUEST_URL, vm, params)
            this.searchMethod = 'キーワード検索結果：'
        },
        getPrefixSearchSuggestion: function(url, vm, params) {
            axios.get(url, { params })
              .then(function(response){
                  vm.prefix.titles = response.data.product.items
                  vm.prefix.authors = response.data.author.items
              })
              .catch(function (error) {
                  vm.message = 'Error!' + error
              })
              .finally(function () {
                  vm.message = ''
              })
        },
        getPartialSearchSuggestion: function(url, vm, params) {
            axios.get(url, { params })
              .then(function(response){
                  vm.partial.titles = response.data.product.items
                  vm.partial.authors = response.data.author.items
              })
              .catch(function (error) {
                  vm.message = 'Error!' + error
              })
              .finally(function () {
                  vm.message = ''
              })
        },
        getPrefixPartialSearchSuggestion: function(url, vm, params) {
            axios.get(url, { params })
              .then(function(response){
                  vm.prefixAndPartial.titles = response.data.product.items
                  vm.prefixAndPartial.authors = response.data.author.items
              })
              .catch(function (error) {
                  vm.message = 'Error!' + error
              })
              .finally(function () {
                  vm.message = ''
              })
        },
        getWebSearchSuggestion: function(url, vm, params) {
            axios.get(url, { params })
              .then(function(response){
                  vm.web.titles = response.data.product.items
                  vm.web.authors = response.data.author.items
              })
              .catch(function (error) {
                  vm.message = 'Error!' + error
              })
              .finally(function () {
                  vm.message = ''
              })
        },
        getPrefixSearch: function(url, vm, params) {
            axios.get(url, { params })
              .then(function(response){
                  vm.prefix.titles = response.data.items
                  vm.prefix.authors = null
              })
              .catch(function (error) {
                  vm.message = 'Error!' + error
              })
              .finally(function () {
                  vm.message = ''
              })
        },
        getPartialSearch: function(url, vm, params) {
            axios.get(url, { params })
              .then(function(response){
                  vm.partial.titles = response.data.items
                  vm.partial.authors = null
              })
              .catch(function (error) {
                  vm.message = 'Error!' + error
              })
              .finally(function () {
                  vm.message = ''
              })
        },
        getPrefixPartialSearch: function(url, vm, params) {
            axios.get(url, { params })
              .then(function(response){
                  vm.prefixAndPartial.titles = response.data.items
                  vm.prefixAndPartial.authors = null
              })
              .catch(function (error) {
                  vm.message = 'Error!' + error
              })
              .finally(function () {
                  vm.message = ''
              })
        },
        getWebSearch: function(url, vm, params) {
            axios.get(url, { params })
              .then(function(response){
                  vm.web.titles = response.data.items
                  vm.web.authors = null
              })
              .catch(function (error) {
                  vm.message = 'Error!' + error
              })
              .finally(function () {
                  vm.message = ''
              })
        },
    }
})