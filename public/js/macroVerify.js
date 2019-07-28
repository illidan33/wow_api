let vue = new Vue({
	el: "#container",
	delimiters: ['${', '}$'],
	data: {
		professions: [
			{"id": 1, "name": "猎人"},
			{"id": 10, "name": "德鲁伊"},
			{"id": 20, "name": "死亡骑士"},
			{"id": 30, "name": "法师"},
			{"id": 40, "name": "圣骑士"},
			{"id": 50, "name": "牧师"},
			{"id": 60, "name": "盗贼"},
			{"id": 70, "name": "萨满祭司"},
			{"id": 80, "name": "术士"},
			{"id": 90, "name": "战士"},
			{"id": 100, "name": "武僧"},
			{"id": 110, "name": "恶魔猎手"},
			{"id": 120, "name": "其他"},
		],
		macros: [],
	},
	methods: {
		search: function() {
			let verify = document.getElementById("verify");
			isVerify = verify.value;
			let professionId = document.getElementById("profession").value;
			axios.get('/macro/macroList', {
				params: {
					isVerify: isVerify,
					professionId: professionId,
				}
			})
				.then(function(response) {
					// console.log(response);
					vue.$data.macros = response.data.data;
				})
				.catch(function(error) {
					// console.log(error);
				});
		},
		update: function(id) {
			// console.log(id);
			let isVerify = document.getElementById("verify_" + id).value;
			let author = document.getElementById("author_" + id).value;
			let title = document.getElementById("title_" + id).value;
			let macro = document.getElementById("macro_" + id).value;

			axios.put('/macro/updateMacro', {
					id: parseInt(id),
					isVerify: parseInt(isVerify),
					author: author,
					title: title,
					macro: macro,
				}
			).then(function(response) {
				// console.log(response);
				if ((response.data.result == "true" || response.data.result == true) &&
					(isVerify === 1 || isVerify === "1")) {
					document.getElementById("verify_" +
						id).parentElement.parentElement.style.display = "none";
				}
			});

			// console.log(id);
		}
	},
})