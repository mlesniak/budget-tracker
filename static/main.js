Vue.component('budget-display', {
    template: "#budget",
    props: ["balance", "daily", "remainingDays"]
});


Vue.component('cookie-input', {
    template: "#cookie-input",
    data: function() {
        return {
            secret: undefined
        };
    }, methods: {
        submitCookie(e) {
            console.log("Storing cookie", this.secret);
        }
    }
});


Vue.component('transactions', {
    template: "#transactions",
    props: ["transactions"],
    data: function() {
        return {
            active: false
        }
    },
    methods: {
        getFormattedDate(date) {
            var year = date.getFullYear();
            var month = (1 + date.getMonth()).toString();
            month = month.length > 1 ? month : '0' + month;
            var day = date.getDate().toString();
            day = day.length > 1 ? day : '0' + day;
            return day + "." + month;
        }
    }
});


Vue.component('transaction-input', {
    template: "#transaction-input",
    data: function () {
        return {
            description: undefined,
            amount: undefined,
            show: true
        }
    },
    methods: {
        add() {
            var self = this;
            axios.post('/api/transaction', {
                description: this.description,
                amount: this.amount
            }).then(function (response) {
                app.$emit('update');
                self.clear();
            });
        },
        sub() {
            var self = this;
            axios.post('/api/transaction', {
                description: this.description,
                amount: "-" + this.amount
            }).then(function (response) {
                app.$emit('update');
                self.clear();
            });
        },
        clear() {
            this.description = undefined;
            this.amount = undefined;
            // Trick to reset/clear native browser form validation state.
            // See https://bootstrap-vue.js.org/docs/components/form/
            this.show = false;
            this.$nextTick(() => { 
                this.show = true;
                this.$nextTick(() => {
                    this.$refs.input.focus();
                });
            });
        }
    }
});

var app = new Vue({
    el: '#app',
    data: {
        transactions: [],
        budget: {}
    },
    created() {
        this.$on('update', function() {
            // TODO ML Move to transactions component!
            this.fetchTransactions();
            // TODO ML Move to budget component!
            this.fetchBudget();
        });
        this.$emit('update');
    },

    methods: {
        getDatePath() {
            var d = new Date();
            return (d.getYear() + 1900) + "/" + (d.getMonth() + 1);
        },
        fetchTransactions() {
            axios.get('/api/transaction/' + this.getDatePath()).then(response => {
                this.transactions = response.data;
                this.transactions.reverse();
            });
        },
        fetchBudget() {
            axios.get('/api/transaction/' + this.getDatePath() + '/budget').then(response => {
                this.budget = response.data;
            });
        },
    }
})

