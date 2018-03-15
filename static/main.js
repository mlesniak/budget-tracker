var app = new Vue({
    el: '#app',
    data: {
        transactions: [],
        budget: {},

        // For entering a new transaction.
        amount: "",
        category: ""
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
        },

        add() {
            var self = this;
            axios.post('/api/transaction', {
                category: this.category,
                amount: this.amount
            }).then(function(response) {
                self.fetchBudget();
            });
        },

        sub() {
            var self = this;
            axios.post('/api/transaction', {
                category: this.category,
                amount: "-" + this.amount
            }).then(function(response) {
                self.fetchBudget();
            });
        }
    }
})

