let vue = new Vue({
    el: "#container",
	delimiters: ['${', '}$'],
    data: {
        professions: [
            {'id': 1, 'name': '猎人'},
            {'id': 10, 'name': '德鲁伊'},
            {'id': 20, 'name': '死亡骑士'},
            {'id': 30, 'name': '法师'},
            {'id': 40, 'name': '圣骑士'},
            {'id': 50, 'name': '牧师'},
            {'id': 60, 'name': '盗贼'},
            {'id': 70, 'name': '萨满祭司'},
            {'id': 80, 'name': '术士'},
            {'id': 90, 'name': '战士'},
            {'id': 100, 'name': '武僧'},
            {'id': 110, 'name': '恶魔猎手'},
            {'id': 120, 'name': '其他'},
        ],
        select_pid: '',
        macros: [],
        select_macro: []
    },
    methods: {
        search: function() {
            let pid = document.getElementById("profession").value;
            if (!pid) {
                alert('职业或专精未选择！');
                return false;
            }
            vue.$data.select_macro = [];
            vue.$data.macros = [];
            document.getElementById("first_macro").innerHTML = "请选择一条宏模板";

            axios.get('/macro/macroList', {
                params: {
                    professionId: pid,
                }
            })
                .then(function(response) {
                    vue.$data.macros = response.data.data;
                })
                .catch(function(error) {
                    console.log(error);
                })
                .then(function() {
                    // always executed
                });
        },
        changeMacro: function() {
            let id = document.getElementById("macro_id").value;
            if (!id) {
                vue.$data.select_macro = [];
                return;
            }
            axios.get('/macro/macroList', {
                params: {
                    id: id,
                }
            })
                .then(function(response) {
                    vue.$data.select_macro = response.data.data[0];
                })
                .catch(function(error) {
                    console.log(error);
                })
                .then(function() {
                    // always executed
                });
        },
        copyMacro: function() {
            document.getElementById("macro_text").select();
            document.execCommand("Copy");
            alert("已复制^-^！");
        },
    },
});
