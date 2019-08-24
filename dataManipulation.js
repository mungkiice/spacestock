// Working with Arrays
//1.1 Simple with sort
let husband = ["alphard", "civic", "terios"];
let wife = ["terios", "jazz"];
function merge(arr1, arr2){
	let arr3 = arr1.concat(arr2).sort();
	let result = [];
	arr3.forEach(function(i){
		if(!result.includes(i)){
			result.push(i);
		}
	});
	return result;
}

// 1.2 Object Merge
var invoices = [
{"to": "John Doe", "amount": 120000},
{"to": "John Doe", "amount": 80000},
{"to": "Jane Doe", "amount": 100000}
]

function groupInvoiceByPerson(invoices) {
	var result = invoices.reduce(function(pass, val){
		var item = pass.filter(function(o){
			return o.to == val.to;
		}).pop() || {to:val.to, amount:0};
		item.amount += val.amount;
		pass.push(item);
		return pass;
	}, []);
	return result.filter(function(o, i, arr){
		return i == arr.indexOf(o);
	});
}

// 1.3 Finding similar string
var bannedWords = ["scam", "spam", "dirty"]
var input1 = "this is a scammer"
var input2 = "this is a spam message"
var input3 = "this contains dirty word"
var input4 = "this is a valid message"

function isBanned(words, bannedWords) {
    var result = false;
	bannedWords.forEach(function(bannedWord){
		if (words.includes(bannedWord)) {
           result = true;
		}
	})
    return result;
}
