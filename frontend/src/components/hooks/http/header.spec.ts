import { testingExports } from './header';

describe('Testing of non-public methods.', () => {
  describe('IsHeaders', () => {
    it('Should be false', () => {
      expect(testingExports.isHeaders({})).toBeFalsy();
    });
    it('Should be true', () => {
      const headers = new Headers();
      headers.append('X-AAA', 'aaa');
      expect(testingExports.isHeaders(headers)).toBeTruthy();
    });
  });

  describe('arrToObject', () => {
    it('Should be Array to Object', () => {
      const arr = [
        ['X-AAA', 'aaa'],
        ['X-BBB', 'bbb'],
        ['X-CCC', ''],
      ];
      expect(testingExports.arrToObject(arr)).toEqual({
        'x-aaa': 'aaa',
        'x-bbb': 'bbb',
      });
    });
  });

  describe('headersToObject', () => {
    it('Should be HTTP Headers to Object', () => {
      const headers = new Headers();
      headers.append('X-AAA', 'aaa');
      headers.append('X-BBB', 'bbb');
      headers.append('X-CCC', '');
      expect(testingExports.headersToObject(headers)).toEqual({
        'x-aaa': 'aaa',
        'x-bbb': 'bbb',
      });
    });
  });

  describe('objectToObjectWithValueFilter', () => {
    it('Should be record to object', () => {
      const obj: Record<string, string> = {
        'X-AAA': 'aaa',
        'X-BBB': 'bbb',
        'X-CCC': '',
      };
      expect(testingExports.objectToObjectWithValueFilter(obj)).toMatchObject({
        'x-aaa': 'aaa',
        'x-bbb': 'bbb',
      });
    });
  });
});
