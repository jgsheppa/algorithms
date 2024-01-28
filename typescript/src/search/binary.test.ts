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
      expected: 1,
    },
    {
      name: 'binary search where log n === 4',
      input: {
        list: [1, 2, 3, 4],
        target: 4,
      },
      expected: 2,
    },
    {
      name: 'binary search where log n === 8',
      input: {
        list: [1, 2, 3, 4, 5, 6, 7, 8],
        target: 7,
      },
      expected: 3,
    },
    {
      name: 'binary search where log n === 16',
      input: {
        list: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16],
        target: 13,
      },
      expected: 4,
    },
  ];

  for (let i = 0; i < testCases.length; i++) {
    const { name, input, expected } = testCases[i];
    await t.test(name, () => {
      const searchResult = binarySearch(input.list, input.target);
      assert.strictEqual(searchResult, expected);
    });
  }
});
