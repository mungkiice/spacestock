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