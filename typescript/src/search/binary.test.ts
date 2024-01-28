import { test } from 'node:test';
import assert from 'node:assert';
import { binarySearch } from './binary';

test('binary search returns correct number of guesses', async (t) => {
  const testCases = [
    {
      name: 'binary search where log n === 2',
      input: {
        list: [1, 2],
        target: 1,
      },
      expected: { numberOfGuesses: 1, targetFound: true },
    },
    {
      name: 'binary search where log n === 4',
      input: {
        list: [1, 2, 3, 4],
        target: 3,
      },
      expected: { numberOfGuesses: 2, targetFound: true },
    },
    {
      name: 'binary search where log n === 8',
      input: {
        list: [1, 2, 3, 4, 5, 6, 7, 8],
        target: 7,
      },
      expected: { numberOfGuesses: 3, targetFound: true },
    },
    {
      name: 'binary search where log n === 16',
      input: {
        list: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16],
        target: 13,
      },
      expected: { numberOfGuesses: 4, targetFound: true },
    },
    {
      name: 'binary search where number is not in list',
      input: {
        list: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16],
        target: 17,
      },
      expected: { numberOfGuesses: 5, targetFound: false },
    },
  ];

  for (let i = 0; i < testCases.length; i++) {
    const { name, input, expected } = testCases[i];
    await t.test(name, () => {
      const searchResult = binarySearch(input.list, input.target);
      assert.deepEqual(searchResult, expected);
    });
  }
});
