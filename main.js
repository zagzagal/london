
if (!WebAssembly.instantiateStreaming) { // polyfill
	WebAssembly.instantiateStreaming = async (resp, importObject) => {
		const source = await (await resp).arrayBuffer();
		return await WebAssembly.instantiate(source, importObject);
	};
}
const go = new Go();
let mod, inst;
WebAssembly.instantiateStreaming(fetch("lib.wasm"), go.importObject).then(
	async result => {
		mod = result.module;
		inst = result.instance;
		await go.run(inst);
	}
);


function process() {
	var max = document.getElementById('trys').value;
	var p = document.getElementById('pool').value;
	var e = document.getElementById('eval').value;
	var mHand = document.getElementById('minHand').value;
	var bar = document.getElementById('progBar');

	bar.max = max;
	var s = 0;
	var tick = max/100;

	/*      for (i=0;i<max;i++) {
		var re = stuff(p, e);
		if (re==true) {
			s++;
		}
	}*/
	var i = 0;
	if (validateInput(p, e, mHand)){
		(function doSort() {
			if (i%tick==0) {
				bar.value=i;
			}
			var re = deckCheck(p,e,mHand);
			if (re==true) {
				s++;
			}
			i++;
			if (i<max) {
				setTimeout(doSort,0);
			} else {
				bar.value=max;
				outProgress(s, max);
			}
		})();
	} else {
		document.getElementById("output").innerHTML = 
			`Invalid Input`;
	}

}

function outProgress(s, max) {
	var percent = s/max*100;
	document.getElementById("output").innerHTML = `${s} of ${max} [${percent}%]`;
}

