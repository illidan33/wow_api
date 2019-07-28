let vue = new Vue({
	el: "#container",
	delimiters: ['${', '}$'],
	data: {
		commands: [],
		conditions: [],
		templates: [],
		condition_radio: 2,
		selectCommandKey: '',
		selectCommand: '',
		select_item: '',//多选条件，需要分割开
		selectConditions: [],
		select_template: '',//组合模板是很多条件
		selectTemplates: [],
		skillName: '',
		macro: '',//组合条件组合成的宏
		allMacro: ''//加上头部tip的宏
	},
	methods: {
		insertMacro: function() {
			// console.log('insertMacro_'+vue.$datacondition_radio);
			vue.$data.skillName = vue.$data.skillName.replace(/(^\s*)|(\s*$)/g, "");
			//自己组合条件
			if (vue.$data.condition_radio === 2) {
				//分割多选条件
				vue.$data.selectConditions = '[' + vue.$data.select_item + ']';

				let changeLine = '';//是否换行
				//如果是castsequence是特殊的
				if (vue.$data.selectCommand === '/castsequence') {
					let s = ' reset=12/combat/target';
					if (vue.$data.macro) {
						changeLine = "\r\n";
					}
					vue.$data.macro += changeLine + vue.$data.selectCommand + s + ' ' +
						vue.$data.skillName;
				} else {
					if (vue.$data.macro) {
						let ok = confirm('是否换行');
						if (ok) {
							changeLine = "\r\n";
							vue.$data.macro += changeLine + vue.$data.selectCommand + ' ' +
								vue.$data.selectConditions + vue.$data.skillName;
						} else {
							changeLine = ';';
							vue.$data.macro += changeLine + vue.$data.selectConditions +
								vue.$data.skillName;//不换行不加命令
						}
					} else {
						if (!vue.$data.selectCommand) {
							alert('请选择命令');
						} else {
							vue.$data.macro = changeLine + vue.$data.selectCommand + ' ' +
								vue.$data.selectConditions + vue.$data.skillName;
						}
					}
				}
			} else {
				vue.$data.macro = '';
				vue.$data.selectTemplates = vue.$data.select_template.split(';');
				// console.log(vue.$data.selectTemplates);
				let len = vue.$data.selectTemplates.length;
				let splitStr = '';
				for (let i = 0; i < len; i++) {
					if (vue.$data.macro) {
						splitStr = ';';
					}
					// console.log(vue.$data.selectTemplates[i]);
					vue.$data.macro += splitStr + vue.$data.selectTemplates[i] + vue.$data.skillName;
				}
			}
			vue.$data.allMacro = "#showtooltip\r\n" + vue.$data.macro;
		},
		copyHong: function() {
			document.getElementById("hong_text").select();
			document.execCommand("Copy");
			alert("已复制，可以进入游戏在宏命令中新建宏，然后CTRL+V粘贴进去了哦^-^!");
		},
		changeCommand: function() {
			vue.$data.selectCommand = vue.$data.commands[vue.$data.selectCommandKey].name;
			let key = vue.$data.selectCommandKey;
			let val = vue.$data.commands[key];

			let con = document.getElementById("condition");
			if (val.type !== '1') {
				con.style.display = "none";
			} else {
				con.style.display = "block";
			}
		}
	},
});

vue.$data.commands = [
	{index: 0, desc: '释放技能或使用物品(同use)', name: '/cast', type: '1'},
	{index: 1, desc: '随机施放', name: '/castrandom', type: '1'},
	{index: 2, desc: '顺序施放', name: '/castsequence', type: '1'},
	{index: 3, desc: '使用物品', name: '/use', type: '1'},
	{index: 4, desc: '随机使用物品', name: '/userandom', type: '1'},
	{index: 5, desc: '停止执行宏', name: '/stopmacro', type: '1'},
	{index: 6, desc: '切换技能条', name: '/changeactionbar', type: '1'},
	{index: 7, desc: '切换敌方目标', name: '/targetenemy', type: '1'},
	{index: 8, desc: '切换目标', name: '/target', type: '1'},
	{index: 9, desc: '设置上个目标为目标', name: '/targetlasttarget', type: '1'},
	{index: 10, desc: '设置某个目标为焦点', name: '/focus', type: '1'},
	{index: 11, desc: '装备武器', name: '/equip', type: '0'},
	{index: 12, desc: '取消状态/技能', name: '/cancelaura', type: '0'},
	{index: 13, desc: '取消变形形态', name: '/cancelform', type: '0'},
	{index: 14, desc: '清除焦点', name: '/clearfocus', type: '0'},
	{index: 15, desc: '开始攻击', name: '/startattack', type: '0'},
	{index: 16, desc: '停止攻击', name: '/stopattack', type: '0'},
	{index: 17, desc: '清除目标', name: '/cleartarget', type: '0'},
	{index: 18, desc: '协助', name: '/assist', type: '0'},
	{index: 19, desc: '宠物开始攻击', name: '/petattack', type: '0'},
	{index: 20, desc: '宠物跟随', name: '/petfollow', type: '0'},
	{index: 21, desc: '宠物协助', name: '/petpassive', type: '0'},
	{index: 22, desc: '宠物防御', name: '/petdefensive', type: '0'},
	{index: 23, desc: '宠物待在某地', name: '/petstay', type: '0'},
	{index: 24, desc: '宠物技能开启自动施放', name: '/petautocaston', type: '0'},
	{index: 25, desc: '宠物技能关闭自动施放', name: '/petautocastoff', type: '0'},
	{index: 26, desc: '召唤坐骑', name: '/mount', type: '0'},
	{index: 27, desc: '解散坐骑', name: '/dismount', type: '0'},
	{index: 28, desc: '在当前频道用白字说', name: '/S', type: '0'},
	{index: 29, desc: '在当前频道用红字喊话', name: '/Y', type: '0'},
	{index: 30, desc: '在小队频道说', name: '/P', type: '0'},
	{index: 31, desc: '在团队频道说', name: '/RA', type: '0'},
	{index: 32, desc: '表情命令', name: '/E', type: '0'},
	{index: 33, desc: '延迟9秒大喊说话', name: '/IN 9 /Y', type: '0'},
];

