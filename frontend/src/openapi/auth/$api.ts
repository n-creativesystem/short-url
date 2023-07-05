import type { AspidaClient, BasicHeaders } from 'aspida'
import type { Methods as Methods0 } from './_provider/authorize'
import type { Methods as Methods1 } from './_provider/callback'
import type { Methods as Methods2 } from './enabled'
import type { Methods as Methods3 } from './logout'
import type { Methods as Methods4 } from './userinfo'

const api = <T>({ baseURL, fetch }: AspidaClient<T>) => {
  const prefix = (baseURL === undefined ? '/api' : baseURL).replace(/\/$/, '')
  const PATH0 = '/auth'
  const PATH1 = '/authorize'
  const PATH2 = '/callback'
  const PATH3 = '/auth/enabled'
  const PATH4 = '/auth/logout'
  const PATH5 = '/auth/userinfo'
  const GET = 'GET'

  return {
    _provider: (val0: number | string) => {
      const prefix0 = `${PATH0}/${val0}`

      return {
        authorize: {
          get: (option?: { config?: T | undefined } | undefined) =>
            fetch<void, Methods0['get']['resHeaders'], Methods0['get']['status']>(prefix, `${prefix0}${PATH1}`, GET, option).send(),
          $get: (option?: { config?: T | undefined } | undefined) =>
            fetch<void, Methods0['get']['resHeaders'], Methods0['get']['status']>(prefix, `${prefix0}${PATH1}`, GET, option).send().then(r => r.body),
          $path: () => `${prefix}${prefix0}${PATH1}`
        },
        callback: {
          get: (option?: { config?: T | undefined } | undefined) =>
            fetch<void, Methods1['get']['resHeaders'], Methods1['get']['status']>(prefix, `${prefix0}${PATH2}`, GET, option).send(),
          $get: (option?: { config?: T | undefined } | undefined) =>
            fetch<void, Methods1['get']['resHeaders'], Methods1['get']['status']>(prefix, `${prefix0}${PATH2}`, GET, option).send().then(r => r.body),
          $path: () => `${prefix}${prefix0}${PATH2}`
        }
      }
    },
    enabled: {
      /**
       * @returns OK
       */
      get: (option?: { config?: T | undefined } | undefined) =>
        fetch<Methods2['get']['resBody'], BasicHeaders, Methods2['get']['status']>(prefix, PATH3, GET, option).json(),
      /**
       * @returns OK
       */
      $get: (option?: { config?: T | undefined } | undefined) =>
        fetch<Methods2['get']['resBody'], BasicHeaders, Methods2['get']['status']>(prefix, PATH3, GET, option).json().then(r => r.body),
      $path: () => `${prefix}${PATH3}`
    },
    logout: {
      get: (option?: { config?: T | undefined } | undefined) =>
        fetch<void, Methods3['get']['resHeaders'], Methods3['get']['status']>(prefix, PATH4, GET, option).send(),
      $get: (option?: { config?: T | undefined } | undefined) =>
        fetch<void, Methods3['get']['resHeaders'], Methods3['get']['status']>(prefix, PATH4, GET, option).send().then(r => r.body),
      $path: () => `${prefix}${PATH4}`
    },
    userinfo: {
      /**
       * @returns OK
       */
      get: (option?: { config?: T | undefined } | undefined) =>
        fetch<Methods4['get']['resBody'], BasicHeaders, Methods4['get']['status']>(prefix, PATH5, GET, option).json(),
      /**
       * @returns OK
       */
      $get: (option?: { config?: T | undefined } | undefined) =>
        fetch<Methods4['get']['resBody'], BasicHeaders, Methods4['get']['status']>(prefix, PATH5, GET, option).json().then(r => r.body),
      $path: () => `${prefix}${PATH5}`
    }
  }
}

export type ApiInstance = ReturnType<typeof api>
export default api
