var app = new Vue({
    el: '#app',
    data: {
        transactions: []
    },
    created() {
        this.fetchData();
    },

    methods: {
        fetchData() {
            axios.get('/api/transaction/2018/3').then(response => {
                console.log(response.data)
                this.transactions = response.data;
            });
        }
    }
})

