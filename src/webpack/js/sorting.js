function sortObjectByValue(list) {
  let keysSorted = Object.keys(list)
                   .sort((a, b) => {
                     return list[b] - list[a]
                   })

  return keysSorted
}

function stringRepeat(string, number) {
  let arr = []
  for (let i = 1; i <= number; i ++) {
    arr.push(string)
  }

  return arr.join('')
}

export {sortObjectByValue, stringRepeat}
