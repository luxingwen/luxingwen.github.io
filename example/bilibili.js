var imgs = document.getElementsByClassName("img-content");
var reg = /[\w\/\.]+.jpg/;
var str = "";
for(var i = 0; i < imgs.length; i ++){
	var bgImg = imgs[i].style.backgroundImage;
	var imgReg = reg.exec(bgImg);
	if (imgReg==null || imgReg[0]==undefined){
		continue;
	}
	var address = "http:"+imgReg[0];
	str += "<img width =\"500\" src=\"" + address + "\"> \n";
}
console.log(str);
