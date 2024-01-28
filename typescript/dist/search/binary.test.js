"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const node_test_1 = require("node:test");
const node_assert_1 = __importDefault(require("node:assert"));
const binary_1 = require("./binary");
(0, node_test_1.test)('binary search returns correct number of guesses', (t) => __awaiter(void 0, void 0, void 0, function* () {
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
        yield t.test(name, () => {
            const searchResult = (0, binary_1.binarySearch)(input.list, input.target);
            node_assert_1.default.deepEqual(searchResult, expected);
        });
    }
}));
