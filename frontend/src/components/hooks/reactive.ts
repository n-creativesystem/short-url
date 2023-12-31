import { useEffect, useReducer } from 'react';

export type ReactiveVar<T> = {
  (newValue?: T | ((value: T) => T)): T;
  subscribe: (handler: Function) => () => void;
  unsubscribe: (handler: Function) => void;
};

type EqualsFunc<T> = (a: T, b: T) => boolean;

export const makeVar = <T extends unknown>(
  initialValue: T,
  equalsFunc?: EqualsFunc<T>
): ReactiveVar<T> => {
  let value = initialValue;
  const subscribers = new Set<Function>();
  const reactiveVar = (newValue?: T | ((value: T) => T)) => {
    if (newValue !== undefined) {
      let nextValue = value;
      let valueChanged;

      if (newValue instanceof Function) {
        nextValue = newValue(value);
      } else {
        nextValue = newValue;
      }

      valueChanged = equalsFunc
        ? !equalsFunc(nextValue, value)
        : nextValue !== value;
      value = nextValue;

      if (valueChanged) {
        subscribers.forEach((handler) => handler(value));
      }
    }
    return value;
  };
  reactiveVar.subscribe = (handler: Function) => {
    subscribers.add(handler);
    return () => subscribers.delete(handler);
  };
  reactiveVar.unsubscribe = (handler: Function) => {
    subscribers.delete(handler);
  };
  return reactiveVar;
};

export const useReactiveVar = <T extends unknown>(
  reactiveVar: ReactiveVar<T>
) => {
  const handler = useReducer((x) => x + 1, 0)[1] as () => void;

  useEffect(() => {
    reactiveVar.subscribe(handler);
    return () => {
      reactiveVar.unsubscribe(handler);
    };
  }, [reactiveVar]);

  return reactiveVar();
};
