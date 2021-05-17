// vue
const vm = new Vue({
    delimiters: ['@{', '}'],
    el: '#app',
    data() {
      return {
          a: 'a',
          name: 'zhou'
      }
    },
    mounted() {
        this.getList()
    },
    methods: {
        getList() {
            var that = this
            axios.get('/api/ping')
            .then(function (response) {
                
                const data = response.data
                if (data.code != 0) {
                    console.log(data.message)
                }
                that.name = data.content
            })
            .catch(function (error) {
                console.log(error);
            });
        },
        hello() {
            alert("hello")
        },
        goScan() {
            window.location.href = '/scan'
        },
        goCurr() {
            window.location.href = '/curr'
        },
        pcFeature() {
            window.location.href = '/pc/feature'
        }
    }
  })

  //console.log(vm.name)