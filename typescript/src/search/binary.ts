export function binarySearch<T extends number>(
  list: Array<number>,
  target: T,
): number {
  let numberOfGuesses = 0;
  let low = 0;
  let high = list.length;

  while (low <= high) {
    const middle = Math.floor((low + high) / 2);
    const guess = list[low];
    if (guess === target) {
      numberOfGuesses++;
      console.log('Found target! Number of guesses:', numberOfGuesses);
      return numberOfGuesses;
    } else if (guess < target) {
      numberOfGuesses++;
      high = middle - 1;
    } else {
      numberOfGuesses++;
      low = middle + 1;
    }
  }

  return numberOfGuesses;
}
