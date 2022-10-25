function myFunction() {
	// Get the text field
	var copyText = document.getElementById("search");

	// Select the text field
	copyText.select();
	copyText.setSelectionRange(0, 99999); // For mobile devices

	// Copy the text inside the text field
	navigator.clipboard.writeText(copyText.value);

	// Alert the copied text
	alert("Copied the text: " + copyText.value);
}

function copy() {
	var copyText = document.getElementById("search").innerText;
	var elem = document.createElement("textarea");
	document.body.appendChild(elem);
	elem.value = copyText;
	elem.select();
	document.execCommand("copy");
	document.body.removeChild(elem);
	alert("Copied the text");
}
//# sourceMappingURL=choices.js.map