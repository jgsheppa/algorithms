/**
 * `binarySearch` finds a target number within an
 * ordered list of numbers. The number of guesses
 * required to find the target is returned.
 *
 * @param list
 * @param target
 * @returns number
 */
export function binarySearch<T extends number>(
  list: Array<number>,
  target: T,
): { numberOfGuesses: number; targetFound: boolean } {
  let numberOfGuesses = 0;
  let low = 0;
  let high = list.length - 1;

  while (low <= high) {
    const middle = Math.floor((low + high) / 2);
    const guess = list[middle];

    if (guess === target) {
      numberOfGuesses++;
      return { numberOfGuesses, targetFound: true };
    } else if (guess > target) {
      numberOfGuesses++;
      high = middle - 1;
    } else {
      numberOfGuesses++;
      low = middle + 1;
    }
  }

  return { numberOfGuesses, targetFound: false };
}
