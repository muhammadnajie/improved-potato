const NRP = require('./noRepeatPlease');

test('permAlone("aab") should return 2', () => {
    expect(NRP("aab").toBe(2))
});

test('permAlone("aaa") should return 2', () => {
    expect(NRP("aaa").toBe(0))
});

test('permAlone("aabb") should return 8', () => {
    expect(NRP("aabb").toBe(8))
});

test('permAlone("abcdefa") should return 3600', () => {
    expect(NRP("abcdefa").toBe(3600))
});

test('permAlone("abfdefa") should return 2640', () => {
    expect(NRP("abfdefa").toBe(2640))
});

test('permAlone("zzzzzzzz") should return 0', () => {
    expect(NRP("zzzzzzzz").toBe(0))
});

test('permAlone("a") should return 1', () => {
    expect(NRP("a").toBe(1))
});

test('permAlone("aaab") should return 0', () => {
    expect(NRP("aaab").toBe(0))
});

test('permAlone("aaabb") should return 12', () => {
    expect(NRP("aaabb").toBe(12))
});