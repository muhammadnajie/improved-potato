function updateInventory(arr1, arr2) {
    const inventory = [...arr1];

    let exist;
    arr2.forEach(arr2Inv => {
        exist = false;

        inventory.filter(inv => {
            if (inv[1] === arr2Inv[1]) {
                inv[0] += arr2Inv[0];
                exist = true;
            }
        });
        
        if (!exist) {
            inventory.push(arr2Inv);
        }
    });

    //sorting
    inventory.sort((a, b) => {
        //assume each item unique
        return a[1] < b[1] ? -1 : 1;
    });
    
    return inventory;
}

// Example inventory lists
var curInv = [
    [21, "Bowling Ball"],
    [2, "Dirty Sock"],
    [1, "Hair Pin"],
    [5, "Microphone"]
];

var newInv = [
    [2, "Hair Pin"],
    [3, "Half-Eaten Apple"],
    [67, "Bowling Ball"],
    [7, "Toothpaste"]
];

updateInventory(curInv, newInv);