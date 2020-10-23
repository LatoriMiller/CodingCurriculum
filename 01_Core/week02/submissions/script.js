   //Notes: Make an input field to push the items into an array
   
   // Stack overflow solution
    //https://stackoverflow.com/questions/39050438/how-to-sort-and-format-last-names-in-an-array-alphabetically-in-javascript
    var nameArray = ["Latori Miller", "Lucy Yang", "Reggie Johnson", "Dyl Trimble", "Leland Capps"];

    var namesLastFirst = [];
    function reverseName(){
        for (var i=0; i < nameArray.length; i++){
            console.log(nameArray[i]);
            //split each Name first step
            var newName = nameArray[i].split(" ");
            // reversing the names below
            namesLastFirst.push(`${newName[1]} ${newName[0]}`);            
        }
    }

    reverseName();

    function showNames(){
        // sort the new namesLastFirst using the .sort
        namesLastFirst.sort();

        for (i=0; i < namesLastFirst.length; i++){
            console.log(namesLastFirst[i]); 
            var node = document.createElement("LI");
            var textnode = document.createTextNode(`${namesLastFirst[i]}`);       
            node.appendChild(textnode);                  
            document.getElementById("namesList").appendChild(node);    
        }
    }
    showNames();

   
   
    
    // // function alphabetizer(names) {
    // //     return names.map(function(name) {
    // //         var full = name.split(" "),
    // //         last = full.pop();
    // //         return " " + last + " " + full.join(" ");
    // //         //return last + ", " + full.join(" ");
    // //     }).sort();
    // // }  

    // // go break into arrays for each 1st and last name e.g. ["Latori", "Miller"]
    // // remove the last item from each array using pop
    // // then combine and sort
    // function alphabetizer(names) {
    //     return names.map(function(name) {
    //         var full = name.split(" ");
    //         last = full.pop();
    //         return  last + " " + full.join();
    //     }).sort();
    // } 

    // // console.log(alphabetizer(nameArray));
    // var sortedArray = alphabetizer(nameArray);

    // // need to reverse to display names correctly
    // // this function will reverse the names 
    // //Question - Could I have combined the .sort and the .reverse in one function 
    // function displayName(names) {
    //     return names.map(function(name) {
    //         var full = name.split(" ");
    //         full = full.reverse();
    //         last = full.pop();
    //         return  " " + full.join() + " " + last;
    //     })
    // }
    // var correctNames = displayName(sortedArray);
    // // console.log(correctNames);
     

    // // Solution that works for first names only
    // // https://davidwells.io/snippets/sort-an-array-alphabetically-in-javascript
    // // var nonSortedNames = ['latori', 'nastassja', 'artie', 'amanda', 'melanie'];
    // // var sortedNames = nonSortedNames.sort(function (a, b) {
    // //   if (a < b) return -1;
    // //   else if (a > b) return 1;
    // //   return 0;
    // // });
    // // console.log(sortedNames); // ["amanda", "artie", "latori", "melanie", "nastassja"]

    // document.getElementById("names").innerHTML = correctNames;