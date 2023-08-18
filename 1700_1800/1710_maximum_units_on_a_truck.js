var maximumUnits = function(boxTypes, truckSize) {
    let ans = 0;

    boxTypes.sort((box1, box2) => box1[1] - box2[1]);
    for (let i = boxTypes.length - 1; i >= 0 && truckSize > 0; i--) {
        if (boxTypes[i][0] <= truckSize) {
            ans += boxTypes[i][0] * boxTypes[i][1];
            truckSize -= boxTypes[i][0];
        } else {
            ans += truckSize * boxTypes[i][1];
            truckSize = 0;
        }
    }

    return ans;
};
