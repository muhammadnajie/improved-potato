function sym(args) {
  let result = [];
  for(let i = 0; i < arguments.length; i++) {
    arguments[i] = set(arguments[i]);
    result = [...result, ...arguments[i]].sort();

    const unique = set(result);

    unique.forEach(val => {
      let firstIdx = result.indexOf(val);
      let lastIdx = result.lastIndexOf(val);
      if (firstIdx != lastIdx)
      {
        result.splice(firstIdx, lastIdx - firstIdx + 1);
      }
    });
  }

  return result;
}

function set(arr) {
  return arr.filter((currValue, idx, self) => self.indexOf(currValue) === idx);
}

console.log(sym([1, 2, 3, 3], [5, 2, 1, 4]))