export class TestingExports implements Record<string, Function> {
  [x: string]: Function;
  constructor(...methods: Function[]) {
    if (process.env.NODE_ENV === 'test') methods.forEach((fn) => this.add(fn));
  }

  public add(method: Function) {
    this[method.name] = method;
  }
}
