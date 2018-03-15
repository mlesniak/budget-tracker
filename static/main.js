var app = new Vue({
    el: '#app',
    data: {
        transactions: [],
        budget: {}
    },
    created() {
        this.fetchTransactions();
        this.fetchBudget();
    },

    methods: {
        fetchTransactions() {
            axios.get('/api/transaction/2018/3').then(response => {
                this.transactions = response.data;
            });
        },

        fetchBudget() {
            axios.get('/api/transaction/2018/3/budget').then(response => {
                this.budget = response.data;
            });
        }
    }
})

