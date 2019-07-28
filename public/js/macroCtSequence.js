let vue = new Vue({
    el: "#content",
	delimiters: ['${', '}$'],
    data: {
        macros: [],
        temp: {
            "html": '<div class="col-md-3 col-sm-6 col-xs-12 temp" id="" style="border: 1px solid #c6c8ca;border-radius: 20px;margin: 10px 0px;padding: 10px 0px;">\n' +
            '            <div class="form-group col-md-12">\n' +
            '                <label>技能名称(不要带空格)：</label>\n' +
            '                <input type="text" class="form-control skillName" name="skillName">\n' +
            '            </div>\n' +
            '            <div class="form-group col-md-12">\n' +
            '                <label>冷却时间(单位：秒，比如6.4，顺发技能请填：0)：</label>\n' +
            '                <input type="text" class="form-control cooldown" name="cooldown">\n' +
            '            </div>\n' +
            '            <div class="form-group col-md-12">\n' +
            '                <label>技能优先级(数字越小，优先级越高！优先级越高，越先施放！)：</label>\n' +
            '                <select class="form-control level" name="level">\n' +
            '                    <option value="1">1</option>\n' +
            '                    <option value="2">2</option>\n' +
            '                    <option value="3">3</option>\n' +
            '                    <option value="4">4</option>\n' +
            '                    <option value="5">5</option>\n' +
            '                    <option value="6">6</option>\n' +
            '                    <option value="7">7</option>\n' +
            '                    <option value="8">8</option>\n' +
            '                    <option value="9">9</option>\n' +
            '                    <option value="10">10</option>\n' +
            '                </select>\n' +
            '            </div>\n' +
            '            <div class="form-group col-md-12">\n' +
            '                <button class="btn btn-info remove" v-on:click="remove(this)">删除</button>\n' +
            '            </div>' +
            '        </div>',
        },
        tempNum: 1,
    },
    methods: {
        create: function() {
            let temps = document.getElementsByClassName("temp")

            let len = temps.length;
            let sequences = [];
            let isExist = [];
            for (let i = 0; i < len; i++) {
                let name = temps[i].getElementsByClassName("skillName")[0].value;
                let cooldown = temps[i].getElementsByClassName("cooldown")[0].value;
                let level = parseInt(temps[i].getElementsByClassName("level")[0].value);

                name = name.replace(/s|xA0/g, "");
                if (name === "") {
                    alert("第" + (i + 1) + "个技能名称为空，请修改！");
                    return;
                }
                if (cooldown === "" || cooldown === 0) {
                    alert("第" + (i + 1) + "个技能冷却时间为空，请修改！");
                    return;
                }
                if (isExist.indexOf(level) !== -1) {
                    alert("第" + (i + 1) + "个技能优先级重复，请修改！");
                    return;
                } else {
                    isExist.push(level)
                }

                let sequence = {
                    "skillName": name,
                    "cooldown": parseInt(cooldown * 100),
                    "level": level,
                };
                sequences.push(sequence);
            }

            axios.post('/macro/createSequence', sequences)
                .then(function(response) {
                    vue.$data.macros = [];
                    let text = {
                        "id": 1,
                        "text": response.data.text,
                        "desc": response.data.desc,
                    };
                    vue.$data.macros.push(text);
                });
        },
        addTemp: function() {
            if (vue.$data.tempNum >= 10) {
                alert("技能放的够多了，别贪心哦！！！");
            } else {
                let temp = document.createElement("div");
                temp.innerHTML = vue.$data.temp.html;
                let btn = temp.getElementsByClassName("remove")[0];
                btn.id = "remove_btn" + vue.$data.tempNum;
                btn.onclick = function() {
                    this.parentNode.parentNode.parentNode.removeChild(this.parentNode.parentNode);
                }
                vue.$data.tempNum += 1;

                document.getElementById("content").insertBefore(temp, document.getElementById("btn_div"));
            }

        }
    },
    created: function() {
    }
})
vue.addTemp();