vue.$data.conditions = [
	{desc: '战斗状态', name: 'combat', type: '1'},
	{desc: '非战斗状态', name: 'nocombat', type: '1'},
	{desc: '潜行', name: 'stealth', type: '2'},
	{desc: '非潜行', name: 'nostealth', type: '2'},
	{desc: '房间内', name: 'indoors', type: '3'},
	{desc: '房间外', name: 'outdoors', type: '3'},
	{desc: '可以飞行', name: 'flyable', type: '4'},
	{desc: '不可飞行', name: 'noflyable', type: '4'},
	{desc: '游泳中', name: 'swimming', type: '5'},
	{desc: '非游泳中', name: 'noswimming', type: '5'},
	{desc: 'shift按键', name: 'mod:shift', type: '6'},
	{desc: 'ctrl按键', name: 'mod:ctrl', type: '6'},
	{desc: 'alt按键', name: 'mod:alt', type: '6'},
	{desc: '无mod状态（没有按shift、alt、ctrl等键）', name: 'nomod', type: '6'},
	{desc: '目标：友方', name: 'help', type: '7'},
	{desc: '目标：敌方', name: 'harm', type: '7'},
	{desc: '目标：存活', name: 'nodead', type: '8'},
	{desc: '目标：死亡', name: 'dead', type: '8'},
	{desc: '目标：存在', name: 'exists', type: '9'},
	{desc: '以自己为目标', name: '@player', type: '10'},
	{desc: '以焦点为目标', name: '@focus', type: '10'},
	{desc: '以鼠标指向为目标', name: '@mouseover', type: '10'},
	{desc: '在鼠标位置施放', name: '@cursor', type: '10'},
	{desc: '当前目标', name: '@target', type: '10'},
	{desc: '以目标的目标为目标', name: '@targettarget', type: '10'},
	{desc: '以宠物为目标', name: '@pet', type: '10'},
	{desc: '当前变形形态:xx', name: 'stance:xx', type: '11'},
	{desc: '不是当前变形形态:xx', name: 'nostance:xx', type: '11'},
	{desc: '正在施放技能xx', name: 'channeling:技能xx', type: '12'},
	{desc: '没有施放技能xx', name: 'nochanneling:技能xx', type: '12'},
	{desc: '没有施放任何技能', name: 'nochanneling', type: '12'},
	{desc: '当前动作条是否是xx', name: 'actionbar:xx', type: '13'},
	{desc: '当前专精是否是xx', name: 'spec:xx', type: '14'},
	{desc: '是否选择天赋第x层第xx个技能', name: 'talent:x/xx', type: '15'},
	{desc: '是否选择pvp天赋第x层第xx个技能', name: 'pvptalent:x/xx', type: '15'},
];

vue.$data.templates = [
	{
		desc: '( 伤害技能 )鼠标指向是敌方，或目标是敌方',
		name: '/cast [@mouseover,harm,nodead];[@target,harm,nodead];[]'
	},
	{
		desc: '( 治疗技能 )鼠标指向是友方，或目标是友方，或给自己使用',
		name: '/cast [@mouseover,help,nodead];[@target,help,nodead];[@player]'
	},
	{
		desc: '（伤害跟治疗技能一键）优先鼠标指向，鼠标指向敌人就用技能2，指向友人就用技能1，鼠标指向没有目标就是对当前目标用技能2(敌对)或者技能1(友善)，都没有目标就对自己用技能1',
		name: "/cast [@mouseover,help,nodead];[@mouseover,harm,nodead][harm,nodead];[]"
	},
	{
		desc: '当前天赋第x层随技能选择,施放选择的技能( 需要自己修改x为对应的天赋层1-6 )',
		name: '/cast [talent:x/1];[talent:x/2];[talent:x/3]'
	},
	{
		desc: '当前荣誉天赋第x层随技能选择,施放选择的技能( 需要自己修改x为对应的天赋层1-6 )',
		name: '/cast [pvptalent:x/1];[pvptalent:x/2];[pvptalent:x/3]'
	},
	{desc: '对目标施放范围技能', name: "/cast [@cursor];[]"},
];