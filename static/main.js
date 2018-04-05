Vue.component('budget-display', {
    template: "#budget",
    props: ["balance", "daily", "remainingDays"]
});


var cookieComponent = Vue.component('cookie-input', {
    template: "#cookie-input",
    data: function () {
        return {
            secret: undefined
        };
    }, methods: {
        submitCookie(e) {
            // TODO ML Set secure and enforce https.
            this.createCookie("auth", this.secret, 365);
            console.log("TODO redirect to homepage");
        },
        createCookie(name, value, days) {
            var date = new Date();
            date.setTime(date.getTime() + (days * 24 * 60 * 60 * 1000));
            var expires = "; expires=" + date.toUTCString();
            document.cookie = name + "=" + value + expires + "; path=/";
        }
    }
});


Vue.component('transactions', {
    template: "#transactions",
    // props: ["transactions"],
    data: function () {
        return {
            active: false,
            transactions: []
        }
    },
    created: function() {
        this.$on('update', function () {
            console.log("update received");
            this.fetchTransactions();
        });
        this.fetchTransactions();
    },
    methods: {
        getDatePath() {
            var d = new Date();
            return (d.getYear() + 1900) + "/" + (d.getMonth() + 1);
        },
        getFormattedDate(date) {
            var year = date.getFullYear();
            var month = (1 + date.getMonth()).toString();
            month = month.length > 1 ? month : '0' + month;
            var day = date.getDate().toString();
            day = day.length > 1 ? day : '0' + day;
            return day + "." + month;
        },
        fetchTransactions() {
            axios.get('/api/transaction/' + this.getDatePath()).then(response => {
                this.transactions = response.data;
                this.transactions.reverse();
            });
        },
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

var Homepage = Vue.component('homepage', {
    template: "#homepage"
});


var routes = [
    { path: '/cookie', component: cookieComponent },
    { path: '/', component: Homepage },
];

var router = new VueRouter({
    routes: routes
});

var app = new Vue({
    el: '#app',
    router: router,
    data: {
        budget: {}
    },
    created: function() {
        this.$on('update', function () {
            // TODO ML Move to transactions component!
            // this.fetchTransactions();
            // TODO ML Move to budget component!
            this.fetchBudget();
        });
        this.$emit('update');
        console.log("update emitted");
    },

    methods: {
        getDatePath() {
            var d = new Date();
            return (d.getYear() + 1900) + "/" + (d.getMonth() + 1);
        },
        fetchBudget() {
            axios.get('/api/transaction/' + this.getDatePath() + '/budget').then(response => {
                this.budget = response.data;
            });
        },
    }
})

